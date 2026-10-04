//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestCeltDecodePostfilterFirstAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	for _, gain := range []float32{0, .3, .7} {
		for tap := int32(0); tap < 3; tap++ {
			state := opuscc.OpusT_OpusCustomDecoder{Fpostfilter_period_old: 45, Fpostfilter_period: 100, Fpostfilter_gain_old: gain, Fpostfilter_gain: .3, Fpostfilter_tapset_old: tap, Fpostfilter_tapset: 2}
			h := make([]float32, 256+120+2)
			for i := range h {
				h[i] = float32(i%13-6) / 128
			}
			c := slices.Clone(h)
			opuscc.CompareCeltDecodePostfilterFirst(&state, mode, &h[257], 120)
			nativeComb(&c[257], &c[257], 45, 100, 120, gain, .3, tap, 2, mode.Fwindow, 120)
			for i := range h {
				if math.Float32bits(h[i]) != math.Float32bits(c[i]) {
					t.Fatal("first postfilter", gain, tap, i)
				}
			}
		}
	}
}
func TestCeltDecodePostfilterClampAgainstC(t *testing.T) {
	for _, a := range []int32{-2147483648, -1, 0, 15, 16, 100, 2147483647} {
		for _, b := range []int32{-1, 0, 15, 16, 2147483647} {
			g := opuscc.OpusT_OpusCustomDecoder{Fpostfilter_period: a, Fpostfilter_period_old: b, Frng: 123, Fpostfilter_gain: .5}
			c := g
			opuscc.CompareCeltDecodePostfilterClamp(&g)
			nativeCeltDecodePostfilterClamp(&c)
			if g != c {
				t.Fatal("postfilter clamp", a, b)
			}
		}
	}
}
func TestCeltDecodePacketFinishAgainstC(t *testing.T) {
	for _, v := range []int32{-2147483648, -1, 0, 1, 40, 2147483647} {
		g := opuscc.OpusT_OpusCustomDecoder{Floss_duration: v, Fplc_duration: v, Flast_frame_type: v, Fprefilter_and_fold: v, Frng: 123, Ferror1: 7, Fpostfilter_period: 45}
		c := g
		opuscc.CompareCeltDecodePacketFinish(&g)
		nativeCeltDecodePacketFinish(&c)
		if g != c {
			t.Fatal("packet finish", v)
		}
	}
}
func TestCeltDecodeRecoverEnergyAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, loss := range []int32{0, 1, 10, 40} {
			for _, intra := range []int32{0, 1, -1} {
				for _, start := range []int32{0, 2, 3} {
					e, l, p := make([]float32, 8), make([]float32, 8), make([]float32, 8)
					for i := range e {
						e[i] = float32(i) - 10
						l[i] = float32(i) - 8
						p[i] = float32(i) - 6
					}
					c := slices.Clone(e)
					state := opuscc.OpusT_OpusCustomDecoder{Floss_duration: loss}
					opuscc.CompareCeltDecodeRecoverEnergy(&state, &e[1], &l[1], &p[1], 3, start, 3, LM, intra)
					nativeCeltDecodeRecoverEnergy(c[1:], l[1:], p[1:], 3, start, 3, LM, intra, loss)
					for i := range e {
						if math.Float32bits(e[i]) != math.Float32bits(c[i]) {
							t.Fatal("energy recovery loop", LM, loss, intra, start, i)
						}
					}
				}
			}
		}
	}
}
func TestCeltDecodeRecoveryBandAgainstC(t *testing.T) {
	values := []float32{-25, -20, -10, -1, 0, math.Float32frombits(0x80000000), .1, 2, 4, math.Float32frombits(0x7fc12345), float32(math.Inf(1)), float32(math.Inf(-1))}
	for _, e := range values {
		for _, l := range values {
			for _, p := range values {
				for _, missing := range []int32{0, 1, 10} {
					for _, safety := range []float32{0, .5, 1.5} {
						g, c := e, e
						opuscc.CompareCeltDecodeRecoveryBand(&g, &l, &p, missing, safety)
						nativeCeltDecodeRecoveryBand(&c, &l, &p, missing, safety)
						if math.Float32bits(g) != math.Float32bits(c) {
							t.Fatal("recovery band", e, l, p, missing, safety, g, c)
						}
					}
				}
			}
		}
	}
}
func TestCeltDecodeRecoverySafetyAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, loss := range []int32{-2147483648, -1, 0, 1, 10, 11, 40, 2147483647} {
			state := opuscc.OpusT_OpusCustomDecoder{Floss_duration: loss}
			m, s := opuscc.CompareCeltDecodeRecoverySafety(&state, LM)
			cm, cs := nativeCeltDecodeRecoverySafety(loss, LM)
			if m != cm || math.Float32bits(s) != math.Float32bits(cs) {
				t.Fatal("recovery safety", LM, loss)
			}
		}
	}
}
func TestCeltDecodePostfilterFinishAgainstC(t *testing.T) {
	for _, LM := range []int32{-1, 0, 1, 3} {
		for _, period := range []int32{-2147483648, -1, 0, 15, 2147483647} {
			for _, bits := range []uint32{0, 0x80000000, 0x3f400000, 0x7fc12345, 0x7f800000} {
				g := opuscc.OpusT_OpusCustomDecoder{Fpostfilter_period: 45, Fpostfilter_period_old: 99, Fpostfilter_gain: math.Float32frombits(0x80000000), Fpostfilter_gain_old: 1, Fpostfilter_tapset: 2, Fpostfilter_tapset_old: 1, Frng: 123, Floss_duration: 7}
				c := g
				gain := math.Float32frombits(bits)
				opuscc.CompareCeltDecodePostfilterFinish(&g, period, gain, -3, LM)
				nativeCeltDecodePostfilterFinish(&c, period, gain, -3, LM)
				if g.Fpostfilter_period != c.Fpostfilter_period || g.Fpostfilter_period_old != c.Fpostfilter_period_old || g.Fpostfilter_tapset != c.Fpostfilter_tapset || g.Fpostfilter_tapset_old != c.Fpostfilter_tapset_old || math.Float32bits(g.Fpostfilter_gain) != math.Float32bits(c.Fpostfilter_gain) || math.Float32bits(g.Fpostfilter_gain_old) != math.Float32bits(c.Fpostfilter_gain_old) || g.Frng != c.Frng || g.Floss_duration != c.Floss_duration {
					t.Fatal("postfilter state", LM, period, bits)
				}
			}
		}
	}
}
func TestCeltDecodeBoostsAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			for _, budget := range []int32{0, 8, 100, 500, 992} {
				for _, start := range []int32{0, 2, 5} {
					for _, pattern := range []byte{0, 255, 71} {
						bands := []int16{0, 1, 3, 6, 10, 18}
						caps := []int32{0, 8, 40, 128, 500}
						out := []int32{901, -1, -2, -3, -4, -5, 902}
						co := slices.Clone(out)
						data := make([]byte, 128)
						for i := range data {
							data[i] = pattern
						}
						var ec opuscc.OpusT_ec_ctx
						opuscc.Opus_ec_dec_init(nil, &ec, &data[0], 128)
						c := ec
						r, tell := opuscc.CompareCeltDecodeBoosts(&bands[0], &caps[0], &out[1], start, 5, channels, LM, budget, &ec)
						cr, ct := nativeCeltDecodeBoosts(&c, data, bands, caps, co[1:], start, 5, channels, LM, budget)
						if ec != c || r != cr || tell != ct || !slices.Equal(out, co) {
							t.Fatal("dynamic boosts", LM, channels, budget, start, pattern, r, cr, tell, ct)
						}
					}
				}
			}
		}
	}
}
func TestCeltDecodeSilenceEnergyAgainstC(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, channels := range []int32{1, 2} {
			e := make([]float32, bands*channels+2)
			for i := range e {
				e[i] = math.Float32frombits(0x7fc12345)
			}
			e[0] = 901
			e[len(e)-1] = -902
			c := slices.Clone(e)
			opuscc.CompareCeltDecodeSilenceEnergy(&e[1], bands, channels)
			nativeCeltDecodeSilenceEnergy(&c[1], bands, channels)
			for i := range e {
				if math.Float32bits(e[i]) != math.Float32bits(c[i]) {
					t.Fatal("silence energy", bands, channels, i)
				}
			}
		}
	}
}
func TestCeltDecodeHistoryMoveAgainstC(t *testing.T) {
	for _, N := range []int32{0, 1, 120, 240, 960} {
		for _, length := range []int32{0, 1, 128, 2048} {
			h := make([]float32, N+length+2)
			for i := range h {
				h[i] = math.Float32frombits(uint32(i)*7717 + 0x80000000)
			}
			c := slices.Clone(h)
			opuscc.CompareCeltDecodeHistoryMove(&h[1], N, length)
			nativeCeltDecodeHistoryMove(&c[1], N, length)
			for i := range h {
				if math.Float32bits(h[i]) != math.Float32bits(c[i]) {
					t.Fatal("history move", N, length, i)
				}
			}
		}
	}
}
func TestCeltDecodeMaskStorageAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	bands := unsafe.Slice(mode.FeBands, 22)
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			N := int32(120) << LM
			m := opuscc.CompareCeltDecodeMaskStorage(21, channels)
			for i := range m {
				if i%2 != 0 {
					m[i] = byte((1 << (1 << LM)) - 1)
				}
			}
			s := make([]float32, N*channels)
			c := slices.Clone(s)
			energy, previous, older := make([]float32, 42), make([]float32, 42), make([]float32, 42)
			for i := range energy {
				energy[i] = -12
				previous[i] = -10
				older[i] = -11
			}
			pulses := opuscc.CompareCeltDecodePulseStorage(21)
			for i := range pulses {
				pulses[i] = int32(i)*32 + 8
			}
			opuscc.Opus_anti_collapse(nil, mode.FeBands, 21, &s[0], &m[0], LM, channels, N, 0, 21, &energy[0], &previous[0], &older[0], &pulses[0], 123, 0, 0)
			nativeAntiCollapse(bands, 21, c, m, LM, channels, N, 0, 21, energy, previous, older, pulses, 123, 0)
			for i := range s {
				if math.Float32bits(s[i]) != math.Float32bits(c[i]) {
					t.Fatal("owned mask anti-collapse", LM, channels, i)
				}
			}
		}
	}
}
func TestCeltDecodeSpectrumStorageAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			N := int32(120) << LM
			s := opuscc.CompareCeltDecodeSpectrumStorage(N, channels)
			for i := range s {
				s[i] = float32(i%17-8) / 128
			}
			energy := make([]float32, 42)
			for i := range energy {
				energy[i] = -12
			}
			left, right := make([]float32, N+120), make([]float32, N+120)
			cl, cr := slices.Clone(left), slices.Clone(right)
			opuscc.CompareCeltSynthesis(nil, &s[0], &energy[0], &left[0], &right[0], 0, 21, channels, channels, 0, LM, 1, 0)
			nativeCeltSynthesis(s, energy, cl, cr, 0, 21, channels, channels, 0, LM, 1, 0)
			for i := range left {
				if math.Float32bits(left[i]) != math.Float32bits(cl[i]) || math.Float32bits(right[i]) != math.Float32bits(cr[i]) {
					t.Fatal("owned spectral synthesis", LM, channels, i)
				}
			}
		}
	}
}
func TestCeltDecodeFineStorageAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	for _, channels := range []int32{1, 2} {
		q := opuscc.CompareCeltDecodeFineStorage(21)
		priority := opuscc.CompareCeltDecodePriorityStorage(21)
		for i := range q {
			q[i] = int32(i % 9)
			priority[i] = int32(i % 2)
		}
		energy := make([]float32, 42)
		for i := range energy {
			energy[i] = -12
		}
		cEnergy := slices.Clone(energy)
		data := make([]byte, 128)
		for i := range data {
			data[i] = byte(i*71 + 13)
		}
		var ec opuscc.OpusT_ec_ctx
		opuscc.Opus_ec_dec_init(nil, &ec, &data[0], 128)
		c := ec
		opuscc.Opus_unquant_fine_energy(nil, mode, 0, 21, &energy[0], nil, &q[0], &ec, channels)
		opuscc.Opus_unquant_energy_finalise(nil, mode, 0, 21, &energy[0], &q[0], &priority[0], 12, &ec, channels)
		nativeEnergyDecode(&c, data, cEnergy, 21, 0, 21, channels, 1, 0, 0, nil, q)
		nativeEnergyDecode(&c, data, cEnergy, 21, 0, 21, channels, 2, 12, 0, q, priority)
		if ec != c {
			t.Fatal("owned fine entropy", channels)
		}
		for i := range energy {
			if math.Float32bits(energy[i]) != math.Float32bits(cEnergy[i]) {
				t.Fatal("owned fine energy", channels, i)
			}
		}
	}
}
func TestCeltDecodeOffsetsStorageAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	for LM := int32(0); LM <= 3; LM++ {
		offsets := opuscc.CompareCeltDecodeOffsetsStorage(21)
		for i := range offsets {
			offsets[i] = int32(i%3) * 16
		}
		caps := opuscc.CompareCeltDecodeCapsStorage(mode, 21, LM, 2)
		var a [7][23]int32
		copy(a[0][1:22], offsets)
		copy(a[3][1:22], caps)
		cfg := [12]int32{0, 21, 5, 512, 0, 0, 0, 2, LM, 0, 0, 0}
		data := make([]byte, 128)
		for i := range data {
			data[i] = byte(i*71 + 13)
		}
		var ec opuscc.OpusT_ec_ctx
		opuscc.Opus_ec_dec_init(nil, &ec, &data[0], uint32(len(data)))
		c := ec
		g, values, out := opuscc.CompareCeltDecodeOffsetsAllocation(mode, offsets, caps, LM, &ec)
		var cv [3]int32
		n := nativeAllocationDriver(&c, data, &a, &cv, &cfg)
		if g != n || values != cv || ec != c {
			t.Fatal("owned offsets allocation", LM, g, n, values, cv)
		}
		for k := range out {
			if !slices.Equal(out[k][:], a[k+4][1:22]) {
				t.Fatal("owned offsets results", LM, k)
			}
		}
	}
}
func TestCeltDecodeCapsStorageAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	bands := unsafe.Slice(mode.FeBands, mode.FnbEBands+1)
	cache := unsafe.Slice(mode.Fcache.Fcaps, 8*mode.FnbEBands)
	for LM := int32(0); LM <= 3; LM++ {
		for _, channels := range []int32{1, 2} {
			got := opuscc.CompareCeltDecodeCapsStorage(mode, mode.FnbEBands, LM, channels)
			want := make([]int32, len(got))
			nativeCaps(bands, cache, want, LM, channels)
			if !slices.Equal(got, want) {
				t.Fatal("owned caps", LM, channels)
			}
		}
	}
}
func TestCeltDecodeTFStorageAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, transient := range []int32{0, 1} {
			for _, start := range []int32{0, 5} {
				data := make([]byte, 64)
				for i := range data {
					data[i] = byte(i*71 + 13)
				}
				var ec opuscc.OpusT_ec_ctx
				opuscc.Opus_ec_dec_init(nil, &ec, &data[0], uint32(len(data)))
				c := ec
				want := make([]int32, 21)
				got := opuscc.CompareCeltDecodeTFStorage(21, start, 21, transient, LM, &ec)
				nativeTFDecode(&c, data, start, 21, transient, want, LM)
				if ec != c || !slices.Equal(got, want) {
					t.Fatal("owned TF storage", LM, transient, start)
				}
			}
		}
	}
}
func TestCeltDecodeEnergyClearAgainstC(t *testing.T) {
	for _, bands := range []int32{1, 3, 21, 25} {
		for _, start := range []int32{0, 1, bands} {
			for _, end := range []int32{0, bands - 1, bands} {
				e, l, p := make([]float32, 2*bands+2), make([]float32, 2*bands+2), make([]float32, 2*bands+2)
				for i := range e {
					e[i] = float32(i + 1)
					l[i] = float32(i + 2)
					p[i] = float32(i + 3)
				}
				ce, cl, cp := slices.Clone(e), slices.Clone(l), slices.Clone(p)
				opuscc.CompareCeltDecodeEnergyClear(&e[1], &l[1], &p[1], bands, start, end)
				nativeCeltDecodeEnergyClear(&ce[1], &cl[1], &cp[1], bands, start, end)
				for i := range e {
					if math.Float32bits(e[i]) != math.Float32bits(ce[i]) || l[i] != cl[i] || p[i] != cp[i] {
						t.Fatal("energy clearing", bands, start, end, i)
					}
				}
			}
		}
	}
}
func TestCeltDecodeEnergyBackgroundAgainstC(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, loss := range []int32{-1, 0, 40, 160, 10000} {
			for _, M := range []int32{1, 2, 4, 8} {
				b, e := make([]float32, 2*bands+2), make([]float32, 2*bands+2)
				for i := range b {
					b[i] = float32(i%9 - 4)
					e[i] = float32(i%7 - 3)
				}
				c := slices.Clone(b)
				state := opuscc.OpusT_OpusCustomDecoder{Floss_duration: loss}
				opuscc.CompareCeltDecodeEnergyBackground(&state, &b[1], &e[1], bands, M)
				nativeCeltDecodeEnergyBackground(&c[1], &e[1], bands, loss, M)
				for i := range b {
					if math.Float32bits(b[i]) != math.Float32bits(c[i]) {
						t.Fatal("background energy", bands, loss, M, i)
					}
				}
			}
		}
	}
}
func TestCeltDecodeEnergyLogsAgainstC(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		for _, transient := range []int32{-1, 0, 1, 7} {
			e, l, p := make([]float32, 2*bands+2), make([]float32, 2*bands+2), make([]float32, 2*bands+2)
			for i := range e {
				e[i] = float32(i%7 - 3)
				l[i] = float32(i%9 - 4)
				p[i] = float32(i%11 - 5)
			}
			cl, cp := slices.Clone(l), slices.Clone(p)
			opuscc.CompareCeltDecodeEnergyLogs(&e[1], &l[1], &p[1], bands, transient)
			nativeCeltDecodeEnergyLogs(&e[1], &cl[1], &cp[1], bands, transient)
			for i := range l {
				if math.Float32bits(l[i]) != math.Float32bits(cl[i]) || math.Float32bits(p[i]) != math.Float32bits(cp[i]) {
					t.Fatal("energy logs", bands, transient, i)
				}
			}
		}
	}
}
func TestCeltDecodeEnergyExceptionalAgainstC(t *testing.T) {
	e := []float32{3, float32(math.NaN()), math.Float32frombits(0x80000000), 0}
	l := []float32{float32(math.NaN()), 1, 0, math.Float32frombits(0x80000000)}
	c := slices.Clone(l)
	opuscc.CompareCeltDecodeEnergyLogs(&e[0], &l[0], nil, 2, 1)
	nativeCeltDecodeEnergyLogs(&e[0], &c[0], nil, 2, 1)
	for i := range l {
		if math.Float32bits(l[i]) != math.Float32bits(c[i]) {
			t.Fatal("MING NaN/zero selection", i)
		}
	}
	b := []float32{float32(math.NaN()), 1, 0, float32(math.Inf(1))}
	c = slices.Clone(b)
	state := opuscc.OpusT_OpusCustomDecoder{}
	opuscc.CompareCeltDecodeEnergyBackground(&state, &b[0], &e[0], 2, 1)
	nativeCeltDecodeEnergyBackground(&c[0], &e[0], 2, 0, 1)
	for i := range b {
		if math.Float32bits(b[i]) != math.Float32bits(c[i]) {
			t.Fatal("background NaN/zero/infinity", i)
		}
	}
}
func TestCeltDecodeEnergyMonoAgainstC(t *testing.T) {
	for _, bands := range []int32{0, 1, 3, 21, 25} {
		a := make([]float32, 2*bands+2)
		a[0], a[len(a)-1] = 77, 88
		for i := int32(0); i < bands; i++ {
			a[1+i] = math.Float32frombits(uint32(i)*0x1234567 + 0x80000000)
		}
		b := slices.Clone(a)
		opuscc.CompareCeltDecodeEnergyMono(&a[1], bands)
		nativeCeltDecodeEnergyMono(&b[1], bands)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("mono energy", bands, i)
			}
		}
	}
}
func TestCeltPLCLostAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, scenario := range []int{0, 1, 2, 3} {
				noise := scenario == 1 || scenario == 3
				size := int(opuscc.CompareCustomDecoderSize(mode, channels))
				data := make([]byte, size+16)
				st := (*opuscc.OpusT_OpusCustomDecoder)(unsafe.Pointer(&data[0]))
				st.Fchannels = channels
				st.Fstream_channels = channels
				st.Foverlap = 120
				st.Fdownsample = 1
				st.Fend = 21
				st.Flast_frame_type = opuscc.FRAME_PLC_PERIODIC
				st.Flast_pitch_index = 100
				st.Frng = 0xdeadbeef
				if noise {
					st.Fskip_plc = 1
				}
				if scenario == 2 {
					st.Flast_frame_type = 0
				}
				if scenario == 3 {
					st.Fprefilter_and_fold = 1
					st.Fpostfilter_period_old = 80
					st.Fpostfilter_period = 96
					st.Fpostfilter_gain_old = .13
					st.Fpostfilter_gain = .2
					st.Fpostfilter_tapset_old = 1
					st.Fpostfilter_tapset = 2
				}
				history := unsafe.Slice(&st.F_decode_mem[0], (2048+120)*channels)
				for i := range history {
					history[i] = float32(math.Sin(float64(i)*.17) * .03)
				}
				energies := unsafe.Slice((*float32)(unsafe.Add(unsafe.Pointer(&st.F_decode_mem[0]), int((2048+120)*channels)*4)), 168)
				for i := range energies {
					energies[i] = -12
				}
				for i := size; i < len(data); i++ {
					data[i] = 165
				}
				c := slices.Clone(data)
				for call := 0; call < 3; call++ {
					st.Fmode = mode
					opuscc.CompareCeltPLCLost(nil, st, 120<<LM, LM)
					st.Fmode = nil
					if ret := nativeCeltLost(c, 120<<LM, LM); ret != 0 {
						t.Fatal("native concealment", ret)
					}
					if !slices.Equal(data, c) {
						for i := range data {
							if data[i] != c[i] {
								t.Fatal("whole concealment", channels, LM, scenario, call, i, data[i], c[i])
							}
						}
					}
				}
			}
		}
	}
}
func TestCeltPLCDispatchAgainstC(t *testing.T) {
	for _, duration := range []int32{-1, 0, 39, 40, 10000} {
		for _, start := range []int32{0, 1, 20} {
			for _, skip := range []int32{-1, 0, 1} {
				state := opuscc.OpusT_OpusCustomDecoder{Floss_duration: 123, Fplc_duration: duration, Fstart: start, Fskip_plc: skip}
				loss, s, kind := opuscc.CompareCeltPLCDispatch(&state)
				if loss != 123 || s != start || (kind == opuscc.FRAME_PLC_NOISE) != nativeCeltPLCDispatch(duration, start, skip) {
					t.Fatal("dispatch", duration, start, skip)
				}
			}
		}
	}
}
func TestCeltPLCFIRStorageAgainstC(t *testing.T) {
	for _, length := range []int32{80, 200, 1024} {
		input := make([]float32, length+24)
		for i := range input {
			input[i] = float32(math.Sin(float64(i) * .17))
		}
		coefficients := make([]float32, 24)
		coefficients[0], coefficients[23] = .125, -.03125
		a := make([]float32, length+2)
		a[0], a[len(a)-1] = 77, 88
		b := slices.Clone(a)
		opuscc.Opus_celt_fir_c(nil, &input[24], &coefficients[0], &a[1], length, 24, 0)
		nativeFIR(input, coefficients, b[1:len(b)-1], length, 24)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("concealment FIR spans", length, i, a[i], b[i])
			}
		}
	}
}
func TestCeltPLCNoiseAgainstC(t *testing.T) {
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, seed := range []uint32{0, 1, 0xffffffff, 0xdeadbeef} {
				bands := []int16{0, 1, 3, 6, 12, 20}
				N := int32(20) << LM
				a := make([]float32, N*channels+2)
				a[0], a[len(a)-1] = 77, 88
				b := slices.Clone(a)
				state := opuscc.OpusT_OpusCustomDecoder{Frng: seed}
				opuscc.CompareCeltPLCNoise(&state, &bands[0], &a[1], N, 1, 5, LM, channels)
				s := nativeCeltPLCNoise(seed, &bands[0], &b[1], N, 1, 5, LM, channels)
				if state.Frng != s {
					t.Fatal("noise seed", channels, LM, seed)
				}
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("noise spectrum", channels, LM, seed, i, a[i], b[i])
					}
				}
			}
		}
	}
}
func TestCeltPLCFinishAgainstC(t *testing.T) {
	for _, loss := range []int32{-10, 0, 1, 9999, 10000} {
		for _, plc := range []int32{-10, 0, 1, 9999, 10000} {
			for LM := int32(0); LM <= 3; LM++ {
				for _, frameType := range []int32{opuscc.FRAME_PLC_PERIODIC, opuscc.FRAME_PLC_NOISE, opuscc.FRAME_PLC_NEURAL, opuscc.FRAME_DRED} {
					a := opuscc.OpusT_OpusCustomDecoder{Floss_duration: 123, Fplc_duration: plc, Fprefilter_and_fold: 1, Fskip_plc: 1, Frng: 0xdeadbeef}
					b := a
					opuscc.CompareCeltPLCFinish(&a, loss, LM, frameType)
					nativeCeltPLCFinish(&b, loss, LM, frameType)
					if a != b {
						t.Fatal("finish", loss, plc, LM, frameType)
					}
				}
			}
		}
	}
}
func TestCeltPLCLPCHistoryAgainstC(t *testing.T) {
	for _, N := range []int32{0, 120, 240, 960} {
		h := make([]float32, 2048)
		for i := range h {
			h[i] = math.Float32frombits(uint32(i) * uint32(7919))
		}
		a := [26]float32{}
		a[0], a[25] = 77, 88
		b := a
		opuscc.CompareCeltPLCLPCHistory((*[24]float32)(unsafe.Pointer(&a[1])), &h[0], 2048, N)
		nativeCeltPLCLPCHistory((*[24]float32)(unsafe.Pointer(&b[1])), &h[0], 2048, N)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("LPC history", N, i)
			}
		}
	}
}
func TestCeltPLCExtrapolateAgainstC(t *testing.T) {
	for _, pitch := range []int32{40, 100, 511, 1024} {
		for _, N := range []int32{120, 240, 960} {
			for _, decay := range []float32{0, .923, 1} {
				a := make([]float32, 2170)
				x := make([]float32, 1024)
				a[0], a[len(a)-1] = 77, 88
				for i := 1; i < len(a)-1; i++ {
					a[i] = float32((i*37)%79-39) / 13
				}
				for i := range x {
					x[i] = float32((i*43)%89-44) / 17
				}
				b := slices.Clone(a)
				g := opuscc.CompareCeltPLCExtrapolate(&a[1], &x[0], 2048, 1024, N, 120, pitch, .8, decay)
				c := nativeCeltPLCExtrapolate(&b[1], &x[0], 2048, 1024, N, 120, pitch, .8, decay)
				if math.Float32bits(g) != math.Float32bits(c) {
					t.Fatal("extrapolation energy", pitch, N, decay, g, c)
				}
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("extrapolation", pitch, N, decay, i)
					}
				}
			}
		}
	}
}
func TestCeltPLCExcitationHistoryAgainstC(t *testing.T) {
	for _, period := range []int32{0, 1, 512, 1024} {
		h := make([]float32, 2050)
		for i := range h {
			h[i] = math.Float32frombits(uint32(i) * uint32(7919))
		}
		before := slices.Clone(h)
		a := make([]float32, period+26)
		a[0], a[len(a)-1] = 77, 88
		b := slices.Clone(a)
		opuscc.CompareCeltPLCExcitationHistory(&a[1], &h[1], 2048, period)
		nativeCeltPLCExcitationHistory(&b[1], &h[1], 2048, period)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("excitation history", period, i)
			}
		}
		if !slices.Equal(h, before) {
			t.Fatal("history input changed")
		}
	}
}
func TestCeltPLCSynthesisAttenuateAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2453))
	for _, length := range []int32{1, 120, 240, 1080} {
		for _, overlap := range []int32{0, 1, min(120, length)} {
			for _, factor := range []float32{0, .1, .2, .21, .5, 1, 2, float32(math.NaN())} {
				a := make([]float32, length+2)
				a[0], a[len(a)-1] = 77, 88
				w := make([]float32, overlap)
				for i := int32(0); i < length; i++ {
					a[i+1] = float32(rng.NormFloat64() * 7)
				}
				for i := range w {
					w[i] = float32(i+1) / float32(len(w)+1)
				}
				b := slices.Clone(a)
				s2 := float32(0)
				for _, v := range a[1 : len(a)-1] {
					s2 += float32(v * v)
				}
				s1 := float32(factor * s2)
				opuscc.CompareCeltPLCSynthesisAttenuate(&a[1], unsafe.SliceData(w), length, overlap, s1)
				nativeCeltPLCSynthesisAttenuate(&b[1], unsafe.SliceData(w), length, overlap, s1)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("attenuation", length, overlap, factor, i, a[i], b[i])
					}
				}
			}
		}
	}
}
func TestCeltPLCExcitationDecayAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2452))
	for _, length := range []int32{0, 1, 2, 31, 120, 512, 1024} {
		for trial := 0; trial < 80; trial++ {
			a := make([]float32, 1024)
			for i := range a {
				a[i] = float32(rng.NormFloat64() * 1e4)
			}
			g := opuscc.CompareCeltPLCExcitationDecay(&a[0], 1024, length)
			c := nativeCeltPLCExcitationDecay(&a[0], 1024, length)
			if math.Float32bits(g) != math.Float32bits(c) {
				t.Fatal("excitation decay", length, trial, g, c)
			}
		}
	}
}
func TestCeltPLCLagWindowAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2451))
	for trial := 0; trial < 1000; trial++ {
		var a [25]float32
		for i := range a {
			a[i] = float32(rng.NormFloat64() * 1e10)
		}
		b := a
		opuscc.CompareCeltPLCLagWindow(&a)
		nativeCeltPLCLagWindow(&b)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("lag window", trial, i, a[i], b[i])
			}
		}
	}
}
func TestCeltPLCDecayAgainstC(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, loss := range []int32{0, 1, 99} {
			for _, alias := range []int{0, 1, 2} {
				a := make([]float32, 66)
				b := make([]float32, 66)
				for i := range a {
					a[i] = float32(i) - 22
					b[i] = float32(i%5) - 12
				}
				a[0], a[65] = 77, 88
				want := slices.Clone(a)
				wb := slices.Clone(b)
				gp, cp := &b[1], &wb[1]
				if alias == 1 {
					gp, cp = &a[1], &want[1]
				}
				if alias == 2 {
					gp, cp = &a[2], &want[2]
				}
				opuscc.CompareCeltPLCDecay(&a[1], gp, 21, 1, 20, channels, loss)
				nativeCeltPLCDecay(&want[1], cp, 21, 1, 20, channels, loss)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
						t.Fatal("PLC decay", channels, loss, alias, i)
					}
				}
			}
		}
	}
}
func TestPLCPitchSearchAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1316))
	for _, channels := range []int32{0, 1, 2, 3} {
		for trial := 0; trial < 80; trial++ {
			left := make([]float32, opuscc.DEC_PITCH_BUF_SIZE)
			right := make([]float32, opuscc.DEC_PITCH_BUF_SIZE)
			for i := range left {
				left[i] = float32(rng.NormFloat64())
				right[i] = float32(rng.NormFloat64())
			}
			if trial == 0 {
				clear(left)
				clear(right)
			}
			if trial%3 == 1 {
				for i := range left {
					left[i] = float32(math.Sin(float64(i) * 0.17))
					right[i] = float32(math.Sin(float64(i) * 0.13))
				}
			}
			beforeL := slices.Clone(left)
			beforeR := slices.Clone(right)
			var rp *float32
			if channels == 2 {
				rp = &right[0]
			}
			g := opuscc.ComparePLCPitchSearch(&left[0], rp, channels)
			c := nativePLCPitchSearch(left, right, channels)
			if g != c || !sameFloatBits(left, beforeL) || !sameFloatBits(right, beforeR) {
				t.Fatal(channels, trial, g, c)
			}
		}
	}
}

func TestRemoveDoublingAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1215))
	for _, maxPeriod := range []int{16, 31, 64, 128} {
		for _, minPeriod := range []int{4, 5, 8} {
			for _, n := range []int{8, 17, 64, 240} {
				for trial := 0; trial < 100; trial++ {
					x := make([]float32, maxPeriod/2+n/2)
					for i := range x {
						x[i] = float32(rng.NormFloat64())
					}
					if trial == 0 {
						clear(x)
					}
					if trial%3 == 1 {
						for i := range x {
							x[i] = float32(math.Sin(float64(i) * 0.3))
						}
					}
					before := append([]float32(nil), x...)
					initial := minPeriod + rng.Intn(maxPeriod-minPeriod+8)
					previous := rng.Intn(maxPeriod)
					previousGain := float32(rng.Float64())
					g := [3]int32{77, int32(initial), 88}
					c := g
					gg := opuscc.Opus_remove_doubling(nil, &x[0], int32(maxPeriod), int32(minPeriod), int32(n), &g[1], int32(previous), previousGain, 0)
					cg := nativeRemoveDoubling(x, int32(maxPeriod), int32(minPeriod), int32(n), &c[1], int32(previous), previousGain)
					if g != c || math.Float32bits(gg) != math.Float32bits(cg) || !sameFloatBits(x, before) {
						t.Fatal(maxPeriod, minPeriod, n, trial, initial, previous, g, c, gg, cg)
					}
				}
			}
		}
	}
}

func TestPitchSearchAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1114))
	for _, n := range []int{4, 8, 12, 17, 32, 128, 240} {
		for _, maxPitch := range []int{4, 8, 12, 16, 31, 128} {
			if maxPitch>>2 > 3 && n>>2 < 3 {
				continue
			}
			for trial := 0; trial < 60; trial++ {
				x := make([]float32, n/2)
				y := make([]float32, (n+maxPitch)/2)
				for i := range x {
					x[i] = float32(rng.NormFloat64())
				}
				for i := range y {
					y[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					clear(x)
					clear(y)
				}
				if trial == 1 {
					for i := range x {
						x[i] = float32(math.Sin(float64(i) * 0.7))
					}
					for i := range y {
						y[i] = float32(math.Sin(float64(i) * 0.7))
					}
				}
				x0 := append([]float32(nil), x...)
				y0 := append([]float32(nil), y...)
				g := [3]int32{77, -1, 88}
				c := g
				opuscc.Opus_pitch_search(nil, &x[0], &y[0], int32(n), int32(maxPitch), &g[1], 0)
				nativePitchSearch(x, y, int32(n), int32(maxPitch), &c[1])
				if g != c || !sameFloatBits(x, x0) || !sameFloatBits(y, y0) {
					t.Fatal(n, maxPitch, trial, g, c)
				}
			}
		}
	}
}

func TestPitchDownsampleAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(1013))
	for _, n := range []int{7, 8, 17, 64, 240} {
		for _, factor := range []int{1, 2, 3, 4} {
			for _, channels := range []int{0, 1, 2, 3} {
				for trial := 0; trial < 12; trial++ {
					left := make([]float32, n*factor)
					right := make([]float32, n*factor)
					for i := range left {
						left[i] = float32(rng.NormFloat64())
						right[i] = float32(rng.NormFloat64())
					}
					if trial == 0 {
						clear(left)
						clear(right)
					}
					l0 := append([]float32(nil), left...)
					r0 := append([]float32(nil), right...)
					g := make([]float32, n+2)
					for i := range g {
						g[i] = 77
					}
					c := append([]float32(nil), g...)
					var rp *float32
					if channels == 2 {
						rp = &right[0]
					}
					opuscc.Opus_pitch_downsample(nil, &left[0], rp, &g[1], int32(n), int32(channels), int32(factor), 0)
					nativePitchDownsample(left, right, c[1:], int32(n), int32(channels), int32(factor))
					if !sameFloatBits(g, c) || !sameFloatBits(left, l0) || !sameFloatBits(right, r0) {
						t.Fatal(n, factor, channels, trial, g, c)
					}
				}
			}
		}
	}
}

func TestAutocorrAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(912))
	for _, n := range []int{1, 4, 5, 8, 17, 64, 128} {
		for _, lag := range []int{0, 1, 3, 4, 8, 24} {
			// The four-lag scalar C kernel requires at least three samples.
			if lag >= n || (lag >= 3 && n-lag < 3) {
				continue
			}
			for _, overlap := range []int{0, 1, n / 2, n} {
				for trial := 0; trial < 12; trial++ {
					input := make([]float32, n)
					window := make([]float32, overlap)
					for i := range input {
						input[i] = float32(rng.NormFloat64() * 3)
					}
					for i := range window {
						window[i] = float32(rng.Float64())
					}
					if trial == 0 {
						for i := range input {
							input[i] = math.Float32frombits(uint32(i%2) << 31)
						}
					}
					if trial == 1 {
						for i := range input {
							input[i] = math.Float32frombits(uint32(i + 1))
						}
					}
					before := append([]float32(nil), input...)
					weights := append([]float32(nil), window...)
					goOut := make([]float32, lag+3)
					for i := range goOut {
						goOut[i] = 77
					}
					cOut := append([]float32(nil), goOut...)
					r := opuscc.Opus__celt_autocorr(nil, &input[0], &goOut[1], unsafe.SliceData(window), int32(overlap), int32(lag), int32(n), 0)
					c := nativeAutocorr(input, cOut[1:], window, int32(overlap), int32(lag), int32(n))
					if r != c || !sameFloatBits(goOut, cOut) || !sameFloatBits(input, before) || !sameFloatBits(window, weights) {
						t.Fatal(n, lag, overlap, trial, goOut, cOut)
					}
				}
			}
		}
	}
	// Input/output aliasing has C's correlation-then-tail update order.
	for _, overlap := range []int{0, 4} {
		g := make([]float32, 32)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		window := []float32{0.2, 0.4, 0.6, 0.8}
		opuscc.Opus__celt_autocorr(nil, &g[0], &g[1], &window[0], int32(overlap), 4, 32, 0)
		nativeAutocorr(c, c[1:], window, int32(overlap), 4, 32)
		if !sameFloatBits(g, c) {
			t.Fatal("alias", overlap, g, c)
		}
	}
}

func TestIIRAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(811))
	for _, ord := range []int{4, 8, 24} {
		for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 31, 64} {
			for trial := 0; trial < 12; trial++ {
				input := make([]float32, n+2)
				coeff := make([]float32, ord)
				mem := make([]float32, ord+2)
				for i := range input {
					input[i] = float32(rng.NormFloat64() * 3)
				}
				for i := range coeff {
					coeff[i] = float32(rng.NormFloat64() * 0.02)
				}
				for i := range mem {
					mem[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i%2) << 31)
					}
				}
				if trial == 1 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i + 1))
					}
				}
				before := append([]float32(nil), input...)
				goOut := make([]float32, ord+n+2)
				for i := range goOut {
					goOut[i] = 77
				}
				cOut := append([]float32(nil), goOut...)
				cMem := append([]float32(nil), mem...)
				opuscc.Opus_celt_iir(nil, &input[0], &coeff[0], &goOut[ord], int32(n), int32(ord), &mem[1], 0)
				nativeIIR(input, coeff, cOut[ord:], cMem[1:], int32(n), int32(ord))
				if !sameFloatBits(goOut, cOut) || !sameFloatBits(mem, cMem) || !sameFloatBits(input, before) {
					t.Fatal(ord, n, trial, goOut, cOut, mem, cMem)
				}
			}
		}
	}
	// In-place and partial overlaps preserve block read-ahead and store order.
	for _, offset := range []int{3, 4, 5} {
		g := make([]float32, 40)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		coeff := []float32{0.1, 0.2, 0.3, 0.4}
		gm := []float32{1, 2, 3, 4}
		cm := append([]float32(nil), gm...)
		opuscc.Opus_celt_iir(nil, &g[4], &coeff[0], &g[offset], 24, 4, &gm[0], 0)
		nativeIIR(c[4:], coeff, c[offset:], cm, 24, 4)
		if !sameFloatBits(g, c) || !sameFloatBits(gm, cm) {
			t.Fatal("overlap", offset, g, c)
		}
	}
}

func TestFIRAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(710))
	for _, ord := range []int{3, 4, 5, 7, 8, 24} {
		for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 31, 64} {
			for trial := 0; trial < 12; trial++ {
				input := make([]float32, ord+n+2)
				coeff := make([]float32, ord)
				for i := range input {
					input[i] = float32(rng.NormFloat64() * 3)
				}
				for i := range coeff {
					coeff[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i%2) << 31)
					}
				}
				if trial == 1 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i + 1))
					}
				}
				before := append([]float32(nil), input...)
				goOut := make([]float32, n+2)
				for i := range goOut {
					goOut[i] = 77
				}
				cOut := append([]float32(nil), goOut...)
				opuscc.Opus_celt_fir_c(nil, &input[ord], &coeff[0], &goOut[1], int32(n), int32(ord), 0)
				nativeFIR(input, coeff, cOut[1:], int32(n), int32(ord))
				if !sameFloatBits(goOut, cOut) || !sameFloatBits(input, before) {
					t.Fatal(ord, n, trial, goOut, cOut)
				}
			}
		}
	}
	// Partial overlaps are permitted (only exact x==y is asserted against).
	for _, offset := range []int{3, 5} {
		g := make([]float32, 40)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		coeff := []float32{0.1, 0.2, 0.3, 0.4}
		opuscc.Opus_celt_fir_c(nil, &g[4], &coeff[0], &g[offset], 24, 4, 0)
		nativeFIR(c, coeff, c[offset:], 24, 4)
		if !sameFloatBits(g, c) {
			t.Fatal("overlap", offset, g, c)
		}
	}
}

func TestLPCAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, p := range []int{1, 4, 16, 24} {
		for trial := 0; trial < 100; trial++ {
			signal := make([]float32, 128)
			for i := range signal {
				signal[i] = float32(rng.NormFloat64())
			}
			ac := make([]float32, p+1)
			for lag := range ac {
				for i := lag; i < len(signal); i++ {
					ac[lag] += float32(signal[i] * signal[i-lag])
				}
			}
			if trial == 0 {
				clear(ac)
			}
			if trial == 1 {
				for i := range ac {
					ac[i] = 1
				}
			}
			g, c := make([]float32, p), make([]float32, p)
			opuscc.Opus__celt_lpc(nil, &g[0], &ac[0], int32(p))
			nativeLPC(c, ac)
			for i := range g {
				if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
					t.Fatalf("p=%d trial=%d tap=%d Go=%g C=%g", p, trial, i, g[i], c[i])
				}
			}
		}
	}
}
