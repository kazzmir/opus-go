package opuscc

import "testing"

func TestFFTFactorPointers(t *testing.T) {
	for _, n := range []int32{1, 2, 3, 4, 5, 7, 16, 60, 120, 480, 1024, 65536, 1 << 30, 2147483647} {
		guard := struct {
			before  int32
			factors [64]int32
			after   int32
		}{before: 77, after: 88}
		for i := range guard.factors {
			guard.factors[i] = -9
		}
		used := kf_factor(nil, n, &guard.factors)
		product := int64(1)
		remaining := n
		for i := 0; i < used; i += 2 {
			radix := guard.factors[i]
			remaining /= radix
			product *= int64(radix)
			if remaining != guard.factors[i+1] {
				t.Fatal(n, guard)
			}
		}
		if product != int64(n) || remaining != 1 || guard.before != 77 || guard.after != 88 {
			t.Fatal(n, guard)
		}
		for _, v := range guard.factors[used:] {
			if v != -9 {
				t.Fatal("unused tail", n)
			}
		}
	}
	var f [64]int32
	used := kf_factor(nil, 60, &f)
	if used != 6 || f[0] != 4 || f[1] != 15 || f[2] != 3 || f[3] != 5 || f[4] != 5 || f[5] != 1 {
		t.Fatal("order", f)
	}
}
