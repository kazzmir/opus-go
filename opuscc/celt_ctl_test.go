package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestAmp2Log2LocalUnion(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	bandE := []OpusT_celt_ener{0.75, 1.5, 3.25, 0.625, 2.25, 4.5}
	bandLogE := make([]OpusT_celt_glog, len(bandE))

	Opus_amp2Log2(nil, uintptr(unsafe.Pointer(&mode)), 2, 3, uintptr(unsafe.Pointer(&bandE[0])), uintptr(unsafe.Pointer(&bandLogE[0])), 2)

	want := []OpusT_celt_glog{-6.8525376, -5.6650376, -14, -7.115572, -5.0800753, -14}
	for i, expected := range want {
		if got := bandLogE[i]; got != expected {
			t.Fatalf("bandLogE[%d]: got %v, want %v", i, got, expected)
		}
	}
}

func TestQuantCoarseEnergyLocalQI(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	eBands := [6]OpusT_celt_glog{1.3, -0.8, 2.1, 0.4, 1.8, -1.2}
	oldEBands := [6]OpusT_celt_glog{0.1, -0.5, 1.4, 0.3, 1.1, -0.9}
	errors := [6]OpusT_celt_glog{}
	buffer := make([]byte, 32)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(tls, uintptr(unsafe.Pointer(&encoder)), entropyBufferPointer(buffer), uint32(len(buffer)))

	badness := quant_coarse_energy_impl(
		tls,
		uintptr(unsafe.Pointer(&mode)),
		0,
		3,
		uintptr(unsafe.Pointer(&eBands[0])),
		uintptr(unsafe.Pointer(&oldEBands[0])),
		120,
		0,
		uintptr(unsafe.Pointer(&e_prob_model[0][0][0])),
		uintptr(unsafe.Pointer(&errors[0])),
		uintptr(unsafe.Pointer(&encoder)),
		2,
		0,
		0,
		16,
		0,
	)
	Opus_ec_enc_done(tls, uintptr(unsafe.Pointer(&encoder)))

	if got, want := badness, int32(0); got != want {
		t.Fatalf("badness: got %d, want %d", got, want)
	}
	if got, want := oldEBands, [6]OpusT_celt_glog{1.0898438, -0.36923218, 2.337799, 0.26953125, 1.9882812, -0.7286072}; got != want {
		t.Fatalf("reconstructed energies: got %v, want %v", got, want)
	}
	if got, want := errors, [6]OpusT_celt_glog{0.2101562, -0.43076783, -0.23779917, 0.13046876, -0.1882813, -0.47139287}; got != want {
		t.Fatalf("quantization errors: got %v, want %v", got, want)
	}
	if got, want := buffer[:encoder.Foffs], []byte{0x69, 0x06}; string(got) != string(want) {
		t.Fatalf("encoded bytes: got % x, want % x", got, want)
	}
	if got, want := encoder.Fnbits_total, int32(41); got != want {
		t.Fatalf("encoded bit count: got %d, want %d", got, want)
	}
}
