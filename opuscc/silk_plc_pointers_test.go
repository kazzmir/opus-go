package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"runtime"
	"slices"
	"testing"
	"unsafe"
	"weak"
)

// From plc_energy_pointers_test.go.
func TestPLCEnergyPointers(t *testing.T) {
	for _, n := range []int32{20, 40, 60, 80} {
		for _, nb := range []int32{2, 4} {
			exc := make([]int32, n*nb)
			for i := range exc {
				exc[i] = 1 << 14
			}
			for i := int32(0); i < (nb-2)*n; i++ {
				exc[i] = 1 << 30
			}
			before := slices.Clone(exc)
			gains := [2]int32{1024, 2048}
			out := [6]int32{123, 0, 0, 0, 0, 456}
			silk_PLC_energy(nil, &out[1], &out[2], &out[3], &out[4], &exc[0], &gains, n, nb)
			if out != [6]int32{123, n, 0, 4 * n, 0, 456} || !slices.Equal(exc, before) || gains != [2]int32{1024, 2048} {
				t.Fatalf("n=%d nb=%d out=%v", n, nb, out)
			}
			var aliased int32
			silk_PLC_energy(nil, &aliased, &aliased, &aliased, &aliased, &exc[0], &gains, n, nb)
			if aliased != 4*n {
				t.Fatalf("aliased output=%d want=%d", aliased, 4*n)
			}
			for i := range exc {
				exc[i] = 1 << 28
			}
			gains = [2]int32{1 << 20, 1 << 20}
			silk_PLC_energy(nil, &out[1], &out[2], &out[3], &out[4], &exc[0], &gains, n, nb)
			if out != [6]int32{123, 0, 0, 0, 0, 456} {
				t.Fatal("SMULWW did not narrow before shifting")
			}
		}
	}
}

// From plc_glue_pointers_test.go.
func TestPLCGluePointers(t *testing.T) {
	var d OpusT_silk_decoder_state
	Opus_silk_PLC_glue_frames(nil, &d, nil, 0)
	var lost [16]int16
	for i := range lost {
		lost[i] = 100
	}
	d.FlossCnt = 1
	Opus_silk_PLC_glue_frames(nil, &d, &lost[0], 16)
	if d.FsPLC.Fconc_energy != 160000 || d.FsPLC.Flast_frame_lost != 1 {
		t.Fatal(d.FsPLC)
	}
	d.FlossCnt = 0
	good := [18]int16{0: 77, 17: 88}
	for i := 1; i < 17; i++ {
		good[i] = 1000
	}
	Opus_silk_PLC_glue_frames(nil, &d, &good[1], 16)
	if good[0] != 77 || good[17] != 88 || good[1] >= 1000 || good[16] != 1000 || d.FsPLC.Flast_frame_lost != 0 {
		t.Fatal(d.FsPLC, good)
	}
}

