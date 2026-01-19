package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kazzmir/opus-go/ogg"
)

// Writes length-prefixed Opus packets (u32le length + bytes) for audio packets only.
func main() {
	var (
		noCRC = flag.Bool("no-crc", false, "skip Ogg CRC verification")
		out   = flag.String("out", "", "output file (default stdout)")
	)
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: oggopusextract [--no-crc] [--out packets.bin] file.ogg")
		os.Exit(2)
	}

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		fatal(err)
	}
	defer in.Close()

	r, err := ogg.NewOpusReader(in)
	if err != nil {
		fatal(err)
	}
	r.SetVerifyCRC(!*noCRC)

	var w io.Writer = os.Stdout
	var f *os.File
	if *out != "" {
		f, err = os.Create(*out)
		if err != nil {
			fatal(err)
		}
		defer f.Close()
		w = f
	}
	bw := bufio.NewWriterSize(w, 128*1024)
	defer bw.Flush()

	for {
		pkt, err := r.ReadAudioPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal(err)
		}
		var lenBuf [4]byte
		binary.LittleEndian.PutUint32(lenBuf[:], uint32(len(pkt.Data)))
		if _, err := bw.Write(lenBuf[:]); err != nil {
			fatal(err)
		}
		if _, err := bw.Write(pkt.Data); err != nil {
			fatal(err)
		}
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
