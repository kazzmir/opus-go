//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func extensionTestOffset(base, p *byte) int32 {
	if p == nil {
		return -1
	}
	return int32(uintptr(unsafe.Pointer(p)) - uintptr(unsafe.Pointer(base)))
}
func extensionTestState(st *opuscc.OpusT_OpusExtensionIterator) [19]int32 {
	return [19]int32{extensionTestOffset(st.Fdata, st.Fdata), extensionTestOffset(st.Fdata, st.Fcurr_data), extensionTestOffset(st.Fdata, st.Frepeat_data), extensionTestOffset(st.Fdata, st.Flast_long), extensionTestOffset(st.Fdata, st.Fsrc_data), st.Flen1, st.Fcurr_len, st.Frepeat_len, st.Fsrc_len, st.Ftrailing_short_len, st.Fnb_frames, st.Fframe_max, st.Fcurr_frame, st.Frepeat_frame, int32(st.Frepeat_l)}
}
func extensionTestWithOutput(st *opuscc.OpusT_OpusExtensionIterator, ext *opuscc.OpusT_opus_extension_data) [19]int32 {
	v := extensionTestState(st)
	v[15] = ext.Fid
	v[16] = ext.Fframe
	v[17] = extensionTestOffset(st.Fdata, ext.Fdata)
	v[18] = ext.Flen1
	return v
}
func TestExtensionNextAgainstC(t *testing.T) {
	fixtures := [][]byte{nil, {0}, {1}, {2}, {3, 0}, {3, 255}, {4}, {5}, {6}, {7}, {7, 44}, {65, 255}, {65, 255, 0}, {65, 2, 11, 12, 7, 99, 4, 2, 21, 22, 23, 31, 32, 33}, {65, 1, 19, 3, 0, 7, 20, 5, 1, 21, 22, 1, 23, 24}, {6, 6, 6, 4}, {7, 11, 2, 7, 22}}
	fixtures = append(fixtures, append([]byte{65, 255, 0}, make([]byte, 255)...))
	rng := rand.New(rand.NewSource(21897))
	for i := 0; i < 4000; i++ {
		data := make([]byte, rng.Intn(65))
		rng.Read(data)
		fixtures = append(fixtures, data)
	}
	for fi, data := range fixtures {
		for _, frames := range []int32{0, 1, 3, 48} {
			for _, scenario := range []int32{0, 1, 2} {
				var st opuscc.OpusT_OpusExtensionIterator
				base := unsafe.SliceData(data)
				opuscc.Opus_opus_extension_iterator_init(nil, &st, base, int32(len(data)), frames)
				ext := opuscc.OpusT_opus_extension_data{Fid: 91, Fframe: 92, Fdata: base, Flen1: 93}
				original := slices.Clone(data)
				limit := (len(data)+1)*(int(frames)+1) + 1
				for step := 0; step < limit; step++ {
					if scenario == 1 {
						opuscc.Opus_opus_extension_iterator_set_frame_max(nil, &st, 0)
					} else if scenario == 2 && step == 1 {
						opuscc.Opus_opus_extension_iterator_set_frame_max(nil, &st, 1)
					}
					v := extensionTestWithOutput(&st, &ext)
					want := int32((fi + step) % 2)
					var out *opuscc.OpusT_opus_extension_data
					if want != 0 {
						out = &ext
					}
					result := opuscc.Opus_opus_extension_iterator_next(nil, &st, out)
					native := nativeExtensionIterator(data, 0, 0, &v, 2, want)
					if result != native || extensionTestWithOutput(&st, &ext) != v || !slices.Equal(data, original) {
						t.Fatal(fi, frames, scenario, step, result, native, extensionTestWithOutput(&st, &ext), v)
					}
					if result <= 0 {
						break
					}
					if step == limit-1 {
						t.Fatal("next did not stop")
					}
				}
			}
		}
	}
}

