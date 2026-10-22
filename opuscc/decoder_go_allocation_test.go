package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"testing"
	"unsafe"
	"weak"
)

// Explicit fixture only: normal creation never registers a root or codec views.
func legacyRegisteredDecoderTestStorage(tls *libc.TLS, rate, channels int32) (*OpusT_OpusDecoder, error) {
	owned := opusDecoderOwnedStorage(channels)
	st := libc.RegisterGoAllocation(tls, owned.Decoder)
	if st == nil {
		return nil, opusErrorFromCode(-7)
	}
	if ret := opusDecoderInitWithStorage(tls, st, rate, channels, &owned); ret != 0 {
		libc.RemoveGoAllocation(tls, st)
		return nil, opusErrorFromCode(ret)
	}
	opusBindCodecViews(tls, owned)
	return st, nil
}

//go:noinline
func droppedDecoderGoOwner(t *testing.T, tls *libc.TLS, channels int32) weak.Pointer[OpusT_OpusDecoder] {
	st, err := Opus_opus_decoder_create_typed(tls, 48000, channels)
	if err != nil {
		t.Fatal(err)
	}
	if libc.RegisteredGoAllocation[OpusT_OpusDecoder](tls, uintptr(unsafe.Pointer(st))) != nil {
		t.Fatal("factory registered Go owner")
	}
	if _, ok := libc.GoAllocationMetadata[OpusT_OpusDecoder, opusCodecBinding](tls, st); ok {
		t.Fatal("factory registered codec views")
	}
	// No destructor: dropping Go pointers is sufficient, even with TLS still live.
	return weak.Make(st)
}

func TestDecoderGoCollectionPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	for _, channels := range []int32{1, 2} {
		w := droppedDecoderGoOwner(t, tls, channels)
		for i := 0; i < 10; i++ {
			runtime.GC()
		}
		if w.Value() != nil {
			t.Fatal("Go owner retained without pointers/destructor")
		}
	}
	runtime.KeepAlive(tls)
}

func TestDecoderGoNoFreePointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		tls := libc.NewTLS()
		st, err := Opus_opus_decoder_create_typed(nil, 48000, channels)
		if err != nil {
			t.Fatal(err)
		}
		entropyInitGrowStack(12)
		runtime.GC()
		Opus_opus_decoder_destroy_typed(tls, st)
		tls.Close()
		runtime.GC()
		pcm := make([]float32, 120*channels+2)
		pcm[0], pcm[len(pcm)-1] = 77, 88
		if n := Opus_opus_decode_float_typed(nil, st, nil, 0, &pcm[1], 120, 0); n != 120 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("decode after no-op destroy/TLS close", n)
		}
		for _, sample := range pcm[1 : len(pcm)-1] {
			if sample != 0 {
				t.Fatal("initial PLC silence")
			}
		}
	}
}
