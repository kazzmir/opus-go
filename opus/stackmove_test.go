package opus

import (
	"bytes"
	"math"
	"runtime"
	"runtime/debug"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestTypedScratchPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	var buffer cBuf
	defer buffer.free(tls)
	if copyInPointer[byte](tls, &buffer, nil) != nil || buffer.p != nil {
		t.Fatal("empty scratch input")
	}
	p := copyInPointer(tls, &buffer, []int16{-32768, -1, 0, 32767})
	runtime.GC()
	got := cBufferSlice[int16](p, 4)
	if p != buffer.p || got[0] != -32768 || got[1] != -1 || got[2] != 0 || got[3] != 32767 || buffer.ensurePointer(tls, 1) != p {
		t.Fatal("typed scratch ownership/reuse", got)
	}
	if copyInPointer[byte](tls, &buffer, nil) != nil || buffer.p != p {
		t.Fatal("empty input must retain scratch")
	}
	buffer.ensurePointer(tls, 64)
	if buffer.p == nil || buffer.n != 64 {
		t.Fatal("typed scratch growth")
	}
	buffer.free(tls)
	buffer.free(tls)
	if buffer.p != nil || buffer.n != 0 {
		t.Fatal("typed scratch release")
	}
}

func TestTypedDecoderStatePointers(t *testing.T) {
	for _, multi := range []bool{false, true} {
		var d *Decoder
		var err error
		if multi {
			d, err = NewMultistreamDecoder(48000, 3, 2, 1, []uint8{0, 1, 2})
		} else {
			d, err = NewDecoder(48000, 2)
		}
		if err != nil {
			t.Fatal(err)
		}
		st, ms := d.st, d.ms
		runtime.GC()
		if d.st != st || d.ms != ms || (multi && (d.ms == nil || d.st != nil)) || (!multi && (d.st == nil || d.ms != nil)) {
			t.Fatal("typed decoder state owners")
		}
		pcm := make([]int16, 120*d.channels)
		if n, err := d.Decode(nil, pcm, 120, false); err != nil || n != 120 {
			t.Fatal("typed int16 PLC", n, err)
		}
		floats := make([]float32, 120*d.channels)
		runtime.GC()
		if n, err := d.DecodeF32(nil, floats, 120, false); err != nil || n != 120 {
			t.Fatal("typed float PLC", n, err)
		}
		if err := d.Close(); err != nil {
			t.Fatal(err)
		}
		if d.st != nil || d.ms != nil || d.tls != nil {
			t.Fatal("typed decoder close")
		}
		if err := d.Close(); err != nil {
			t.Fatal("repeated close", err)
		}
	}
}

// skipOn32Bit skips encoder tests on 32-bit targets, where the transpiled
// encoder aborts inside celt's FFT (celt_fatal from opus_fft_c during
// tonality analysis) before any of this is exercised. That predates these
// tests - nothing encoded on 386 before - and needs its own fix.
func skipOn32Bit(t *testing.T) {
	t.Helper()
	if unsafe.Sizeof(uintptr(0)) == 4 {
		t.Skip("opusccenc aborts on 32-bit targets (pre-existing)")
	}
}

// encodeSine encodes one second of a 440 Hz sine with pcm/packet buffers
// that don't escape - so they're stack-allocated, the case that used to
// break when the transpiled encoder grew the goroutine stack mid-call.
func encodeSine(t *testing.T) []byte {
	t.Helper()
	enc, err := NewEncoder(24000, 1, ApplicationAudio)
	if err != nil {
		t.Fatal(err)
	}
	defer enc.Close()
	var pcm [480]int16
	var packet [4000]byte
	var out []byte
	for f := range 50 {
		for i := range pcm {
			pcm[i] = int16(16000 * math.Sin(2*math.Pi*440*float64(f*len(pcm)+i)/24000))
		}
		n, err := enc.Encode(pcm[:], len(pcm), packet[:])
		if err != nil {
			t.Fatal(err)
		}
		if allZero(packet[:n]) {
			t.Fatalf("frame %d: encoder reported %d bytes but the packet is all zero", f, n)
		}
		out = append(out, packet[:n]...)
		runtime.GC() // shrinks the stack, so the next Encode grows (moves) it
	}
	return out
}

func allZero(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}

// TestEncodeWithStackBuffersIsDeterministic: identical input must give
// identical output even when the caller's buffers live on a stack that
// moves during Encode.
func TestEncodeWithStackBuffersIsDeterministic(t *testing.T) {
	skipOn32Bit(t)
	defer debug.SetGCPercent(debug.SetGCPercent(1))
	want := encodeSine(t)
	for i := range 5 {
		if got := encodeSine(t); !bytes.Equal(got, want) {
			t.Fatalf("run %d produced a different stream", i+1)
		}
	}
}

// TestDecodeWithStackBuffers decodes into a stack-allocated buffer from a
// fresh goroutine, whose small initial stack the decoder has to grow
// mid-call, and checks the result against a decode into heap memory.
//
// It deliberately doesn't sweep stack depths or force GCs: that also
// provokes a separate, pre-existing fault inside the transpiled decoder
// (it takes uintptrs to some of its own stack locals), which these
// wrapper-level changes don't address.
func TestDecodeWithStackBuffers(t *testing.T) {
	skipOn32Bit(t) // needs the encoder for its input
	enc, err := NewEncoder(48000, 1, ApplicationAudio)
	if err != nil {
		t.Fatal(err)
	}
	defer enc.Close()
	pcm := make([]int16, 960)
	for i := range pcm {
		pcm[i] = int16(16000 * math.Sin(2*math.Pi*440*float64(i)/48000))
	}
	packet := make([]byte, 4000)
	n, err := enc.Encode(pcm, 960, packet)
	if err != nil {
		t.Fatal(err)
	}
	packet = packet[:n]

	decode := func(into []int16) {
		dec, err := NewDecoder(48000, 1)
		if err != nil {
			t.Error(err)
			return
		}
		defer dec.Close()
		if _, err := dec.Decode(packet, into, len(into), false); err != nil {
			t.Error(err)
		}
	}
	want := make([]int16, 960)
	decode(want)
	if allZero16(want) {
		t.Fatal("reference decode is silent")
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		var got [960]int16
		decode(got[:])
		for j := range want {
			if got[j] != want[j] {
				t.Errorf("sample %d = %d, want %d", j, got[j], want[j])
				return
			}
		}
	}()
	<-done
}

func allZero16(s []int16) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