func TestExtensionRepeatAgainstC(t *testing.T) {
	fixtures := []struct {
		data                                []byte
		source, current, lastLong, trailing int32
		flag                                byte
	}{
		{[]byte{7, 11, 5, 22, 33}, 2, 3, -1, 1, 1}, {[]byte{6, 5, 0, 0}, 1, 2, -1, 0, 1},
		{[]byte{0, 1, 2, 0, 3, 0, 7, 11, 5, 22, 33}, 8, 9, -1, 1, 1},
		{[]byte{65, 2, 11, 12, 7, 99, 4, 2, 21, 22, 23, 31, 32, 33}, 6, 7, 4, 1, 0},
		{[]byte{7, 11, 5}, 2, 3, -1, 1, 1}, {[]byte{65, 2, 11, 12, 7, 99, 4}, 6, 7, 4, 1, 0},
	}
	for fi, f := range fixtures {
		for _, frames := range []int32{1, 2, 3} {
			for _, max := range []int32{0, 1, 2, 3, 48} {
				for _, want := range []int32{0, 1} {
					base := unsafe.SliceData(f.data)
					var st opuscc.OpusT_OpusExtensionIterator
					opuscc.Opus_opus_extension_iterator_init(nil, &st, base, int32(len(f.data)), frames)
					st.Fcurr_data = (*byte)(unsafe.Add(unsafe.Pointer(base), f.current))
					st.Fcurr_len = int32(len(f.data)) - f.current
					st.Frepeat_len = f.source
					st.Fsrc_data = base
					st.Fsrc_len = f.source
					st.Frepeat_frame = 1
					st.Frepeat_l = f.flag
					st.Fframe_max = max
					st.Ftrailing_short_len = f.trailing
					if f.lastLong >= 0 {
						st.Flast_long = (*byte)(unsafe.Add(unsafe.Pointer(base), f.lastLong))
					}
					ext := opuscc.OpusT_opus_extension_data{Fid: 91, Fframe: 92, Fdata: base, Flen1: 93}
					for step := 0; step < 20; step++ {
						v := extensionTestWithOutput(&st, &ext)
						var out *opuscc.OpusT_opus_extension_data
						if want != 0 {
							out = &ext
						}
						result := opuscc.CompareExtensionRepeat(&st, out)
						native := nativeExtensionIterator(f.data, 0, 0, &v, 1, want)
						if result != native || extensionTestWithOutput(&st, &ext) != v {
							t.Fatal(fi, frames, max, want, step, result, native, extensionTestWithOutput(&st, &ext), v)
						}
						if result <= 0 {
							break
						}
						if step == 19 {
							t.Fatal("repeat did not stop")
						}
					}
				}
			}
		}
	}
}

func TestExtensionIteratorInitAgainstC(t *testing.T) {
	for _, data := range [][]byte{nil, {7, 99, 0, 1, 2, 3, 4, 5}} {
		for _, length := range []int32{-1, 0, 1, int32(len(data))} {
			for _, frames := range []int32{-1, 0, 1, 3, 48, 49} {
				base := unsafe.SliceData(data)
				st := opuscc.OpusT_OpusExtensionIterator{Fdata: base, Fcurr_data: base, Frepeat_data: base, Fsrc_data: base, Flast_long: base, Flen1: 9, Fcurr_len: 8, Frepeat_len: 7, Fsrc_len: 6, Ftrailing_short_len: 5, Fnb_frames: 4, Fframe_max: 3, Fcurr_frame: 2, Frepeat_frame: 1, Frepeat_l: 9}
				v := extensionTestState(&st)
				result := int32(0)
				func() {
					defer func() {
						if recover() != nil {
							result = -99
						}
					}()
					opuscc.Opus_opus_extension_iterator_init(nil, &st, base, length, frames)
				}()
				native := nativeExtensionIterator(data, length, frames, &v)
				if result != native || extensionTestState(&st) != v {
					t.Fatal(length, frames, result, native, extensionTestState(&st), v)
				}
			}
		}
	}
}

