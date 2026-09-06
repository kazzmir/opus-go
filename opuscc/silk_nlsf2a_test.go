package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func runNLSF2A(t *testing.T, tls *libc.TLS, nlsf []OpusT_opus_int16) []OpusT_opus_int16 {
	t.Helper()
	a_Q12 := make([]OpusT_opus_int16, len(nlsf))
	Opus_silk_NLSF2A(tls, uintptr(unsafe.Pointer(&a_Q12[0])), uintptr(unsafe.Pointer(&nlsf[0])), int32(len(nlsf)), int32(0))
	return a_Q12
}

func checkNLSF2A(t *testing.T, label string, got, want []OpusT_opus_int16) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("%s: length mismatch: got %d, want %d", label, len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("%s: a_Q12[%d]: got %d, want %d", label, i, got[i], want[i])
		}
	}
}

func TestNLSF2ALocalArrays(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	/* expected values from the C reference implementation (silk/NLSF2A.c) */
	nlsf10 := []OpusT_opus_int16{468, 1398, 2410, 3514, 4701, 5985, 7381, 8912, 10606, 12500}
	want10 := []OpusT_opus_int16{15962, -29088, 32698, -25141, 13825, -5508, 1570, -306, 37, -2}
	checkNLSF2A(t, "d=10", runNLSF2A(t, tls, nlsf10), want10)

	nlsf16 := []OpusT_opus_int16{468, 1001, 1550, 2114, 2698, 3302, 3930, 4583, 5265, 5977, 6723, 7506, 8330, 9200, 10122, 11100}
	want16 := []OpusT_opus_int16{15556, -28280, 32673, -26853, 16648, -8054, 3102, -961, 241, -48, 8, -1, 0, 0, 0, 0}
	checkNLSF2A(t, "d=16", runNLSF2A(t, tls, nlsf16), want16)

	/* tightly clustered NLSFs exercise the bandwidth-expansion stabilization loop */
	tight10 := []OpusT_opus_int16{2000, 2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900}
	wantTight10 := []OpusT_opus_int16{16490, -30061, 32671, -23442, 11602, -4012, 957, -151, 14, -1}
	checkNLSF2A(t, "tight d=10", runNLSF2A(t, tls, tight10), wantTight10)
}
