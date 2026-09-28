package opuscc

import (
	"testing"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDenormaliseBandsLocalExp2Union(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	eBands := [3]int16{0, 2, 4}
	x := [4]OpusT_celt_norm{0.25, -0.5, 0.75, -1}
	freq := [4]OpusT_celt_sig{}
	bandLogE := [2]OpusT_celt_glog{-5.5, -6.25}

	Opus_denormalise_bands(
		tls,
		&eBands[0],
		4,
		&x[0],
		&freq[0],
		&bandLogE[0],
		0,
		2,
		1,
		1,
		0,
	)

	if got, want := freq, [4]OpusT_celt_sig{0.47880167, -0.95760334, 0.74999994, -0.99999994}; got != want {
		t.Fatalf("denormalised spectrum: got %v, want %v", got, want)
	}
}
