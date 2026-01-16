package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"opusgo/oggopus"
	"opusgo/opus"
	"opusgo/wav"
)

func main() {
	var out = flag.String("out", "out.wav", "output wav file")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: oggopus2wav --out out.wav input.ogg")
		os.Exit(2)
	}

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		fatal(err)
	}
	defer in.Close()

	r, err := oggopus.NewReader(in)
	if err != nil {
		fatal(err)
	}

	dec, err := opus.NewDecoderFromHead(r.Head)
	if err != nil {
		fatal(err)
	}
	defer dec.Close()

	outf, err := os.Create(*out)
	if err != nil {
		fatal(err)
	}
	defer outf.Close()

	ww, err := wav.NewWriter(outf, 48000, int(r.Head.Channels))
	if err != nil {
		fatal(err)
	}
	defer ww.Close()

	// Maximum Opus frame size at 48kHz is 120ms = 5760 samples/channel.
	maxFrame := 5760
	pcm := make([]int16, maxFrame*int(r.Head.Channels))

	preSkipRemaining := int(r.Head.PreSkip)

	for {
		pkt, err := r.ReadAudioPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal(err)
		}

		n, err := dec.Decode(pkt.Data, pcm, maxFrame, false)
		if err != nil {
			fatal(err)
		}

		// n is samples per channel.
		frames := pcm[:n*int(r.Head.Channels)]

		// Apply OpusHead pre-skip in samples per channel.
		if preSkipRemaining > 0 {
			skip := preSkipRemaining
			if skip > n {
				skip = n
			}
			preSkipRemaining -= skip

			// Drop 'skip' samples per channel from interleaved PCM.
			drop := skip * int(r.Head.Channels)
			if drop >= len(frames) {
				continue
			}
			frames = frames[drop:]
		}

		if len(frames) > 0 {
			if err := ww.WriteInt16PCM(frames); err != nil {
				fatal(err)
			}
		}
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
