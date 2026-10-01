package opuscc

import (
	"runtime"
	"testing"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecodeIndicesPointers(t *testing.T) {
	for _, fs := range []int32{8, 12, 16} {
		for _, sub := range []int32{2, 4} {
			var st OpusT_silk_decoder_state
			st.Fnb_subfr = sub
			Opus_silk_decoder_set_fs(nil, &st, fs, 16000)
			st.FVAD_flags[0] = 1
			packet := []byte{0x93, 0x57, 0xc1, 0x2a, 0xee, 0x44, 0x18, 0xb7}
			var dec OpusT_ec_dec
			Opus_ec_dec_init(nil, &dec, &packet[0], uint32(len(packet)))
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_silk_decode_indices(nil, &st, &dec, 0, 0, CODE_INDEPENDENTLY)
			if st.Findices.FsignalType < 1 || st.Findices.FsignalType > 2 || st.Findices.FSeed < 0 || st.Findices.FSeed > 3 || dec.Fbuf != &packet[0] {
				t.Fatal("indices/ownership")
			}
		}
	}
}

func TestDecodeIndicesFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder, 8, 8000); got != OPUS_OK {
		t.Fatalf("set decoder sample rate: got %d", got)
	}
	data := []byte{0x93, 0x57, 0xc1, 0x2a, 0xee, 0x44, 0x18, 0xb7, 0x6d, 0x09, 0xfa, 0x35, 0x81, 0x62, 0xdc, 0x4e}
	var rangeDecoder OpusT_ec_dec
	Opus_ec_dec_init(tls, &rangeDecoder, &data[0], uint32(len(data)))
	Opus_silk_decode_indices(tls, &decoder, &rangeDecoder, 0, 0, CODE_INDEPENDENTLY)
	if got, want := decoder.Findices.FsignalType, int8(0); got != want {
		t.Fatalf("signal type: got %d, want %d", got, want)
	}
	if got, want := decoder.Findices.FGainsIndices, [4]OpusT_opus_int8{15, 4, 4, 3}; got != want {
		t.Fatalf("gain indices: got %v, want %v", got, want)
	}
	if got, want := decoder.Findices.FNLSFIndices[:10], []int8{12, -1, -2, 1, -2, -2, 0, -1, 0, 0}; !equalInt8s(got, want) {
		t.Fatalf("NLSF indices: got %v, want %v", got, want)
	}
	if got, want := decoder.Findices.FSeed, int8(1); got != want {
		t.Fatalf("seed: got %d, want %d", got, want)
	}
}

func equalInt8s(got, want []int8) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
