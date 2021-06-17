package main

import (
	"encoding/json"
	"io"
	"log"
	"sync/atomic"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtspv2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

var (
	outboundVideoTracks map[string]*webrtc.TrackLocalStaticSample = map[string]*webrtc.TrackLocalStaticSample{}
	peerConnectionCount int64
)

// HTTP Handler that accepts an Offer and returns an Answer
// adds outboundVideoTrack to PeerConnection
func doSignaling(ctx echo.Context) error {
	var clientReq = &struct {
		CamID string                    `json:"camId"`
		Offer webrtc.SessionDescription `json:"offer"`
	}{}

	if err := json.NewDecoder(ctx.Request().Body).Decode(&clientReq); err != nil {
		return err
	}

	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{
					"stun:stun.l.google.com:19302",
				},
			},
		},
	})
	if err != nil {
		return err
	}

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		if connectionState == webrtc.ICEConnectionStateDisconnected {
			atomic.AddInt64(&peerConnectionCount, -1)
			if err := peerConnection.Close(); err != nil {
				panic(err)
			}
		} else if connectionState == webrtc.ICEConnectionStateConnected {
			atomic.AddInt64(&peerConnectionCount, 1)
		}
	})

	if _, err = peerConnection.AddTrack(outboundVideoTracks[clientReq.CamID]); err != nil {
		return err
	}

	if err = peerConnection.SetRemoteDescription(clientReq.Offer); err != nil {
		return err
	}

	gatherCompletePromise := webrtc.GatheringCompletePromise(peerConnection)

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		return err
	} else if err = peerConnection.SetLocalDescription(answer); err != nil {
		return err
	}

	<-gatherCompletePromise

	response, err := json.Marshal(*peerConnection.LocalDescription())
	if err != nil {
		return err
	}

	if _, err := ctx.Response().Write(response); err != nil {
		return err
	}

	return nil
}

var rtspURLs map[string]string = map[string]string{
	// "c238": "rtsp://admin:tg12346789g@14.161.28.68:50238/Streaming/channels/101",
	// "c239": "rtsp://admin:tg12346789g@14.161.28.68:50239/Streaming/channels/101",
	// "c240": "rtsp://admin:tg12346789g@14.161.28.68:50240/Streaming/channels/101",
	// "c241": "rtsp://admin:tg12346789g@14.161.28.68:50241/Streaming/channels/101",
	// "c242": "rtsp://admin:tg12346789g@14.161.28.68:50242/Streaming/channels/101",
	// "c238": "rtsp://admin:tg12346789g@10.10.13.238:554/Streaming/channels/101",
	// "c239": "rtsp://admin:tg12346789g@10.10.13.239:554/Streaming/channels/101",
	// "c240": "rtsp://admin:tg12346789g@10.10.13.240:554/Streaming/channels/101",
	// "c241": "rtsp://admin:tg12346789g@10.10.13.241:554/Streaming/channels/101",
	// "c242": "rtsp://admin:tg12346789g@10.10.13.242:554/Streaming/channels/101",
	"c238": "rtsp://10.10.14.60:8554/c238",
	"c239": "rtsp://10.10.14.60:8554/c239",
	"c240": "rtsp://10.10.14.60:8554/c240",
	"c241": "rtsp://10.10.14.60:8554/c241",
	"c242": "rtsp://10.10.14.60:8554/c242",
}

func main() {
	var err error

	for cam := range rtspURLs {
		outboundVideoTracks[cam], err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
			MimeType: "video/h264",
		}, "pion-rtsp", "pion-rtsp")
		if err != nil {
			log.Fatalln(err)
		}
	}

	go rtspConsumer()

	sv := echo.New()
	sv.Use(middleware.Logger())
	sv.Use(middleware.Recover())
	sv.Use(middleware.CORS())

	sv.Static("/", "./static")
	sv.POST("/signaling", doSignaling)

	log.Fatalln(sv.Start(":3030"))
}

// Convert H264 to Annex-B, then write to outboundVideoTrack which sends to all PeerConnections
func rtspConsumer() {
	annexbNALUStartCode := func() []byte { return []byte{0x00, 0x00, 0x00, 0x01} }

	for cam, rtspURL := range rtspURLs {
		go func(cam, rtspURL string) {
			for {
				rtspClient, err := rtspv2.Dial(rtspv2.RTSPClientOptions{
					URL:              rtspURL,
					DialTimeout:      5 * time.Second,
					ReadWriteTimeout: 5 * time.Second,
					DisableAudio:     true,
				})
				if err != nil {
					panic(err)
				}

				codecs := rtspClient.CodecData
				for i, t := range codecs {
					log.Println(cam, "- track", i, "-", t.Type().String())
				}
				if codecs[0].Type() != av.H264 {
					panic("RTSP feed must begin with a H264 codec")
				}
				if len(codecs) != 1 {
					log.Println("Ignoring all but the first stream.")
				}

				var previousTime time.Duration

				for {
					select {
					case pkg := <-rtspClient.OutgoingPacketQueue:
						if pkg.Idx != 0 {
							continue
						}

						pkg.Data = pkg.Data[4:]
						if pkg.IsKeyFrame {
							pkg.Data = append(annexbNALUStartCode(), pkg.Data...)
							pkg.Data = append(codecs[0].(h264parser.CodecData).PPS(), pkg.Data...)
							pkg.Data = append(annexbNALUStartCode(), pkg.Data...)
							pkg.Data = append(codecs[0].(h264parser.CodecData).SPS(), pkg.Data...)
							pkg.Data = append(annexbNALUStartCode(), pkg.Data...)
						}
						bufferDuration := pkg.Time - previousTime
						previousTime = pkg.Time

						if err = outboundVideoTracks[cam].WriteSample(media.Sample{Data: pkg.Data, Duration: bufferDuration}); err != nil && err != io.ErrClosedPipe {
							panic(err)
						}
					case sig := <-rtspClient.Signals:
						if sig == rtspv2.SignalStreamRTPStop {
							goto exit
						}
					}
				}

			exit:
				rtspClient.Close()
				time.Sleep(5 * time.Second)
				log.Println(cam, "timeout")
			}
		}(cam, rtspURL)
	}
}
