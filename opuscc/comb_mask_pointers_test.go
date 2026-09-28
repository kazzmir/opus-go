package opuscc

import "testing"

func TestCombFilterPointers(t *testing.T) {
	for _, period := range []int{15, 32} {
		for _, n := range []int{1, 17, 64} {
			for _, inPlace := range []bool{false, true} {
				offset := period + 2
				x := make([]float32, offset+n)
				for i := range x {
					x[i] = float32((i*7)%31-15) * 0.125
				}
				wantX := append([]float32(nil), x...)
				out := make([]float32, n+2)
				out[0], out[n+1] = 99, 88
				got := out[1 : n+1]
				if inPlace {
					got = x[offset:]
				}
				want := make([]float32, n)
				for i := range want {
					p := offset + i
					want[i] = wantX[p] + float32(0.5*wantX[p-period]) + float32(0.25*(wantX[p-period+1]+wantX[p-period-1])) + float32(-0.125*(wantX[p-period+2]+wantX[p-period-2]))
					if inPlace {
						wantX[p] = want[i]
					}
				}
				comb_filter_const_c(nil, &got[0], &x[0], int32(period), int32(n), 0.5, 0.25, -0.125)
				for i := range want {
					if got[i] != want[i] {
						t.Fatalf("period=%d n=%d inplace=%t i=%d: got %g want %g", period, n, inPlace, i, got[i], want[i])
					}
				}
				if out[0] != 99 || out[n+1] != 88 {
					t.Fatal("output sentinel changed")
				}
				for i := 0; i < offset; i++ {
					if x[i] != wantX[i] {
						t.Fatal("history changed")
					}
				}
			}
		}
	}
	comb_filter_const_c(nil, nil, nil, 15, 0, 0.5, 0.25, 0.125)
}

func TestCollapseMaskPointers(t *testing.T) {
	for _, blocks := range []int{2, 4, 8, 16, 32} {
		for _, width := range []int{1, 3} {
			x := make([]int32, blocks*width)
			if got := extract_collapse_mask(nil, &x[0], int32(len(x)), int32(blocks)); got != 0 {
				t.Fatal("zero mask")
			}
			for b := 0; b < blocks; b++ {
				x[b*width+width-1] = -2147483648
				if got := extract_collapse_mask(nil, &x[0], int32(len(x)), int32(blocks)); got != uint32(1)<<b {
					t.Fatalf("blocks=%d width=%d active=%d got=%x", blocks, width, b, got)
				}
				x[b*width+width-1] = 0
			}
		}
	}
	if got := extract_collapse_mask(nil, nil, 0, 1); got != 1 {
		t.Fatal("single block special case")
	}
	// Integer division ignores a trailing remainder just as C does.
	x := [5]int32{1, -1, 0, 0, 99}
	if got := extract_collapse_mask(nil, &x[0], 5, 2); got != 1 {
		t.Fatalf("remainder/cancellation: %x", got)
	}
}
