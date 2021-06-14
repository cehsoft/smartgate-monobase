package bmff

import (
	"bytes"
	"io"
)

// References:
// ISO/IEC 14496 Part 12
// ISO/IEC 14496 Part 15

// ISO/IEC 14496 Part 14 is not used.

func WriteInt(w io.Writer, v int, n int) {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[n-i-1] = byte(v & 0xff)
		v >>= 8
	}
	w.Write(b)
}

func WriteString(w io.Writer, s string) {
	w.Write([]byte(s))
}

func WriteTag(w io.Writer, tag string, cb func(w io.Writer)) {
	var b bytes.Buffer
	cb(&b)                    // callback
	WriteInt(w, b.Len()+8, 4) // box size
	WriteString(w, tag)       // box type
	w.Write(b.Bytes())        // box content
}

func WriteFTYP(w io.Writer) {
	WriteTag(w, "ftyp", func(w io.Writer) {
		WriteString(w, "isom")                 // major brand
		WriteInt(w, 0x200, 4)                  // minor version
		WriteString(w, "isomiso2iso5avc1mp41") // compatible brands
	})
}

func WriteMOOV(w io.Writer, width, height uint16) {
	WriteTag(w, "moov", func(w io.Writer) {
		WriteMVHD(w)
		WriteTRAK(w, width, height)
		WriteMVEX(w)
	})
}

func WriteMVHD(w io.Writer) {
	WriteTag(w, "mvhd", func(w io.Writer) {
		WriteInt(w, 0, 4)          // version and flags
		WriteInt(w, 0, 4)          // creation time
		WriteInt(w, 0, 4)          // modification time
		WriteInt(w, 1000, 4)       // timescale
		WriteInt(w, 0, 4)          // duration (all 1s == unknown)
		WriteInt(w, 0x00010000, 4) // rate (1.0 == normal)
		WriteInt(w, 0x0100, 2)     // volume (1.0 == normal)
		WriteInt(w, 0, 2)          // reserved
		WriteInt(w, 0, 4)          // reserved
		WriteInt(w, 0, 4)          // reserved
		WriteInt(w, 0x00010000, 4) // matrix
		WriteInt(w, 0x0, 4)        // matrix
		WriteInt(w, 0x0, 4)        // matrix
		WriteInt(w, 0x0, 4)        // matrix
		WriteInt(w, 0x00010000, 4) // matrix
		WriteInt(w, 0x0, 4)        // matrix
		WriteInt(w, 0x0, 4)        // matrix
		WriteInt(w, 0x0, 4)        // matrix
		WriteInt(w, 0x40000000, 4) // matrix
		WriteInt(w, 0, 4)          // pre-defined
		WriteInt(w, 0, 4)          // pre-defined
		WriteInt(w, 0, 4)          // pre-defined
		WriteInt(w, 0, 4)          // pre-defined
		WriteInt(w, 0, 4)          // pre-defined
		WriteInt(w, 0, 4)          // pre-defined
		WriteInt(w, -1, 4)         // next track id
	})
}

func WriteTRAK(w io.Writer, width, height uint16) {
	WriteTag(w, "trak", func(w io.Writer) {
		WriteTKHD(w, width, height)
		WriteMDIA(w, width, height)
	})
}

func WriteTKHD(w io.Writer, width, height uint16) {
	WriteTag(w, "tkhd", func(w io.Writer) {
		WriteInt(w, 7, 4)               // version and flags (track enabled)
		WriteInt(w, 0, 4)               // creation time
		WriteInt(w, 0, 4)               // modification time
		WriteInt(w, 1, 4)               // track id
		WriteInt(w, 0, 4)               // reserved
		WriteInt(w, 0, 4)               // duration
		WriteInt(w, 0, 4)               // reserved
		WriteInt(w, 0, 4)               // reserved
		WriteInt(w, 0, 2)               // layer
		WriteInt(w, 0, 2)               // alternate group
		WriteInt(w, 0, 2)               // volume (ignored for video tracks)
		WriteInt(w, 0, 2)               // reserved
		WriteInt(w, 0x00010000, 4)      // matrix
		WriteInt(w, 0x0, 4)             // matrix
		WriteInt(w, 0x0, 4)             // matrix
		WriteInt(w, 0x0, 4)             // matrix
		WriteInt(w, 0x00010000, 4)      // matrix
		WriteInt(w, 0x0, 4)             // matrix
		WriteInt(w, 0x0, 4)             // matrix
		WriteInt(w, 0x0, 4)             // matrix
		WriteInt(w, 0x40000000, 4)      // matrix
		WriteInt(w, int(width)<<16, 4)  // width (fixed-point 16.16 format)
		WriteInt(w, int(height)<<16, 4) // height (fixed-point 16.16 format)
	})
}

