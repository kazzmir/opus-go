//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
)

func TestPLCUpdateAgainstC(t *testing.T) {
	for _, nb := range []int32{2, 4} {
		for _, order := range []int32{10, 16} {
			for signal := int8(0); signal <= 2; signal++ {
				for _, gain := range []int16{-100, 0, 500, 3000, 5000, 15000} {
					g := opuscc.OpusT_silk_decoder_state{Ffs_kHz: 16, Fsubfr_length: 80, Fnb_subfr: nb, FLPC_order: order, FprevSignalType: 7}
					g.Findices.FsignalType = signal
					g.FsPLC.FpitchL_Q8 = 12345
					g.FsPLC.Frand_seed = 777
					for i := range g.FsPLC.FprevLPC_Q12 {
						g.FsPLC.FprevLPC_Q12[i] = 123
					}
					ctrl := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{65536, 123456, 2147483647, 76543}, FLTP_scale_Q14: 23456, FpitchL: [4]int32{100, 180, 250, 300}}
					for i := range ctrl.FLTPCoef_Q14 {
						ctrl.FLTPCoef_Q14[i] = gain + int16(i%5)
					}
					for i := range ctrl.FPredCoef_Q12[1] {
						ctrl.FPredCoef_Q12[1][i] = int16(i * 137)
					}
					c := g
					before := ctrl
					opuscc.ComparePLCUpdate(&g, &ctrl)
					nativePLCUpdate(&c, &ctrl)
					if g != c || ctrl != before {
						t.Fatal(nb, order, signal, gain, g.FsPLC, c.FsPLC)
					}
				}
			}
		}
	}
}
