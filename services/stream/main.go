package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"sync/atomic"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtspv2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/sethvargo/go-envconfig"
)

var (
	outboundVideoTracks map[string]*webrtc.TrackLocalStaticSample = map[string]*webrtc.TrackLocalStaticSample{}
	peerConnectionCount int64
)

type Config struct {
	TLS          bool   `env:"TLS,default=false"`
	HTTP         string `env:"HTTP_PORT,default=:3030"`
	CertPath     string `env:"CERT_PATH,default=cert/"`
	TurnURL      string `env:"TURN_URL"`
	TurnPassword string `env:"TURN_PASSWORD"`
	TurnUsername string `env:"TURN_USERNAME"`
}

// HTTP Handler that accepts an Offer and returns an Answer
// adds outboundVideoTrack to PeerConnection
func doSignaling(ctx echo.Context, config Config) error {
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
					config.TurnURL,
				},
				Username:       config.TurnUsername,
				Credential:     config.TurnPassword,
				CredentialType: webrtc.ICECredentialTypePassword,
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

func main() {

	// var rtspURLs map[string]string = map[string]string{
	// 	// "c238": "rtsp://admin:tg12346789g@14.161.28.68:50238/Streaming/channels/101",
	// 	// "c239": "rtsp://admin:tg12346789g@14.161.28.68:50239/Streaming/channels/101",
	// 	// "c240": "rtsp://admin:tg12346789g@14.161.28.68:50240/Streaming/channels/101",
	// 	// "c241": "rtsp://admin:tg12346789g@14.161.28.68:50241/Streaming/channels/101",
	// 	// "c242": "rtsp://admin:tg12346789g@14.161.28.68:50242/Streaming/channels/101",
	// 	// "c238": "rtsp://admin:tg12346789g@10.10.13.238:554/Streaming/channels/101",
	// 	// "c239": "rtsp://admin:tg12346789g@10.10.13.239:554/Streaming/channels/101",
	// 	// "c240": "rtsp://admin:tg12346789g@10.10.13.240:554/Streaming/channels/101",
	// 	// "c241": "rtsp://admin:tg12346789g@10.10.13.241:554/Streaming/channels/101",
	// 	// "c242": "rtsp://admin:tg12346789g@10.10.13.242:554/Streaming/channels/101",
	// 	"c238": "rtsp://10.10.14.60:8554/c238",
	// 	"c239": "rtsp://10.10.14.60:8554/c239",
	// 	"c240": "rtsp://10.10.14.60:8554/c240",
	// 	"c241": "rtsp://10.10.14.60:8554/c241",
	// 	"c242": "rtsp://10.10.14.60:8554/c242",
	// }

	rtspConfig, err := ioutil.ReadFile("rtsp_config.json")
	if err != nil {
		log.Fatalln(err)
	}

	rtspURLs := map[string]string{}
	err = json.Unmarshal(rtspConfig, &rtspURLs)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	var config Config
	l := envconfig.PrefixLookuper("STREAM_WEBRTC_", envconfig.OsLookuper())
	if err := envconfig.ProcessWith(ctx, &config, l); err != nil {
		log.Fatalln(err)
	}
	cf, _ := json.Marshal(config)
	log.Println("Loaded config:", string(cf))

	for cam := range rtspURLs {
		outboundVideoTracks[cam], err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
			MimeType: "video/h264",
		}, "pion-rtsp", "pion-rtsp")
		if err != nil {
			log.Fatalln(err)
		}
	}

	go rtspConsumer(rtspURLs)

	sv := echo.New()
	sv.Use(middleware.Logger())
	sv.Use(middleware.Recover())
	sv.Use(middleware.CORS())

	sv.POST("/signaling", func(c echo.Context) error {
		return doSignaling(c, config)
	})
	sv.GET("*", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	if config.TLS {
		err = sv.StartTLS(config.HTTP, config.CertPath+"server.crt", config.CertPath+"server.key")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err = sv.Start(config.HTTP)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// Convert H264 to Annex-B, then write to outboundVideoTrack which sends to all PeerConnections
func rtspConsumer(rtspURLs map[string]string) {
	annexbNALUStartCode := func() []byte { return []byte{0x00, 0x00, 0x00, 0x01} }

	for cam, rtspURL := range rtspURLs {
		go func(cam, rtspURL string) {
			for {
				connected := false
				var previousTime time.Duration
				var codecs []av.CodecData

				rtspClient, err := rtspv2.Dial(rtspv2.RTSPClientOptions{
					URL:              rtspURL,
					DialTimeout:      5 * time.Second,
					ReadWriteTimeout: 5 * time.Second,
					DisableAudio:     true,
				})
				if err != nil {
					log.Println(err)
					goto exit
				}

				connected = true

				codecs = rtspClient.CodecData
				for i, t := range codecs {
					log.Println(cam, "- track", i, "-", t.Type().String())
				}
				if codecs[0].Type() != av.H264 {
					log.Println("RTSP feed must begin with a H264 codec")
					goto exit
				}

				if len(codecs) != 1 {
					log.Println("Ignoring all but the first stream.")
				}

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
							log.Println(err)
							goto exit
						}
					case sig := <-rtspClient.Signals:
						if sig == rtspv2.SignalStreamRTPStop {
							goto exit
						}
					}
				}

			exit:
				if connected {
					rtspClient.Close()
				}
				time.Sleep(5 * time.Second)
				log.Println(cam, "restarting")
			}
		}(cam, rtspURL)
	}
}
