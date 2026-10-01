package opuscc

import (
	"runtime"
	"testing"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestAntiCollapsePointers(t *testing.T) {
	bands := [2]int16{0, 2}
	x := [4]float32{77, 0, 0, 88}
	mask := byte(0)
	energy := float32(2.5)
	p1, p2 := [2]float32{.25, 0}, [2]float32{.75, 0}
	pulse := int32(3)
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_anti_collapse(nil, &bands[0], 1, &x[1], &mask, 0, 1, 2, 0, 1, &energy, &p1[0], &p2[0], &pulse, 123456, 0, 0)
	if x != [4]float32{77, .7071068, .7071068, 88} {
		t.Fatal(x)
	}
	Opus_anti_collapse(nil, nil, 0, nil, nil, 0, 1, 0, 0, 0, nil, nil, nil, nil, 0, 0, 0)
}

func TestAntiCollapseLocalExp2Union(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	eBands := [2]int16{0, 2}
	x := [2]OpusT_celt_norm{0, 0}
	collapseMasks := [1]byte{0}
	logE := [1]OpusT_celt_glog{2.5}
	prev1LogE := [2]OpusT_celt_glog{0.25, 0}
	prev2LogE := [2]OpusT_celt_glog{0.75, 0}
	pulses := [1]int32{3}

	Opus_anti_collapse(
		tls,
		&eBands[0], 1,
		&x[0],
		&collapseMasks[0],
		0,
		1,
		2,
		0,
		1,
		&logE[0],
		&prev1LogE[0],
		&prev2LogE[0],
		&pulses[0],
		123456,
		0,
		0,
	)

	if got, want := x, [2]OpusT_celt_norm{0.7071068, 0.7071068}; got != want {
		t.Fatalf("repaired spectrum: got %v, want %v", got, want)
	}
}
