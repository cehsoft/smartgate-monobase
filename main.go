package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/base"
	"github.com/deepch/vdk/av"
	"github.com/deepch/vdk/codec/h264parser"
	"github.com/deepch/vdk/format/mp4f"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pion/rtp"
	"golang.org/x/net/websocket"
)

func main() {

	channelMP4f := make(chan []byte, 1024)

	sv := echo.New()
	sv.Use(middleware.Logger())
	sv.Use(middleware.Recover())
	sv.Static("/", "web/static")
	sv.GET("/ws", func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				select {
				case mp4f := <-channelMP4f:
					err := websocket.Message.Send(ws, mp4f)
					if err != nil {
						c.Logger().Error(err)
					}

					fmt.Println("sent", len(mp4f), "bytes")
				default:
					// fmt.Println("nothing to sent")
				}
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	go func() {
		code, err := stream(channelMP4f)
		if err != nil {
			log.Println(err)
		}
		os.Exit(code)
	}()

	err := sv.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

var ErrNoH264Track = errors.New("h264 track not found")

func stream(channelMP4f chan []byte) (int, error) {
	tcpProtocal := base.StreamProtocolTCP
	c := &gortsplib.Client{
		StreamProtocol: &tcpProtocal,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	conn, err := c.DialRead("rtsp://admin:tg12346789g@14.161.28.68:50239/Streaming/channels/101")
	if err != nil {
		return 1, err
	}
	defer conn.Close()

	h264TrackID := -1

	for _, track := range conn.Tracks() {
		if track.IsH264() {
			h264TrackID = track.ID
			sps, pps, err := track.ExtractDataH264()
			if err != nil {
				return 1, err
			}
			codec, err := h264parser.NewCodecDataFromSPSAndPPS(sps, pps)
			if err != nil {
				return 1, err
			}
			log.Println(codec.Width(), codec.Height())
		}
	}

	if h264TrackID < 0 {
		return 1, ErrNoH264Track
	}

	startedFUA := false
	bufferFUA := bytes.Buffer{}
	var previousPacketTS uint32 = 0
	// previousPacketSeq := 0
	var sps, pps []byte
	var mp4muxer *mp4f.Muxer

	channelAV := make(chan *av.Packet, 100)
	go func() {
		isStarted := false

		for {
			pkgAV := <-channelAV

			if mp4muxer != nil {
				if pkgAV.IsKeyFrame {
					isStarted = true
				}

				if !isStarted {
					continue
				}

				av := *pkgAV
				av.Time = av.Time + av.Duration

				ready, buf, err := mp4muxer.WritePacket(av, false)
				if err != nil {
					log.Println("&&& mp4muxer.WritePacket err", err)
					continue
				}

				if ready {
					channelMP4f <- buf
				}
			} else if sps != nil && pps != nil {
				codec, err := h264parser.NewCodecDataFromSPSAndPPS(sps, pps)
				if err != nil {
					log.Println("&&& NewCodecDataFromSPSAndPPS err", err)
					continue
				}

				fmt.Println(codec.Height(), codec.Width())

				mp4muxer = &mp4f.Muxer{}
				err = mp4muxer.WriteHeader([]av.CodecData{
					codec,
				})
				if err != nil {
					log.Println("&&& mp4muxer.WriteHeader err", err)
					continue
				}

				meta, init := mp4muxer.GetInit([]av.CodecData{
					codec,
				})

				fmt.Println(">>>>> meta: ", meta)
				channelMP4f <- init
			}
		}
	}()

	err = conn.ReadFrames(func(trackID int, typ gortsplib.StreamType, buf []byte) {
		if h264TrackID == trackID && typ == gortsplib.StreamTypeRTP {
			pkg := &rtp.Packet{}
			err := pkg.Unmarshal(buf)
			if err != nil {
				log.Println(err.Error())
				return
			}

			if previousPacketTS == 0 {
				previousPacketTS = pkg.Timestamp
			}

			nalus, _ := h264parser.SplitNALUs(pkg.Payload)
			avPackets := []*av.Packet{}

			for _, nal := range nalus {
				naluType := nal[0] & 0x1f

				switch {
				case naluType >= 1 && naluType <= 5:
					// fmt.Print(1, 5, " ")
					avPackets = append(avPackets, &av.Packet{
						Data:            append(binSize(len(nal)), nal...),
						CompositionTime: time.Duration(1) * time.Millisecond,
						Idx:             int8(trackID),
						IsKeyFrame:      naluType == 5,
						Duration:        time.Duration(float32(pkg.Timestamp-previousPacketTS)/90) * time.Millisecond,
						Time:            time.Duration(pkg.Timestamp/90) * time.Millisecond,
					})
				case naluType == 7:
					// fmt.Print("SPS", " + ")
					sps = nal
				case naluType == 8:
					// fmt.Print("PPS", " - ")
					pps = nal
				case naluType == 24:
					fmt.Print(24, " ")
				case naluType == 28:
					// fmt.Print(28, " ")

					fuIndicator := pkg.Payload[0]
					fuHeader := pkg.Payload[1]
					isStart := fuHeader&0x80 != 0
					isEnd := fuHeader&0x40 != 0

					if isStart {
						// fmt.Print("isStart [ ")
						startedFUA = true

						bufferFUA.Reset()
						bufferFUA.Write([]byte{fuIndicator&0xe0 | fuHeader&0x1f})
					}

					if startedFUA {
						bufferFUA.Write(pkg.Payload[2:])

						if isEnd {
							// fmt.Print("] isEnd ")
							startedFUA = false

							naluTypeFUA := bufferFUA.Bytes()[0] & 0x1f
							if naluTypeFUA == 7 || naluTypeFUA == 9 {
								// TODO
							}

							avPackets = append(avPackets, &av.Packet{
								Data:            append(binSize(bufferFUA.Len()), bufferFUA.Bytes()...),
								CompositionTime: time.Duration(1) * time.Millisecond,
								Duration:        time.Duration(float32(pkg.Timestamp-previousPacketTS)/90) * time.Millisecond,
								Idx:             int8(trackID),
								IsKeyFrame:      naluTypeFUA == 5,
								Time:            time.Duration(pkg.Timestamp/90) * time.Millisecond,
							})
						}
					}
				}

				if len(avPackets) > 0 {
					previousPacketTS = pkg.Timestamp
				}
			}

			for _, av := range avPackets {
				channelAV <- av
			}
			// naluType := avc.GetNaluType(nalu[0])
			// if naluType != 28 {
			// 	log.Println("nalType", naluType)
			// }

			// if naluType == avc.NALU_SPS {
			// 	if sps == nil {
			// 		sps, err = avc.ParseSPSNALUnit(nalu, false)
			// 		if err != nil {
			// 			log.Println(err.Error())
			// 			return
			// 		}
			// 		codecs := avc.CodecString("avc1", sps)
			// 		log.Println(codecs)
			// 	}
			// }
			// log.Println("pkg.Marker:", pkg.Marker)
			// builder.Push(pkg)

			// log.Println("here", sample)
			// if sample != nil {
			// 	log.Println("here", sample.Data)
			// }

			// h264 := codecs.H264Packet{IsAVC: true}
			// _, err = h264.Unmarshal(pkg.Payload)
			// if err != nil {
			// 	log.Println(err.Error())
			// 	return
			// }
			// log.Println(len(b), h264.IsAVC)
			// log.Println(pkg.SequenceNumber, pkg.PayloadOffset, pkg.SSRC, pkg.PayloadType)
		}
	})

	if err != nil {
		return 1, err
	}

	return 0, nil
}

func binSize(val int) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(val))
	return buf
}
