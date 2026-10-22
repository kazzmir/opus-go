package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// From nlsf2a_pointers_test.go.
func TestNLSF2APointers(t *testing.T) {
	for _, d := range []int{10, 16} {
		for _, tight := range []bool{false, true} {
			in := make([]int16, d)
			for i := range in {
				in[i] = int16((i + 1) * 32768 / (d + 1))
				if tight {
					in[i] = int16(2000 + 100*i)
				}
			}
			before := slices.Clone(in)
			out := make([]int16, d+2)
			out[0] = 123
			out[d+1] = 456
			Opus_silk_NLSF2A(nil, &out[1], &in[0], int32(d), 0)
			if out[0] != 123 || out[d+1] != 456 || !slices.Equal(in, before) {
				t.Fatal("buffer boundary changed")
			}
			Opus_silk_NLSF2A(nil, &in[0], &in[0], int32(d), 0)
			if !slices.Equal(in, out[1:d+1]) {
				t.Fatal("in-place output differs")
			}
		}
	}
}

// From nlsf_decode_pointers_test.go.
func TestNLSFDecodePointers(t *testing.T) {
	for _, cb := range []*OpusT_silk_NLSF_CB_struct{&Opus_silk_NLSF_CB_NB_MB, &Opus_silk_NLSF_CB_WB} {
		n := int(cb.Forder)
		delta := unsafe.Slice((*int16)(unsafe.Pointer(cb.FdeltaMin_Q15)), n+1)
		for index := 0; index < int(cb.FnVectors); index++ {
			for _, residual := range []int8{-10, 0, 10} {
				var indices [17]int8
				indices[0] = int8(index)
				for i := 1; i <= n; i++ {
					indices[i] = residual
				}
				before := indices
				var output [18]int16
				output[0], output[n+1] = 111, 222
				Opus_silk_NLSF_decode(nil, &output[1], &indices[0], cb)
				if output[0] != 111 || output[n+1] != 222 || indices != before {
					t.Fatal("sentinels or indices changed")
				}
				if output[1] < delta[0] || 32768-int32(output[n]) < int32(delta[n]) {
					t.Fatal("endpoint spacing violated")
				}
				for i := 1; i < n; i++ {
					if int32(output[i+1])-int32(output[i]) < int32(delta[i]) {
						t.Fatalf("order=%d index=%d output=%v", n, index, output)
					}
				}
			}
		}
	}
}

// From nlsf_stabilize_pointers_test.go.
func TestNLSFStabilizePointers(t *testing.T) {
	for _, input := range [][4]int16{{1000, 2000, 3000, 4000}, {0, 0, 0, 0}, {32767, 32767, 32767, 32767}, {30000, 20000, 10000, -32768}} {
		values := [6]int16{111, input[0], input[1], input[2], input[3], 222}
		delta := [5]int16{50, 100, 200, 300, 400}
		before := delta
		Opus_silk_NLSF_stabilize(nil, &values[1], &delta[0], 4)
		if values[0] != 111 || values[5] != 222 || delta != before || values[1] < 50 || values[4] > 32768-400 {
			t.Fatal("bounds or sentinels changed")
		}
		for i := 1; i < 4; i++ {
			if int32(values[i+1])-int32(values[i]) < int32(delta[i]) {
				t.Fatal("spacing violated")
			}
		}
		stable := values
		Opus_silk_NLSF_stabilize(nil, &values[1], &delta[0], 4)
		if values != stable {
			t.Fatal("stable vector changed")
		}
	}
}

// From nlsf_unpack_pointers_test.go.
func cloneTestNLSFCodebook(src *OpusT_silk_NLSF_CB_struct) *OpusT_silk_NLSF_CB_struct {
	cb := *src
	n, v := int(cb.Forder), int(cb.FnVectors)
	cb.FCB1_NLSF_Q8 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FCB1_NLSF_Q8, n*v)))
	cb.FCB1_Wght_Q9 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FCB1_Wght_Q9, n*v)))
	cb.FCB1_iCDF = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FCB1_iCDF, 2*v)))
	cb.Fpred_Q8 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fpred_Q8, 2*(n-1))))
	cb.Fec_sel = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fec_sel, n*v/2)))
	cb.Fec_iCDF = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fec_iCDF, 72)))
	cb.Fec_Rates_Q5 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fec_Rates_Q5, 72)))
	cb.FdeltaMin_Q15 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FdeltaMin_Q15, n+1)))
	return &cb
}

