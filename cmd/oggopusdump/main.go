package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kazzmir/opus-go/ogg"
)

func main() {
	var (
		noCRC = flag.Bool("no-crc", false, "skip Ogg CRC verification")
		max   = flag.Int("max", 0, "stop after N audio packets (0 = all)")
	)
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: oggopusdump [--no-crc] [--max N] file.ogg")
		os.Exit(2)
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fatal(err)
	}
	defer f.Close()

	r, err := ogg.NewOpusReader(f)
	if err != nil {
		fatal(err)
	}
	r.SetVerifyCRC(!*noCRC)

	fmt.Printf("OpusHead: version=%d channels=%d preSkip=%d inputRate=%d mapping=%d\n",
		r.Head.Version, r.Head.Channels, r.Head.PreSkip, r.Head.InputSampleRate, r.Head.ChannelMappingFamily)
	fmt.Printf("OpusTags: vendor=%q comments=%d\n", r.Tags.Vendor, len(r.Tags.Comments))

	count := 0
	for {
		pkt, err := r.ReadAudioPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal(err)
		}
		count++
		if pkt.GranuleValid {
			fmt.Printf("pkt=%d bytes=%d granule=%d page=%d eos=%v\n", count, len(pkt.Data), pkt.GranulePos, pkt.PageSequence, pkt.EOS)
		} else {
			fmt.Printf("pkt=%d bytes=%d granule=? (%d) page=%d eos=%v\n", count, len(pkt.Data), pkt.GranulePos, pkt.PageSequence, pkt.EOS)
		}
		if *max > 0 && count >= *max {
			break
		}
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
