package opus

import (
	"bytes"
	"math"
	"runtime"
	"runtime/debug"
	"testing"
	"unsafe"
)

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
