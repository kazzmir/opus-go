package opuscc

import (
	"slices"
	"testing"
)

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
