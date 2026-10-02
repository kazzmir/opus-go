//go:build compareopus && cgo

package main

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestDecodeCoreTransitionAgainstC(t *testing.T) {
	for _, loss := range []int32{0, 1, -1} {
		for _, prev := range []int32{1, 2} {
			for _, signal := range []int8{0, 1, 2} {
				for k := int32(0); k < 4; k++ {
					d := opuscc.OpusT_silk_decoder_state{FlossCnt: loss, FprevSignalType: prev, FlagPrev: 77}
					d.Findices.FsignalType = signal
					ctrl := opuscc.OpusT_silk_decoder_control{FpitchL: [4]int32{101, 102, 103, 104}}
					for i := range ctrl.FLTPCoef_Q14 {
						ctrl.FLTPCoef_Q14[i] = int16(i*71 - 900)
					}
					c, cc := d, ctrl
					g := opuscc.CompareDecodeCoreTransition(&d, &ctrl, k)
					n := nativeDecodeCoreTransition(&c, &cc, k)
					if g != n || d != c || ctrl != cc {
						t.Fatal("transition", loss, prev, signal, k, g, n)
					}
				}
			}
		}
	}
}
func TestDecodeCoreHistoryAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		d := opuscc.OpusT_silk_decoder_state{Fltp_mem_length: rate * 20, Fsubfr_length: rate * 5}
		for i := range d.FoutBuf {
			d.FoutBuf[i] = int16(i*3 - 900)
		}
		c := d
		frame := make([]int16, 2*d.Fsubfr_length)
		for i := range frame {
			frame[i] = int16(i*31 - 900)
		}
		opuscc.CompareDecodeCoreHistory(&d, frame)
		nativeDecodeCoreHistory(&c, frame)
		if d != c {
			t.Fatal("history", rate)
		}
	}
}
func TestDecodeCoreAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, signal := range []int8{0, 1, 2} {
				for _, interp := range []int8{0, 4} {
					for _, loss := range []int32{0, 1} {
						order := int32(10)
						if rate == 16 {
							order = 16
						}
						d := opuscc.OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, Fltp_mem_length: rate * 20, FLPC_order: order, Fprev_gain_Q16: 65536, FprevSignalType: 2, FlagPrev: rate * 5, FlossCnt: loss}
						d.Findices.FsignalType = signal
						d.Findices.FquantOffsetType = 1
						d.Findices.FNLSFInterpCoef_Q2 = interp
						d.Findices.FSeed = 17
						for i := range d.FoutBuf {
							d.FoutBuf[i] = int16((i*37)%1000 - 500)
						}
						for i := range d.FsLPC_Q14_buf {
							d.FsLPC_Q14_buf[i] = int32(i*97 - 800)
						}
						ctrl := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{65536, 72000, 68000, 76000}, FLTP_scale_Q14: 12288}
						for k := range ctrl.FPredCoef_Q12 {
							ctrl.FPredCoef_Q12[k] = [16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
						}
						for k := int32(0); k < nb; k++ {
							ctrl.FpitchL[k] = rate * 5
							copy(ctrl.FLTPCoef_Q14[k*5:k*5+5], []int16{300, -150, 1200, -100, 75})
						}
						c, cc := d, ctrl
						pulses := make([]int16, d.Fframe_length)
						for i := range pulses {
							pulses[i] = int16((i*7)%9 - 4)
						}
						before := slices.Clone(pulses)
						out := make([]int16, d.Fframe_length+2)
						out[0], out[len(out)-1] = 77, 88
						want := slices.Clone(out)
						tls := libc.NewTLS()
						ps := libc.XmallocPointer(tls, uint64(unsafe.Sizeof(opuscc.OpusT_opus_ccgo_pseudostack_state{})))
						scratch := libc.XmallocPointer(tls, opuscc.GLOBAL_STACK_SIZE)
						*(*opuscc.OpusT_opus_ccgo_pseudostack_state)(ps) = opuscc.OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: uintptr(scratch), Fglobal_stack: uintptr(scratch)}
						libc.Xpthread_setspecific(tls, 0x6f707573, uintptr(ps))
						opuscc.CompareDecodeCore(tls, &d, &ctrl, uintptr(unsafe.Pointer(&out[1])), uintptr(unsafe.Pointer(&pulses[0])))
						tls.Close()
						r := nativeDecodeCore(&c, &cc, want[1:len(want)-1], pulses)
						if r != 0 || d != c || ctrl != cc || !slices.Equal(out, want) || !slices.Equal(pulses, before) {
							t.Fatal("decode core", rate, nb, signal, interp, loss, r, d.Fprev_gain_Q16, c.Fprev_gain_Q16)
						}
					}
				}
			}
		}
	}
}
