package opuscc

import (
	"runtime"
	"testing"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func newDecodeFrameTestState(rate, nb int32) *OpusT_silk_decoder_state {
	d := new(OpusT_silk_decoder_state)
	Opus_silk_init_decoder(nil, d)
	d.Fnb_subfr = nb
	Opus_silk_decoder_set_fs(nil, d, rate, rate*1000)
	d.FpsNLSF_CB = cloneTestNLSFCodebook(d.FpsNLSF_CB)
	d.FVAD_flags[0] = 1
	d.FLBRR_flags[0] = 1
	return d
}
func TestDecodeFramePulseScratchPointers(t *testing.T) {
	for _, length := range []int32{1, 16, 17, 80, 120, 160, 240, 320} {
		pulses := silkDecodeFramePulses(length)
		entropyInitGrowStack(12)
		runtime.GC()
		if int32(len(pulses)) != (length+15)&^15 {
			t.Fatal("pulse padding", length, len(pulses))
		}
		for _, v := range pulses {
			if v != 0 {
				t.Fatal("scratch zero")
			}
		}
	}
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, flag := range []int32{0, 2} {
				for _, cond := range []int32{CODE_INDEPENDENTLY, CODE_CONDITIONALLY} {
					d := newDecodeFrameTestState(rate, nb)
					buf := make([]byte, 900)
					for i := range buf {
						buf[i] = byte(i*71 + 13)
					}
					var ec OpusT_ec_ctx
					Opus_ec_dec_init(nil, &ec, &buf[0], uint32(len(buf)))
					frame := make([]int16, d.Fframe_length+2)
					frame[0], frame[len(frame)-1] = 77, 88
					before := append([]byte(nil), buf...)
					count := int32(-99)
					entropyInitGrowStack(12)
					runtime.GC()
					ret := silk_decode_frame(nil, d, &ec, &frame[1], &count, flag, cond, 0)
					if ret != 0 || count != d.Fframe_length || d.FlossCnt != 0 || frame[0] != 77 || frame[len(frame)-1] != 88 || d.FpsNLSF_CB.FCB1_NLSF_Q8 == nil || ec.Fbuf != &buf[0] {
						t.Fatal("typed normal/fec frame", rate, nb, flag, cond, ret, count)
					}
					for i := range buf {
						if buf[i] != before[i] {
							t.Fatal("frame wrote entropy input")
						}
					}
				}
			}
		}
	}
	d := newDecodeFrameTestState(12, 2)
	buf := make([]byte, 900)
	ec := new(OpusT_ec_ctx)
	Opus_ec_dec_init(nil, ec, &buf[0], uint32(len(buf)))
	frame := make([]int16, d.Fframe_length)
	var count int32
	tls := libc.NewTLS()
	defer tls.Close()
	libc.Xpthread_setspecific(tls, 0x6f707573, 123)
	silk_decode_frame(tls, d, ec, &frame[0], &count, 0, CODE_INDEPENDENTLY, 0)
	if libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
		t.Fatal("normal frame touched TLS")
	}
	// Length validation precedes entropy/PCM/count consumption.
	d.Fframe_length = 0
	count = 123
	panicked := false
	func() {
		defer func() { panicked = recover() != nil }()
		silk_decode_frame(nil, d, nil, nil, &count, 0, 0, 0)
	}()
	if !panicked || count != 123 {
		t.Fatal("frame validation/count ordering", panicked, count)
	}
}

func TestDecodeFrameLossPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, flag := range []int32{1, 2, -1, 7} {
				d := newPLCConcealTestDecoder(rate, nb)
				frame := make([]int16, d.Fframe_length+2)
				frame[0], frame[len(frame)-1] = 77, 88
				count := int32(-99)
				entropyInitGrowStack(12)
				runtime.GC()
				ret := silk_decode_frame(nil, d, nil, &frame[1], &count, flag, CODE_INDEPENDENTLY, 0)
				if ret != 0 || count != d.Fframe_length || d.FlossCnt != 1 || frame[0] != 77 || frame[len(frame)-1] != 88 || d.FpsNLSF_CB.FCB1_NLSF_Q8 == nil {
					t.Fatal("loss frame", rate, nb, flag, ret, count, d.FlossCnt)
				}
			}
		}
	}
	d := newPLCConcealTestDecoder(8, 2)
	tls := libc.NewTLS()
	defer tls.Close()
	libc.Xpthread_setspecific(tls, 0x6f707573, 123)
	frame := make([]int16, d.Fframe_length)
	var count int32
	silk_decode_frame(tls, d, nil, &frame[0], &count, 1, CODE_INDEPENDENTLY, 0)
	if libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
		t.Fatal("loss frame touched TLS")
	}
}