func WriteMDIA(w io.Writer, width, height uint16) {
	WriteTag(w, "mdia", func(w io.Writer) {
		WriteMDHD(w)
		WriteHDLR(w)
		WriteMINF(w, width, height)
	})
}

func WriteMDHD(w io.Writer) {
	WriteTag(w, "mdhd", func(w io.Writer) {
		WriteInt(w, 0, 4)      // version and flags
		WriteInt(w, 0, 4)      // creation time
		WriteInt(w, 0, 4)      // modification time
		WriteInt(w, 10000, 4)  // timescale
		WriteInt(w, 0, 4)      // duration
		WriteInt(w, 0x55c4, 2) // language ('und' == undefined)
		WriteInt(w, 0, 2)      // pre-defined
	})
}

func WriteHDLR(w io.Writer) {
	WriteTag(w, "hdlr", func(w io.Writer) {
		WriteInt(w, 0, 4)                        // version and flags
		WriteInt(w, 0, 4)                        // pre-defined
		WriteString(w, "vide")                   // handler type
		WriteInt(w, 0, 4)                        // reserved
		WriteInt(w, 0, 4)                        // reserved
		WriteInt(w, 0, 4)                        // reserved
		WriteString(w, "MicroMSE Video Handler") // name
		WriteInt(w, 0, 1)                        // null-terminator
	})
}

func WriteMINF(w io.Writer, width, height uint16) {
	WriteTag(w, "minf", func(w io.Writer) {
		WriteVMHD(w)
		WriteDINF(w)
		WriteSTBL(w, width, height)
	})
}

func WriteVMHD(w io.Writer) {
	WriteTag(w, "vmhd", func(w io.Writer) {
		WriteInt(w, 1, 4) // version and flags
		WriteInt(w, 0, 2) // graphics mode
		WriteInt(w, 0, 2) // opcolor
		WriteInt(w, 0, 2) // opcolor
		WriteInt(w, 0, 2) // opcolor
	})
}

func WriteDINF(w io.Writer) {
	WriteTag(w, "dinf", func(w io.Writer) {
		WriteDREF(w)
	})
}

func WriteDREF(w io.Writer) {
	WriteTag(w, "dref", func(w io.Writer) {
		WriteInt(w, 0, 4) // version and flags
		WriteInt(w, 1, 4) // entry count
		WriteTag(w, "url ", func(w io.Writer) {
			WriteInt(w, 1, 4) // version and flags
		})
	})
}

func WriteSTBL(w io.Writer, width, height uint16) {
	WriteTag(w, "stbl", func(w io.Writer) {
		WriteSTSD(w, width, height)
		WriteSTSZ(w)
		WriteSTSC(w)
		WriteSTTS(w)
		WriteSTCO(w)
	})
}

// Sample Table Box
func WriteSTSD(w io.Writer, width, height uint16) {
	WriteTag(w, "stsd", func(w io.Writer) {
		WriteInt(w, 0, 6) // reserved
		WriteInt(w, 1, 2) // deta reference index
		WriteTag(w, "avc1", func(w io.Writer) {
			WriteInt(w, 0, 6)           // reserved
			WriteInt(w, 1, 2)           // data reference index
			WriteInt(w, 0, 2)           // pre-defined
			WriteInt(w, 0, 2)           // reserved
			WriteInt(w, 0, 4)           // pre-defined
			WriteInt(w, 0, 4)           // pre-defined
			WriteInt(w, 0, 4)           // pre-defined
			WriteInt(w, int(width), 2)  // width
			WriteInt(w, int(height), 2) // height
			WriteInt(w, 0x00480000, 4)  // horizontal resolution: 72 dpi
			WriteInt(w, 0x00480000, 4)  // vertical resolution: 72 dpi
			WriteInt(w, 0, 4)           // data size: 0
			WriteInt(w, 1, 2)           // frame count: 1
			w.Write(make([]byte, 32))   // compressor name
			WriteInt(w, 0x18, 2)        // depth
			WriteInt(w, 0xffff, 2)      // pre-defined

			// Raspberry Pi 3B+ SPS/PPS for H.264 Main 4.0
			sps := []byte{
				0x27, 0x64, 0x00, 0x28, 0xac, 0x2b, 0x40, 0x28,
				0x02, 0xdd, 0x00, 0xf1, 0x22, 0x6a,
			}
			pps := []byte{
				0x28, 0xee, 0x02, 0x5c, 0xb0, 0x00,
			}

			// MPEG-4 Part 15 extension
			// See ISO/IEC 14496-15:2004 5.3.4.1.2
			WriteTag(w, "avcC", func(w io.Writer) {
				WriteInt(w, 1, 1)    // configuration version
				WriteInt(w, 0x64, 1) // H.264 profile (0x64 == high)
				WriteInt(w, 0x00, 1) // H.264 profile compatibility
				WriteInt(w, 0x28, 1) // H.264 level (0x28 == 4.0)
				WriteInt(w, 0xff, 1) // nal unit length - 1 (upper 6 bits == 1)
				WriteInt(w, 0xe1, 1) // number of sps (upper 3 bits == 1)
				WriteInt(w, len(sps), 2)
				w.Write(sps)
				WriteInt(w, 1, 1) // number of pps
				WriteInt(w, len(pps), 2)
				w.Write(pps)
			})
		})
	})
}

