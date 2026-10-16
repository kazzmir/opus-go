package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestAntiCollapseOwnedPointers(t *testing.T) {
	bands := [3]int16{0, 1, 3}
	mode := OpusT_OpusCustomMode{FeBands: &bands[0], FnbEBands: 2}
	for LM := int32(0); LM <= 3; LM++ {
		for _, C := range []int32{1, 2} {
			for _, encode := range []int32{0, 1} {
				for _, full := range []bool{false, true} {
					for _, aliasHistory := range []bool{false, true} {
						for _, seed := range []uint32{0, 0xdeadbeef} {
							size := int32(3) << LM
							a := make([]float32, C*size+2)
							for i := range a {
								a[i] = float32(i%7-3) / 17
							}
							a[0], a[len(a)-1] = 77, 88
							b := slices.Clone(a)
							masks := make([]byte, 2*C)
							if full {
								for i := range masks {
									masks[i] = byte((uint32(1) << uint32(int32(1)<<LM)) - 1)
								}
							}
							pulse := [2]int32{0, 16}
							energy := [4]float32{.5, 1, 1.5, 2}
							h1, h2 := [4]float32{-6, -5, -4, -3}, [4]float32{-5, -4, -3, -2}
							oldHistory, ownedHistory := &h1[0], &h1[0]
							if aliasHistory {
								oldHistory, ownedHistory = &a[1], &b[1]
							}
							entropyInitGrowStack(12)
							runtime.GC()
							Opus_anti_collapse(nil, &bands[0], 2, &a[1], &masks[0], LM, C, size, 0, 2, &energy[0], oldHistory, &h2[0], &pulse[0], seed, encode, 0)
							antiCollapseOwned(nil, &bands[0], 2, b[1:len(b)-1], masks, LM, C, size, 0, 2, &energy[0], ownedHistory, &h2[0], &pulse[0], seed, encode, 0)
							for i := range a {
								if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
									t.Fatal("owned anti-collapse bits/history aliases/guards", LM, C, encode, full, aliasHistory, seed, i)
								}
							}
						}
					}
				}
			}
		}
	}
	state := OpusT_OpusCustomDecoder{Frng: 123}
	before := state
	a, b := [5]float32{77, 0, 0, 0, 88}, [5]float32{77, 0, 0, 0, 88}
	masks, pulse := [2]byte{}, [2]int32{0, 16}
	energy, history := [2]float32{.5, 1}, [4]float32{-6, -5, -4, -3}
	celtDecodeAntiCollapse(nil, &state, &mode, &a[1], &masks[0], &pulse[0], &energy[0], &history[0], &history[0], 3, 0, 1, 0, 2, 1)
	celtDecodeAntiCollapseOwned(nil, &state, &mode, b[1:4], masks[:], &pulse[0], &energy[0], &history[0], &history[0], 3, 0, 1, 0, 2, 1)
	if a != b || state != before {
		t.Fatal("decode anti-collapse spectrum and unchanged seed/state")
	}
	antiCollapseOwned(nil, nil, 0, nil, nil, 0, 1, 0, 0, 0, nil, nil, nil, nil, 0, 0, 0)
	celtDecodeAntiCollapseOwned(nil, nil, nil, nil, nil, nil, nil, nil, nil, 120, 0, 1, 0, 21, 0)
}

func TestAntiCollapseMaskAliasPointers(t *testing.T) {
	bands := [3]int16{0, 1, 3}
	a, b := [5]float32{77, 0, 0, 0, 88}, [5]float32{77, 0, 0, 0, 88}
	// The first repaired float changes bytes used by the next band's mask.
	oldMask := unsafe.Slice((*byte)(unsafe.Pointer(&a[1])), 4)[2:4]
	ownedMask := unsafe.Slice((*byte)(unsafe.Pointer(&b[1])), 4)[2:4]
	pulse, energy, history := [2]int32{0, 16}, [2]float32{.5, 1}, [4]float32{-6, -5, -4, -3}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_anti_collapse(nil, &bands[0], 2, &a[1], &oldMask[0], 0, 1, 3, 0, 2, &energy[0], &history[0], &history[0], &pulse[0], 0, 0, 0)
	antiCollapseOwned(nil, &bands[0], 2, b[1:4], ownedMask, 0, 1, 3, 0, 2, &energy[0], &history[0], &history[0], &pulse[0], 0, 0, 0)
	for i := range a {
		if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
			t.Fatal("live mask/spectrum alias", i)
		}
	}
}

func TestAntiCollapsePointers(t *testing.T) {
	bands := [2]int16{0, 2}
	x := [4]float32{77, 0, 0, 88}
	span := x
	mask := byte(0)
	energy := float32(2.5)
	p1, p2 := [2]float32{.25, 0}, [2]float32{.75, 0}
	pulse := int32(3)
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_anti_collapse(nil, &bands[0], 1, &x[1], &mask, 0, 1, 2, 0, 1, &energy, &p1[0], &p2[0], &pulse, 123456, 0, 0)
	antiCollapseSpans(nil, bands[:], 1, span[1:3], []byte{mask}, []int32{pulse}, []float32{energy}, p1[:], p2[:], 0, 1, 2, 0, 1, 123456, 0, 0)
	if span != x {
		t.Fatal("anti-collapse slice kernel/guards")
	}
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
