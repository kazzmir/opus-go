package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// Only ABI comparison fixtures need a contiguous transport layout. These are
// actual Go fields, never registered MS allocations and never explicitly freed.
func flatMSComparisonStorage(streams, coupled int32) opusMSStorageViews {
	return opusMSStorageViewsAt(opusDecoderAlignedStorage(opusMSDecoderStorageType(streams, coupled)))
}

func flatMSTestDecoder(tls *libc.TLS, rate, channels, streams, coupled int32, mapping *byte) (*OpusT_OpusMSDecoder, error) {
	owned := flatMSComparisonStorage(streams, coupled)
	if ret := opusMSDecoderInitWithCodecs(tls, owned.Decoder, rate, channels, streams, coupled, mapping, owned.Children, owned.Codecs); ret != OPUS_OK {
		return nil, opusErrorFromCode(ret)
	}
	return owned.Decoder, nil
}

//go:noinline
func unregisteredMSOwner(t *testing.T) (*OpusMSDecoder, [2]weak.Pointer[OpusT_OpusCustomMode], [2][2]weak.Pointer[OpusT_silk_NLSF_CB_struct], [2][2]weak.Pointer[[32]int16]) {
	mapping := [3]byte{0, 1, 2}
	st, err := Opus_opus_multistream_decoder_create_typed(48000, 3, 2, 1, &mapping[0])
	if err != nil {
		t.Fatal(err)
	}
	var modes [2]weak.Pointer[OpusT_OpusCustomMode]
	var books [2][2]weak.Pointer[OpusT_silk_NLSF_CB_struct]
	var coefs [2][2]weak.Pointer[[32]int16]
	for i := range modes {
		var child *OpusT_OpusDecoder
		if Opus_opus_multistream_decoder_ctl_typed(nil, st, OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, OpusDecoderCtlArgs{Value: int32(i), Decoder: &child}) != 0 {
			t.Fatal("child")
		}
		modes[i], books[i], coefs[i] = decoderStorageForeignPointers(child)
	}
	return st, modes, books, coefs
}

func TestMSDecoderGoOwnershipPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	st, modes, books, coefs := unregisteredMSOwner(t)
	if libc.RegisteredGoAllocation[OpusMSDecoder](tls, uintptr(unsafe.Pointer(st))) != nil {
		t.Fatal("Go factory registered its parent")
	}
	if _, ok := libc.GoAllocationMetadata[OpusT_OpusMSDecoder, opusMSCodecBinding](tls, st.opusMSHeader()); ok {
		t.Fatal("Go factory requires TLS metadata")
	}
	entropyInitGrowStack(12)
	for trial := 0; trial < 3; trial++ {
		runtime.GC()
		for i := range modes {
			if modes[i].Value() == nil {
				t.Fatal("Go parent lost child mode")
			}
			for j := range books[i] {
				if books[i][j].Value() == nil || coefs[i][j].Value() == nil {
					t.Fatal("Go parent lost child SILK owner")
				}
			}
		}
	}
	// Destroy is a no-op, and closing runtime TLS cannot destroy Go storage.
	Opus_opus_multistream_decoder_destroy_typed(tls, st)
	tls.Close()
	var child *OpusT_OpusDecoder
	if Opus_opus_multistream_decoder_ctl_typed(nil, st, OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, OpusDecoderCtlArgs{Value: 1, Decoder: &child}) != 0 {
		t.Fatal("TLS-independent child")
	}
	st = nil
	runtime.GC()
	if child.FFs != 48000 || modes[0].Value() == nil || modes[1].Value() == nil {
		t.Fatal("child back reference must retain Go parent/siblings")
	}
	var rate int32
	if Opus_opus_decoder_ctl_typed(nil, child, OPUS_GET_SAMPLE_RATE_REQUEST, OpusDecoderCtlArgs{I32: &rate}) != 0 || rate != 48000 {
		t.Fatal("CTL after parent/TLS dropped")
	}
	runtime.KeepAlive(child)
}

func TestMSDecoderGoDecodePointers(t *testing.T) {
	mapping := [3]byte{0, 1, 2}
	st, err := Opus_opus_multistream_decoder_create_typed(48000, 3, 2, 1, &mapping[0])
	if err != nil {
		t.Fatal(err)
	}
	Opus_opus_multistream_decoder_destroy_typed(nil, st)
	pcm := make([]float32, 362)
	pcm[0], pcm[361] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_multistream_decode_float_typed(nil, st, nil, 0, &pcm[1], 120, 0) != 120 || pcm[0] != 77 || pcm[361] != 88 {
		t.Fatal("decode without TLS allocation/registration")
	}
	for _, value := range pcm[1:361] {
		if value != 0 {
			t.Fatal("initial PLC must be silent")
		}
	}
}

//go:noinline
func droppedMSGoOwner(t *testing.T) weak.Pointer[OpusMSDecoder] {
	mapping := [1]byte{0}
	st, err := Opus_opus_multistream_decoder_create_typed(48000, 1, 1, 0, &mapping[0])
	if err != nil {
		t.Fatal(err)
	}
	return weak.Make(st)
}

func TestMSDecoderGoCollectionPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	w := droppedMSGoOwner(t)
	for trial := 0; trial < 10; trial++ {
		runtime.GC()
	}
	if w.Value() != nil {
		t.Fatal("Go multistream retained without a Go owner")
	}
	runtime.KeepAlive(tls)
}
