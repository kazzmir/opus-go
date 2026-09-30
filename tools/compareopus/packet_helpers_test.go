//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"
	"unsafe"

	"github.com/kazzmir/opus-go/opuscc"
)

func goPacketParse(data *byte, length, self, mask int32, public ...bool) packetParseResult {
	r := packetParseResult{Toc: 0xa5, Payload: -77, Packet: -78, PaddingLen: -79}
	for i := range r.Sizes {
		r.Sizes[i] = 1234
	}
	f := [48]*byte{}
	for i := range f {
		f[i] = data
	}
	pad := data
	var toc *byte
	var frames *[48]*byte
	var size *[48]int16
	var payload, offset, padlen *int32
	var padding **byte
	if mask&1 != 0 {
		toc = &r.Toc
	}
	if mask&2 != 0 {
		frames = &f
	}
	if mask&4 != 0 {
		size = (*[48]int16)(unsafe.Pointer(&r.Sizes[1]))
	}
	if mask&8 != 0 {
		payload = &r.Payload
	}
	if mask&16 != 0 {
		offset = &r.Packet
	}
	if mask&32 != 0 {
		padding = &pad
		padlen = &r.PaddingLen
	}
	if len(public) != 0 && public[0] {
		r.Count = opuscc.Opus_opus_packet_parse(nil, data, length, toc, frames, size, payload)
	} else {
		r.Count = opuscc.Opus_opus_packet_parse_impl(nil, data, length, self, toc, frames, size, payload, offset, padding, padlen)
	}
	for i, p := range f {
		r.Frames[i] = -1
		if p != nil {
			r.Frames[i] = int32(uintptr(unsafe.Pointer(p)) - uintptr(unsafe.Pointer(data)))
		}
	}
	r.Padding = -1
	if pad != nil {
		r.Padding = int32(uintptr(unsafe.Pointer(pad)) - uintptr(unsafe.Pointer(data)))
	}
	return r
}

func packetParserFixtures() [][]byte {
	fixtures := [][]byte{nil, {0}, {1}, {2}, {3}, {0x83, 0xc2, 2, 1, 0xaa, 0xbb, 0xcc, 0xdd, 0xee}, {0x81, 3, 11, 12, 13, 14, 15, 16, 99, 99}, {0x83, 48}, {0x83, 49}, {0x83, 0x41, 255, 0}, {0x83, 0x81, 252}, {0x82, 255, 255}}
	rng := rand.New(rand.NewSource(173))
	for i := 0; i < 4000; i++ {
		data := make([]byte, rng.Intn(300))
		rng.Read(data)
		fixtures = append(fixtures, data)
	}
	for _, n := range []int{251, 252, 1275, 1276, 2550, 2551, 65536} {
		for _, toc := range []byte{0, 1, 2, 0x83} {
			data := make([]byte, n+1)
			data[0] = toc
			if toc == 0x83 {
				data[1] = 2
			}
			fixtures = append(fixtures, data)
		}
	}
	return fixtures
}

func TestPacketParseAgainstC(t *testing.T) {
	for _, packet := range packetParserFixtures() {
		for _, mask := range []int32{63, 62, 61, 55, 4, 0} {
			data := append(slices.Clone(packet), 0, 0)
			before := slices.Clone(data)
			g := goPacketParse(&data[0], int32(len(packet)), 0, mask, true)
			c := nativePacketParse(&data[0], int32(len(packet)), 0, mask, 1)
			if g != c || !slices.Equal(data, before) {
				t.Fatalf("packet=%x mask=%d Go=%+v C=%+v", packet, mask, g, c)
			}
		}
	}
}

func TestPacketParseImplAgainstC(t *testing.T) {
	for _, packet := range packetParserFixtures() {
		for _, self := range []int32{0, 1, -1} {
			for _, mask := range []int32{63, 62, 61, 55, 47, 31, 4, 0} {
				// Extra backing bytes permit comparisons of valid one-past-end cursors.
				data := append(slices.Clone(packet), 0, 0)
				before := slices.Clone(data)
				g := goPacketParse(&data[0], int32(len(packet)), self, mask)
				c := nativePacketParse(&data[0], int32(len(packet)), self, mask, 0)
				if g != c || !slices.Equal(data, before) {
					t.Fatalf("packet=%x self=%d mask=%d Go=%+v C=%+v", packet, self, mask, g, c)
				}
			}
		}
	}
	for _, length := range []int32{-1, 0} {
		g := goPacketParse(nil, length, 0, 63)
		c := nativePacketParse(nil, length, 0, 63, 0)
		if g != c {
			t.Fatal(length, g, c)
		}
	}
}

func TestPacketHelpersAgainstC(t *testing.T) {
	for toc := 0; toc < 256; toc++ {
		b := byte(toc)
		if got, want := opuscc.Opus_opus_packet_get_bandwidth(nil, &b), nativePacketBandwidth(&b); got != want {
			t.Fatalf("TOC %#x bandwidth: Go=%d C=%d", toc, got, want)
		}
		if got, want := opuscc.Opus_opus_packet_get_nb_channels(nil, &b), nativePacketChannels(&b); got != want {
			t.Fatalf("TOC %#x channels: Go=%d C=%d", toc, got, want)
		}
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
			if got, want := opuscc.Opus_opus_packet_get_samples_per_frame(nil, &b, rate), nativePacketSamplesPerFrame(&b, rate); got != want {
				t.Fatalf("TOC %#x rate %d samples/frame: Go=%d C=%d", toc, rate, got, want)
			}
		}
		if got, want := opuscc.Opus_opus_packet_get_nb_frames(nil, &b, 1), nativePacketFrames(&b, 1); got != want {
			t.Fatalf("TOC %#x length 1 frames: Go=%d C=%d", toc, got, want)
		}
		for second := 0; second < 256; second++ {
			packet := [2]byte{b, byte(second)}
			if got, want := opuscc.Opus_opus_packet_get_nb_frames(nil, &packet[0], 2), nativePacketFrames(&packet[0], 2); got != want {
				t.Fatalf("packet %x frames: Go=%d C=%d", packet, got, want)
			}
		}
	}
	for _, length := range []int32{-1, 0} {
		if got, want := opuscc.Opus_opus_packet_get_nb_frames(nil, nil, length), nativePacketFrames(nil, length); got != want {
			t.Fatalf("length %d frames: Go=%d C=%d", length, got, want)
		}
	}
}
