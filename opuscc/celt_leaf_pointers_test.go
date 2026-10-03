package opuscc

import (
	"math"
	"runtime"
	"testing"
	"unsafe"
)

func TestCeltPLCLagWindowPointers(t *testing.T) {
	for trial := 0; trial < 80; trial++ {
		storage := new([27]float32)
		storage[0], storage[26] = 77, 88
		for i := 1; i < 26; i++ {
			storage[i] = float32((i*7919+trial*997)%65536-32768) / 37
		}
		ac := (*[25]float32)(unsafe.Pointer(&storage[1]))
		want := *ac
		want[0] *= 1.0001
		for i := int32(1); i <= 24; i++ {
			want[i] -= float32(float32(float32(want[i]*float32(float32(.008)*float32(.008)))*float32(i)) * float32(i))
		}
		entropyInitGrowStack(12)
		runtime.GC()
		celtPLCLagWindow(ac)
		for i := range ac {
			if math.Float32bits(ac[i]) != math.Float32bits(want[i]) {
				t.Fatal("rounded lag window", trial, i)
			}
		}
		if storage[0] != 77 || storage[26] != 88 {
			t.Fatal("lag window guards")
		}
	}
}
func TestCeltPLCDecayPointers(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, loss := range []int32{0, 1, 99} {
			for _, alias := range []bool{false, true} {
				a := make([]float32, 44)
				b := make([]float32, 44)
				for i := range a {
					a[i] = float32(i) - 22
					b[i] = float32(i%5) - 12
				}
				a[0], a[43] = 77, 88
				if alias {
					b = a
				}
				want := append([]float32(nil), a...)
				wb := append([]float32(nil), b...)
				if alias {
					wb = want
				}
				for c := int32(0); c < max(channels, 1); c++ {
					for i := int32(1); i < 20; i++ {
						index := 1 + c*21 + i
						decay := float32(.5)
						if loss == 0 {
							decay = 1.5
						}
						if wb[index] > want[index]-decay {
							want[index] = wb[index]
						} else {
							want[index] -= decay
						}
					}
				}
				entropyInitGrowStack(12)
				runtime.GC()
				celtPLCDecay(&a[1], &b[1], 21, 1, 20, channels, loss)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(want[i]) {
						t.Fatal("decay order", channels, loss, alias, i)
					}
				}
			}
		}
	}
	celtPLCDecay(nil, nil, 21, 5, 5, 2, 0)
	a := []float32{float32(math.NaN()), 3}
	b := []float32{2, float32(math.NaN())}
	celtPLCDecay(&a[0], &b[0], 2, 0, 2, 1, 0)
	if !math.IsNaN(float64(a[0])) || a[1] != 1.5 {
		t.Fatal("MAXG NaN selection", a)
	}
}
func TestEntropyWritePointers(t *testing.T) {
	buffer := [3]byte{}
	enc := OpusT_ec_enc{
		Fbuf:     &buffer[0],
		Fstorage: uint32(len(buffer)),
	}
	if got := ec_write_byte(nil, &enc, 0x123); got != 0 {
		t.Fatalf("front write: %d", got)
	}
	if got := ec_write_byte_at_end(nil, &enc, 0x245); got != 0 {
		t.Fatalf("back write: %d", got)
	}
	if got := ec_write_byte(nil, &enc, 0x67); got != 0 {
		t.Fatalf("last free byte: %d", got)
	}
	if want := [3]byte{0x23, 0x67, 0x45}; buffer != want {
		t.Fatalf("buffer: got %x, want %x", buffer, want)
	}
	before := enc
	if got := ec_write_byte(nil, &enc, 0xff); got != -1 {
		t.Fatalf("full front write: %d", got)
	}
	if got := ec_write_byte_at_end(nil, &enc, 0xff); got != -1 {
		t.Fatalf("full back write: %d", got)
	}
	if enc != before || buffer != [3]byte{0x23, 0x67, 0x45} {
		t.Fatal("failed write modified state or buffer")
	}
	empty := OpusT_ec_enc{}
	if ec_write_byte(nil, &empty, 1) != -1 || ec_write_byte_at_end(nil, &empty, 1) != -1 || empty != (OpusT_ec_enc{}) {
		t.Fatal("zero-capacity writer should fail without accessing a buffer")
	}
}

func TestStereoSplitPointers(t *testing.T) {
	// Interior pointers and sentinels check both ends of the requested span.
	x := [...]float32{99, 1, -2, 0, 0.125, 88}
	y := [...]float32{77, -1, 3, 0, 0.25, 66}
	wantX, wantY := x, y
	for i := 1; i < len(x)-1; i++ {
		l := float32(float32(0.70710678) * x[i])
		r := float32(float32(0.70710678) * y[i])
		wantX[i], wantY[i] = l+r, r-l
	}
	stereo_split(nil, &x[1], &y[1], 4)
	if x != wantX || y != wantY {
		t.Fatalf("got X=%v Y=%v, want X=%v Y=%v", x, y, wantX, wantY)
	}
	stereo_split(nil, nil, nil, 0)
}
