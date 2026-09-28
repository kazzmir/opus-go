package opuscc

import "testing"

func TestPulseDecodePointers(t *testing.T) {
	for _, length := range []int32{16, 80, 120, 160, 240, 320} {
		for signal := int32(0); signal <= 2; signal++ {
			for offset := int32(0); offset <= 1; offset++ {
				n := ((length + 15) / 16) * 16
				out := make([]int16, n+2)
				for i := range out {
					out[i] = 123
				}
				dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 7}
				Opus_silk_decode_pulses(nil, &dec, &out[1], signal, offset, length)
				if out[0] != 123 || out[n+1] != 123 || dec.Fext != 123 || dec.Fend_window != 456 || dec.Fnend_bits != 7 {
					t.Fatal("guard or unrelated entropy fields changed")
				}
				for _, v := range out[1 : n+1] {
					if v != 0 {
						t.Fatalf("zero packet produced nonzero pulse: %d", v)
					}
				}
			}
		}
	}
}