// From plc_update_pointers_test.go.
func newPLCConcealTestDecoder(rate, nb int32) *OpusT_silk_decoder_state {
	order := int32(10)
	cb := &Opus_silk_NLSF_CB_NB_MB
	if rate == 16 {
		order = 16
		cb = &Opus_silk_NLSF_CB_WB
	}
	d := &OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, Fltp_mem_length: rate * 20, FLPC_order: order, FprevSignalType: TYPE_VOICED, FpsNLSF_CB: cloneTestNLSFCodebook(cb)}
	d.FsPLC = OpusT_silk_PLC_struct{Ffs_kHz: rate, FpitchL_Q8: rate * 5 << 8, FLTPCoef_Q14: [5]int16{300, -150, 1200, -100, 75}, FprevLPC_Q12: [16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}, FprevGain_Q16: [2]int32{98304, 131072}, FprevLTP_scale_Q14: 13000, FrandScale_Q14: 11000, Frand_seed: 12345, Fsubfr_length: rate * 5, Fnb_subfr: nb}
	for i := range d.Fexc_Q14 {
		d.Fexc_Q14[i] = int32((i*71)%3000000 - 1500000)
	}
	for i := range d.FoutBuf {
		d.FoutBuf[i] = int16((i*37)%1000 - 500)
	}
	return d
}
func TestPLCConcealScratchPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, signal := range []int32{TYPE_UNVOICED, TYPE_VOICED} {
				for _, loss := range []int32{0, 1, 3} {
					for _, lost := range []int32{1, -1} {
						for _, reset := range []bool{false, true} {
							d := newPLCConcealTestDecoder(rate, nb)
							d.FprevSignalType = signal
							d.FlossCnt = loss
							if reset {
								d.FsPLC.Ffs_kHz = 0
								d.Ffirst_frame_after_reset = 1
							}
							w := *d
							w.FpsNLSF_CB = nil
							control, expected := OpusT_silk_decoder_control{}, OpusT_silk_decoder_control{}
							frame := make([]int16, d.Fframe_length+2)
							frame[0], frame[len(frame)-1] = 77, 88
							want := make([]int16, d.Fframe_length)
							entropyInitGrowStack(12)
							runtime.GC()
							silk_PLC(nil, d, &control, &frame[1], lost, 0)
							silk_PLC(nil, &w, &expected, &want[0], lost, 0)
							g := *d
							g.FpsNLSF_CB = nil
							if g != w || control != expected || d.FlossCnt != loss+1 || frame[0] != 77 || frame[len(frame)-1] != 88 || d.FpsNLSF_CB == nil {
								t.Fatal("typed complete concealment", rate, nb, signal, loss, lost, reset)
							}
							for i := range want {
								if frame[i+1] != want[i] {
									t.Fatal("concealment PCM", rate, nb, signal, loss, i)
								}
							}
						}
					}
				}
			}
		}
	}
	d := newPLCConcealTestDecoder(16, 4)
	control := OpusT_silk_decoder_control{}
	tls := libc.NewTLS()
	defer tls.Close()
	libc.Xpthread_setspecific(tls, 0x6f707573, 123)
	frame := make([]int16, d.Fframe_length)
	silk_PLC(tls, d, &control, &frame[0], 1, 0)
	if libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
		t.Fatal("concealment touched TLS scratch")
	}
	// Failed rewhitening leaves reset/bandwidth expansion committed, but cannot
	// increment lossCnt or overwrite PCM/control. This is Go-only error ordering.
	d = newPLCConcealTestDecoder(8, 2)
	d.FsPLC.Ffs_kHz = 0
	d.Fltp_mem_length = 16
	frame = make([]int16, d.Fframe_length)
	for i := range frame {
		frame[i] = 77
	}
	before := control
	panicked := false
	func() { defer func() { panicked = recover() != nil }(); silk_PLC(nil, d, &control, &frame[0], 1, 0) }()
	if !panicked || d.FlossCnt != 0 || control != before || d.FsPLC.Ffs_kHz != 8 || d.FsPLC.FprevLPC_Q12[0] == 120 {
		t.Fatal("failed concealment ordering", panicked, d.FlossCnt)
	}
	for _, v := range frame {
		if v != 77 {
			t.Fatal("failed concealment PCM")
		}
	}
}
func TestPLCConcealHistoryAliasPointers(t *testing.T) {
	d := newPLCConcealTestDecoder(16, 4)
	w := *d
	control, expected := OpusT_silk_decoder_control{}, OpusT_silk_decoder_control{}
	frame := make([]int16, d.Fframe_length)
	silk_PLC(nil, &w, &expected, &frame[0], 1, 0)
	entropyInitGrowStack(12)
	runtime.GC()
	silk_PLC(nil, d, &control, &d.FoutBuf[0], 1, 0)
	copy(w.FoutBuf[:], frame)
	if *d != w || control != expected {
		t.Fatal("whitening-before-aliased-output order")
	}
}

func TestPLCSynthesisScratchPointers(t *testing.T) {
	for _, shape := range [][2]int32{{160, 80}, {160, 160}, {240, 120}, {240, 240}, {320, 160}, {320, 320}, {320, 0}} {
		d := &OpusT_silk_decoder_state{Fltp_mem_length: shape[0], Fframe_length: shape[1], FLPC_order: 16}
		buffer := silkPLCSynthesisBuffer(d)
		entropyInitGrowStack(12)
		runtime.GC()
		if len(buffer) != int(shape[0]+shape[1]) {
			t.Fatal("synthesis geometry", shape, len(buffer))
		}
		for _, v := range buffer {
			if v != 0 {
				t.Fatal("owned scratch not zero")
			}
		}
		for i := range buffer {
			buffer[i] = int32(i * 1003)
		}
		index := int32(len(buffer) - 1)
		coefficients := [5]int16{0, 0, 16384, 0, 0}
		if p := silkPLCLTPPrediction(buffer, index, &coefficients); p != 2+(buffer[index-2]>>2) {
			t.Fatal("owned scratch prediction", shape, p)
		}
		frame := make([]int16, shape[1])
		silkPLCLPC(nil, d, buffer[shape[0]-16:], new([16]int16), unsafe.SliceData(frame), 1024)
		for i := range d.FsLPC_Q14_buf {
			if d.FsLPC_Q14_buf[i] != buffer[int(shape[0]+shape[1])-16+i] {
				t.Fatal("owned scratch history", shape, i)
			}
		}
	}
}

