package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/rtsp"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

var (
	outboundVideoTrack  *webrtc.TrackLocalStaticSample
	peerConnectionCount int64
)

// HTTP Handler that accepts an Offer and returns an Answer
// adds outboundVideoTrack to PeerConnection
func doSignaling(w http.ResponseWriter, r *http.Request) {
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
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

	if _, err = peerConnection.AddTrack(outboundVideoTrack); err != nil {
		panic(err)
	}

	var offer webrtc.SessionDescription
	if err = json.NewDecoder(r.Body).Decode(&offer); err != nil {
		panic(err)
	}

	if err = peerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}

	gatherCompletePromise := webrtc.GatheringCompletePromise(peerConnection)

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	} else if err = peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}

	<-gatherCompletePromise

	response, err := json.Marshal(*peerConnection.LocalDescription())
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	outboundVideoTrack, err = webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
		MimeType: "video/h264",
	}, "pion-rtsp", "pion-rtsp")
	if err != nil {
		panic(err)
	}

	go rtspConsumer()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/doSignaling", doSignaling)

	fmt.Println("Open http://localhost:8080 to access this demo")
	panic(http.ListenAndServe(":8080", nil))
}

const rtspURL = "rtsp://admin:tg12346789g@14.161.28.68:50238/Streaming/channels/101"

// Convert H264 to Annex-B, then write to outboundVideoTrack which sends to all PeerConnections
func rtspConsumer() {
	annexbNALUStartCode := func() []byte { return []byte{0x00, 0x00, 0x00, 0x01} }

	for {
		session, err := rtsp.Dial(rtspURL)
		if err != nil {
			panic(err)
		}
		session.RtpKeepAliveTimeout = 10 * time.Second

		codecs, err := session.Streams()
		if err != nil {
			panic(err)
		}
		for i, t := range codecs {
			log.Println("Stream", i, "is of type", t.Type().String())
		}
		if codecs[0].Type() != av.H264 {
			panic("RTSP feed must begin with a H264 codec")
		}
		if len(codecs) != 1 {
			log.Println("Ignoring all but the first stream.")
		}

		var previousTime time.Duration
		for {
			pkt, err := session.ReadPacket()
			if err != nil {
				break
			}

			if pkt.Idx != 0 {
				//audio or other stream, skip it
				continue
			}

			pkt.Data = pkt.Data[4:]

			// For every key-frame pre-pend the SPS and PPS
			if pkt.IsKeyFrame {
				pkt.Data = append(annexbNALUStartCode(), pkt.Data...)
				pkt.Data = append(codecs[0].(h264parser.CodecData).PPS(), pkt.Data...)
				pkt.Data = append(annexbNALUStartCode(), pkt.Data...)
				pkt.Data = append(codecs[0].(h264parser.CodecData).SPS(), pkt.Data...)
				pkt.Data = append(annexbNALUStartCode(), pkt.Data...)
			}

			bufferDuration := pkt.Time - previousTime
			previousTime = pkt.Time
			if err = outboundVideoTrack.WriteSample(media.Sample{Data: pkt.Data, Duration: bufferDuration}); err != nil && err != io.ErrClosedPipe {
				panic(err)
			}
		}

		if err = session.Close(); err != nil {
			log.Println("session Close error", err)
		}

		time.Sleep(5 * time.Second)
	}
}
