package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecodeCoreLTPStoragePointers(t *testing.T) {
	for _, shape := range [][3]int32{{160, 160, 240}, {240, 360, 480}, {320, 640, 640}} {
		for _, lag := range []int32{0, 1, 40, 80} {
			for _, gain := range []int32{-2147483648, -65536, 0, 65536, 2147483647} {
				samples := make([]int16, shape[0])
				for i := range samples {
					samples[i] = int16(i*137 - 20000)
				}
				h := make([]int32, shape[2]+2)
				for i := range h {
					h[i] = int32(i*17000003 - 2100000000)
				}
				want := append([]int32(nil), h...)
				for i := int32(0); i < lag+2; i++ {
					want[1+shape[1]-i-1] = int32(int64(gain) * int64(samples[shape[0]-i-1]) >> 16)
				}
				entropyInitGrowStack(12)
				runtime.GC()
				silkDecodeCoreLTPWhiten(h[1:len(h)-1], samples, shape[1], shape[0], lag, gain)
				if !equalInt32s(h, want) {
					t.Fatal("LTP rewhitening", shape, lag, gain)
				}
				for i := int32(0); i < lag+2; i++ {
					index := 1 + shape[1] - i - 1
					want[index] = int32(int64(gain) * int64(want[index]) >> 16)
				}
				silkDecodeCoreLTPScale(h[1:len(h)-1], shape[1], lag, gain)
				if !equalInt32s(h, want) {
					t.Fatal("LTP scaling", shape, lag, gain)
				}
			}
		}
	}
	h := []int32{-2147483648, 2147483647, -1, 0, 100000003, 12345, -54321, 2147483647}
	b := &[5]int16{-32768, 32767, -1, 16384, 12345}
	entropyInitGrowStack(12)
	runtime.GC()
	for index := int32(4); index < int32(len(h)); index++ {
		want := int32(2)
		for j := int32(0); j < 5; j++ {
			want = int32(int64(want) + (int64(h[index-j]) * int64(b[j]) >> 16))
		}
		if got := silkDecodeCoreLTPPrediction(h, index, b); got != want {
			t.Fatal("Q13 MAC narrowing", index, got, want)
		}
	}
}

func TestDecodeCoreWhiteningPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, order := range []int32{10, 16} {
			for _, k := range []int32{0, 2} {
				d := newPLCConcealTestDecoder(rate, 4)
				d.FLPC_order = order
				for i := range d.FoutBuf {
					d.FoutBuf[i] = int16(i*71 - 20000)
				}
				for _, start := range []int32{1, 17, d.Fltp_mem_length - order - 1} {
					a := &[16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
					samples := make([]int16, d.Fltp_mem_length+2)
					for i := range samples {
						samples[i] = 123
					}
					samples[0], samples[len(samples)-1] = 77, 88
					want := append([]int16(nil), samples...)
					Opus_silk_LPC_analysis_filter(nil, &want[start+1], &d.FoutBuf[start+k*d.Fsubfr_length], &a[0], d.Fltp_mem_length-start, order, 0)
					before := *d
					entropyInitGrowStack(12)
					runtime.GC()
					silkDecodeCoreWhiten(nil, d, samples[1:len(samples)-1], a, start, k, 0)
					for i := range samples {
						if samples[i] != want[i] {
							t.Fatal("whitening", rate, order, k, start, i)
						}
					}
					if *d != before {
						t.Fatal("whitening changed decoder")
					}
				}
			}
		}
	}
}