func WriteSTTS(w io.Writer) {
	WriteTag(w, "stts", func(w io.Writer) {
		WriteInt(w, 0, 4) // version and flags
		WriteInt(w, 0, 4) // entry count
	})
}

func WriteSTSC(w io.Writer) {
	WriteTag(w, "stsc", func(w io.Writer) {
		WriteInt(w, 0, 4) // version and flags
		WriteInt(w, 0, 4) // entry count
	})
}

func WriteSTSZ(w io.Writer) {
	WriteTag(w, "stsz", func(w io.Writer) {
		WriteInt(w, 0, 4) // version and flags
		WriteInt(w, 0, 4) // sample size
		WriteInt(w, 0, 4) // sample count
	})
}

func WriteSTCO(w io.Writer) {
	WriteTag(w, "stco", func(w io.Writer) {
		WriteInt(w, 0, 4) // version and flags
		WriteInt(w, 0, 4) // entry count
	})
}

// Movie Extends Box
func WriteMVEX(w io.Writer) {
	WriteTag(w, "mvex", func(w io.Writer) {
		WriteMEHD(w)
		WriteTREX(w)
	})
}

// Movie Extends Header Box
func WriteMEHD(w io.Writer) {
	WriteTag(w, "mehd", func(w io.Writer) {
		WriteInt(w, 0, 4) // version and flags
		WriteInt(w, 0, 4) // fragment duration
	})
}

// Track Extends Box
func WriteTREX(w io.Writer) {
	WriteTag(w, "trex", func(w io.Writer) {
		WriteInt(w, 0, 4)          // version and flags
		WriteInt(w, 1, 4)          // track id
		WriteInt(w, 1, 4)          // default sample description index
		WriteInt(w, 0, 4)          // default sample duration
		WriteInt(w, 0, 4)          // default sample size
		WriteInt(w, 0x00010000, 4) // default sample flags
	})
}

// Movie Fragment Box
func WriteMOOF(w io.Writer, seq int, data []byte) {
	WriteTag(w, "moof", func(w io.Writer) {
		WriteMFHD(w, seq)
		WriteTRAF(w, seq, data)
	})
}

// Movie Fragment Header Box
func WriteMFHD(w io.Writer, seq int) {
	WriteTag(w, "mfhd", func(w io.Writer) {
		WriteInt(w, 0, 4)   // version and flags
		WriteInt(w, seq, 4) // sequence number
	})
}

// Track Fragment Box
func WriteTRAF(w io.Writer, seq int, data []byte) {
	WriteTag(w, "traf", func(w io.Writer) {
		WriteTFHD(w)
		WriteTFDT(w, seq)
		WriteTRUN(w, data)
	})
}

// Track Fragment Header Box
func WriteTFHD(w io.Writer) {
	WriteTag(w, "tfhd", func(w io.Writer) {
		WriteInt(w, 0x020020, 4)   // version and flags
		WriteInt(w, 1, 4)          // track ID
		WriteInt(w, 0x01010000, 4) // default sample flags
	})
}

// Track Fragment Base Media Decode Time Box
func WriteTFDT(w io.Writer, seq int) {
	WriteTag(w, "tfdt", func(w io.Writer) {
		WriteInt(w, 0x01000000, 4) // version and flags
		WriteInt(w, 330*seq, 8)    // base media decode time
	})
}

// Track Run Box
func WriteTRUN(w io.Writer, data []byte) {
	WriteTag(w, "trun", func(w io.Writer) {
		WriteInt(w, 0x00000305, 4) // version and flags
		WriteInt(w, 1, 4)          // sample count
		WriteInt(w, 0x70, 4)       // data offset
		if (len(data) > 0) && (data[0]&0x1f == 0x5) {
			WriteInt(w, 0x02000000, 4) // first sample flags (i-frame)
		} else {
			WriteInt(w, 0x01010000, 4) // first sample flags (not i-frame)
		}
		WriteInt(w, 330, 4)         // sample duration
		WriteInt(w, 4+len(data), 4) // sample size
	})
}

// Media Data Box
func WriteMDAT(w io.Writer, data []byte) {
	WriteTag(w, "mdat", func(w io.Writer) {
		WriteInt(w, len(data), 4)
		w.Write(data)
	})
}
