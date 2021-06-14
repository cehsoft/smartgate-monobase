package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/aler9/gortsplib"
	"github.com/aler9/gortsplib/pkg/base"
	"github.com/edgeware/mp4ff/avc"
	"github.com/gorilla/websocket"
	"github.com/markbates/pkger"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	"github.com/init-tech-solution/spitc-streamer/bmff"
)

// Command line flag parameters
var (
	flagListen string
)

func init() {
	flag.StringVar(
		&flagListen, "l", "localhost:2020", "listen on host:port",
	)
}

const (
	nalTypeNonIDRCodedSlice = 1
	nalTypeIDRCodedSlice    = 5

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

var ErrNoH264Track = errors.New("h264 track not found")

// Websocket parameters
var upgrader = websocket.Upgrader{
	// Tune read buffers for short acknowledgement messages
	ReadBufferSize: 256,

	// Tune write buffers to comfortably fit most all B and P frames.
	WriteBufferSize: 8192,

	// Allow any origin for demo purposes
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// client structure
type client struct {
	hub     *hub
	conn    *websocket.Conn // Websocket connection
	frags   chan []byte     // Buffered channel of outbound MP4 fragments
	n       int             // Frame number
	haveIDR bool            // Received i-frame?
}

// hub maintains a set of active clients and broadcasts video to clients
type hub struct {
	clients    map[*client]bool // registered clients
	nals       chan []byte      // NAL units from camera source
	register   chan *client     // register requests from clients
	unregister chan *client     // unregister requests from clients
}

// newHub instantiates a new hub
func newHub() *hub {
	return &hub{
		clients:    make(map[*client]bool),
		nals:       make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

// run processes register and unregister requests, and nal units
func (h *hub) run() {
	for {
		select {
		// Register request
		case c := <-h.register:
			h.clients[c] = true

			var frag bytes.Buffer
			bmff.WriteFTYP(&frag)
			bmff.WriteMOOV(&frag, 1920, 1080)
			c.frags <- frag.Bytes()

		// Unregister request
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.frags)
			}

		// New NAL from source
		case nal := <-h.nals:
			nal = bytes.TrimPrefix(nal, []byte{0, 0, 0, 1})
			if len(nal) == 0 {
				break
			}
			nalType := (nal[0] & 0x1F)
			log.Println(nalType)

			for c := range h.clients {
				var frag bytes.Buffer

				switch nalType {
				case nalTypeIDRCodedSlice:
					c.haveIDR = true
					fallthrough
				case nalTypeNonIDRCodedSlice:
					if c.haveIDR {
						bmff.WriteMOOF(&frag, c.n, nal)
						bmff.WriteMDAT(&frag, nal)
						c.n++

						select {
						// Write MP4 fragment
						case c.frags <- frag.Bytes():

						// Buffered channel full. Drop client.
						default:
							close(c.frags)
							delete(h.clients, c)
						}
					}
				default:
					// noop
				}
			}
		}
	}
}

type source struct {
	con *gortsplib.ClientConn
	hub *hub
}

func newSource(h *hub) (*source, error) {
	// Open device
	// dev, err := v4l2.Open("/dev/video0")
	// if nil != err {
	// 	log.Fatal(err)
	// }

	// // Set pixel format
	// if err := dev.SetPixelFormat(
	// 	1280,
	// 	720,
	// 	v4l2.V4L2_PIX_FMT_H264,
	// ); nil != err {
	// 	log.Fatal(err)
	// }

	// // Set bitrate
	// if err := dev.SetBitrate(1500000); nil != err {
	// 	log.Fatal(err)
	// }

	tcpProtocal := base.StreamProtocolTCP
	c := &gortsplib.Client{
		StreamProtocol: &tcpProtocal,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	conn, err := c.DialRead("rtsp://admin:tg12346789g@14.161.28.68:50240/Streaming/channels/101")
	if err != nil {
		return nil, err
	}

	return &source{
		// device: dev,
		con: conn,
		hub: h,
	}, nil
}

func (s *source) run() error {
	// Start stream
	// if err := s.device.Start(); nil != err {
	// 	log.Fatal(err)
	// }
	// defer s.device.Stop()

	// for {
	// 	select {
	// 	case b := <-s.device.C:
	// 		s.hub.nals <- b.Data
	// 		b.Release()
	// 	}
	// }
	defer s.con.Close()

	h264TrackID := -1

	for _, track := range s.con.Tracks() {
		if track.IsH264() {
			h264TrackID = track.ID
		}

	}

	if h264TrackID < 0 {
		return ErrNoH264Track
	}

	builder := samplebuilder.New(10, &codecs.H264Packet{}, 90000)
	go func() {
		time.Sleep(1 * time.Second)
		for {
			s := builder.Pop()
			if s != nil {
				log.Print(s.Timestamp)
			}
		}
	}()
	var sps *avc.SPS
	err := s.con.ReadFrames(func(trackID int, typ gortsplib.StreamType, buf []byte) {
		if h264TrackID == trackID && typ == gortsplib.StreamTypeRTP {
			pkg := &rtp.Packet{}
			err := pkg.Unmarshal(buf)
			if err != nil {
				log.Println(err.Error())
				return
			}

			nalu := pkg.Payload
			naluType := avc.GetNaluType(nalu[0])
			if naluType != 28 {
				log.Println("nalType", naluType)
			}

			if naluType == avc.NALU_SPS {
				if sps == nil {
					sps, err = avc.ParseSPSNALUnit(nalu, false)
					if err != nil {
						log.Println(err.Error())
						return
					}
					log.Println(sps.Profile)
				}
			}

			builder.Push(pkg)
			// sample := builder.Pop()
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
		return err
	}

	return nil
}

// Handle websocket client connections
func serveWs(h *hub, w http.ResponseWriter, r *http.Request) {
	// Upgrade websocket connection from HTTP to TCP
	conn, err := upgrader.Upgrade(w, r, nil)
	if nil != err {
		log.Println(err)
		return
	}

	// Instantiate client
	c := &client{hub: h, conn: conn, frags: make(chan []byte, 30), n: 1}
	c.hub.register <- c

	// Go routine writes new MP4 fragment to client websocket
	go func(c *client) {
		defer func() {
			c.conn.Close()
		}()

		for {
			select {
			case frag, ok := <-c.frags:
				// log.Println(frag)
				c.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if !ok {
					// Hub closed the channel
					c.conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				// Write next segment
				nw, err := c.conn.NextWriter(websocket.BinaryMessage)
				if nil != err {
					return
				}
				nw.Write(frag)

				// Close writer
				if err := nw.Close(); nil != err {
					return
				}
			}
		}
	}(c)
}

func main() {
	flag.Parse()

	// Parse host:port into host and port
	host, port, err := net.SplitHostPort(flagListen)
	if nil != err {
		log.Fatal(err)
	}

	// One-to-many hub broadcasts NAL units as MP4 fragments to clients
	hub := newHub()
	go hub.run()

	// Open source
	src, err := newSource(hub)
	if err != nil {
		log.Println(err)
	}
	go src.run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.Handle("/", http.FileServer(pkger.Dir("/web/static")))

	// Start server
	fmt.Printf("Listening on http://%v:%v\n", host, port)
	log.Fatal(http.ListenAndServe(flagListen, nil))
}
