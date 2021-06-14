package streamer

type Streamer struct {
	chanMediaFragments chan []byte
}

func NewStreamer() *Streamer {
	return &Streamer{
		chanMediaFragments: make(chan []byte),
	}
}