func TestWriteExtensionAgainstC(t *testing.T) {
	for _, id := range []int32{3, 31, 32, 127} {
		for _, length := range []int32{-1, 0, 1, 2, 254, 255, 256, 510, 511} {
			for _, last := range []int32{0, 1, -1} {
				for _, capacity := range []int32{0, 1, 2, 255, 257, 600} {
					for _, pos := range []int32{0, 1} {
						for _, sizeOnly := range []bool{false, true} {
							payload := make([]byte, 512)
							for i := range payload {
								payload[i] = byte(7*i + 11)
							}
							before := slices.Clone(payload)
							g := make([]byte, 602)
							for i := range g {
								g[i] = 77
							}
							c := slices.Clone(g)
							gb, cb := g, c
							if sizeOnly {
								gb = nil
								cb = nil
							}
							r := opuscc.CompareWriteExtension(unsafe.SliceData(gb), capacity, pos, id, length, &payload[0], last)
							want := nativeWriteExtension(cb, capacity, pos, id, length, payload, last)
							if r != want || !slices.Equal(g, c) || !slices.Equal(payload, before) {
								t.Fatal(id, length, last, capacity, pos, sizeOnly, r, want)
							}
						}
					}
				}
			}
		}
	}
}

func TestWritePayloadAgainstC(t *testing.T) {
	for _, id := range []int32{3, 31, 32, 127} {
		for _, length := range []int32{-1, 0, 1, 2, 254, 255, 256, 510, 511} {
			for _, last := range []int32{0, 1, -1} {
				for _, capacity := range []int32{0, 1, 2, 255, 257, 600} {
					for _, pos := range []int32{0, 1} {
						for _, sizeOnly := range []bool{false, true} {
							payload := make([]byte, 512)
							for i := range payload {
								payload[i] = byte(7*i + 11)
							}
							before := slices.Clone(payload)
							g := make([]byte, 602)
							for i := range g {
								g[i] = 77
							}
							c := slices.Clone(g)
							gb, cb := g, c
							if sizeOnly {
								gb = nil
								cb = nil
							}
							r := opuscc.CompareWritePayload(unsafe.SliceData(gb), capacity, pos, id, length, &payload[0], last)
							want := nativeWritePayload(cb, capacity, pos, id, length, payload, last)
							if r != want || !slices.Equal(g, c) || !slices.Equal(payload, before) {
								t.Fatal(id, length, last, capacity, pos, sizeOnly, r, want)
							}
						}
					}
				}
			}
		}
	}
}

func TestSkipExtensionAgainstC(t *testing.T) {
	for id := 0; id < 256; id++ {
		for _, length := range []int32{-1, 0, 1, 2, 3, 20, 257, 600} {
			for _, lace := range []byte{0, 1, 254, 255} {
				data := make([]byte, 601)
				for i := range data {
					data[i] = byte(11*i + 3)
				}
				data[0] = byte(id)
				data[1] = lace
				before := slices.Clone(data)
				p := &data[0]
				h := int32(77)
				g := opuscc.CompareSkipExtension(&p, length, &h)
				c, offset, ch := nativeSkipPayload(data, length, 0, 0, 77, 1)
				if g != c || h != ch || unsafe.Pointer(p) != unsafe.Add(unsafe.Pointer(&data[0]), offset) || !slices.Equal(data, before) {
					t.Fatal(id, length, lace, g, c, h, ch, offset)
				}
			}
		}
	}
}

func TestSkipPayloadAgainstC(t *testing.T) {
	for id := int32(0); id < 256; id++ {
		for _, length := range []int32{0, 1, 2, 20, 256, 257, 600} {
			for _, trailing := range []int32{0, 1, 5} {
				for _, lace := range []byte{0, 1, 254, 255} {
					data := make([]byte, 601)
					for i := range data {
						data[i] = byte(11*i + 3)
					}
					data[0] = lace
					before := slices.Clone(data)
					p := &data[0]
					h := int32(77)
					g := opuscc.CompareSkipPayload(&p, length, &h, id, trailing)
					c, offset, ch := nativeSkipPayload(data, length, id, trailing, 77, 0)
					if g != c || h != ch || unsafe.Pointer(p) != unsafe.Add(unsafe.Pointer(&data[0]), offset) || !slices.Equal(data, before) {
						t.Fatal(id, length, trailing, lace, g, c, h, ch, offset)
					}
				}
			}
		}
	}
}
