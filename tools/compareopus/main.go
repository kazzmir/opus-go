//go:build compareopus && cgo

// Command compareopus compares the native C and Go Opus codecs.
package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../../opus/include
#cgo LDFLAGS: ${SRCDIR}/../../../opus/.libs/libopus.a -lm
#include <opus.h>
#include <opus_multistream.h>
*/
import "C"

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
)

const maxFrame = 5760 // 120 ms at 48 kHz, samples per channel

type nativeDecoder struct {
	single *C.OpusDecoder
	multi  *C.OpusMSDecoder
}

func newNative(h ogg.OpusHead) (*nativeDecoder, error) {
	d := new(nativeDecoder)
	var code C.int
	if h.ChannelMappingFamily == 0 {
		d.single = C.opus_decoder_create(48000, C.int(h.Channels), &code)
	} else {
		if len(h.ChannelMapping) != int(h.Channels) || len(h.ChannelMapping) == 0 {
			return nil, fmt.Errorf("invalid channel mapping")
		}
		// libopus copies the mapping; it does not retain this Go pointer.
		d.multi = C.opus_multistream_decoder_create(48000, C.int(h.Channels), C.int(h.StreamCount), C.int(h.CoupledStreamCount), (*C.uchar)(unsafe.Pointer(&h.ChannelMapping[0])), &code)
	}
	if code != C.OPUS_OK || (d.single == nil && d.multi == nil) {
		d.close()
		return nil, fmt.Errorf("C decoder create: %s (%d)", C.GoString(C.opus_strerror(code)), code)
	}
	return d, nil
}

func (d *nativeDecoder) close() {
	if d.single != nil {
		C.opus_decoder_destroy(d.single)
	}
	if d.multi != nil {
		C.opus_multistream_decoder_destroy(d.multi)
	}
}

func (d *nativeDecoder) decode(packet []byte, pcm []int16) (int, error) {
	// Both buffers contain no Go pointers and are borrowed only for this call.
	data := (*C.uchar)(unsafe.Pointer(&packet[0]))
	out := (*C.opus_int16)(unsafe.Pointer(&pcm[0]))
	var n C.int
	if d.single != nil {
		n = C.opus_decode(d.single, data, C.opus_int32(len(packet)), out, maxFrame, 0)
	} else {
		n = C.opus_multistream_decode(d.multi, data, C.opus_int32(len(packet)), out, maxFrame, 0)
	}
	if n < 0 {
		return 0, fmt.Errorf("C decode: %s (%d)", C.GoString(C.opus_strerror(n)), n)
	}
	return int(n), nil
}

func compare(path string, tolerance, limit int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := ogg.NewOpusReader(f)
	if err != nil {
		return err
	}
	// Projection (family 3) needs a different API; do not silently misdecode it.
	if h := r.Head; h.ChannelMappingFamily != 0 && h.ChannelMappingFamily != 1 && h.ChannelMappingFamily != 2 && h.ChannelMappingFamily != 255 {
		return fmt.Errorf("unsupported mapping family %d", h.ChannelMappingFamily)
	}
	gd, err := opus.NewDecoderFromHead(r.Head)
	if err != nil {
		return err
	}
	defer gd.Close()
	cd, err := newNative(r.Head)
	if err != nil {
		return err
	}
	defer cd.close()
	channels := int(r.Head.Channels)
	gp, cp := make([]int16, maxFrame*channels), make([]int16, maxFrame*channels)
	var packets, samples, different, exceeded int64
	var maxError int
	var squaredError float64
	for limit == 0 || packets < int64(limit) {
		pkt, err := r.ReadAudioPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("packet %d: %w", packets+1, err)
		}
		if len(pkt.Data) == 0 {
			return fmt.Errorf("packet %d: empty audio packet (not a PLC test)", packets+1)
		}
		gn, ge := gd.Decode(pkt.Data, gp, maxFrame, false)
		cn, ce := cd.decode(pkt.Data, cp)
		packets++
		if ge != nil || ce != nil {
			return fmt.Errorf("packet %d: Go error=%v; C error=%v", packets, ge, ce)
		}
		if gn != cn {
			return fmt.Errorf("packet %d: frame lengths differ: Go=%d C=%d", packets, gn, cn)
		}
		for i := 0; i < gn*channels; i++ {
			delta := int(gp[i]) - int(cp[i])
			if delta < 0 {
				delta = -delta
			}
			if delta != 0 {
				different++
			}
			if delta > tolerance {
				if exceeded == 0 {
					fmt.Printf("%s: first mismatch: packet=%d frame=%d channel=%d absolute-frame=%d Go=%d C=%d error=%d\n", path, packets, i/channels, i%channels, (samples+int64(i))/int64(channels), gp[i], cp[i], delta)
				}
				exceeded++
			}
			if delta > maxError {
				maxError = delta
			}
			squaredError += float64(delta) * float64(delta)
		}
		samples += int64(gn * channels)
	}
	if samples == 0 {
		return fmt.Errorf("no PCM samples decoded")
	}
	status := "PASS"
	if exceeded != 0 {
		status = "FAIL"
	}
	fmt.Printf("%s %s: channels=%d packets=%d samples=%d different=%d outside-tolerance=%d max-error=%d RMS-error=%.6g\n", status, path, channels, packets, samples, different, exceeded, maxError, math.Sqrt(squaredError/float64(samples)))
	if exceeded != 0 {
		return fmt.Errorf("%d samples exceed tolerance %d", exceeded, tolerance)
	}
	return nil
}

