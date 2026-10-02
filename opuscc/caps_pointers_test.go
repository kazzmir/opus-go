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

func TestModePulseIndexPointers(t *testing.T) {
	makeMode := func() *OpusT_OpusCustomMode {
		m := mode48000_960_120
		index := slices.Clone(cache_index50[:])
		m.Fcache.Findex = &index[0]
		return &m
	}
	m := makeMode()
	entropyInitGrowStack(12)
	runtime.GC()
	for i, want := range cache_index50 {
		if got := modePulseIndex(m, int32(i)); got != want {
			t.Fatal(i, got, want)
		}
	}
	signed := []int16{-32768, -1, 0, 32767}
	m.Fcache.Findex = &signed[0]
	for i, want := range signed {
		if modePulseIndex(m, int32(i)) != want {
			t.Fatal("signed index", i)
		}
	}
	runtime.KeepAlive(m)
}

func TestModePulseBitsPointers(t *testing.T) {
	makeMode := func() *OpusT_OpusCustomMode {
		m := mode48000_960_120
		index := slices.Clone(cache_index50[:])
		bits := slices.Clone(cache_bits50[:])
		m.Fcache.Findex = &index[0]
		m.Fcache.Fbits = &bits[0]
		return &m
	}
	m := makeMode()
	entropyInitGrowStack(12)
	runtime.GC()
	if !slices.Equal(unsafe.Slice(m.Fcache.Fbits, int(m.Fcache.Fsize)), cache_bits50[:]) {
		t.Fatal("owned pulse bits")
	}
	for i, offset := range cache_index50 {
		if offset < 0 {
			continue
		}
		cache := modePulseCache(m, int32(i))
		max := int32(*cache)
		for pulse := int32(0); pulse <= max; pulse++ {
			want := int32(0)
			if pulse != 0 {
				want = int32(cache_bits50[int(offset)+int(pulse)]) + 1
			}
			if modePulses2Bits(cache, pulse) != want {
				t.Fatal(i, pulse)
			}
		}
		for bits := int32(-2); bits <= 260; bits++ {
			q := modeBits2Pulses(cache, bits)
			if q < 0 || q > max {
				t.Fatal(i, bits, q)
			}
		}
	}
	// Signed index -1 must remain a backwards interior offset, not an unsigned load.
	guarded := []byte{0, 77}
	index := []int16{-1}
	m.Fcache.Fbits = &guarded[1]
	m.Fcache.Findex = &index[0]
	if cache := modePulseCache(m, 0); cache != &guarded[0] || modeBits2Pulses(cache, 1) != 0 || modePulses2Bits(cache, 0) != 0 {
		t.Fatal("signed cache offset / zero pulses")
	}
	runtime.KeepAlive(m)
}

func TestModeBandPointers(t *testing.T) {
	makeMode := func() *OpusT_OpusCustomMode {
		m := mode48000_960_120
		bands := slices.Clone(eband5ms[:])
		log := slices.Clone(logN400[:])
		index := slices.Clone(cache_index50[:])
		bits := slices.Clone(cache_bits50[:])
		caps := slices.Clone(cache_caps50[:])
		window := slices.Clone(window120[:])
		vectors := slices.Clone(band_allocation[:])
		m.FeBands = &bands[0]
		m.FlogN = &log[0]
		m.Fcache.Findex = &index[0]
		m.Fcache.Fbits = &bits[0]
		m.Fcache.Fcaps = &caps[0]
		m.Fwindow = &window[0]
		m.FallocVectors = &vectors[0]
		return &m
	}
	m := makeMode()
	entropyInitGrowStack(12)
	runtime.GC()
	for i, want := range eband5ms {
		if modeBand(m, int32(i)) != want {
			t.Fatal("owned band", i)
		}
	}
	if !slices.Equal(unsafe.Slice(m.FlogN, len(logN400)), logN400[:]) || !slices.Equal(unsafe.Slice(m.Fcache.Findex, len(cache_index50)), cache_index50[:]) || !slices.Equal(unsafe.Slice(m.Fcache.Fbits, len(cache_bits50)), cache_bits50[:]) || !slices.Equal(unsafe.Slice(m.Fwindow, len(window120)), window120[:]) || !slices.Equal(unsafe.Slice(m.FallocVectors, len(band_allocation)), band_allocation[:]) {
		t.Fatal("complete mode table owners")
	}
	for lm := int32(0); lm <= 3; lm++ {
		for channels := int32(1); channels <= 2; channels++ {
			g, c := [23]int32{77}, [23]int32{77}
			g[22] = 88
			c[22] = 88
			Opus_init_caps(nil, m.FeBands, m.Fcache.Fcaps, &g[1], 21, lm, channels)
			Opus_init_caps(nil, &eband5ms[0], &cache_caps50[0], &c[1], 21, lm, channels)
			if g != c {
				t.Fatal("owned bands/caps", lm, channels)
			}
		}
	}
	signed := []int16{-32768, -1, 0, 32767}
	m.FeBands = &signed[0]
	for i, want := range signed {
		if modeBand(m, int32(i)) != want {
			t.Fatal("signed band", i)
		}
	}
	runtime.KeepAlive(m)
}