func TestPLCLTPHistoryPointers(t *testing.T) {
	h := []int32{-2147483648, 2147483647, -1, 0, 100000003, 12345, -54321, 2147483647}
	b := &[5]int16{-32768, 32767, -1, 16384, 12345}
	entropyInitGrowStack(12)
	runtime.GC()
	for index := int32(4); index < int32(len(h)); index++ {
		want := int32(2)
		for j := int32(0); j < 5; j++ {
			want = int32(int64(want) + (int64(h[index-j]) * int64(b[j]) >> 16))
		}
		if got := silkPLCLTPPrediction(h, index, b); got != want {
			t.Fatal("LTP history", index, got, want)
		}
	}
	b[0] = 0
	h[4] = 1
	if got := silkPLCLTPPrediction(h, 4, b); got != int32(int64(2)+(int64(h[3])*int64(b[1])>>16)+(int64(h[2])*int64(b[2])>>16)+(int64(h[1])*int64(b[3])>>16)+(int64(h[0])*int64(b[4])>>16)) {
		t.Fatal("live taps/history", got)
	}
}

func TestPLCWhiteningPointers(t *testing.T) {
	for _, length := range []int32{160, 240, 320} {
		for _, order := range []int32{10, 16} {
			for _, index := range []int32{1, 17, length - order - 1} {
				d := &OpusT_silk_decoder_state{Fltp_mem_length: length, FLPC_order: order, FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)}
				for i := range d.FoutBuf {
					d.FoutBuf[i] = int16((i*71)%60000 - 30000)
				}
				a := new([16]int16)
				for i := range a {
					a[i] = int16(i*137 - 700)
				}
				samples := make([]int16, length+2)
				for i := range samples {
					samples[i] = 123
				}
				samples[0], samples[len(samples)-1] = 77, 88
				want := append([]int16(nil), samples...)
				Opus_silk_LPC_analysis_filter(nil, &want[1+index], &d.FoutBuf[index], &a[0], length-index, order, 0)
				before := *d
				entropyInitGrowStack(12)
				runtime.GC()
				silkPLCWhiten(nil, d, samples[1:len(samples)-1], a, index, 0)
				for i := range samples {
					if samples[i] != want[i] {
						t.Fatal("whitening cursor", length, order, index, i, samples[i], want[i])
					}
				}
				if *d != before {
					t.Fatal("whitening changed decoder")
				}
			}
		}
	}
}

