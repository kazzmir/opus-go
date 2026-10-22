package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"
)

func TestDecoderStaticGoClearExtentPointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		owned := opusDecoderOwnedStorage(channels)
		if uintptr(unsafe.Pointer(owned.Decoder))&15 != 0 {
			t.Fatal("static single decoder alignment", channels)
		}
		if opusDecoderInitWithStorage(nil, owned.Decoder, 48000, channels, &owned) != 0 {
			t.Fatal("init")
		}
		decoderStorageForeignPointers(owned.Decoder)
		needed := int(Opus_opus_decoder_get_size(nil, channels))
		bodySize := int(unsafe.Sizeof(opusGoMonoBody{}))
		if channels == 2 {
			bodySize = int(unsafe.Sizeof(opusGoStereoBody{}))
		}
		// Only poison final scalar Go padding, never a Go pointer or the cell Owner.
		image := unsafe.Slice((*byte)(unsafe.Pointer(owned.Decoder)), bodySize)
		for i := needed; i < len(image); i++ {
			image[i] = 0xa7
		}
		entropyInitGrowStack(12)
		runtime.GC()
		owned.Clear()
		for i, b := range image {
			want := byte(0)
			if i >= needed {
				want = 0xa7
			}
			if b != want {
				t.Fatal("typed clear ABI extent/Go padding", channels, i, b, want)
			}
		}
		runtime.GC()
		runtime.KeepAlive(owned)
	}
}

//go:noinline
func singleStaticGoOwner(t *testing.T, channels int32) (*OpusT_OpusDecoder, weak.Pointer[OpusT_OpusCustomMode], [2]weak.Pointer[OpusT_silk_NLSF_CB_struct], [2]weak.Pointer[[32]int16]) {
	st, err := Opus_opus_decoder_create_typed(nil, 48000, channels)
	if err != nil {
		t.Fatal(err)
	}
	mode, books, coefs := decoderStorageForeignPointers(st)
	return st, mode, books, coefs
}

func TestDecoderStaticGoOwnershipPointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		st, mode, books, coefs := singleStaticGoOwner(t, channels)
		entropyInitGrowStack(12)
		for trial := 0; trial < 3; trial++ {
			runtime.GC()
			if mode.Value() == nil {
				t.Fatal("header lost static Go CELT owner")
			}
			for i := range books {
				if books[i].Value() == nil || coefs[i].Value() == nil {
					t.Fatal("header lost static Go SILK owner", i)
				}
			}
		}
		var rate int32
		if Opus_opus_decoder_ctl_typed(nil, st, OPUS_GET_SAMPLE_RATE_REQUEST, OpusDecoderCtlArgs{I32: &rate}) != 0 || rate != 48000 {
			t.Fatal("typed header after GC")
		}
		runtime.KeepAlive(st)
	}
}
