package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/base"
	"github.com/aler9/gortsplib/pkg/rtph264"
	"github.com/edgeware/mp4ff/avc"
)

func main() {

	// Main
	code, err := Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(code)
}

var ErrNoH264Track = errors.New("h264 track not found")

func Run() (int, error) {
	tcpProtocal := base.StreamProtocolTCP
	c := &gortsplib.Client{
		StreamProtocol: &tcpProtocal,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	conn, err := c.DialRead("rtsp://admin:tg12346789g@14.161.28.68:50240/Streaming/channels/101")
	if err != nil {
		return 1, err
	}
	defer conn.Close()

	h264TrackID := -1

	for _, track := range conn.Tracks() {
		if track.IsH264() {
			h264TrackID = track.ID
		}
	}

	if h264TrackID < 0 {
		return 1, ErrNoH264Track
	}

	dec := rtph264.NewDecoder()
	err = conn.ReadFrames(func(trackID int, typ gortsplib.StreamType, buf []byte) {
		if h264TrackID == trackID {
			// convert RTP frames into H264 NALUs

			nalus, dur, err := dec.Decode(buf)
			if err != nil {
				return
			}

			if len(nalus) > 0 {
				nalu := nalus[0]

				// sliceType, err := avc.GetSliceTypeFromNALU(nalu)
				// if err != nil {
				// 	log.Println(err)
				// 	return
				// }

				nalType := avc.GetNaluType(nalu[0])

				// log.Println(len(nalus), dur.String(), sliceType.String(), nalType.String())
			}
		}
	})
	if err != nil {
		return 1, err
	}

	return 0, nil
}