func TestDecodeFrameFinishPointers(t *testing.T) {
	for _, nb := range []int32{2, 4} {
		for _, alias := range []int32{0, 1, 2} {
			d := &OpusT_silk_decoder_state{Fnb_subfr: nb, FlagPrev: 17}
			c := &OpusT_silk_decoder_control{FpitchL: [4]int32{41, 42, 43, 44}}
			count := int32(-77)
			p := &count
			if alias == 1 {
				p = &d.FlagPrev
			} else if alias == 2 {
				p = &c.FpitchL[nb-1]
			}
			entropyInitGrowStack(12)
			runtime.GC()
			silkDecodeFrameFinish(d, c, p, 160)
			lag := int32(40 + nb)
			if alias == 1 {
				lag = 160
			}
			if *p != 160 || d.FlagPrev != lag || alias == 2 && c.FpitchL[nb-1] != 160 {
				t.Fatal("lag/count ordering", nb, alias, d.FlagPrev, *p)
			}
		}
	}
}

func TestDecodeFrameHistoryPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			d := newPLCConcealTestDecoder(rate, nb)
			w := *d
			w.FpsNLSF_CB = nil
			frame := make([]int16, d.Fframe_length)
			for i := range frame {
				frame[i] = int16(i*71 - 9000)
			}
			move := d.Fltp_mem_length - d.Fframe_length
			copy(w.FoutBuf[:move], w.FoutBuf[d.Fframe_length:d.Fframe_length+move])
			copy(w.FoutBuf[move:d.Fltp_mem_length], frame)
			entropyInitGrowStack(12)
			runtime.GC()
			silkDecodeFrameHistory(d, &frame[0])
			g := *d
			g.FpsNLSF_CB = nil
			if g != w || d.FpsNLSF_CB.FCB1_NLSF_Q8 == nil {
				t.Fatal("frame history", rate, nb)
			}
		}
	}
	// Go-only overlapping memcpy source: shift must happen before frame reload.
	d := newPLCConcealTestDecoder(8, 2)
	w := d.FoutBuf
	move := d.Fltp_mem_length - d.Fframe_length
	copy(w[:move], w[d.Fframe_length:d.Fltp_mem_length])
	copy(w[move:d.Fltp_mem_length], w[:d.Fframe_length])
	silkDecodeFrameHistory(d, &d.FoutBuf[0])
	if d.FoutBuf != w {
		t.Fatal("history clear/read alias")
	}
}

func TestDecodeFrameFieldAccesses(t *testing.T) {
	var tls *libc.TLS
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder, 8, 8000); got != OPUS_OK {
		t.Fatalf("set fs: got %d", got)
	}
	Opus_silk_PLC_Reset(tls, &decoder)
	decoder.FsPLC.FprevGain_Q16 = [2]OpusT_opus_int32{65536, 65536}
	decoder.FsPLC.FpitchL_Q8 = 20 << 8
	decoder.FsPLC.FLTPCoef_Q14 = [5]OpusT_opus_int16{300, -150, 1200, -100, 75}
	decoder.FsPLC.FprevLPC_Q12 = [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}
	decoder.FsPLC.FprevLTP_scale_Q14 = 13000
	decoder.FsPLC.FrandScale_Q14 = 11000
	decoder.FsPLC.Frand_seed = 12345
	for i := range decoder.Fexc_Q14 {
		decoder.Fexc_Q14[i] = int32((i*71)%3000 - 1500)
	}
	for i := range decoder.FoutBuf {
		decoder.FoutBuf[i] = int16((i*37)%1000 - 500)
	}
	output := make([]int16, decoder.Fframe_length)
	var samples int32
	if got := silk_decode_frame(tls, &decoder, nil, &output[0], &samples, 1, CODE_INDEPENDENTLY, 0); got != OPUS_OK {
		t.Fatalf("decode result: got %d", got)
	}
	if got, want := samples, int32(160); got != want {
		t.Fatalf("sample count: got %d, want %d", got, want)
	}
	if got, want := output[:8], []int16{20, 32, -38, -29, -31, -28, -25, -22}; !equalInt16s(got, want) {
		t.Fatalf("output prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.FoutBuf[decoder.Fltp_mem_length-8:decoder.Fltp_mem_length], []int16{0, 1, 1, 1, 1, 1, 2, 2}; !equalInt16s(got, want) {
		t.Fatalf("output history suffix: got %v, want %v", got, want)
	}
}