func TestBandContextModePointers(t *testing.T) {
	makeContext := func() *band_ctx {
		m := mode48000_960_120
		bands := slices.Clone(eband5ms[:])
		log := slices.Clone(logN400[:])
		index := slices.Clone(cache_index50[:])
		bits := slices.Clone(cache_bits50[:])
		m.FeBands = &bands[0]
		m.FlogN = &log[0]
		m.Fcache.Findex = &index[0]
		m.Fcache.Fbits = &bits[0]
		return &band_ctx{Fm: &m, Fi: 3}
	}
	ctx := makeContext()
	entropyInitGrowStack(12)
	runtime.GC()
	if modeBand(ctx.Fm, ctx.Fi) != eband5ms[3] || modeLogN(ctx.Fm, ctx.Fi) != logN400[3] {
		t.Fatal("context-owned mode tables")
	}
	for i, offset := range cache_index50 {
		if offset < 0 {
			continue
		}
		if cache := modePulseCache(ctx.Fm, int32(i)); *cache != cache_bits50[offset] {
			t.Fatal("context-owned pulse cache", i)
		}
	}
	saved := *ctx
	ctx.Fm = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if saved.Fm == nil || modeBand(saved.Fm, 21) != eband5ms[21] || modeLogN(saved.Fm, 20) != logN400[20] {
		t.Fatal("copied context ownership")
	}
	*ctx = saved
	runtime.GC()
	if ctx.Fm != saved.Fm {
		t.Fatal("context restore")
	}
	runtime.KeepAlive(ctx)
}

func TestBandContextEnergyPointers(t *testing.T) {
	makeContext := func() *band_ctx {
		energy := []float32{.25, 1, 0, 2, .5, 3}
		mode := OpusT_OpusCustomMode{FnbEBands: 3}
		return &band_ctx{Fm: &mode, FbandE: &energy[0], Fi: 1}
	}
	ctx := makeContext()
	saved := *ctx
	ctx.FbandE = nil
	entropyInitGrowStack(12)
	runtime.GC()
	*ctx = saved
	want := [6]float32{.25, 1, 0, 2, .5, 3}
	for i, w := range want {
		if bandContextEnergy(ctx, int32(i)) != w {
			t.Fatal("owned energy", i)
		}
	}
	x, y := [2]float32{.3, -.4}, [2]float32{-.2, .1}
	rx, ry := x, y
	intensity_stereo(nil, ctx.Fm, &x[0], &y[0], ctx.FbandE, ctx.Fi, 2)
	intensity_stereo(nil, ctx.Fm, &rx[0], &ry[0], &want[0], ctx.Fi, 2)
	if x != rx || y != ry {
		t.Fatal("owned intensity energy")
	}
	runtime.KeepAlive(ctx)
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
