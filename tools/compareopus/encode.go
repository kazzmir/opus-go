//go:build compareopus && cgo

package main

/*
#include <opus.h>

// cgo cannot call variadic opus_encoder_ctl directly.
static int configure_encoder(OpusEncoder *enc, int bitrate, int complexity, int vbr) {
    int ret = opus_encoder_ctl(enc, OPUS_SET_BITRATE(bitrate));
    if (ret != OPUS_OK) return ret;
    ret = opus_encoder_ctl(enc, OPUS_SET_COMPLEXITY(complexity));
    if (ret != OPUS_OK) return ret;
    return opus_encoder_ctl(enc, OPUS_SET_VBR(vbr));
}
static int encoder_lookahead(OpusEncoder *enc, int *out) {
    return opus_encoder_ctl(enc, OPUS_GET_LOOKAHEAD(out));
}
*/
import "C"

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"unsafe"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
	"github.com/kazzmir/opus-go/wav"
)

type encodeOptions struct {
	bitrate, complexity   int
	vbr, exact            bool
	maxNRMSE, maxSizeDiff float64
}

func validThreshold(v float64) bool {
	return v >= 0 && !math.IsNaN(v) && !math.IsInf(v, 0)
}

func compareEncode(path string, opt encodeOptions, limit int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	wr, err := wav.NewReader(f)
	if err != nil {
		return err
	}
	channels := wr.Channels()
	if wr.SampleRate() != 48000 || (channels != 1 && channels != 2) {
		return fmt.Errorf("encode requires 48 kHz mono/stereo 16-bit PCM WAV; got %d Hz, %d channels", wr.SampleRate(), channels)
	}
	ge, err := opus.NewEncoder(48000, channels, opus.ApplicationAudio)
	if err != nil {
		return err
	}
	defer ge.Close()
	if err := ge.SetBitrate(opt.bitrate); err != nil {
		return err
	}
	if err := ge.SetComplexity(opt.complexity); err != nil {
		return err
	}
	if err := ge.SetVBR(opt.vbr); err != nil {
		return err
	}
	var code C.int
	ce := C.opus_encoder_create(48000, C.int(channels), C.OPUS_APPLICATION_AUDIO, &code)
	if ce == nil || code != C.OPUS_OK {
		return fmt.Errorf("C encoder create: %s", C.GoString(C.opus_strerror(code)))
	}
	defer C.opus_encoder_destroy(ce)
	vbr := C.int(0)
	if opt.vbr {
		vbr = 1
	}
	if code = C.configure_encoder(ce, C.int(opt.bitrate), C.int(opt.complexity), vbr); code != C.OPUS_OK {
		return fmt.Errorf("C encoder configure: %s", C.GoString(C.opus_strerror(code)))
	}
	lookahead, err := ge.Lookahead()
	if err != nil {
		return err
	}
	var cLookahead C.int
	if code = C.encoder_lookahead(ce, &cLookahead); code != C.OPUS_OK {
		return fmt.Errorf("C encoder lookahead: %s", C.GoString(C.opus_strerror(code)))
	}
	if lookahead != int(cLookahead) {
		return fmt.Errorf("lookahead differs: Go=%d C=%d", lookahead, cLookahead)
	}
	// Use the SAME decoder implementation for both streams, with independent
	// state, so this test isolates encoder differences from Go decoder errors.
	head := ogg.OpusHead{Channels: uint8(channels)}
	gd, err := newNative(head)
	if err != nil {
		return err
	}
	defer gd.close()
	cd, err := newNative(head)
	if err != nil {
		return err
	}
	defer cd.close()
	const frame = 960 // 20 ms
	pcm := make([]int16, frame*channels)
	gp, cp := make([]byte, 4000), make([]byte, 4000)
	gpcm, cpcm := make([]int16, maxFrame*channels), make([]int16, maxFrame*channels)
	var packets, unequal, sizeUnequal, goBytes, cBytes int64
	var fed, realFrames, samples, pcmUnequal int64
	var squaredError, referenceEnergy float64
	maxError := 0
	eof := false
	for {
		clear(pcm)
		if !eof {
			// -max-packets limits source audio, then still flushes lookahead.
			if limit > 0 && packets >= int64(limit) {
				eof = true
			} else {
				n, err := wr.ReadInt16PCM(pcm)
				if err != nil && err != io.EOF {
					return err
				}
				if n%channels != 0 {
					return fmt.Errorf("partial interleaved WAV frame")
				}
				realFrames += int64(n / channels)
				eof = err == io.EOF || n < len(pcm)
			}
		}
		if realFrames == 0 && eof {
			return fmt.Errorf("empty WAV")
		}
		if eof && fed >= realFrames+int64(lookahead) {
			break
		}
		gn, gerr := ge.Encode(pcm, frame, gp)
		cn := int(C.opus_encode(ce, (*C.opus_int16)(unsafe.Pointer(&pcm[0])), frame, (*C.uchar)(unsafe.Pointer(&cp[0])), C.opus_int32(len(cp))))
		packets++
		if gerr != nil || cn < 0 {
			return fmt.Errorf("packet %d: Go encode error=%v; C encode return=%d", packets, gerr, cn)
		}
		if gn <= 0 || cn == 0 {
			return fmt.Errorf("packet %d: empty encoded packet", packets)
		}
		goBytes += int64(gn)
		cBytes += int64(cn)
		if gn != cn {
			sizeUnequal++
		}
		if !bytes.Equal(gp[:gn], cp[:cn]) {
			if unequal == 0 {
				fmt.Printf("%s: first unequal payload: packet=%d Go-bytes=%d C-bytes=%d\n", path, packets, gn, cn)
			}
			unequal++
		}
		gframes, gerr := gd.decode(gp[:gn], gpcm)
		cframes, cerr := cd.decode(cp[:cn], cpcm)
		if gerr != nil || cerr != nil || gframes != frame || cframes != frame {
			return fmt.Errorf("packet %d: roundtrip lengths Go=%d C=%d; errors=%v, %v", packets, gframes, cframes, gerr, cerr)
		}
		for i := 0; i < frame*channels; i++ {
			pos := fed + int64(i/channels)
			// Compare only real audio, excluding codec delay and tail padding.
			if pos < int64(lookahead) || pos >= realFrames+int64(lookahead) {
				continue
			}
			delta := int(gpcm[i]) - int(cpcm[i])
			if delta < 0 {
				delta = -delta
			}
			if delta != 0 {
				pcmUnequal++
			}
			if delta > maxError {
				maxError = delta
			}
			squaredError += float64(delta) * float64(delta)
			referenceEnergy += float64(cpcm[i]) * float64(cpcm[i])
			samples++
		}
		fed += frame
	}
	if samples != realFrames*int64(channels) {
		return fmt.Errorf("roundtrip sample count %d, expected %d", samples, realFrames*int64(channels))
	}
	nrmse := 0.0
	if referenceEnergy > 0 {
		nrmse = math.Sqrt(squaredError / referenceEnergy)
	} else if squaredError > 0 {
		nrmse = math.Inf(1)
	}
	sizeDiff := math.Abs(float64(goBytes-cBytes)) / float64(cBytes)
	pass := nrmse <= opt.maxNRMSE && sizeDiff <= opt.maxSizeDiff && (!opt.exact || unequal == 0)
	status := "PASS"
	if !pass {
		status = "FAIL"
	}
	fmt.Printf("%s %s: packets=%d unequal-packets=%d unequal-lengths=%d Go-bytes=%d C-bytes=%d size-diff=%.6g samples=%d PCM-different=%d max-error=%d RMS-error=%.6g NRMSE=%.6g\n", status, path, packets, unequal, sizeUnequal, goBytes, cBytes, sizeDiff, samples, pcmUnequal, maxError, math.Sqrt(squaredError/float64(samples)), nrmse)
	if !pass {
		return fmt.Errorf("encode comparison exceeds thresholds (exact=%t, max-nrmse=%g, max-size-diff=%g)", opt.exact, opt.maxNRMSE, opt.maxSizeDiff)
	}
	return nil
}
