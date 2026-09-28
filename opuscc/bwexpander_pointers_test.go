package opuscc

import "testing"

func TestBWExpanderPointers(t *testing.T) {
	for _, n := range []int{1, 10, 16} {
		for _, chirp := range []int32{0, 32768, 64881, 65535, 65536} {
			a16, want16 := make([]int16, n+2), make([]int16, n+2)
			a32, want32 := make([]int32, n+2), make([]int32, n+2)
			for i := range a16 {
				a16[i] = int16(32767 - i*5000)
				a32[i] = int32(2147483647 - int64(i)*300000000)
			}
			copy(want16, a16)
			copy(want32, a32)
			c := chirp
			for i := 1; i <= n; i++ {
				product := int32(int64(c) * int64(want16[i]))
				want16[i] = int16((int64(product) + 32768) >> 16)
				want32[i] = int32((int64(c) * int64(want32[i])) >> 16)
				product = int32(int64(c) * int64(chirp-65536))
				c += int32((int64(product) + 32768) >> 16)
			}
			Opus_silk_bwexpander(nil, &a16[1], int32(n), chirp)
			Opus_silk_bwexpander_32(nil, &a32[1], int32(n), chirp)
			for i := range a16 {
				if a16[i] != want16[i] || a32[i] != want32[i] {
					t.Fatalf("n=%d chirp=%d index=%d: got (%d,%d), want (%d,%d)", n, chirp, i, a16[i], a32[i], want16[i], want32[i])
				}
			}
		}
	}
}
