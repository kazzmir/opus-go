package opuscc

import (
	"testing"
	"unsafe"
)

func TestCNGExcPointers(t *testing.T) {
	for _, tc := range []struct{ length, mask int }{
		{1, 1}, {2, 1}, {3, 3}, {7, 7}, {8, 7}, {15, 15},
		{31, 31}, {32, 31}, {127, 127}, {128, 127}, {255, 255}, {320, 255},
	} {
		buffer := make([]int32, tc.mask+1)
		for i := range buffer {
			buffer[i] = int32(i*7919 - 123456)
		}
		for _, initial := range []int32{0, 3176576, -1, -2147483648} {
			out := make([]int32, tc.length+2)
			out[0], out[len(out)-1] = 111, 222
			seed := initial
			// Verify continued seed updates across successive calls too.
			wantSeed := uint64(uint32(initial))
			for pass := 0; pass < 2; pass++ {
				silk_CNG_exc(nil, &out[1], &buffer[0], int32(tc.length), &seed)
				for i := 0; i < tc.length; i++ {
					wantSeed = (907633515 + wantSeed*196314165) % (1 << 32)
					idx := int(wantSeed/(1<<24)) & tc.mask
					if out[i+1] != buffer[idx] {
						t.Fatalf("length=%d seed=%d pass=%d sample=%d: got %d want %d", tc.length, initial, pass, i, out[i+1], buffer[idx])
					}
				}
				if seed != int32(uint32(wantSeed)) || out[0] != 111 || out[len(out)-1] != 222 {
					t.Fatalf("length=%d: seed or sentinels changed incorrectly", tc.length)
				}
			}
		}
	}
	seed := int32(12345)
	silk_CNG_exc(nil, nil, nil, 0, &seed)
	if seed != 12345 {
		t.Fatal("empty excitation changed seed")
	}
}

func TestNLSFPolyPointers(t *testing.T) {
	input := [16]int32{131000, -130123, 65537, -32767, 1, 0, 12345, -45678, -99123, 54321, 721, 15, -4321, 23456, -120000, 99999}
	// Golden vectors use C's signed fixed-point arithmetic, rounding each
	// product as floor((product + 32768) / 65536), including negative values.
	for _, tc := range []struct {
		order, offset int
		want          []int32
	}{
		{1, 0, []int32{65536, -131000}},
		{1, 1, []int32{65536, 130123}},
		{5, 0, []int32{65536, -109760, 179772, -209579, 174309, -199637}},
		{5, 1, []int32{65536, 154247, 333397, 514303, 634924, 720112}},
		{8, 0, []int32{65536, 13840, 175919, 36343, 125602, -59340, -36320, -262802, -103077}},
		{8, 1, []int32{65536, 30777, 275223, 186278, 499590, 515209, 549405, 829676, 519002}},
	} {
		// Exact-size input is important for odd-offset interleaved access.
		lsf := append([]int32(nil), input[tc.offset:2*tc.order]...)
		out := make([]int32, tc.order+3)
		out[0], out[len(out)-1] = 111, 222
		silk_NLSF2A_find_poly(nil, &out[1], &lsf[0], int32(tc.order))
		for i, want := range tc.want {
			if out[i+1] != want {
				t.Fatalf("order=%d offset=%d coefficient=%d: got %d want %d", tc.order, tc.offset, i, out[i+1], want)
			}
		}
		if out[0] != 111 || out[len(out)-1] != 222 {
			t.Fatal("output sentinel overwritten")
		}
		for i, value := range lsf {
			if value != input[i+tc.offset] {
				t.Fatal("input modified")
			}
		}
	}
}

func TestDecoderSizePointer(t *testing.T) {
	out := [3]int32{111, -1, 222}
	if ret := Opus_silk_Get_Decoder_Size(nil, &out[1]); ret != SILK_NO_ERROR {
		t.Fatalf("return code: %d", ret)
	}
	if out != [3]int32{111, int32(unsafe.Sizeof(OpusT_silk_decoder{})), 222} {
		t.Fatalf("size/sentinels: %v", out)
	}
}