func main() {
	mode := flag.String("mode", "decode", "comparison: decode (Opus inputs) or encode (WAV inputs)")
	bitrate := flag.Int("bitrate", 64000, "encode bitrate in bits/sec")
	complexity := flag.Int("complexity", 10, "encode complexity (0-10)")
	vbr := flag.Bool("vbr", true, "encode with variable bitrate")
	exact := flag.Bool("exact", false, "encode: require identical packet payloads")
	nrmse := flag.Float64("max-nrmse", 0.01, "encode: maximum RMS PCM difference / C-output RMS (0.01 = 1%)")
	sizeDiff := flag.Float64("max-size-diff", 0.05, "encode: maximum relative total payload size difference (0.05 = 5%)")
	tolerance := flag.Int("tolerance", 0, "decode: maximum allowed absolute int16 sample difference (0 = exact)")
	limit := flag.Int("max-packets", 0, "compare at most this many packets per file (0 = all)")
	flag.Parse()
	if (*mode != "decode" && *mode != "encode") || *tolerance < 0 || *tolerance > 65535 || *limit < 0 || *bitrate < 500 || *bitrate > 512000 || *complexity < 0 || *complexity > 10 || !validThreshold(*nrmse) || !validThreshold(*sizeDiff) {
		fmt.Fprintln(os.Stderr, "invalid mode, tolerance, packet limit, encoder settings, or thresholds")
		os.Exit(2)
	}
	paths := flag.Args()
	if len(paths) == 0 {
		var err error
		pattern := "*.opus"
		if *mode == "encode" {
			pattern = "*.wav"
		}
		paths, err = filepath.Glob(pattern)
		if err != nil || len(paths) == 0 {
			fmt.Fprintf(os.Stderr, "provide input files, or run in a directory containing %s\n", pattern)
			os.Exit(2)
		}
	}
	fmt.Printf("C library: %s; mode=%s PCM=int16 rate=48000 tolerance=%d max-packets=%d\n", C.GoString(C.opus_get_version_string()), *mode, *tolerance, *limit)
	opt := encodeOptions{*bitrate, *complexity, *vbr, *exact, *nrmse, *sizeDiff}
	if *mode == "encode" {
		fmt.Printf("encode: application=audio frame-ms=20 bitrate=%d complexity=%d vbr=%t exact=%t max-nrmse=%g max-size-diff=%g\n", *bitrate, *complexity, *vbr, *exact, *nrmse, *sizeDiff)
	}
	failed := false
	for _, path := range paths {
		var err error
		if *mode == "encode" {
			err = compareEncode(path, opt, *limit)
		} else {
			err = compare(path, *tolerance, *limit)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", path, err)
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
}