func TestDecodeCoreCoefficientPointers(t *testing.T) {
	for _, order := range []int32{0, 10, 16} {
		for k := int32(0); k < 4; k++ {
			d := &OpusT_silk_decoder_state{FLPC_order: order}
			c := new(OpusT_silk_decoder_control)
			for row := range c.FPredCoef_Q12 {
				for i := range c.FPredCoef_Q12[row] {
					c.FPredCoef_Q12[row][i] = int16(row*1000 + i*71 - 900)
				}
			}
			for i := range c.FLTPCoef_Q14 {
				c.FLTPCoef_Q14[i] = int16(i*31 - 400)
			}
			before := *c
			var snapshot [16]int16
			for i := range snapshot {
				snapshot[i] = 123
			}
			want := snapshot
			copy(want[:order], c.FPredCoef_Q12[k>>1][:order])
			entropyInitGrowStack(12)
			runtime.GC()
			a, b := silkDecodeCoreCoefficients(d, c, k, &snapshot)
			if a != &c.FPredCoef_Q12[k>>1] || b[0] != c.FLTPCoef_Q14[k*5] || snapshot != want || *c != before {
				t.Fatal("coefficient views/snapshot", order, k)
			}
			c.FPredCoef_Q12[k>>1][0] = 999
			c.FLTPCoef_Q14[k*5] = 888
			if a[0] != 999 || b[0] != 888 || snapshot != want {
				t.Fatal("live coefficients vs snapshot", order, k)
			}
		}
	}
	a, b, owner := func() (*[16]int16, *[5]int16, weak.Pointer[OpusT_silk_decoder_control]) {
		c := new(OpusT_silk_decoder_control)
		c.FPredCoef_Q12[1][0] = 77
		c.FLTPCoef_Q14[15] = 88
		var snapshot [16]int16
		a, b := silkDecodeCoreCoefficients(&OpusT_silk_decoder_state{FLPC_order: 16}, c, 3, &snapshot)
		return a, b, weak.Make(c)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if owner.Value() == nil || a[0] != 77 || b[0] != 88 {
		t.Fatal("sole coefficient interiors lost control")
	}
	a[1] = 99
	b[1] = 100
	if owner.Value().FPredCoef_Q12[1][1] != 99 || owner.Value().FLTPCoef_Q14[16] != 100 {
		t.Fatal("live coefficient stores")
	}
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	// Go-only snapshot overlap. Rewhitening retains the original live A, even
	// when the snapshot writes overlap its source.
	c := new(OpusT_silk_decoder_control)
	for i := range c.FPredCoef_Q12[0] {
		c.FPredCoef_Q12[0][i] = int16(i + 1)
	}
	want := c.FPredCoef_Q12
	flat := unsafe.Slice(&want[0][0], 32)
	copy(flat[1:11], flat[:10])
	a, _ = silkDecodeCoreCoefficients(&OpusT_silk_decoder_state{FLPC_order: 10}, c, 0, (*[16]int16)(unsafe.Pointer(&c.FPredCoef_Q12[0][1])))
	if c.FPredCoef_Q12 != want || a != &c.FPredCoef_Q12[0] {
		t.Fatal("snapshot alias ordering")
	}
}

func TestDecodeCoreExcitationPointers(t *testing.T) {
	for _, length := range []int32{0, 1, 17, 80, 160, 320} {
		for _, seed := range []int8{-128, -1, 0, 17, 127} {
			d := newPLCConcealTestDecoder(16, 4)
			d.Fframe_length = length
			d.Findices.FSeed = seed
			for i := range d.Fexc_Q14 {
				d.Fexc_Q14[i] = 1234567
			}
			w := *d
			w.FpsNLSF_CB = nil
			pulses := make([]int16, length)
			edges := []int16{-32768, 32767, -1, 0, 1, 13, -14}
			s := int32(seed)
			for i := range pulses {
				pulses[i] = edges[i%len(edges)]
				s = int32(uint32(RAND_INCREMENT) + uint32(s)*uint32(RAND_MULTIPLIER))
				v := int32(uint32(int32(pulses[i])) << 14)
				if v > 0 {
					v -= QUANT_LEVEL_ADJUST_Q10 << 4
				} else if v < 0 {
					v += QUANT_LEVEL_ADJUST_Q10 << 4
				}
				v += 1600
				if s < 0 {
					v = -v
				}
				w.Fexc_Q14[i] = v
				s = int32(uint32(s) + uint32(int32(pulses[i])))
			}
			entropyInitGrowStack(12)
			runtime.GC()
			got := silkDecodeCoreExcitation(d, unsafe.SliceData(pulses), 100)
			g := *d
			g.FpsNLSF_CB = nil
			if got != s || g != w || d.FpsNLSF_CB.FCB1_NLSF_Q8 == nil {
				t.Fatal("excitation/ownership", length, seed, got, s)
			}
		}
	}
	pulse, owner := func() (*int16, weak.Pointer[OpusT_silk_decoder_state]) {
		source := newPLCConcealTestDecoder(16, 4)
		return &source.FoutBuf[1], weak.Make(source)
	}()
	target := &OpusT_silk_decoder_state{Fframe_length: 320}
	entropyInitGrowStack(12)
	runtime.GC()
	silkDecodeCoreExcitation(target, pulse, 100)
	source := owner.Value()
	if source == nil || source.FpsNLSF_CB == nil || source.FpsNLSF_CB.FCB1_NLSF_Q8 == nil {
		t.Fatal("sole pulse interior lost scanned backing")
	}
	runtime.KeepAlive(pulse)
	// Go-only effective-type alias. The second pulse load must follow excitation
	// writes, rather than use a cached original pulse.
	d := &OpusT_silk_decoder_state{Fframe_length: 1}
	d.Fexc_Q14[0] = 7
	seed := int32(uint32(RAND_INCREMENT))
	sample := int32(7<<14) - (QUANT_LEVEL_ADJUST_Q10 << 4) + 1600
	if seed < 0 {
		sample = -sample
	}
	want := int32(uint32(seed) + uint32(int32(int16(sample))))
	got := silkDecodeCoreExcitation(d, (*int16)(unsafe.Pointer(&d.Fexc_Q14[0])), 100)
	if got != want || d.Fexc_Q14[0] != sample {
		t.Fatal("live pulse reload", got, want, d.Fexc_Q14[0], sample)
	}
}

func TestDecodeCoreControlPointers(t *testing.T) {
	for _, loss := range []int32{0, 1, -1} {
		for _, prev := range []int32{TYPE_UNVOICED, TYPE_VOICED} {
			for _, signal := range []int8{0, TYPE_UNVOICED, TYPE_VOICED} {
				for k := int32(0); k < 4; k++ {
					d := &OpusT_silk_decoder_state{FlossCnt: loss, FprevSignalType: prev, FlagPrev: 77}
					d.Findices.FsignalType = signal
					control := new(OpusT_silk_decoder_control)
					for i := range control.FLTPCoef_Q14 {
						control.FLTPCoef_Q14[i] = int16(i*71 - 900)
					}
					control.FpitchL = [4]int32{101, 102, 103, 104}
					want := *control
					before := *d
					active := loss != 0 && prev == TYPE_VOICED && signal != TYPE_VOICED && k < 2
					if active {
						clear(want.FLTPCoef_Q14[k*5 : k*5+5])
						want.FLTPCoef_Q14[k*5+2] = 4096
						want.FpitchL[k] = 77
					}
					entropyInitGrowStack(12)
					runtime.GC()
					if got := silkDecodeCoreTransition(d, control, k); got != active || *control != want || *d != before {
						t.Fatal("transition control", loss, prev, signal, k, got)
					}
				}
			}
		}
	}
	for _, d := range []*OpusT_silk_decoder_state{{}, {FlossCnt: 1}, {FlossCnt: 1, FprevSignalType: TYPE_VOICED, Findices: OpusT_SideInfoIndices{FsignalType: TYPE_VOICED}}} {
		if silkDecodeCoreTransition(d, nil, 0) {
			t.Fatal("unused control")
		}
	}
	d := &OpusT_silk_decoder_state{FlossCnt: 1, FprevSignalType: TYPE_VOICED}
	if silkDecodeCoreTransition(d, nil, 2) {
		t.Fatal("unused late control")
	}
}

func TestDecodeCoreHistoryPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, length := range []int32{0, rate * 5} {
			d := newPLCConcealTestDecoder(rate, 4)
			d.Fsubfr_length = length
			before := *d
			before.FpsNLSF_CB = nil
			owner := weak.Make(d)
			frame := make([]int16, 2*length)
			for i := range frame {
				frame[i] = int16(i*31 - 900)
			}
			entropyInitGrowStack(12)
			runtime.GC()
			silkDecodeCoreHistory(d, unsafe.SliceData(frame))
			copy(before.FoutBuf[d.Fltp_mem_length:d.Fltp_mem_length+2*length], frame)
			g := *d
			g.FpsNLSF_CB = nil
			if g != before || owner.Value() == nil || d.FpsNLSF_CB.FCB1_NLSF_Q8 == nil {
				t.Fatal("history staging/ownership", rate, length)
			}
		}
	}
	// Go-only overlapping copy: C memcpy does not define this alias.
	d := newPLCConcealTestDecoder(8, 4)
	w := d.FoutBuf
	copy(w[d.Fltp_mem_length:d.Fltp_mem_length+2*d.Fsubfr_length], w[d.Fltp_mem_length-1:d.Fltp_mem_length+2*d.Fsubfr_length-1])
	silkDecodeCoreHistory(d, &d.FoutBuf[d.Fltp_mem_length-1])
	if d.FoutBuf != w {
		t.Fatal("history overlap order")
	}
}

func TestDecodeCoreFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder, 8, 8000); got != OPUS_OK {
		t.Fatalf("set fs: got %d", got)
	}
	decoder.Fprev_gain_Q16 = 65536
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FquantOffsetType = 1
	decoder.Findices.FSeed = 17
	var control OpusT_silk_decoder_control
	control.FGains_Q16 = [4]OpusT_opus_int32{65536, 72000, 68000, 76000}
	for i := range control.FPredCoef_Q12 {
		control.FPredCoef_Q12[i] = [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}
	}
	pulses := make([]int16, decoder.Fframe_length)
	for i := range pulses {
		pulses[i] = int16((i*7)%9 - 4)
	}
	output := make([]int16, decoder.Fframe_length)
	Opus_silk_decode_core(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&pulses[0])), 0)
	/* expected values verified against the C reference implementation (silk/decode_core.c) */
	if got, want := output[:8], []int16{4, 3, 1, 1, -3, -4, 2, 0}; !equalInt16s(got, want) {
		t.Fatalf("output prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.Fexc_Q14[:8], []int32{60416, 51712, 18944, 11264, -44032, -68096, 35328, 3840}; !equalInt32s(got, want) {
		t.Fatalf("excitation prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.Fprev_gain_Q16, int32(76000); got != want {
		t.Fatalf("previous gain: got %d, want %d", got, want)
	}
}

func TestDecodeCoreLocalLPCArray(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder, 16, 16000); got != OPUS_OK {
		t.Fatalf("set fs: got %d", got)
	}
	decoder.Fprev_gain_Q16 = 65536
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FquantOffsetType = 1
	decoder.Findices.FSeed = 17
	var control OpusT_silk_decoder_control
	control.FGains_Q16 = [4]OpusT_opus_int32{65536, 72000, 68000, 76000}
	for i := range control.FPredCoef_Q12 {
		control.FPredCoef_Q12[i] = [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
	}
	pulses := make([]int16, decoder.Fframe_length)
	for i := range pulses {
		pulses[i] = int16((i*7)%9 - 4)
	}
	output := make([]int16, decoder.Fframe_length)
	Opus_silk_decode_core(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&pulses[0])), 0)

	/* expected values from the C reference implementation (silk/decode_core.c) */
	if got, want := output[:16], []int16{4, 3, 1, 1, -3, -4, 2, 0, 2, -4, -3, -1, -1, -3, -4, -2}; !equalInt16s(got, want) {
		t.Fatalf("output prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.Fexc_Q14[:8], []int32{60416, 51712, 18944, 11264, -44032, -68096, 35328, 3840}; !equalInt32s(got, want) {
		t.Fatalf("excitation prefix: got %v, want %v", got, want)
	}
	wantSLPC := []int32{-3536, 27216, 61472, -50592, -21152, 12160, -44400, -69136, 34048, -2080, -28944, 60576, -50080, 16336, 13584, 42736}
	if got := decoder.FsLPC_Q14_buf[:]; !equalInt32s(got, wantSLPC) {
		t.Fatalf("LPC state: got %v, want %v", got, wantSLPC)
	}
	if got, want := decoder.Fprev_gain_Q16, int32(76000); got != want {
		t.Fatalf("previous gain: got %d, want %d", got, want)
	}
}

func equalInt32s(got, want []int32) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
