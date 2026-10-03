package opuscc

import (
	"math"
	"runtime"
	"testing"
	"unsafe"
)

func TestCeltPLCExcitationStoragePointers(t *testing.T) {
	for _, period := range []int32{512, 1024} {
		storage := celtPLCExcitationStorage(period)
		if len(storage) != int(period+24) {
			t.Fatal("excitation geometry")
		}
		history := make([]float32, 2048)
		for i := range history {
			history[i] = float32(math.Sin(float64(i) * .17))
		}
		celtPLCExcitationHistory(unsafe.SliceData(storage), &history[0], 2048, period)
		exc := storage[24:]
		storage = nil
		mode := newSynthesisTestMode()
		var ac [25]float32
		var coefficients [24]float32
		entropyInitGrowStack(12)
		runtime.GC()
		Opus__celt_autocorr(nil, unsafe.SliceData(exc), &ac[0], mode.Fwindow, 120, 24, period, 0)
		celtPLCLagWindow(&ac)
		Opus__celt_lpc(nil, &coefficients[0], &ac[0], 24)
		filtered := make([]float32, period)
		Opus_celt_fir_c(nil, &exc[0], &coefficients[0], &filtered[0], period, 24, 0)
		copy(exc, filtered)
		runtime.GC()
		decay := celtPLCExcitationDecay(&exc[0], period, period)
		if math.IsNaN(float64(decay)) || decay < 0 || decay > 1 {
			t.Fatal("owned excitation pipeline", period, decay)
		}
	}
	storage := celtPLCExcitationStorage(0)
	if len(storage) != 24 || len(storage[24:]) != 0 {
		t.Fatal("prefix-only excitation")
	}
}
func TestCeltPLCNoiseStoragePointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			mode := newSynthesisTestMode()
			N := mode.FshortMdctSize << LM
			storage := celtPLCNoiseStorage(N, channels)
			if len(storage) != int(N*channels) {
				t.Fatal("noise storage geometry")
			}
			state := &OpusT_OpusCustomDecoder{Fmode: mode, Frng: 0xdeadbeef}
			celtPLCNoise(nil, state, mode.FeBands, unsafe.SliceData(storage), N, 0, mode.FeffEBands, LM, channels)
			var outputs [2]*float32
			buffers := make([][]float32, channels)
			for c := range buffers {
				buffers[c] = make([]float32, N+mode.Foverlap+2)
				buffers[c][0], buffers[c][len(buffers[c])-1] = 77, 88
				outputs[c] = &buffers[c][1]
			}
			energy := make([]float32, 2*mode.FnbEBands)
			for i := range energy {
				energy[i] = -12
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celt_synthesis(nil, mode, unsafe.SliceData(storage), &outputs[0], &energy[0], 0, mode.FeffEBands, channels, channels, 0, LM, 1, 0, 0)
			for _, buffer := range buffers {
				if buffer[0] != 77 || buffer[len(buffer)-1] != 88 {
					t.Fatal("owned noise synthesis guards")
				}
				for _, sample := range buffer[1 : len(buffer)-1] {
					if math.IsNaN(float64(sample)) || math.IsInf(float64(sample), 0) {
						t.Fatal("owned noise synthesis")
					}
				}
			}
		}
	}
	if len(celtPLCNoiseStorage(0, 2)) != 0 {
		t.Fatal("empty noise storage")
	}
}
func TestCeltPLCNoisePointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, seed := range []uint32{0, 1, 0xffffffff, 0xdeadbeef} {
				bands := []int16{0, 1, 3, 6, 12, 20}
				N := int32(20) << LM
				a := make([]float32, N*channels+2)
				a[0], a[len(a)-1] = 77, 88
				want := append([]float32(nil), a...)
				state := &OpusT_OpusCustomDecoder{Frng: seed, Fmode: newSynthesisTestMode()}
				s := seed
				for c := int32(0); c < channels; c++ {
					for i := int32(1); i < 5; i++ {
						offset := 1 + N*c + int32(bands[i])<<LM
						length := int32(bands[i+1]-bands[i]) << LM
						for j := int32(0); j < length; j++ {
							s = 1664525*s + 1013904223
							want[offset+j] = float32(int32(s) >> 20)
						}
						Opus_renormalise_vector(nil, &want[offset], length, 1, 0)
					}
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celtPLCNoise(nil, state, &bands[0], &a[1], N, 1, 5, LM, channels)
				if state.Frng != s {
					t.Fatal("noise RNG", channels, LM, seed)
				}
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
						t.Fatal("noise spectrum", channels, LM, seed, i)
					}
				}
			}
		}
	}
	state := &OpusT_OpusCustomDecoder{Frng: 123}
	celtPLCNoise(nil, state, nil, nil, 0, 0, 0, 0, 2)
	celtPLCNoise(nil, state, nil, nil, 0, 0, 1, 0, 0)
	if state.Frng != 123 {
		t.Fatal("unused noise")
	}
}
func TestCeltPLCFinishPointers(t *testing.T) {
	for _, loss := range []int32{-10, 0, 1, 9999, 10000} {
		for _, plc := range []int32{-10, 0, 1, 9999, 10000} {
			for LM := int32(0); LM <= 3; LM++ {
				for _, frameType := range []int32{FRAME_PLC_PERIODIC, FRAME_PLC_NOISE, FRAME_PLC_NEURAL, FRAME_DRED} {
					state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Floss_duration: 123, Fplc_duration: plc, Fprefilter_and_fold: 1, Fskip_plc: 1, Frng: 0xdeadbeef}
					before := *state
					want := before
					want.Floss_duration = min(int32(10000), loss+int32(1)<<LM)
					want.Fplc_duration = min(int32(10000), plc+int32(1)<<LM)
					want.Flast_frame_type = frameType
					entropyInitGrowStack(12)
					runtime.GC()
					celtPLCFinish(state, loss, LM, frameType)
					if *state != want || state.Fmode.FeBands == nil {
						t.Fatal("finish state owners/order", loss, plc, LM, frameType)
					}
				}
			}
		}
	}
	// Signed-overflow and out-of-C-domain shift cases are Go-only: retain the
	// translated int32 wrapping rather than widening/saturating before addition.
	state := &OpusT_OpusCustomDecoder{Fplc_duration: math.MaxInt32}
	celtPLCFinish(state, math.MaxInt32, 0, FRAME_PLC_NOISE)
	if state.Floss_duration != math.MinInt32 || state.Fplc_duration != math.MinInt32 {
		t.Fatal("finish wrapping", state.Floss_duration, state.Fplc_duration)
	}
	celtPLCFinish(state, 9, 32, FRAME_PLC_PERIODIC)
	if state.Floss_duration != 9 || state.Fplc_duration != math.MinInt32 {
		t.Fatal("finish large Go shift")
	}
}
func TestCeltPLCLPCHistoryPointers(t *testing.T) {
	for _, N := range []int32{0, 120, 240, 960} {
		h := make([]float32, 2048)
		for i := range h {
			h[i] = math.Float32frombits(uint32(i) * uint32(7919))
		}
		storage := new([26]float32)
		storage[0], storage[25] = 77, 88
		memory := (*[24]float32)(unsafe.Pointer(&storage[1]))
		entropyInitGrowStack(12)
		runtime.GC()
		celtPLCLPCHistory(memory, &h[0], 2048, N)
		if storage[0] != 77 || storage[25] != 88 {
			t.Fatal("LPC history guards")
		}
		for i := range memory {
			if math.Float32bits(memory[i]) != math.Float32bits(h[2048-int(N)-1-i]) {
				t.Fatal("reverse LPC history", N, i)
			}
		}
	}
	// Go-only scratch/history overlap must retain store/load order.
	a := make([]float32, 48)
	for i := range a {
		a[i] = float32(i + 1)
	}
	want := append([]float32(nil), a...)
	for i := 0; i < 24; i++ {
		want[8+i] = want[31-i]
	}
	memory := (*[24]float32)(unsafe.Pointer(&a[8]))
	celtPLCLPCHistory(memory, &a[0], 32, 0)
	for i := range a {
		if a[i] != want[i] {
			t.Fatal("reverse LPC live alias", i)
		}
	}
}
func TestCeltPLCExtrapolatePointers(t *testing.T) {
	for _, pitch := range []int32{40, 100, 511, 1024} {
		for _, N := range []int32{120, 240, 960} {
			a := make([]float32, 2170)
			x := make([]float32, 1024)
			a[0], a[len(a)-1] = 77, 88
			for i := 1; i < len(a)-1; i++ {
				a[i] = float32((i*37)%79-39) / 13
			}
			for i := range x {
				x[i] = float32((i*43)%89-44) / 17
			}
			want := append([]float32(nil), a...)
			attenuation := float32(float32(.8) * float32(.923))
			energy := float32(0)
			j := int32(0)
			for i := int32(0); i < N+120; i++ {
				if j >= pitch {
					j -= pitch
					attenuation = float32(attenuation * float32(.923))
				}
				want[1+2048-N+i] = float32(attenuation * x[1024-pitch+j])
				sample := want[1+2048-N-pitch+j]
				energy += float32(sample * sample)
				j++
			}
			entropyInitGrowStack(12)
			runtime.GC()
			got := celtPLCExtrapolate(&a[1], &x[0], 2048, 1024, N, 120, pitch, .8, .923)
			if math.Float32bits(got) != math.Float32bits(energy) {
				t.Fatal("extrapolation energy", pitch, N, got, energy)
			}
			for i := range a {
				if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
					t.Fatal("extrapolation samples", pitch, N, i)
				}
			}
		}
	}
	if celtPLCExtrapolate(nil, nil, 0, 0, 0, 0, 0, 1, 1) != 0 {
		t.Fatal("unused extrapolation")
	}
	// Go-only scratch/history alias: excitation reads remain live after stores.
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	want := append([]float32(nil), a...)
	got := celtPLCExtrapolate(&a[0], &a[5], 10, 3, 4, 2, 3, 1, .5)
	// Build the exact sequential reference, including period attenuation.
	energy := float32(0)
	atten := float32(.5)
	for i := 0; i < 6; i++ {
		if i == 3 {
			atten = float32(atten * .5)
		}
		j := i % 3
		want[6+i] = float32(atten * want[5+j])
		sample := want[3+j]
		energy += float32(sample * sample)
	}
	if got != energy {
		t.Fatal("aliased energy")
	}
	for i := range a {
		if a[i] != want[i] {
			t.Fatal("live excitation alias", i)
		}
	}
}
func TestCeltPLCExcitationHistoryPointers(t *testing.T) {
	for _, period := range []int32{0, 1, 512, 1024} {
		h := make([]float32, 2050)
		x := make([]float32, period+26)
		h[0], h[2049], x[0], x[len(x)-1] = 77, 88, 99, 66
		for i := 1; i < 2049; i++ {
			h[i] = math.Float32frombits(uint32(i) * 7919)
		}
		before := append([]float32(nil), h...)
		entropyInitGrowStack(12)
		runtime.GC()
		celtPLCExcitationHistory(&x[1], &h[1], 2048, period)
		if x[0] != 99 || x[len(x)-1] != 66 {
			t.Fatal("history destination guards")
		}
		for i := int32(0); i < period+24; i++ {
			if math.Float32bits(x[i+1]) != math.Float32bits(h[1+2048-period-24+i]) {
				t.Fatal("history samples", period, i)
			}
		}
		for i := range h {
			if math.Float32bits(h[i]) != math.Float32bits(before[i]) {
				t.Fatal("source history changed")
			}
		}
	}
	// Go-only outer-scratch/history aliases: preserve forward stores, not copy's
	// memmove semantics for a destination one sample after the source.
	for _, offset := range []int{0, 1, 3} {
		a := make([]float32, 30)
		for i := range a {
			a[i] = float32(i + 1)
		}
		want := append([]float32(nil), a...)
		for i := 0; i < 24; i++ {
			want[offset+i] = want[1+i]
		}
		celtPLCExcitationHistory(&a[offset], &a[0], 25, 0)
		for i := range a {
			if a[i] != want[i] {
				t.Fatal("forward history alias", offset, i)
			}
		}
	}
}
func TestCeltPLCSynthesisAttenuatePointers(t *testing.T) {
	for _, length := range []int32{1, 120, 240, 1080} {
		for _, overlap := range []int32{0, 1, min(120, length)} {
			for _, factor := range []float32{0, .1, .2, .21, .5, 1, 2, float32(math.NaN())} {
				a := make([]float32, length+2)
				a[0], a[len(a)-1] = 77, 88
				w := make([]float32, overlap)
				for i := int32(0); i < length; i++ {
					a[i+1] = float32((i*37)%79-39) / 13
				}
				for i := range w {
					w[i] = float32(i+1) / float32(len(w)+1)
				}
				want := append([]float32(nil), a...)
				s2 := float32(0)
				for _, v := range a[1 : len(a)-1] {
					s2 += float32(v * v)
				}
				s1 := float32(factor * s2)
				if !(s1 > float32(.2*s2)) {
					clear(want[1 : len(want)-1])
				} else if s1 < s2 {
					ratio := float32(math.Sqrt(float64((s1 + 1) / (s2 + 1))))
					for i := int32(0); i < overlap; i++ {
						g := float32(1) - float32(w[i]*(float32(1)-ratio))
						want[i+1] = float32(g * want[i+1])
					}
					for i := overlap; i < length; i++ {
						want[i+1] = float32(ratio * want[i+1])
					}
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celtPLCSynthesisAttenuate(&a[1], unsafe.SliceData(w), length, overlap, s1)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
						t.Fatal("synthesis attenuation", length, overlap, factor, i)
					}
				}
			}
		}
	}
	// Go-only mode-window/output alias: reload each window value after prior
	// output stores rather than snapshotting the window.
	aAlias := []float32{.5, .75, 1, 1.25, 1.5}
	wantAlias := append([]float32(nil), aAlias...)
	s2Alias := float32(0)
	for _, v := range aAlias[1:] {
		s2Alias += float32(v * v)
	}
	s1Alias := float32(.5 * s2Alias)
	ratioAlias := float32(math.Sqrt(float64((s1Alias + 1) / (s2Alias + 1))))
	for i := 0; i < 4; i++ {
		g := float32(1) - float32(wantAlias[i]*(float32(1)-ratioAlias))
		wantAlias[i+1] = float32(g * wantAlias[i+1])
	}
	celtPLCSynthesisAttenuate(&aAlias[1], &aAlias[0], 4, 4, s1Alias)
	for i := range aAlias {
		if math.Float32bits(aAlias[i]) != math.Float32bits(wantAlias[i]) {
			t.Fatal("live attenuation window alias")
		}
	}
	celtPLCSynthesisAttenuate(nil, nil, 0, 0, 0)
	a := []float32{float32(math.NaN()), 1}
	celtPLCSynthesisAttenuate(&a[0], nil, 2, 0, 100)
	if a[0] != 0 || a[1] != 0 {
		t.Fatal("NaN explosion not cleared")
	}
}
func TestCeltPLCExcitationDecayPointers(t *testing.T) {
	for _, length := range []int32{0, 1, 2, 31, 120, 512, 1024} {
		a := make([]float32, 1026)
		a[0], a[1025] = 77, 88
		for i := 1; i < 1025; i++ {
			a[i] = float32((i*7919)%65536-32768) / 37
		}
		before := append([]float32(nil), a...)
		e1, e2 := float32(1), float32(1)
		half := length >> 1
		for i := int32(0); i < half; i++ {
			e := a[1+1024-half+i]
			e1 += float32(e * e)
			e = a[1+1024-2*half+i]
			e2 += float32(e * e)
		}
		if !(e1 < e2) {
			e1 = e2
		}
		want := float32(math.Sqrt(float64(e1 / e2)))
		entropyInitGrowStack(12)
		runtime.GC()
		got := celtPLCExcitationDecay(&a[1], 1024, length)
		if math.Float32bits(got) != math.Float32bits(want) {
			t.Fatal("excitation decay", length, got, want)
		}
		for i := range a {
			if a[i] != before[i] {
				t.Fatal("excitation changed")
			}
		}
	}
	recentNaN := []float32{1, 1, float32(math.NaN()), 1}
	if celtPLCExcitationDecay(&recentNaN[0], 4, 4) != 1 {
		t.Fatal("MIN32 recent NaN selection")
	}
	olderNaN := []float32{float32(math.NaN()), 1, 1, 1}
	if !math.IsNaN(float64(celtPLCExcitationDecay(&olderNaN[0], 4, 4))) {
		t.Fatal("MIN32 older NaN selection")
	}
	if celtPLCExcitationDecay(nil, 0, 0) != 1 {
		t.Fatal("unused excitation")
	}
	a := []float32{100, 100, 1, 1}
	if !(celtPLCExcitationDecay(&a[0], 4, 4) < 1) {
		t.Fatal("decaying waveform")
	}
	a = []float32{1, 1, 100, 100}
	if celtPLCExcitationDecay(&a[0], 4, 4) != 1 {
		t.Fatal("no amplification")
	}
}
func TestCeltPLCLagWindowPointers(t *testing.T) {
	for trial := 0; trial < 80; trial++ {
		storage := new([27]float32)
		storage[0], storage[26] = 77, 88
		for i := 1; i < 26; i++ {
			storage[i] = float32((i*7919+trial*997)%65536-32768) / 37
		}
		ac := (*[25]float32)(unsafe.Pointer(&storage[1]))
		want := *ac
		want[0] *= 1.0001
		for i := int32(1); i <= 24; i++ {
			want[i] -= float32(float32(float32(want[i]*float32(float32(.008)*float32(.008)))*float32(i)) * float32(i))
		}
		entropyInitGrowStack(12)
		runtime.GC()
		celtPLCLagWindow(ac)
		for i := range ac {
			if math.Float32bits(ac[i]) != math.Float32bits(want[i]) {
				t.Fatal("rounded lag window", trial, i)
			}
		}
		if storage[0] != 77 || storage[26] != 88 {
			t.Fatal("lag window guards")
		}
	}
}
func TestCeltPLCDecayPointers(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, loss := range []int32{0, 1, 99} {
			for _, alias := range []bool{false, true} {
				a := make([]float32, 44)
				b := make([]float32, 44)
				for i := range a {
					a[i] = float32(i) - 22
					b[i] = float32(i%5) - 12
				}
				a[0], a[43] = 77, 88
				if alias {
					b = a
				}
				want := append([]float32(nil), a...)
				wb := append([]float32(nil), b...)
				if alias {
					wb = want
				}
				for c := int32(0); c < max(channels, 1); c++ {
					for i := int32(1); i < 20; i++ {
						index := 1 + c*21 + i
						decay := float32(.5)
						if loss == 0 {
							decay = 1.5
						}
						if wb[index] > want[index]-decay {
							want[index] = wb[index]
						} else {
							want[index] -= decay
						}
					}
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celtPLCDecay(&a[1], &b[1], 21, 1, 20, channels, loss)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
						t.Fatal("decay order", channels, loss, alias, i)
					}
				}
			}
		}
	}
	celtPLCDecay(nil, nil, 21, 5, 5, 2, 0)
	a := []float32{float32(math.NaN()), 3}
	b := []float32{2, float32(math.NaN())}
	celtPLCDecay(&a[0], &b[0], 2, 0, 2, 1, 0)
	if !math.IsNaN(float64(a[0])) || a[1] != 1.5 {
		t.Fatal("MAXG NaN selection", a)
	}
}
func TestEntropyWritePointers(t *testing.T) {
	buffer := [3]byte{}
	enc := OpusT_ec_enc{
		Fbuf:     &buffer[0],
		Fstorage: uint32(len(buffer)),
	}
	if got := ec_write_byte(nil, &enc, 0x123); got != 0 {
		t.Fatalf("front write: %d", got)
	}
	if got := ec_write_byte_at_end(nil, &enc, 0x245); got != 0 {
		t.Fatalf("back write: %d", got)
	}
	if got := ec_write_byte(nil, &enc, 0x67); got != 0 {
		t.Fatalf("last free byte: %d", got)
	}
	if want := [3]byte{0x23, 0x67, 0x45}; buffer != want {
		t.Fatalf("buffer: got %x, want %x", buffer, want)
	}
	before := enc
	if got := ec_write_byte(nil, &enc, 0xff); got != -1 {
		t.Fatalf("full front write: %d", got)
	}
	if got := ec_write_byte_at_end(nil, &enc, 0xff); got != -1 {
		t.Fatalf("full back write: %d", got)
	}
	if enc != before || buffer != [3]byte{0x23, 0x67, 0x45} {
		t.Fatal("failed write modified state or buffer")
	}
	empty := OpusT_ec_enc{}
	if ec_write_byte(nil, &empty, 1) != -1 || ec_write_byte_at_end(nil, &empty, 1) != -1 || empty != (OpusT_ec_enc{}) {
		t.Fatal("zero-capacity writer should fail without accessing a buffer")
	}
}

func TestStereoSplitPointers(t *testing.T) {
	// Interior pointers and sentinels check both ends of the requested span.
	x := [...]float32{99, 1, -2, 0, 0.125, 88}
	y := [...]float32{77, -1, 3, 0, 0.25, 66}
	wantX, wantY := x, y
	for i := 1; i < len(x)-1; i++ {
		l := float32(float32(0.70710678) * x[i])
		r := float32(float32(0.70710678) * y[i])
		wantX[i], wantY[i] = l+r, r-l
	}
	stereo_split(nil, &x[1], &y[1], 4)
	if x != wantX || y != wantY {
		t.Fatalf("got X=%v Y=%v, want X=%v Y=%v", x, y, wantX, wantY)
	}
	stereo_split(nil, nil, nil, 0)
}