func TestNLSFCodebookPointers(t *testing.T) {
	for _, src := range []*OpusT_silk_NLSF_CB_struct{&Opus_silk_NLSF_CB_NB_MB, &Opus_silk_NLSF_CB_WB} {
		cb := cloneTestNLSFCodebook(src)
		entropyInitGrowStack(12)
		runtime.GC()
		for index := int32(0); index < int32(cb.FnVectors); index++ {
			g, c := [16]int16{}, [16]int16{}
			gp, cp := [16]byte{}, [16]byte{}
			Opus_silk_NLSF_unpack(nil, &g[0], &gp[0], cb, index)
			Opus_silk_NLSF_unpack(nil, &c[0], &cp[0], src, index)
			if g != c || gp != cp {
				t.Fatal("owned unpack", index)
			}
			indices := [17]int8{}
			indices[0] = int8(index)
			Opus_silk_NLSF_decode(nil, &g[0], &indices[0], cb)
			Opus_silk_NLSF_decode(nil, &c[0], &indices[0], src)
			if g != c {
				t.Fatal("owned decode", index)
			}
		}
		if !slices.Equal(unsafe.Slice(cb.Fec_iCDF, 72), unsafe.Slice(src.Fec_iCDF, 72)) || !slices.Equal(unsafe.Slice(cb.Fec_Rates_Q5, 72), unsafe.Slice(src.Fec_Rates_Q5, 72)) || !slices.Equal(unsafe.Slice(cb.FCB1_iCDF, 2*int(cb.FnVectors)), unsafe.Slice(src.FCB1_iCDF, 2*int(src.FnVectors))) {
			t.Fatal("owned entropy tables")
		}
	}
}

func TestNLSFUnpackPointers(t *testing.T) {
	for _, cb := range []*OpusT_silk_NLSF_CB_struct{&Opus_silk_NLSF_CB_NB_MB, &Opus_silk_NLSF_CB_WB} {
		n := int(cb.Forder)
		before := *cb
		selectors := unsafe.Slice((*uint8)(unsafe.Pointer(cb.Fec_sel)), int(cb.FnVectors)*n/2)
		pred := unsafe.Slice((*uint8)(unsafe.Pointer(cb.Fpred_Q8)), 2*(n-1))
		for index := 0; index < int(cb.FnVectors); index++ {
			ix := make([]int16, n+2)
			p := make([]uint8, n+2)
			ix[0], ix[n+1] = 111, 222
			p[0], p[n+1] = 123, 234
			Opus_silk_NLSF_unpack(nil, &ix[1], &p[1], cb, int32(index))
			for i := 0; i < n; i++ {
				nibble := selectors[index*n/2+i/2] >> uint(4*(i%2))
				if ix[i+1] != int16((nibble&14)/2)*9 || p[i+1] != pred[i+int(nibble&1)*(n-1)] {
					t.Fatalf("order=%d index=%d element=%d", n, index, i)
				}
			}
			if ix[0] != 111 || ix[n+1] != 222 || p[0] != 123 || p[n+1] != 234 || *cb != before {
				t.Fatal("sentinels or codebook changed")
			}
		}
	}
}

// From nlsf_vq_pointers_test.go.
func TestNLSFVQPointers(t *testing.T) {
	input := [2]int16{100, 200}
	codebook := [4]uint8{0, 1, 0, 0}
	weights := [4]int16{2, 3, 4, 5}
	beforeInput, beforeCB, beforeWeights := input, codebook, weights
	output := [4]int32{111, 0, 0, 222}
	Opus_silk_NLSF_VQ(nil, &output[1], &input[0], &codebook[0], &weights[0], 2, 2)
	if output != [4]int32{111, 308, 1100, 222} {
		t.Fatalf("output=%v", output)
	}
	if input != beforeInput || codebook != beforeCB || weights != beforeWeights {
		t.Fatal("input changed")
	}
	input = [2]int16{-32768, 32767}
	codebook = [4]uint8{255, 0}
	weights = [4]int16{1, 1}
	Opus_silk_NLSF_VQ(nil, &output[1], &input[0], &codebook[0], &weights[0], 1, 2)
	if output[1] != 49022 {
		t.Fatalf("signed narrowing: got %d", output[1])
	}
	Opus_silk_NLSF_VQ(nil, nil, nil, nil, nil, 0, 16)
	Opus_silk_NLSF_VQ(nil, &output[1], nil, nil, nil, 2, 0)
	if output != [4]int32{111, 0, 0, 222} {
		t.Fatal("zero-order output differs")
	}
}

// From laroia_pointers_test.go.
func TestLaroiaPointers(t *testing.T) {
	for _, input := range [][]int16{{8192, 16384}, {0, 32767}, {-32768, 32767}, {100, 100, 99, 100}, {1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}} {
		want := make([]int16, len(input))
		for i, v := range input {
			left, right := int64(0), int64(32768)
			if i > 0 {
				left = int64(input[i-1])
			}
			if i+1 < len(input) {
				right = int64(input[i+1])
			}
			want[i] = int16(min(int64(131072)/max(int64(v)-left, 1)+int64(131072)/max(right-int64(v), 1), 32767))
		}
		out := make([]int16, len(input)+2)
		out[0], out[len(out)-1] = 111, 222
		before := slices.Clone(input)
		Opus_silk_NLSF_VQ_weights_laroia(nil, &out[1], &input[0], int32(len(input)))
		if !slices.Equal(out[1:len(out)-1], want) || out[0] != 111 || out[len(out)-1] != 222 || !slices.Equal(input, before) {
			t.Fatalf("input=%v got=%v want=%v", input, out, want)
		}
		alias := slices.Clone(input)
		Opus_silk_NLSF_VQ_weights_laroia(nil, &alias[0], &alias[0], int32(len(alias)))
		if !slices.Equal(alias, want) {
			t.Fatal("in-place result differs")
		}
	}
}