func TestPLCLPCHistoryPointers(t *testing.T) {
	for _, length := range []int32{0, 1, 10, 16, 17, 80, 320} {
		for _, order := range []int32{10, 16} {
			d := &OpusT_silk_decoder_state{Fframe_length: length, FLPC_order: order, FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)}
			for i := range d.FsLPC_Q14_buf {
				d.FsLPC_Q14_buf[i] = int32(i*1000003) - 8000000
			}
			a := new([MAX_LPC_ORDER]int16)
			for i := range a {
				a[i] = int16(i*137 - 700)
			}
			history := make([]int32, 18+length)
			history[0], history[len(history)-1] = 77, 88
			for i := int32(0); i < length; i++ {
				history[17+i] = int32(int64(i) * 1000000003)
			}
			want := append([]int32(nil), history...)
			copy(want[1:17], d.FsLPC_Q14_buf[:])
			wf := make([]int16, length)
			for i := int32(0); i < length; i++ {
				pred := order >> 1
				for j := int32(0); j < order; j++ {
					pred = int32(int64(pred) + (int64(want[17+i-j-1]) * int64(a[j]) >> 16))
				}
				p := int64(int32(uint32(min(max(pred, -2147483648>>4), 2147483647>>4)) << 4))
				want[17+i] = int32(min(max(int64(want[17+i])+p, int64(-2147483648)), int64(2147483647)))
				wf[i] = int16(min(max(((int32(int64(want[17+i])*1024>>16)>>7)+1)>>1, -32768), 32767))
			}
			pcm := make([]int16, length+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			silkPLCLPC(nil, d, history[1:len(history)-1], a, unsafe.SliceData(pcm[1:len(pcm)-1]), 1024)
			for i := range history {
				if history[i] != want[i] {
					t.Fatal("LPC working history", length, order, i)
				}
			}
			for i := range wf {
				if pcm[i+1] != wf[i] {
					t.Fatal("LPC PCM", length, order, i)
				}
			}
			for i, v := range d.FsLPC_Q14_buf {
				if v != want[1+int(length)+i] {
					t.Fatal("history saved after PCM", length, order, i)
				}
			}
			if pcm[0] != 77 || pcm[len(pcm)-1] != 88 || d.FpsNLSF_CB.Forder != 16 {
				t.Fatal("history guard/owner")
			}
		}
	}
	d := OpusT_silk_decoder_state{FLPC_order: 8, Fframe_length: 1, FsLPC_Q14_buf: [16]int32{123}}
	h := make([]int32, 17)
	panicked := false
	func() {
		defer func() { panicked = recover() != nil }()
		silkPLCLPC(nil, &d, h, new([16]int16), nil, 1024)
	}()
	if !panicked || h[0] != 123 || d.FsLPC_Q14_buf[0] != 123 {
		t.Fatal("copy-before-LPC assertion order")
	}
	for _, f := range []struct{ x, p, w int32 }{{2147483647, 1, 2147483647}, {-2147483648, -1, -2147483648}, {0, 2147483647, 2147483632}, {0, -2147483648, -2147483648}} {
		if got := silkPLCAddPrediction(f.x, f.p); got != f.w {
			t.Fatal("saturating prediction", f, got)
		}
	}
}

// This alias is Go-only: short PCM into int32 state violates C effective types.
func TestPLCLPCStatePCMOrderPointers(t *testing.T) {
	d := OpusT_silk_decoder_state{FLPC_order: 16, Fframe_length: 16}
	h := make([]int32, 32)
	for i := 0; i < 16; i++ {
		d.FsLPC_Q14_buf[i] = int32(i + 77)
		h[16+i] = int32(i-8) << 16
	}
	silkPLCLPC(nil, &d, h, new([16]int16), (*int16)(unsafe.Pointer(&d.FsLPC_Q14_buf[0])), 1024)
	for i := 0; i < 16; i++ {
		want := (int32(i-8) << 16) + 128
		if h[16+i] != want || d.FsLPC_Q14_buf[i] != want {
			t.Fatal("PCM-before-history-save order", i, h[16+i], d.FsLPC_Q14_buf[i], want)
		}
	}
}

func TestPLCConcealPCMPointers(t *testing.T) {
	pcm := make([]int16, 3)
	pcm[0], pcm[2] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	for _, sample := range []int32{-2147483648, -8388609, -8388608, -128, -1, 0, 127, 128, 8388352, 8388608, 2147483647} {
		for _, gain := range []int32{-2147483648, -65536, -1, 0, 1, 1024, 65536, 2147483647} {
			scaled := int32(int64(sample) * int64(gain) >> 16)
			want := int16(min(max(((scaled>>7)+1)>>1, -32768), 32767))
			silkPLCPCMStore(pcm[1:2], 0, sample, gain)
			if pcm[1] != want || pcm[0] != 77 || pcm[2] != 88 {
				t.Fatal("PLC PCM", sample, gain, pcm, want)
			}
		}
	}
	for _, f := range []struct {
		s, g int32
		want int16
	}{{128, 65536, 1}, {-128, 65536, 0}, {-129, 65536, -1}, {8388608, 65536, 32767}, {-8388609, 65536, -32768}, {2147483647, 2147483647, -256}} {
		if got := silkPLCPCM(f.s, f.g); got != f.want {
			t.Fatal("PCM narrowing/rounding/saturation", f, got)
		}
	}
}

