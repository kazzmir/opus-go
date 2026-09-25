package opusccenc

import (
	"math/rand"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
)

func shimSlice[T any](tls *libc.TLS, n int) (uintptr, []T) {
	var z T
	p := libc.Xmalloc(tls, uint64(n)*uint64(unsafe.Sizeof(z)))
	return p, unsafe.Slice((*T)(unsafe.Pointer(p)), n)
}

// TestNLSFDecodeMatchesDecoder checks the encoder's SILK NLSF
// dequantizer against the decoder's copy of the same libopus function.
// Negative quantization indices went through a zero-extending
// (uint32)(int8) conversion in libcshim, so the encoder rebuilt
// nonsense spectra for almost every frame and SILK/hybrid output was
// garbage.
func TestNLSFDecodeMatchesDecoder(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	rng := rand.New(rand.NewSource(1))
	for _, cb := range []struct {
		name  string
		e, d  uintptr
		order int
	}{
		{"WB", uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_WB)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_NLSF_CB_WB)), 16},
		{"NB_MB", uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_NB_MB)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_NLSF_CB_NB_MB)), 10},
	} {
		for trial := range 1000 {
			idxP, idx := shimSlice[int8](tls, cb.order+1)
			idx[0] = int8(rng.Intn(32))
			for i := 1; i <= cb.order; i++ {
				idx[i] = int8(rng.Intn(9) - 4)
			}
			encP, encOut := shimSlice[int16](tls, cb.order)
			decP, decOut := shimSlice[int16](tls, cb.order)
			Opus_silk_NLSF_decode(tls, encP, idxP, cb.e)
			opuscc.Opus_silk_NLSF_decode(tls, decP, idxP, cb.d)
			for i := range cb.order {
				if encOut[i] != decOut[i] {
					t.Fatalf("%s trial %d, indices %v: encoder %v, decoder %v", cb.name, trial, idx, encOut, decOut)
				}
			}
		}
	}
}
