package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestModeTablePointers(t *testing.T) {
	makeMode := func() *OpusT_OpusCustomMode {
		m := mode48000_960_120
		window := slices.Clone(window120[:])
		vectors := slices.Clone(band_allocation[:])
		caps := slices.Clone(cache_caps50[:])
		m.Fwindow = &window[0]
		m.FallocVectors = &vectors[0]
		m.Fcache.Fcaps = &caps[0]
		return &m
	}
	m := makeMode()
	entropyInitGrowStack(12)
	runtime.GC()
	if !slices.Equal(unsafe.Slice(m.FallocVectors, len(band_allocation)), band_allocation[:]) || !slices.Equal(unsafe.Slice(m.Fwindow, 120), window120[:]) {
		t.Fatal("owned mode tables")
	}
	for lm := int32(0); lm <= 3; lm++ {
		for channels := int32(1); channels <= 2; channels++ {
			g, c := [23]int32{}, [23]int32{}
			g[0] = 77
			g[22] = 88
			c = g
			Opus_init_caps(nil, &eband5ms[0], m.Fcache.Fcaps, &g[1], 21, lm, channels)
			Opus_init_caps(nil, &eband5ms[0], &cache_caps50[0], &c[1], 21, lm, channels)
			if g != c {
				t.Fatal("owned caps")
			}
		}
	}
	var input [240]float32
	input[17] = 1
	var g, c [120]float32
	st := m.Fmdct.Fkfft[3]
	Opus_clt_mdct_forward_c(nil, &m.Fmdct, st.Fbitrev, st.Ftwiddles, &input[0], &g[0], m.Fwindow, 120, 3, 1, 0)
	Opus_clt_mdct_forward_c(nil, &mode48000_960_120.Fmdct, st.Fbitrev, st.Ftwiddles, &input[0], &c[0], &window120[0], 120, 3, 1, 0)
	if g != c {
		t.Fatal("owned window transform")
	}
}

func TestModeLogPointers(t *testing.T) {
	makeMode := func() *OpusT_OpusCustomMode {
		m := mode48000_960_120
		log := slices.Clone(logN400[:])
		m.FlogN = &log[0]
		return &m
	}
	m := makeMode()
	entropyInitGrowStack(12)
	runtime.GC()
	for i, want := range logN400 {
		if got := modeLogN(m, int32(i)); got != want {
			t.Fatal(i, got, want)
		}
	}
	signed := []int16{-32768, -1, 0, 32767}
	m.FlogN = &signed[0]
	for i, want := range signed {
		if modeLogN(m, int32(i)) != want {
			t.Fatal("signed log", i)
		}
	}
	runtime.KeepAlive(m)
}

func TestCapsPointers(t *testing.T) {
	bands := [4]int16{0, 1, 3, 7}
	var cache [24]uint8
	for i := range cache {
		cache[i] = uint8(i * 7)
	}
	for lm := int32(0); lm < 4; lm++ {
		for channels := int32(1); channels <= 2; channels++ {
			out := [5]int32{123, 0, 0, 0, 456}
			Opus_init_caps(nil, &bands[0], &cache[0], &out[1], 3, lm, channels)
			for i := int32(0); i < 3; i++ {
				want := (int32(cache[3*(2*lm+channels-1)+i]) + 64) * channels * (int32(bands[i+1]-bands[i]) << lm) / 4
				if out[i+1] != want {
					t.Fatalf("LM=%d C=%d i=%d got=%d want=%d", lm, channels, i, out[i+1], want)
				}
			}
			if out[0] != 123 || out[4] != 456 {
				t.Fatal("guard changed")
			}
		}
	}
}
