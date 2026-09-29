package opuscc

import "testing"

func TestCWRSPointers(t *testing.T) {
	for n := int32(2); n <= 6; n++ {
		for k := int32(1); k <= 6; k++ {
			count := celtPVQU(min(n, k), max(n, k)) + celtPVQU(min(n, k+1), max(n, k+1))
			for index := uint32(0); index < count; index++ {
				out := make([]int32, n+2)
				out[0] = 123
				out[n+1] = 456
				got := cwrsi(nil, n, k, index, &out[1])
				var total int32
				var energy float32
				for _, v := range out[1 : n+1] {
					total += max(v, -v)
					energy += float32(v) * float32(v)
				}
				if total != k || got != energy || out[0] != 123 || out[n+1] != 456 {
					t.Fatalf("n=%d k=%d index=%d out=%v energy=%g", n, k, index, out, got)
				}
			}
		}
	}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	out := [4]int32{123, 0, 0, 456}
	if energy := Opus_decode_pulses(nil, &out[1], 2, 1, &dec); energy != 1 || out[0] != 123 || out[3] != 456 {
		t.Fatalf("energy=%g out=%v", energy, out)
	}
}