func TestPLCRandomPointers(t *testing.T) {
	random, owner := func() ([]int32, weak.Pointer[OpusT_silk_NLSF_CB_struct]) {
		cb := cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)
		d := &OpusT_silk_decoder_state{FpsNLSF_CB: cb}
		for i := range d.Fexc_Q14 {
			d.Fexc_Q14[i] = int32(i*1000003) - 160000000
		}
		return silkPLCRandom(d, 192), weak.Make(cb)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if len(random) != 128 || owner.Value() == nil {
		t.Fatal("random interior owner")
	}
	for i := int32(0); i < 128; i++ {
		for _, scale := range []int16{-32768, -1, 0, 16384, 32767} {
			for _, prediction := range []int32{-2147483648, -1, 0, 2147483647} {
				want := int32(uint32(int32(int64(prediction)+(int64(int32((i+192)*1000003)-160000000)*int64(scale)>>16))) << 2)
				if got := silkPLCNoise(prediction, random, i, scale); got != want {
					t.Fatal("random Q14", i, scale, prediction, got, want)
				}
			}
		}
	}
	random[0] = 1 << 20
	if got := silkPLCNoise(2, random, 0, 16384); got != (2+(1<<18))<<2 {
		t.Fatal("live random load", got)
	}
}

func TestPLCLTPCoefficientPointers(t *testing.T) {
	b, owner := func() (*[LTP_ORDER]int16, weak.Pointer[OpusT_silk_NLSF_CB_struct]) {
		cb := cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)
		d := &OpusT_silk_decoder_state{FpsNLSF_CB: cb}
		d.FsPLC.FLTPCoef_Q14 = [5]int16{-32768, 32767, -1, 0, 12345}
		return &d.FsPLC.FLTPCoef_Q14, weak.Make(cb)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if owner.Value() == nil {
		t.Fatal("coefficient interior owner")
	}
	for _, gain := range []int32{-32768, -32767, 0, 16384, 32767, 32768, 65535} {
		before := *b
		silkPLCDecayLTP(b, gain)
		for j := range b {
			want := int16(int32(int16(gain)) * int32(before[j]) >> 15)
			if b[j] != want {
				t.Fatal("LTP decay", gain, j, b[j], want)
			}
		}
	}
	runtime.KeepAlive(b)
}

func TestPLCConcealStateOwnersPointers(t *testing.T) {
	plc, owner := func() (*OpusT_silk_PLC_struct, weak.Pointer[OpusT_silk_NLSF_CB_struct]) {
		cb := cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)
		d := &OpusT_silk_decoder_state{FpsNLSF_CB: cb, FsPLC: OpusT_silk_PLC_struct{Frand_seed: 123, FprevGain_Q16: [2]int32{65536, 131072}}}
		return silkPLCConcealState(d), weak.Make(cb)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	runtime.GC()
	cb := owner.Value()
	if cb == nil || cb.Forder != 16 || plc.Frand_seed != 123 || plc.FprevGain_Q16 != [2]int32{65536, 131072} {
		t.Fatal("sole PLC interior lost scanned decoder owner")
	}
	plc.FpitchL_Q8 = 77
	runtime.KeepAlive(plc)
}

func TestPLCDispatchFramePointers(t *testing.T) {
	d := OpusT_silk_decoder_state{Ffs_kHz: 16, Fframe_length: 320, Fsubfr_length: 80, Fnb_subfr: 4, FLPC_order: 16}
	d.Findices.FsignalType = TYPE_UNVOICED
	c := OpusT_silk_decoder_control{FGains_Q16: [4]int32{1, 2, 3, 4}}
	frame := make([]int16, 322)
	for i := range frame {
		frame[i] = int16(i*17 - 300)
	}
	before := append([]int16(nil), frame...)
	entropyInitGrowStack(12)
	runtime.GC()
	silk_PLC(nil, &d, &c, &frame[1], 0, 0)
	for i := range frame {
		if frame[i] != before[i] {
			t.Fatal("nonloss touched PCM", i)
		}
	}
	// Frame aliases are unconsumed on nonloss: even a numeric state/control
	// interior must not be read or written, and nil is equally valid.
	wd, wc := d, c
	silk_PLC(nil, &d, &c, (*int16)(unsafe.Pointer(&c.FpitchL[0])), 0, 0)
	silk_PLC(nil, &wd, &wc, nil, 0, 0)
	if d != wd || c != wc {
		t.Fatal("unused frame alias")
	}
}

func TestPLCDispatchControlPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, order := range []int32{10, 16} {
				for _, signal := range []int8{0, 1, 2} {
					d := &OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, FLPC_order: order, FlossCnt: 3}
					d.Findices.FsignalType = signal
					c := &OpusT_silk_decoder_control{FGains_Q16: [4]int32{1, 2, 3, 4}, FLTP_scale_Q14: 8192, FpitchL: [4]int32{40, 80, 120, 160}}
					for i := range c.FPredCoef_Q12[1] {
						c.FPredCoef_Q12[1][i] = int16(i * 13)
					}
					for i := range c.FLTPCoef_Q14 {
						c.FLTPCoef_Q14[i] = int16(i * 17)
					}
					w, wc := *d, *c
					silkPLCRate(nil, &w)
					silk_PLC_update(nil, &w, &wc)
					entropyInitGrowStack(12)
					runtime.GC()
					silk_PLC(nil, d, c, nil, 0, -37)
					if *d != w || *c != wc || d.FlossCnt != 3 {
						t.Fatal("typed PLC update", rate, nb, order, signal)
					}
				}
			}
		}
	}
	// Go-only invalid-control fixture: rate reset and the preceding update stores
	// are committed before control is first dereferenced.
	d := OpusT_silk_decoder_state{Ffs_kHz: 16, Fframe_length: 320, Fnb_subfr: 4, Fsubfr_length: 80, FLPC_order: 16}
	d.Findices.FsignalType = TYPE_UNVOICED
	panicked := false
	func() { defer func() { panicked = recover() != nil }(); silk_PLC(nil, &d, nil, nil, 0, 0) }()
	if !panicked || d.FsPLC.Ffs_kHz != 16 || d.FsPLC.FpitchL_Q8 != 16*18*256 || d.FsPLC.FprevGain_Q16 != [2]int32{65536, 65536} || d.FprevSignalType != TYPE_UNVOICED {
		t.Fatal("reset/update-before-control order", panicked, d.FsPLC)
	}
}

func TestPLCDispatchStatePointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		d := func() *OpusT_silk_decoder_state {
			return &OpusT_silk_decoder_state{Fframe_length: rate * 20, Ffs_kHz: rate, FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB), FlossCnt: 7, FsPLC: OpusT_silk_PLC_struct{Ffs_kHz: 0, Frand_seed: 123}}
		}()
		entropyInitGrowStack(12)
		runtime.GC()
		silkPLCRate(nil, d)
		want := d.FsPLC
		if want.Ffs_kHz != rate || want.FpitchL_Q8 != rate*20*128 || want.FprevGain_Q16 != [2]int32{65536, 65536} || want.Fsubfr_length != 20 || want.Fnb_subfr != 2 || want.Frand_seed != 123 || d.FlossCnt != 7 || d.FpsNLSF_CB.Forder != 16 {
			t.Fatal("PLC state owner/reset", rate, want)
		}
		d.FsPLC.FpitchL_Q8 = 77
		silkPLCRate(nil, d)
		if d.FsPLC.FpitchL_Q8 != 77 {
			t.Fatal("matching rate reset")
		}
	}
}

func TestPLCUpdatePointers(t *testing.T) {
	d := OpusT_silk_decoder_state{Ffs_kHz: 16, Fnb_subfr: 4, Fsubfr_length: 80, FLPC_order: 10}
	d.Findices.FsignalType = TYPE_UNVOICED
	for i := range d.FsPLC.FprevLPC_Q12 {
		d.FsPLC.FprevLPC_Q12[i] = 77
	}
	c := OpusT_silk_decoder_control{FGains_Q16: [4]int32{1, 2, 3, 4}, FLTP_scale_Q14: 8192}
	for i := range c.FPredCoef_Q12[1] {
		c.FPredCoef_Q12[1][i] = int16(i)
	}
	before := c
	silk_PLC_update(nil, &d, &c)
	if c != before || d.FprevSignalType != TYPE_UNVOICED || d.FsPLC.FpitchL_Q8 != 16*18*256 || d.FsPLC.FprevGain_Q16 != [2]int32{3, 4} {
		t.Fatal(d.FsPLC)
	}
	for i, v := range d.FsPLC.FprevLPC_Q12 {
		want := int16(77)
		if i < 10 {
			want = int16(i)
		}
		if v != want {
			t.Fatal(i, v, want)
		}
	}
}

