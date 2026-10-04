package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"math"
	"math/bits"
	"runtime"
	"testing"
	"unsafe"
)

func TestCeltPLCHistoryViewsPointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for _, overlap := range []int32{0, 120} {
			for _, N := range []int32{0, 120, 240, 480, 960} {
				owner, _, _ := celtStateTestBuffer(newSynthesisTestMode(), channels)
				h, out, e, b, a := celtPLCHistoryViews(&owner.State, overlap, 21, channels, N)
				owner = nil
				entropyInitGrowStack(12)
				runtime.GC()
				stride := int32(2048) + overlap
				for c := int32(0); c < channels; c++ {
					if len(h[c]) != int(stride) {
						t.Fatal("history span")
					}
					h[c][0] = float32(c + 77)
					if N == 0 && overlap == 0 {
						if out[c] != nil {
							t.Fatal("unused one-past output")
						}
					} else {
						if out[c] != &h[c][2048-N] {
							t.Fatal("output view")
						}
						*out[c] = 99
					}
				}
				if channels == 1 && (h[1] != nil || out[1] != nil) {
					t.Fatal("unused mono channel")
				}
				if uintptr(unsafe.Pointer(b))-uintptr(unsafe.Pointer(e)) != 6*21*4 || uintptr(unsafe.Pointer(a))-uintptr(unsafe.Pointer(e)) != 8*21*4 {
					t.Fatal("energy/LPC offsets")
				}
				*e, *b, *a = -12, -28, .125
			}
		}
	}
}
func TestCeltPLCModePointers(t *testing.T) {
	state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode()}
	m, nb, overlap, bands := celtPLCMode(state)
	state = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if nb != 21 || overlap != 120 || m.FeBands != bands || unsafe.Slice(bands, nb+1)[nb] != 100 || m.Fmdct.Fkfft[0] == nil {
		t.Fatal("typed mode/table owners", nb, overlap)
	}
}
func TestCeltDecodeModePointers(t *testing.T) {
	state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode()}
	want := state.Fmode
	mode, bands, overlap, boundaries := celtDecodeMode(state)
	state.Fmode = nil
	state = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if mode != want || bands != 21 || overlap != 120 || boundaries != mode.FeBands || unsafe.Slice(boundaries, bands+1)[bands] != 100 {
		t.Fatal("retained decode mode owner")
	}
}
func TestCeltDecodeAntiCollapsePointers(t *testing.T) {
	celtDecodeAntiCollapse(nil, nil, nil, nil, nil, nil, nil, nil, nil, 120, 0, 1, 0, 21, 0)
	if celtDecodeAntiCollapseBit(nil, nil, 0) != 0 || celtDecodeAntiCollapseBit(nil, nil, -1) != 0 {
		t.Fatal("unused anti-collapse entropy")
	}
	for _, reserved := range []int32{-1, 0, 8} {
		data := []byte{0, 71, 255, 13}
		var ec OpusT_ec_ctx
		Opus_ec_dec_init(nil, &ec, &data[0], 4)
		ref := ec
		want := int32(0)
		if reserved > 0 {
			want = int32(Opus_ec_dec_bits(nil, &ref, 1))
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := celtDecodeAntiCollapseBit(nil, &ec, reserved)
		if got != want || ec != ref {
			t.Fatal("anti-collapse reservation bit")
		}
	}
}
func TestCeltDecodeFinalEnergyPointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for _, length := range []int32{0, 1, 2, 16} {
			mode := newSynthesisTestMode()
			e := make([]float32, 44)
			for i := range e {
				e[i] = -12
			}
			e[0] = 901
			e[43] = 902
			fine, priority := celtDecodeFineStorage(21), celtDecodePriorityStorage(21)
			for i := range fine {
				fine[i] = int32(i % 9)
				priority[i] = int32(i % 2)
			}
			want := append([]float32(nil), e...)
			data := []byte{0, 71, 255, 13, 40}
			var ec OpusT_ec_ctx
			Opus_ec_dec_init(nil, &ec, &data[0], 5)
			ref := ec
			tell := ref.Fnbits_total - int32(bits.Len32(ref.Frng))
			Opus_unquant_energy_finalise(nil, mode, 0, 21, &want[1], &fine[0], &priority[0], length*8-tell, &ref, channels)
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodeFinalEnergy(nil, mode, &e[1], &fine[0], &priority[0], 0, 21, length, channels, &ec)
			if ec != ref {
				t.Fatal("final energy entropy")
			}
			for i := range e {
				if math.Float32bits(e[i]) != math.Float32bits(want[i]) {
					t.Fatal("final energy forwarding", channels, length, i)
				}
			}
		}
	}
}
func TestCeltDecodeAllocationBudgetPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, tr := range []int32{0, 1, -1} {
			for _, length := range []int32{0, 1, 2, 16, 1275} {
				data := []byte{0, 71, 255, 13}
				var ec OpusT_ec_ctx
				Opus_ec_dec_init(nil, &ec, &data[0], 4)
				before := ec
				want := (length*8)<<BITRES - int32(Opus_ec_tell_frac(nil, &ec)) - 1
				reserve := int32(0)
				if tr != 0 && LM >= 2 && want >= (LM+2)<<BITRES {
					reserve = 1 << BITRES
				}
				entropyInitGrowStack(12)
				runtime.GC()
				budget, r := celtDecodeAllocationBudget(nil, &ec, length, tr, LM)
				if budget != want-reserve || r != reserve || ec != before {
					t.Fatal("allocation reservation", LM, tr, length)
				}
			}
		}
	}
}
func TestCeltDecodeTrimPointers(t *testing.T) {
	for _, tell := range []int32{0, 8, 49} {
		for _, total := range []int32{0, 47, 48, 49, 128} {
			for _, pattern := range []byte{0, 71, 255} {
				data := make([]byte, 16)
				for i := range data {
					data[i] = pattern
				}
				var ec OpusT_ec_ctx
				Opus_ec_dec_init(nil, &ec, &data[0], 16)
				ref := ec
				want := int32(5)
				if tell+48 <= total {
					want = Opus_ec_dec_icdf(nil, &ref, &trim_icdf9[0], 7)
				}
				entropyInitGrowStack(12)
				runtime.GC()
				got := celtDecodeTrim(nil, &ec, tell, total)
				if got != want || ec != ref {
					t.Fatal("trim header", tell, total, pattern)
				}
			}
		}
	}
	if celtDecodeTrim(nil, nil, 1, 0) != 5 {
		t.Fatal("unused trim entropy")
	}
}
func TestCeltDecodeSpreadPointers(t *testing.T) {
	for _, pattern := range []byte{0, 71, 255} {
		for _, total := range []int32{0, 4, 5, 8, 128} {
			data := make([]byte, 16)
			for i := range data {
				data[i] = pattern
			}
			var ec OpusT_ec_ctx
			Opus_ec_dec_init(nil, &ec, &data[0], 16)
			ref := ec
			wantTell := ref.Fnbits_total - int32(bits.Len32(ref.Frng))
			want := int32(SPREAD_NORMAL)
			if wantTell+4 <= total {
				want = Opus_ec_dec_icdf(nil, &ref, &spread_icdf9[0], 5)
			}
			entropyInitGrowStack(12)
			runtime.GC()
			spread, tell := celtDecodeSpread(nil, &ec, total)
			if ec != ref || spread != want || tell != wantTell {
				t.Fatal("spreading cached tell", pattern, total)
			}
		}
	}
}
func TestCeltDecodeGlobalFlagsPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, total := range []int32{0, 3, 4, 8, 128} {
			for _, pattern := range []byte{0, 71, 255} {
				data := make([]byte, 16)
				for i := range data {
					data[i] = pattern
				}
				var ec OpusT_ec_ctx
				Opus_ec_dec_init(nil, &ec, &data[0], 16)
				ref := ec
				wt, ws, wi, wtell := int32(0), int32(0), int32(0), int32(1)
				if LM > 0 && wtell+3 <= total {
					wt = Opus_ec_dec_bit_logp(nil, &ref, 3)
					wtell = ref.Fnbits_total - int32(bits.Len32(ref.Frng))
				}
				if wt != 0 {
					ws = 1 << LM
				}
				if wtell+3 <= total {
					wi = Opus_ec_dec_bit_logp(nil, &ref, 3)
				}
				entropyInitGrowStack(12)
				runtime.GC()
				tr, short, intra, tell := celtDecodeGlobalFlags(nil, &ec, LM, 1<<LM, total, 1)
				if ec != ref || tr != wt || short != ws || intra != wi || tell != wtell {
					t.Fatal("global flags cached tell", LM, total, pattern)
				}
			}
		}
	}
}
func TestCeltDecodePostfilterHeaderPointers(t *testing.T) {
	for _, pattern := range []byte{0, 71, 255} {
		for _, total := range []int32{0, 16, 17, 32, 128} {
			for _, start := range []int32{0, 1} {
				data := make([]byte, 16)
				for i := range data {
					data[i] = pattern
				}
				var ec OpusT_ec_ctx
				Opus_ec_dec_init(nil, &ec, &data[0], 16)
				before := ec
				entropyInitGrowStack(12)
				runtime.GC()
				p, g, tap, tell := celtDecodePostfilterHeader(nil, &ec, start, total, 1)
				if start != 0 || total < 17 {
					if ec != before || p != 0 || g != 0 || tap != 0 || tell != 1 {
						t.Fatal("postfilter header budget guard")
					}
				} else if p < 0 || p > 1022 || g < 0 || g > .75 || tap < 0 || tap > 2 || tell != ec.Fnbits_total-int32(bits.Len32(ec.Frng)) {
					t.Fatal("postfilter header values", p, g, tap, tell)
				}
			}
		}
	}
}
func TestCeltDecodeSilencePointers(t *testing.T) {
	for _, pattern := range []byte{0, 71, 255} {
		for _, total := range []int32{0, 1, 8, 16, 128} {
			data := make([]byte, 16)
			for i := range data {
				data[i] = pattern
			}
			var ec OpusT_ec_ctx
			Opus_ec_dec_init(nil, &ec, &data[0], 16)
			ref := ec
			wantTell := ref.Fnbits_total - int32(bits.Len32(ref.Frng))
			wantSilence := int32(0)
			if wantTell >= total {
				wantSilence = 1
			} else if wantTell == 1 {
				wantSilence = Opus_ec_dec_bit_logp(nil, &ref, 15)
			}
			if wantSilence != 0 {
				wantTell = total
				ref.Fnbits_total += total - (ref.Fnbits_total - int32(bits.Len32(ref.Frng)))
			}
			entropyInitGrowStack(12)
			runtime.GC()
			s, tell := celtDecodeSilence(nil, &ec, total)
			if s != wantSilence || tell != wantTell || ec != ref {
				t.Fatal("silence header", pattern, total)
			}
		}
	}
}
func TestCeltDecodePacketErrorPointers(t *testing.T) {
	for _, nbits := range []int32{0, 1, 7, 8, 9, 100} {
		for _, length := range []int32{0, 1, 8, 1275} {
			for _, flag := range []int32{0, 1, -1} {
				state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Ferror1: 7, Frng: 123}
				ec := &OpusT_ec_ctx{Fnbits_total: nbits, Frng: 1, Ferror1: flag}
				before := *ec
				entropyInitGrowStack(12)
				runtime.GC()
				r := celtDecodePacketError(state, ec, length)
				want := int32(0)
				errorFlag := int32(7)
				if nbits-1 > length*8 {
					want = -3
				} else if flag != 0 {
					errorFlag = 1
				}
				if r != want || state.Ferror1 != errorFlag || state.Frng != 123 || *ec != before {
					t.Fatal("packet error ordering", nbits, length, flag, r)
				}
			}
		}
	}
	// Overflow and zero-range fixtures describe Go semantics, not C signed UB.
	state := OpusT_OpusCustomDecoder{}
	ec := OpusT_ec_ctx{Fnbits_total: -2147483648, Frng: 1, Ferror1: 1}
	if celtDecodePacketError(&state, &ec, 1275) != -3 || state.Ferror1 != 0 {
		t.Fatal("wrapped tell error ordering")
	}
	ec.Fnbits_total = 0
	ec.Frng = 0
	if celtDecodePacketError(&state, &ec, 0) != 0 || state.Ferror1 != 1 {
		t.Fatal("zero-range tell")
	}
}
func TestCeltDecodeDeemphasisPointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for _, factor := range []int32{1, 2, 3, 6} {
			for _, accum := range []int32{0, 1} {
				mode := newSynthesisTestMode()
				state := &OpusT_OpusCustomDecoder{Fmode: mode, Fdownsample: factor, Fpreemph_memD: [2]float32{.1, -.2}, Frng: 123}
				left, right := make([]float32, 120), make([]float32, 120)
				for i := range left {
					left[i] = float32(i%13 - 6)
					right[i] = float32(i%17 - 8)
				}
				in := [2]*float32{&left[0], nil}
				if channels == 2 {
					in[1] = &right[0]
				}
				pcm := make([]float32, (120/factor)*channels+2)
				pcm[0] = 901
				pcm[len(pcm)-1] = 902
				want := append([]float32(nil), pcm...)
				memory := state.Fpreemph_memD
				deemphasis(nil, &in[0], &want[1], 120, channels, factor, &mode.Fpreemph[0], &memory[0], accum)
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodeDeemphasis(nil, state, mode, &in[0], &pcm[1], 120, channels, accum)
				if state.Fpreemph_memD != memory || state.Frng != 123 {
					t.Fatal("deemphasis forwarding state")
				}
				for i := range pcm {
					if math.Float32bits(pcm[i]) != math.Float32bits(want[i]) {
						t.Fatal("decode deemphasis", channels, factor, accum, i)
					}
				}
			}
		}
	}
}
func TestCeltDecodePrefilterPointers(t *testing.T) {
	celtDecodePrefilter(nil, &OpusT_OpusCustomDecoder{}, 120)
	for _, channels := range []int32{1, 2} {
		for _, N := range []int32{120, 960} {
			for _, flag := range []int32{0, 1, -1, 7} {
				storage, image, _ := celtStateTestBuffer(newSynthesisTestMode(), channels)
				st := &storage.State
				st.Fchannels = channels
				st.Foverlap = 120
				st.Fprefilter_and_fold = flag
				st.Fpostfilter_period_old = 31
				st.Fpostfilter_period = 128
				st.Fpostfilter_gain_old = .25
				st.Fpostfilter_gain = .5
				st.Fpostfilter_tapset_old = 1
				st.Fpostfilter_tapset = 2
				st.Farch = 0
				h := unsafe.Slice(&st.F_decode_mem[0], (DEC_PITCH_BUF_SIZE+120)*channels)
				for i := range h {
					h[i] = float32(i%29-14) * 173
				}
				reference := new(celtStateTestStorage)
				*reference = *storage
				if flag != 0 {
					prefilter_and_fold(nil, &reference.State, N)
				}
				want := unsafe.Slice((*byte)(unsafe.Pointer(&reference.State)), len(image))
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodePrefilter(nil, st, N)
				for i := range image {
					if image[i] != want[i] {
						t.Fatal("prefilter dispatch", channels, N, flag, i)
					}
				}
			}
		}
	}
}
func TestCeltDecodeEnergyMergeMonoPointers(t *testing.T) {
	celtDecodeEnergyMergeMono(nil, 0)
	values := []float32{-28, -1, 0, math.Float32frombits(0x80000000), 1, math.Float32frombits(0x7fc12345)}
	for _, a := range values {
		for _, b := range values {
			e := []float32{901, a, b, 902}
			want := b
			if a > b {
				want = a
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodeEnergyMergeMono(&e[1], 1)
			if math.Float32bits(e[1]) != math.Float32bits(want) || math.Float32bits(e[2]) != math.Float32bits(b) || e[0] != 901 || e[3] != 902 {
				t.Fatal("mono energy merge", a, b)
			}
		}
	}
}
func TestCeltDecodePostfilterPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{0, 1, 2} {
			N := int32(120) << LM
			mode := newSynthesisTestMode()
			state := &OpusT_OpusCustomDecoder{Fmode: mode, Fpostfilter_period: -1, Fpostfilter_period_old: 0, Fpostfilter_gain: .3, Fpostfilter_gain_old: .5, Fpostfilter_tapset: 2, Fpostfilter_tapset_old: 1, Frng: 123}
			wantState := *state
			left, right := make([]float32, 256+N+2), make([]float32, 256+N+2)
			for i := range left {
				left[i] = float32(i%13-6) / 128
				right[i] = float32(i%17-8) / 256
			}
			wl, wr := append([]float32(nil), left...), append([]float32(nil), right...)
			out := [2]*float32{&left[257], nil}
			if channels == 2 {
				out[1] = &right[257]
			}
			wo := [2]*float32{&wl[257], &wr[257]}
			for c := int32(0); c < max(int32(1), channels); c++ {
				celtDecodePostfilterClamp(&wantState)
				celtDecodePostfilterFirst(nil, &wantState, mode, wo[c], 120)
				if LM != 0 {
					celtDecodePostfilterTail(nil, &wantState, mode, wo[c], N, 80, .7, 0, 120)
				}
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodePostfilter(nil, state, mode, &out[0], channels, N, LM, 80, .7, 0, 120)
			if *state != wantState {
				t.Fatal("postfilter driver state")
			}
			for i := range left {
				if math.Float32bits(left[i]) != math.Float32bits(wl[i]) || math.Float32bits(right[i]) != math.Float32bits(wr[i]) {
					t.Fatal("postfilter driver", LM, channels, i)
				}
			}
		}
	}
}
func TestCeltDecodePostfilterTailPointers(t *testing.T) {
	for _, N := range []int32{240, 480, 960} {
		for _, gain := range []float32{0, .3, .7} {
			for tap := int32(0); tap < 3; tap++ {
				mode := newSynthesisTestMode()
				state := &OpusT_OpusCustomDecoder{Fmode: mode, Fpostfilter_period: 100, Fpostfilter_gain: .3, Fpostfilter_tapset: 2}
				before := *state
				h := make([]float32, 256+N+2)
				for i := range h {
					h[i] = float32(i%13-6) / 128
				}
				want := append([]float32(nil), h...)
				Opus_comb_filter(nil, &want[377], &want[377], 100, 80, N-120, .3, gain, 2, tap, mode.Fwindow, 120, 0)
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodePostfilterTail(nil, state, mode, &h[257], N, 80, gain, tap, 120)
				if *state != before {
					t.Fatal("tail postfilter state")
				}
				for i := range h {
					if math.Float32bits(h[i]) != math.Float32bits(want[i]) {
						t.Fatal("tail postfilter", N, gain, tap, i)
					}
				}
			}
		}
	}
}
func TestCeltDecodePostfilterFirstPointers(t *testing.T) {
	for _, gain := range []float32{0, .3, .7} {
		for tap := int32(0); tap < 3; tap++ {
			mode := newSynthesisTestMode()
			state := &OpusT_OpusCustomDecoder{Fmode: mode, Fpostfilter_period_old: 45, Fpostfilter_period: 100, Fpostfilter_gain_old: gain, Fpostfilter_gain: .3, Fpostfilter_tapset_old: tap, Fpostfilter_tapset: 2}
			before := *state
			h := make([]float32, 256+120+2)
			for i := range h {
				h[i] = float32(i%13-6) / 128
			}
			want := append([]float32(nil), h...)
			Opus_comb_filter(nil, &want[257], &want[257], 45, 100, 120, gain, .3, tap, 2, mode.Fwindow, 120, 0)
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodePostfilterFirst(nil, state, mode, &h[257], 120)
			if *state != before {
				t.Fatal("first postfilter state")
			}
			for i := range h {
				if math.Float32bits(h[i]) != math.Float32bits(want[i]) {
					t.Fatal("first postfilter", gain, tap, i)
				}
			}
		}
	}
}
func TestCeltDecodePostfilterClampPointers(t *testing.T) {
	for _, a := range []int32{-2147483648, -1, 0, 15, 16, 100, 2147483647} {
		for _, b := range []int32{-1, 0, 15, 16, 2147483647} {
			state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Fpostfilter_period: a, Fpostfilter_period_old: b, Frng: 123, Fpostfilter_gain: .5}
			want := *state
			if a < COMBFILTER_MINPERIOD {
				want.Fpostfilter_period = COMBFILTER_MINPERIOD
			}
			if b < COMBFILTER_MINPERIOD {
				want.Fpostfilter_period_old = COMBFILTER_MINPERIOD
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodePostfilterClamp(state)
			if *state != want {
				t.Fatal("postfilter clamp", a, b)
			}
		}
	}
}
func TestCeltDecodePacketFinishPointers(t *testing.T) {
	for _, v := range []int32{-2147483648, -1, 0, 1, 40, 2147483647} {
		state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Floss_duration: v, Fplc_duration: v, Flast_frame_type: v, Fprefilter_and_fold: v, Frng: 123, Ferror1: 7, Fpostfilter_period: 45}
		want := *state
		want.Floss_duration = 0
		want.Fplc_duration = 0
		want.Flast_frame_type = FRAME_NORMAL
		want.Fprefilter_and_fold = 0
		entropyInitGrowStack(12)
		runtime.GC()
		celtDecodePacketFinish(state)
		if *state != want {
			t.Fatal("packet finish changed unrelated state", v)
		}
	}
}
func TestCeltDecodeRecoverEnergyPointers(t *testing.T) {
	celtDecodeRecoverEnergy(nil, nil, nil, nil, 21, 0, 21, 0, 1)
	celtDecodeRecoverEnergy(&OpusT_OpusCustomDecoder{}, nil, nil, nil, 21, 0, 21, 0, 0)
	celtDecodeRecoverEnergy(&OpusT_OpusCustomDecoder{Floss_duration: 1}, nil, nil, nil, 21, 5, 5, 0, 0)
	for LM := int32(0); LM <= 3; LM++ {
		for _, loss := range []int32{1, 10, 40} {
			for _, start := range []int32{0, 2, 3} {
				e, l, p := make([]float32, 8), make([]float32, 8), make([]float32, 8)
				for i := range e {
					e[i] = float32(i) - 10
					l[i] = float32(i) - 8
					p[i] = float32(i) - 6
				}
				want := append([]float32(nil), e...)
				state := &OpusT_OpusCustomDecoder{Floss_duration: loss}
				m, s := celtDecodeRecoverySafety(state, LM)
				for c := 0; c < 2; c++ {
					for i := int(start); i < 3; i++ {
						idx := 1 + c*3 + i
						celtDecodeRecoveryBand(&want[idx], &l[idx], &p[idx], m, s)
					}
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodeRecoverEnergy(state, &e[1], &l[1], &p[1], 3, start, 3, LM, 0)
				for i := range e {
					if math.Float32bits(e[i]) != math.Float32bits(want[i]) {
						t.Fatal("recovery loop", LM, loss, start, i)
					}
				}
			}
		}
	}
	// Go-only float/count alias: the second channel must reload loss_duration.
	state := &OpusT_OpusCustomDecoder{Floss_duration: 1}
	energy := (*float32)(unsafe.Pointer(&state.Floss_duration))
	l, p := []float32{-1, 1}, []float32{-2, 2}
	celtDecodeRecoverEnergy(state, energy, &l[0], &p[0], 1, 0, 1, 0, 0)
	if math.Float32frombits(uint32(state.Floss_duration)) != -3.5 || math.Float32frombits(uint32(state.Fplc_duration)) != -1.5 {
		t.Fatal("live per-channel recovery controls")
	}
}
func TestCeltDecodeRecoveryBandPointers(t *testing.T) {
	for _, tc := range []struct {
		e, l, p float32
		m       int32
		s, w    float32
	}{{-10, -8, -6, 2, .5, -16.5}, {-1, -4, -5, 1, 1.5, -6.5}, {-19, 0, 1, 10, 1.5, -21.5}} {
		v := []float32{901, tc.e, tc.l, tc.p, 902}
		entropyInitGrowStack(12)
		runtime.GC()
		celtDecodeRecoveryBand(&v[1], &v[2], &v[3], tc.m, tc.s)
		if v[1] != tc.w || v[0] != 901 || v[4] != 902 || v[2] != tc.l || v[3] != tc.p {
			t.Fatal("energy recovery", tc, v)
		}
	}
	x := float32(4)
	celtDecodeRecoveryBand(&x, &x, &x, 3, .5)
	if x != 3.5 {
		t.Fatal("live energy aliases")
	}
	e, l, p := float32(0), float32(0), math.Float32frombits(0x80000000)
	celtDecodeRecoveryBand(&e, &l, &p, 0, 0)
	if math.Float32bits(e) != 0x80000000 {
		t.Fatal("signed-zero selection")
	}
}
func TestCeltDecodeRecoverySafetyPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, loss := range []int32{-2147483648, -1, 0, 1, 10, 11, 40, 2147483647} {
			state := &OpusT_OpusCustomDecoder{Floss_duration: loss, Frng: 123}
			entropyInitGrowStack(12)
			runtime.GC()
			missing, safety := celtDecodeRecoverySafety(state, LM)
			want := loss >> LM
			if want > 10 {
				want = 10
			}
			ws := float32(0)
			if LM == 0 {
				ws = 1.5
			} else if LM == 1 {
				ws = .5
			}
			if missing != want || math.Float32bits(safety) != math.Float32bits(ws) || state.Floss_duration != loss || state.Frng != 123 {
				t.Fatal("recovery safety", LM, loss)
			}
		}
	}
}
func TestCeltDecodePostfilterFinishPointers(t *testing.T) {
	for _, LM := range []int32{-1, 0, 1, 3} {
		for _, period := range []int32{-2147483648, -1, 0, 15, 2147483647} {
			for _, bits := range []uint32{0, 0x80000000, 0x3f400000, 0x7fc12345, 0x7f800000} {
				state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Fpostfilter_period: 45, Fpostfilter_period_old: 99, Fpostfilter_gain: math.Float32frombits(0x80000000), Fpostfilter_gain_old: 1, Fpostfilter_tapset: 2, Fpostfilter_tapset_old: 1, Frng: 123, Floss_duration: 7}
				owner := state.Fmode
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodePostfilterFinish(state, period, math.Float32frombits(bits), -3, LM)
				oldPeriod, oldBits, oldTap := int32(45), uint32(0x80000000), int32(2)
				if LM != 0 {
					oldPeriod, oldBits, oldTap = period, bits, -3
				}
				if state.Fpostfilter_period != period || math.Float32bits(state.Fpostfilter_gain) != bits || state.Fpostfilter_tapset != -3 || state.Fpostfilter_period_old != oldPeriod || math.Float32bits(state.Fpostfilter_gain_old) != oldBits || state.Fpostfilter_tapset_old != oldTap || state.Fmode != owner || state.Frng != 123 || state.Floss_duration != 7 {
					t.Fatal("postfilter finalization", LM, period, bits)
				}
			}
		}
	}
}
func TestCeltDecodeBoostsPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, C := range []int32{1, 2} {
			for _, budget := range []int32{0, 8, 100, 500, 992} {
				for _, start := range []int32{0, 5} {
					bands := []int16{0, 1, 3, 6, 10, 18}
					caps := []int32{0, 8, 40, 128, 500}
					out := []int32{901, -1, -2, -3, -4, -5, 902}
					data := make([]byte, 128)
					var ec OpusT_ec_ctx
					Opus_ec_dec_init(nil, &ec, &data[0], 128)
					before := ec
					entropyInitGrowStack(12)
					runtime.GC()
					remaining, tell := celtDecodeBoosts(nil, &bands[0], &caps[0], &out[1], start, 5, C, LM, budget, &ec)
					if out[0] != 901 || out[6] != 902 || remaining > budget<<BITRES || tell != int32(Opus_ec_tell_frac(nil, &ec)) {
						t.Fatal("boost scalar/guards", LM, C, budget, start)
					}
					if start == 5 && ec != before {
						t.Fatal("empty boosts entropy")
					}
					for i := int(start); i < 5; i++ {
						if out[i+1] < 0 {
							t.Fatal("negative boost")
						}
					}
				}
			}
		}
	}
	var ec OpusT_ec_ctx
	data := []byte{0}
	Opus_ec_dec_init(nil, &ec, &data[0], 1)
	remaining, _ := celtDecodeBoosts(nil, nil, nil, nil, 0, 0, 1, 0, 0x20000000, &ec)
	if remaining != 0 {
		t.Fatal("Go-only total shift wrapping")
	}
}
func TestCeltDecodeSilenceEnergyPointers(t *testing.T) {
	celtDecodeSilenceEnergy(nil, 0, 2)
	celtDecodeSilenceEnergy(nil, 25, 0)
	// The original Go count multiplication wraps; this is not a C-overflow oracle.
	celtDecodeSilenceEnergy(nil, 0x7fffffff, 2)
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, channels := range []int32{1, 2} {
			e := make([]float32, bands*channels+2)
			for i := range e {
				e[i] = math.Float32frombits(0x7fc12345)
			}
			e[0] = 901
			e[len(e)-1] = -902
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodeSilenceEnergy(&e[1], bands, channels)
			if e[0] != 901 || e[len(e)-1] != -902 {
				t.Fatal("silence energy guards")
			}
			for _, v := range e[1 : len(e)-1] {
				if math.Float32bits(v) != math.Float32bits(-28) {
					t.Fatal("silence energy fill")
				}
			}
		}
	}
}
func TestCeltDecodeHistoryMovePointers(t *testing.T) {
	celtDecodeHistoryMove(nil, 960, 0)
	for _, N := range []int32{0, 1, 120, 240, 960} {
		for _, length := range []int32{0, 1, 128, 2048} {
			h := make([]float32, N+length+2)
			for i := range h {
				h[i] = math.Float32frombits(uint32(i)*7717 + 0x80000000)
			}
			want := append([]float32(nil), h...)
			copy(want[1:1+length], want[1+N:1+N+length])
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodeHistoryMove(&h[1], N, length)
			for i := range h {
				if math.Float32bits(h[i]) != math.Float32bits(want[i]) {
					t.Fatal("history memmove", N, length, i)
				}
			}
		}
	}
}
func TestCeltDecodeMaskStoragePointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, channels := range []int32{0, 1, 2} {
			m := celtDecodeMaskStorage(bands, channels)
			entropyInitGrowStack(12)
			runtime.GC()
			if len(m) != int(bands*channels) {
				t.Fatal("collapse mask geometry")
			}
			for i, v := range m {
				if v != 0 {
					t.Fatal("collapse mask initialization")
				}
				m[i] = byte(i)
			}
			runtime.GC()
			for i, v := range m {
				if v != byte(i) {
					t.Fatal("collapse mask retention")
				}
			}
		}
	}
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			N := int32(120) << LM
			mode := newSynthesisTestMode()
			m := celtDecodeMaskStorage(21, channels)
			for i := range m {
				if i%2 != 0 {
					m[i] = byte((1 << (1 << LM)) - 1)
				}
			}
			saved := append([]byte(nil), m...)
			s := make([]float32, N*channels+2)
			s[0] = 901
			s[len(s)-1] = 902
			energy, previous, older := make([]float32, 42), make([]float32, 42), make([]float32, 42)
			for i := range energy {
				energy[i] = -12
				previous[i] = -10
				older[i] = -11
			}
			pulses := celtDecodePulseStorage(21)
			for i := range pulses {
				pulses[i] = int32(i)*32 + 8
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodeAntiCollapse(nil, &OpusT_OpusCustomDecoder{Frng: 123}, mode, &s[1], &m[0], &pulses[0], &energy[0], &previous[0], &older[0], N, LM, channels, 0, 21, 1)
			if s[0] != 901 || s[len(s)-1] != 902 {
				t.Fatal("owned mask anti-collapse guards")
			}
			for i := range m {
				if m[i] != saved[i] {
					t.Fatal("mask consumer wrote input")
				}
			}
		}
	}
}
func TestCeltDecodeSpectrumStoragePointers(t *testing.T) {
	for _, N := range []int32{0, 120, 240, 480, 960} {
		for _, channels := range []int32{1, 2} {
			s := celtDecodeSpectrumStorage(N, channels)
			entropyInitGrowStack(12)
			runtime.GC()
			x, y := celtDecodeSpectrumChannels(s, N, channels)
			if len(s) != int(N*channels) {
				t.Fatal("spectral geometry")
			}
			if N == 0 {
				if x != nil || y != nil {
					t.Fatal("empty spectral views")
				}
				continue
			}
			if x != &s[0] || (channels == 1 && y != nil) || (channels == 2 && y != &s[N]) {
				t.Fatal("spectral channel views")
			}
			for i, v := range s {
				if v != 0 {
					t.Fatal("spectral initialization")
				}
				s[i] = float32(i%17-8) / 128
			}
			runtime.GC()
			if *x != s[0] || (y != nil && *y != s[N]) {
				t.Fatal("spectral view retention")
			}
		}
	}
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			N := int32(120) << LM
			s := celtDecodeSpectrumStorage(N, channels)
			for i := range s {
				s[i] = float32(i%17-8) / 128
			}
			energy := make([]float32, 42)
			for i := range energy {
				energy[i] = -12
			}
			left, right := make([]float32, N+122), make([]float32, N+122)
			left[0] = 901
			left[len(left)-1] = 902
			right[0] = 903
			right[len(right)-1] = 904
			out := [2]*float32{&left[1], nil}
			if channels == 2 {
				out[1] = &right[1]
			}
			mode := newSynthesisTestMode()
			entropyInitGrowStack(12)
			runtime.GC()
			celt_synthesis(nil, mode, &s[0], &out[0], &energy[0], 0, 21, channels, channels, 0, LM, 1, 0, 0)
			if left[0] != 901 || left[len(left)-1] != 902 || right[0] != 903 || right[len(right)-1] != 904 {
				t.Fatal("owned spectrum synthesis guards", LM, channels)
			}
		}
	}
}
func TestCeltDecodePriorityStoragePointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		p := celtDecodePriorityStorage(bands)
		entropyInitGrowStack(12)
		runtime.GC()
		if len(p) != int(bands) {
			t.Fatal("priority geometry")
		}
		for i, v := range p {
			if v != 0 {
				t.Fatal("priority initialization")
			}
			p[i] = int32(i % 2)
		}
		runtime.GC()
		for i, v := range p {
			if v != int32(i%2) {
				t.Fatal("priority retention")
			}
		}
	}
}
func TestCeltDecodePulseStoragePointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		p := celtDecodePulseStorage(bands)
		entropyInitGrowStack(12)
		runtime.GC()
		if len(p) != int(bands) {
			t.Fatal("pulse geometry")
		}
		for i, v := range p {
			if v != 0 {
				t.Fatal("pulse initialization")
			}
			p[i] = int32(i) * 128
		}
		runtime.GC()
		for i, v := range p {
			if v != int32(i)*128 {
				t.Fatal("pulse retention")
			}
		}
	}
}
func TestCeltDecodeFineStoragePointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		q := celtDecodeFineStorage(bands)
		entropyInitGrowStack(12)
		runtime.GC()
		if len(q) != int(bands) {
			t.Fatal("fine storage geometry")
		}
		for _, v := range q {
			if v != 0 {
				t.Fatal("zero-owned fine storage")
			}
		}
	}
	for _, channels := range []int32{1, 2} {
		mode := newSynthesisTestMode()
		q := celtDecodeFineStorage(21)
		priority := celtDecodePriorityStorage(21)
		for i := range q {
			q[i] = int32(i % 9)
			priority[i] = int32(i % 2)
		}
		energy := make([]float32, 42)
		for i := range energy {
			energy[i] = -12
		}
		want := append([]float32(nil), energy...)
		data := make([]byte, 128)
		for i := range data {
			data[i] = byte(i*71 + 13)
		}
		var ec OpusT_ec_ctx
		Opus_ec_dec_init(nil, &ec, &data[0], 128)
		c := ec
		Opus_unquant_fine_energy(nil, mode, 0, 21, &want[0], nil, &q[0], &c, channels)
		Opus_unquant_energy_finalise(nil, mode, 0, 21, &want[0], &q[0], &priority[0], 12, &c, channels)
		entropyInitGrowStack(12)
		runtime.GC()
		Opus_unquant_fine_energy(nil, mode, 0, 21, &energy[0], nil, &q[0], &ec, channels)
		Opus_unquant_energy_finalise(nil, mode, 0, 21, &energy[0], &q[0], &priority[0], 12, &ec, channels)
		if ec != c {
			t.Fatal("owned fine entropy")
		}
		for i := range energy {
			if math.Float32bits(energy[i]) != math.Float32bits(want[i]) {
				t.Fatal("owned fine consumers", channels, i)
			}
		}
	}
}
func TestCeltDecodeOffsetsStoragePointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		got := celtDecodeOffsetsStorage(bands)
		entropyInitGrowStack(12)
		runtime.GC()
		if len(got) != int(bands) {
			t.Fatal("boost offset geometry")
		}
		for i, v := range got {
			if v != 0 {
				t.Fatal("zero-owned offsets")
			}
			got[i] = int32(i) * 16
		}
		runtime.GC()
		for i, v := range got {
			if v != int32(i)*16 {
				t.Fatal("offset owner retention")
			}
		}
	}
	for LM := int32(0); LM <= 3; LM++ {
		mode := newSynthesisTestMode()
		offsets := celtDecodeOffsetsStorage(21)
		for i := range offsets {
			offsets[i] = int32(i%3) * 16
		}
		caps := celtDecodeCapsStorage(nil, mode, 21, LM, 2)
		var intensity, dual, balance int32
		pulses := celtDecodePulseStorage(21)
		priority := celtDecodePriorityStorage(21)
		var fine [21]int32
		data := make([]byte, 128)
		for i := range data {
			data[i] = byte(i*71 + 13)
		}
		var ec OpusT_ec_ctx
		Opus_ec_dec_init(nil, &ec, &data[0], uint32(len(data)))
		entropyInitGrowStack(12)
		runtime.GC()
		ret := clt_compute_allocation(nil, mode, 0, 21, &offsets[0], &caps[0], 5, &intensity, &dual, 512, &balance, &pulses[0], &fine[0], &priority[0], 2, LM, &ec, 0, 0, 0)
		if ret < 0 || ret > 21 {
			t.Fatal("owned offsets allocation consumer", LM, ret)
		}
	}
}
func TestCeltDecodeCapsStoragePointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			mode := newSynthesisTestMode()
			want := make([]int32, mode.FnbEBands)
			Opus_init_caps(nil, mode.FeBands, mode.Fcache.Fcaps, &want[0], mode.FnbEBands, LM, channels)
			entropyInitGrowStack(12)
			runtime.GC()
			got := celtDecodeCapsStorage(nil, mode, mode.FnbEBands, LM, channels)
			mode = nil
			runtime.GC()
			if len(got) != 21 {
				t.Fatal("owned caps geometry")
			}
			for i := range got {
				if got[i] != want[i] {
					t.Fatal("owned caps", LM, channels, i)
				}
			}
		}
	}
	if len(celtDecodeCapsStorage(nil, &OpusT_OpusCustomMode{}, 0, 0, 1)) != 0 {
		t.Fatal("empty caps")
	}
}
func TestCeltDecodeTFStoragePointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, transient := range []int32{0, 1} {
			for _, start := range []int32{0, 5} {
				data := make([]byte, 64)
				for i := range data {
					data[i] = byte(i*71 + 13)
				}
				ec := new(OpusT_ec_ctx)
				Opus_ec_dec_init(nil, ec, &data[0], uint32(len(data)))
				wantEC := *ec
				want := make([]int32, 21)
				tf_decode(nil, start, 21, transient, &want[0], LM, &wantEC)
				entropyInitGrowStack(12)
				runtime.GC()
				got := celtDecodeTFStorage(nil, 21, start, 21, transient, LM, ec)
				runtime.GC()
				if len(got) != 21 || *ec != wantEC {
					t.Fatal("owned TF entropy/shape")
				}
				for i := range got {
					if got[i] != want[i] {
						t.Fatal("owned TF flags", LM, transient, start, i)
					}
				}
			}
		}
	}
	data := []byte{0}
	var ec OpusT_ec_ctx
	Opus_ec_dec_init(nil, &ec, &data[0], 1)
	if len(celtDecodeTFStorage(nil, 0, 0, 0, 0, 0, &ec)) != 0 {
		t.Fatal("empty TF storage")
	}
}
func TestCeltDecodeEnergyClearPointers(t *testing.T) {
	for _, bands := range []int32{1, 3, 21, 25} {
		for _, start := range []int32{0, 1, bands} {
			for _, end := range []int32{0, bands - 1, bands} {
				e, l, p := make([]float32, 2*bands+2), make([]float32, 2*bands+2), make([]float32, 2*bands+2)
				for i := range e {
					e[i] = float32(i + 1)
					l[i] = float32(i + 2)
					p[i] = float32(i + 3)
				}
				we, wl, wp := append([]float32(nil), e...), append([]float32(nil), l...), append([]float32(nil), p...)
				for c := int32(0); c < 2; c++ {
					for i := int32(0); i < bands; i++ {
						if i < start || i >= end {
							index := 1 + c*bands + i
							we[index] = 0
							wp[index] = -28
							wl[index] = -28
						}
					}
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodeEnergyClear(&e[1], &l[1], &p[1], bands, start, end)
				for i := range e {
					if math.Float32bits(e[i]) != math.Float32bits(we[i]) || l[i] != wl[i] || p[i] != wp[i] {
						t.Fatal("energy band clearing", bands, start, end, i)
					}
				}
			}
		}
	}
	celtDecodeEnergyClear(nil, nil, nil, 0, 0, 0)
	celtDecodeEnergyClear(nil, nil, nil, 21, 0, 21)
	// Go-only overlapping history views retain channel/band and nested-assignment
	// order even when start/end ranges overlap.
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := append([]float32(nil), a...)
	for c := 0; c < 2; c++ {
		for i := 0; i < 2; i++ {
			index := c*3 + i
			want[index] = 0
			want[index+2] = -28
			want[index+1] = -28
		}
		for i := 1; i < 3; i++ {
			index := c*3 + i
			want[index] = 0
			want[index+2] = -28
			want[index+1] = -28
		}
	}
	celtDecodeEnergyClear(&a[0], &a[1], &a[2], 3, 2, 1)
	for i := range a {
		if a[i] != want[i] {
			t.Fatal("energy clearing store order", i)
		}
	}
}
func TestCeltDecodeEnergyBackgroundPointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, loss := range []int32{-1, 0, 40, 160, 10000} {
			for _, M := range []int32{1, 2, 4, 8} {
				b, e := make([]float32, 2*bands+2), make([]float32, 2*bands+2)
				for i := range b {
					b[i] = float32(i%9 - 4)
					e[i] = float32(i%7 - 3)
				}
				want := append([]float32(nil), b...)
				increase := float32(min(int32(160), loss+M)) * float32(.001)
				for i := int32(1); i <= 2*bands; i++ {
					if want[i]+increase < e[i] {
						want[i] += increase
					} else {
						want[i] = e[i]
					}
				}
				state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Floss_duration: loss}
				entropyInitGrowStack(12)
				runtime.GC()
				celtDecodeEnergyBackground(state, &b[1], &e[1], bands, M)
				if state.Floss_duration != loss {
					t.Fatal("background changed duration")
				}
				for i := range b {
					if math.Float32bits(b[i]) != math.Float32bits(want[i]) {
						t.Fatal("background energy", bands, loss, M, i)
					}
				}
			}
		}
	}
	state := &OpusT_OpusCustomDecoder{}
	celtDecodeEnergyBackground(state, nil, nil, 0, 1)
	b := []float32{float32(math.NaN()), 1}
	e := []float32{3, float32(math.NaN())}
	celtDecodeEnergyBackground(state, &b[0], &e[0], 1, 1)
	if b[0] != 3 || !math.IsNaN(float64(b[1])) {
		t.Fatal("background MING NaN selection")
	}
	state.Floss_duration = math.MaxInt32
	b = []float32{0, 0}
	e = []float32{1, 1}
	celtDecodeEnergyBackground(state, &b[0], &e[0], 1, 1)
	want := float32(int32(math.MinInt32)) * float32(.001)
	if b[0] != want || b[1] != want {
		t.Fatal("Go-only background int32 wrapping", b, want)
	}
}
func TestCeltDecodeEnergyLogsPointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, transient := range []int32{-1, 0, 1, 7} {
			e, l, p := make([]float32, 2*bands+2), make([]float32, 2*bands+2), make([]float32, 2*bands+2)
			for i := range e {
				e[i] = float32(i%7 - 3)
				l[i] = float32(i%9 - 4)
				p[i] = float32(i%11 - 5)
			}
			wl, wp := append([]float32(nil), l...), append([]float32(nil), p...)
			if transient == 0 {
				copy(wp[1:1+2*bands], wl[1:1+2*bands])
				copy(wl[1:1+2*bands], e[1:1+2*bands])
			} else {
				for i := int32(1); i <= 2*bands; i++ {
					if !(wl[i] < e[i]) {
						wl[i] = e[i]
					}
				}
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celtDecodeEnergyLogs(&e[1], &l[1], &p[1], bands, transient)
			for i := range l {
				if math.Float32bits(l[i]) != math.Float32bits(wl[i]) || math.Float32bits(p[i]) != math.Float32bits(wp[i]) {
					t.Fatal("energy log order", bands, transient, i)
				}
			}
		}
	}
	celtDecodeEnergyLogs(nil, nil, nil, 0, 0)
	e := []float32{3, math.Float32frombits(0x80000000)}
	l := []float32{float32(math.NaN()), 0}
	celtDecodeEnergyLogs(&e[0], &l[0], nil, 1, 1)
	if l[0] != 3 || math.Float32bits(l[1]) != 0x80000000 {
		t.Fatal("MING tie/NaN selection")
	}
	// Go-only overlapping memcpy images: first previous<-log, then log<-energy.
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := append([]float32(nil), a...)
	copy(want[3:7], want[1:5])
	copy(want[1:5], want[:4])
	celtDecodeEnergyLogs(&a[0], &a[1], &a[3], 2, 0)
	for i := range a {
		if a[i] != want[i] {
			t.Fatal("log copy alias order")
		}
	}
}
func TestCeltDecodeEnergyMonoPointers(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		a := make([]float32, 2*bands+2)
		a[0], a[len(a)-1] = 77, 88
		for i := int32(0); i < bands; i++ {
			a[1+i] = math.Float32frombits(uint32(i)*0x1234567 + 0x80000000)
		}
		want := append([]float32(nil), a...)
		copy(want[1+bands:1+2*bands], want[1:1+bands])
		entropyInitGrowStack(12)
		runtime.GC()
		celtDecodeEnergyMono(&a[1], bands)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
				t.Fatal("mono energy ownership", bands, i)
			}
		}
	}
	celtDecodeEnergyMono(nil, 0)
	a := []float32{float32(math.NaN()), float32(math.Inf(1)), math.Float32frombits(0x80000000), 0, 0, 0}
	celtDecodeEnergyMono(&a[0], 3)
	for i := 0; i < 3; i++ {
		if math.Float32bits(a[i]) != math.Float32bits(a[3+i]) {
			t.Fatal("mono bit preservation")
		}
	}
}
func TestCeltPLCDispatchPointers(t *testing.T) {
	for _, duration := range []int32{-1, 0, 39, 40, 10000} {
		for _, start := range []int32{0, 1, 20} {
			for _, skip := range []int32{-1, 0, 1} {
				state := &OpusT_OpusCustomDecoder{Fmode: newSynthesisTestMode(), Floss_duration: 123, Fplc_duration: duration, Fstart: start, Fskip_plc: skip}
				before := *state
				entropyInitGrowStack(12)
				runtime.GC()
				loss, s, kind := celtPLCDispatch(state)
				want := int32(FRAME_PLC_PERIODIC)
				if duration >= 40 || start != 0 || skip != 0 {
					want = FRAME_PLC_NOISE
				}
				if loss != 123 || s != start || kind != want || *state != before {
					t.Fatal("typed dispatch", duration, start, skip)
				}
			}
		}
	}
}
func TestCeltPLCLostPointers(t *testing.T) {
	// Complete active periodic/noise concealment with typed state/history owners.
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, noise := range []bool{false, true} {
				mode := newSynthesisTestMode()
				owner, image, size := celtStateTestBuffer(mode, channels)
				state := &owner.State
				if ret := opus_custom_decoder_init(nil, state, mode, channels); ret != 0 {
					t.Fatal("concealment init", ret)
				}
				state.Frng = 0xdeadbeef
				state.Flast_pitch_index = 100
				state.Flast_frame_type = FRAME_PLC_PERIODIC
				state.Fskip_plc = libc.BoolInt32(noise)
				history := unsafe.Slice(&state.F_decode_mem[0], int(channels)*(DEC_PITCH_BUF_SIZE+int(mode.Foverlap)))
				for i := range history {
					history[i] = float32(math.Sin(float64(i)*.17) * .03)
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celt_decode_lost(nil, state, mode.FshortMdctSize<<LM, LM)
				if state.Floss_duration != 1<<LM || state.Fplc_duration != 1<<LM {
					t.Fatal("concealment durations", channels, LM, noise)
				}
				if noise && state.Flast_frame_type != FRAME_PLC_NOISE || !noise && state.Flast_frame_type != FRAME_PLC_PERIODIC {
					t.Fatal("concealment dispatch")
				}
				for _, v := range image[size : size+16] {
					if v != 0xa5 {
						t.Fatal("concealment guard")
					}
				}
				tls := libc.NewTLS()
				libc.Xpthread_setspecific(tls, 0x6f707573, 123)
				celt_decode_lost(tls, state, mode.FshortMdctSize<<LM, LM)
				if libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
					t.Fatal("concealment touched TLS")
				}
				tls.Close()
				runtime.KeepAlive(owner)
			}
		}
	}
}
func TestCeltPLCFirstLossPointers(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			mode := newSynthesisTestMode()
			owner, image, size := celtStateTestBuffer(mode, channels)
			state := &owner.State
			opus_custom_decoder_init(nil, state, mode, channels)
			state.Fskip_plc = 0
			state.Flast_frame_type = 0
			state.Fpostfilter_period_old = 80
			state.Fpostfilter_period = 96
			state.Fpostfilter_gain_old = .13
			state.Fpostfilter_gain = .2
			state.Fpostfilter_tapset_old = 1
			state.Fpostfilter_tapset = 2
			h := unsafe.Slice(&state.F_decode_mem[0], (2048+120)*channels)
			for i := range h {
				h[i] = float32(math.Sin(float64(i)*.17) * .03)
			}
			entropyInitGrowStack(12)
			runtime.GC()
			celt_decode_lost(nil, state, 120<<LM, LM)
			if state.Flast_frame_type != FRAME_PLC_PERIODIC || state.Flast_pitch_index < PLC_PITCH_LAG_MIN || state.Flast_pitch_index > PLC_PITCH_LAG_MAX || state.Fprefilter_and_fold != 1 {
				t.Fatal("first loss pitch/LPC")
			}
			state.Fskip_plc = 1
			runtime.GC()
			celt_decode_lost(nil, state, 120<<LM, LM)
			if state.Flast_frame_type != FRAME_PLC_NOISE || state.Fprefilter_and_fold != 0 || state.Fskip_plc != 1 {
				t.Fatal("periodic to folded noise")
			}
			for _, v := range image[size : size+16] {
				if v != 165 {
					t.Fatal("first loss guard")
				}
			}
			runtime.KeepAlive(owner)
		}
	}
}
func TestCeltPLCFIRStoragePointers(t *testing.T) {
	for _, length := range []int32{0, 80, 200, 1024} {
		storage := celtPLCFIRStorage(length)
		if len(storage) != int(length) {
			t.Fatal("FIR geometry")
		}
		if length == 0 {
			continue
		}
		exc := celtPLCExcitationStorage(length)
		for i := range exc {
			exc[i] = float32(math.Sin(float64(i) * .17))
		}
		var coef [24]float32
		coef[0] = .125
		coef[23] = -.03125
		want := make([]float32, length)
		Opus_celt_fir_c(nil, &exc[24], &coef[0], &want[0], length, 24, 0)
		entropyInitGrowStack(12)
		runtime.GC()
		Opus_celt_fir_c(nil, &exc[24], &coef[0], unsafe.SliceData(storage), length, 24, 0)
		for i := range storage {
			if math.Float32bits(storage[i]) != math.Float32bits(want[i]) {
				t.Fatal("owned FIR storage", length, i)
			}
		}
		copy(exc[24:], storage)
		for i := range storage {
			if exc[24+i] != storage[i] {
				t.Fatal("FIR copy-back")
			}
		}
	}
}
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