// From cng_poly_pointers_test.go.
func TestCNGExcPointers(t *testing.T) {
	for _, tc := range []struct{ length, mask int }{
		{1, 1}, {2, 1}, {3, 3}, {7, 7}, {8, 7}, {15, 15},
		{31, 31}, {32, 31}, {127, 127}, {128, 127}, {255, 255}, {320, 255},
	} {
		buffer := make([]int32, tc.mask+1)
		for i := range buffer {
			buffer[i] = int32(i*7919 - 123456)
		}
		for _, initial := range []int32{0, 3176576, -1, -2147483648} {
			out := make([]int32, tc.length+2)
			out[0], out[len(out)-1] = 111, 222
			seed := initial
			// Verify continued seed updates across successive calls too.
			wantSeed := uint64(uint32(initial))
			for pass := 0; pass < 2; pass++ {
				silk_CNG_exc(nil, &out[1], &buffer[0], int32(tc.length), &seed)
				for i := 0; i < tc.length; i++ {
					wantSeed = (907633515 + wantSeed*196314165) % (1 << 32)
					idx := int(wantSeed/(1<<24)) & tc.mask
					if out[i+1] != buffer[idx] {
						t.Fatalf("length=%d seed=%d pass=%d sample=%d: got %d want %d", tc.length, initial, pass, i, out[i+1], buffer[idx])
					}
				}
				if seed != int32(uint32(wantSeed)) || out[0] != 111 || out[len(out)-1] != 222 {
					t.Fatalf("length=%d: seed or sentinels changed incorrectly", tc.length)
				}
			}
		}
	}
	seed := int32(12345)
	silk_CNG_exc(nil, nil, nil, 0, &seed)
	if seed != 12345 {
		t.Fatal("empty excitation changed seed")
	}
}

func TestNLSFPolyPointers(t *testing.T) {
	input := [16]int32{131000, -130123, 65537, -32767, 1, 0, 12345, -45678, -99123, 54321, 721, 15, -4321, 23456, -120000, 99999}
	// Golden vectors use C's signed fixed-point arithmetic, rounding each
	// product as floor((product + 32768) / 65536), including negative values.
	for _, tc := range []struct {
		order, offset int
		want          []int32
	}{
		{1, 0, []int32{65536, -131000}},
		{1, 1, []int32{65536, 130123}},
		{5, 0, []int32{65536, -109760, 179772, -209579, 174309, -199637}},
		{5, 1, []int32{65536, 154247, 333397, 514303, 634924, 720112}},
		{8, 0, []int32{65536, 13840, 175919, 36343, 125602, -59340, -36320, -262802, -103077}},
		{8, 1, []int32{65536, 30777, 275223, 186278, 499590, 515209, 549405, 829676, 519002}},
	} {
		// Exact-size input is important for odd-offset interleaved access.
		lsf := append([]int32(nil), input[tc.offset:2*tc.order]...)
		out := make([]int32, tc.order+3)
		out[0], out[len(out)-1] = 111, 222
		silk_NLSF2A_find_poly(nil, &out[1], &lsf[0], int32(tc.order))
		for i, want := range tc.want {
			if out[i+1] != want {
				t.Fatalf("order=%d offset=%d coefficient=%d: got %d want %d", tc.order, tc.offset, i, out[i+1], want)
			}
		}
		if out[0] != 111 || out[len(out)-1] != 222 {
			t.Fatal("output sentinel overwritten")
		}
		for i, value := range lsf {
			if value != input[i+tc.offset] {
				t.Fatal("input modified")
			}
		}
	}
}

func TestDecoderSizePointer(t *testing.T) {
	out := [3]int32{111, -1, 222}
	if ret := Opus_silk_Get_Decoder_Size(nil, &out[1]); ret != SILK_NO_ERROR {
		t.Fatalf("return code: %d", ret)
	}
	if out != [3]int32{111, int32(unsafe.Sizeof(OpusT_silk_decoder{})), 222} {
		t.Fatalf("size/sentinels: %v", out)
	}
}
