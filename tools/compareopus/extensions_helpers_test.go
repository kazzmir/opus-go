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
func extensionCollectionPackets() [][]byte {
	packets := [][]byte{nil, {0}, {1}, {2}, {3, 0}, {3, 255}, {4}, {5}, {6}, {7}, {7, 44, 65}, {7, 11, 5, 22}, {65, 255}, {65, 255, 0}, {65, 2, 11, 12, 7, 99, 4, 2, 21, 22, 23, 31, 32, 33}, {6, 6, 6, 4}, {7, 11, 2, 7, 22}, {7, 11, 3, 2, 7, 22, 2, 7, 33}}
	packets = append(packets, append([]byte{65, 255, 0}, make([]byte, 255)...))
	rng := rand.New(rand.NewSource(89731))
	for i := 0; i < 1000; i++ {
		p := make([]byte, rng.Intn(65))
		rng.Read(p)
		packets = append(packets, p)
	}
	return packets
}
func extensionCollectionResult(fn func() int32) (result int32) {
	defer func() {
		if recover() != nil {
			result = -99
		}
	}()
	return fn()
}
func TestExtensionCountAgainstC(t *testing.T) {
	for i, data := range extensionCollectionPackets() {
		for _, frames := range []int32{0, 1, 3, 48} {
			before := slices.Clone(data)
			got := opuscc.Opus_opus_packet_extensions_count(nil, unsafe.SliceData(data), int32(len(data)), frames)
			native := nativeExtensionCount(data, int32(len(data)), frames)
			if got != native || !slices.Equal(before, data) {
				t.Fatal(i, frames, got, native)
			}
		}
	}
	for _, args := range [][2]int32{{-1, 1}, {1, 1}, {0, -1}, {0, 49}} {
		got := extensionCollectionResult(func() int32 { return opuscc.Opus_opus_packet_extensions_count(nil, nil, args[0], args[1]) })
		if native := nativeExtensionCount(nil, args[0], args[1]); got != native {
			t.Fatal("assert", args, got, native)
		}
	}
}

func TestExtensionCountExtAgainstC(t *testing.T) {
	for i, data := range extensionCollectionPackets() {
		for _, frames := range []int32{0, 1, 3, 48} {
			var g, c [50]int32
			for i := range g {
				g[i] = 77
				c[i] = 77
			}
			got := opuscc.Opus_opus_packet_extensions_count_ext(nil, unsafe.SliceData(data), int32(len(data)), &g[1], frames)
			native := nativeExtensionCountExt(unsafe.SliceData(data), int32(len(data)), &c[1], frames)
			if got != native || g != c {
				t.Fatal(i, frames, got, native, g, c)
			}
		}
	}
	for _, args := range [][2]int32{{-1, 1}, {1, 1}, {0, -1}, {0, 49}} {
		g, c := [50]int32{77, 88}, [50]int32{77, 88}
		got := extensionCollectionResult(func() int32 { return opuscc.Opus_opus_packet_extensions_count_ext(nil, nil, args[0], &g[0], args[1]) })
		native := nativeExtensionCountExt(nil, args[0], &c[0], args[1])
		if got != native || g != c {
			t.Fatal("assert", args, got, native, g, c)
		}
	}
	for _, frames := range []int32{1, 2, 3} {
		g, c := [8]int32{0x2c072c07, 0x2c072c07, 0x2c072c07, 77}, [8]int32{0x2c072c07, 0x2c072c07, 0x2c072c07, 77}
		got := opuscc.Opus_opus_packet_extensions_count_ext(nil, (*byte)(unsafe.Pointer(&g[0])), 12, &g[0], frames)
		native := nativeExtensionCountExt((*byte)(unsafe.Pointer(&c[0])), 12, &c[0], frames)
		if got != native || g != c {
			t.Fatal("alias", frames, got, native, g, c)
		}
	}
}

func extensionOutputWords(base *byte, out []opuscc.OpusT_opus_extension_data) []int32 {
	v := make([]int32, len(out)*4)
	for i, e := range out {
		v[i*4] = e.Fid
		v[i*4+1] = e.Fframe
		v[i*4+2] = extensionTestOffset(base, e.Fdata)
		v[i*4+3] = e.Flen1
	}
	return v
}
func TestExtensionParseAgainstC(t *testing.T) {
	for pi, data := range extensionCollectionPackets() {
		for _, frames := range []int32{0, 1, 3, 48} {
			total := opuscc.Opus_opus_packet_extensions_count(nil, unsafe.SliceData(data), int32(len(data)), frames)
			for _, capacity := range []int32{0, 1, total, total + 1} {
				out := make([]opuscc.OpusT_opus_extension_data, total+3)
				for i := range out {
					out[i] = opuscc.OpusT_opus_extension_data{Fid: 91, Fframe: 92, Flen1: 93}
				}
				words := extensionOutputWords(unsafe.SliceData(data), out)
				g, c := capacity, capacity
				got := opuscc.Opus_opus_packet_extensions_parse(nil, unsafe.SliceData(data), int32(len(data)), &out[0], &g, frames)
				native := nativeExtensionParse(data, int32(len(data)), frames, words, &c, false, false, 0)
				if got != native || g != c || !slices.Equal(extensionOutputWords(unsafe.SliceData(data), out), words) {
					t.Fatal(pi, frames, capacity, got, native, g, c)
				}
			}
		}
	}
	data := []byte{7, 11, 7, 22}
	out := make([]opuscc.OpusT_opus_extension_data, 4)
	out[0].Fframe = 1
	words := extensionOutputWords(&data[0], out)
	c := int32(1)
	got := opuscc.Opus_opus_packet_extensions_parse(nil, &data[0], 4, &out[0], &out[0].Fframe, 1)
	native := nativeExtensionParse(data, 4, 1, words, &c, false, false, 1)
	if got != native || out[0].Fframe != c || !slices.Equal(extensionOutputWords(&data[0], out), words) {
		t.Fatal("alias", got, native, out, c, words)
	}
	for _, test := range []struct {
		data                     []byte
		length, frames, capacity int32
		no, nc                   bool
	}{{nil, 0, 0, 0, true, false}, {data, 4, 1, 0, true, false}, {data, 4, 1, 1, true, false}, {nil, -1, 1, 2, false, false}, {nil, 1, 1, 2, false, false}, {nil, 0, 49, 2, false, false}, {nil, -1, 49, 2, true, true}} {
		g, c := test.capacity, test.capacity
		out := make([]opuscc.OpusT_opus_extension_data, 4)
		words := extensionOutputWords(unsafe.SliceData(test.data), out)
		p := &out[0]
		n := &g
		if test.no {
			p = nil
		}
		if test.nc {
			n = nil
		}
		got := extensionCollectionResult(func() int32 {
			return opuscc.Opus_opus_packet_extensions_parse(nil, unsafe.SliceData(test.data), test.length, p, n, test.frames)
		})
		native := nativeExtensionParse(test.data, test.length, test.frames, words, &c, test.no, test.nc, 0)
		if got != native || g != c || !slices.Equal(extensionOutputWords(unsafe.SliceData(test.data), out), words) {
			t.Fatal("assert/nil", test, got, native, g, c)
		}
	}
}

func TestExtensionParseExtAgainstC(t *testing.T) {
	for pi, data := range extensionCollectionPackets() {
		for _, frames := range []int32{0, 1, 3, 48} {
			counts := make([]int32, frames)
			total := opuscc.Opus_opus_packet_extensions_count_ext(nil, unsafe.SliceData(data), int32(len(data)), unsafe.SliceData(counts), frames)
			for _, capacity := range []int32{0, 1, total, total + 1} {
				out := make([]opuscc.OpusT_opus_extension_data, total+3)
				for i := range out {
					out[i] = opuscc.OpusT_opus_extension_data{Fid: 91, Fframe: 92, Flen1: 93}
				}
				words := extensionOutputWords(unsafe.SliceData(data), out)
				g, c := capacity, capacity
				before := slices.Clone(counts)
				got := opuscc.Opus_opus_packet_extensions_parse_ext(nil, unsafe.SliceData(data), int32(len(data)), &out[0], &g, unsafe.SliceData(counts), frames)
				native := nativeExtensionParse(data, int32(len(data)), frames, words, &c, false, false, 0, counts)
				if got != native || g != c || !slices.Equal(extensionOutputWords(unsafe.SliceData(data), out), words) || !slices.Equal(counts, before) {
					t.Fatal(pi, frames, capacity, got, native, g, c)
				}
			}
		}
	}
	data := []byte{7, 11, 5, 22, 33, 9, 44}
	for alias := int32(1); alias <= 2; alias++ {
		gcounts, ccounts := []int32{2, 1, 1}, []int32{2, 1, 1}
		if alias == 2 {
			gcounts[0] = 5
			ccounts[0] = 5
		}
		out := make([]opuscc.OpusT_opus_extension_data, 10)
		out[0].Fframe = 4
		words := extensionOutputWords(&data[0], out)
		g, c := int32(4), int32(4)
		gp := &out[0].Fframe
		if alias == 2 {
			gp = &gcounts[0]
		}
		got := opuscc.Opus_opus_packet_extensions_parse_ext(nil, &data[0], 7, &out[0], gp, &gcounts[0], 3)
		g = *gp
		native := nativeExtensionParse(data, 7, 3, words, &c, false, false, alias, ccounts)
		if got != native || g != c || !slices.Equal(gcounts, ccounts) || !slices.Equal(extensionOutputWords(&data[0], out), words) {
			t.Fatal("alias", alias, got, native, g, c, gcounts, ccounts)
		}
	}
	for _, test := range []struct {
		data                     []byte
		length, frames, capacity int32
		no, nc                   bool
	}{{nil, 0, 0, 0, true, false}, {data, 7, 3, 0, true, false}, {data, 7, 3, 1, true, false}, {data, 7, 3, 10, false, false}, {nil, -1, 1, 2, false, false}, {nil, 1, 1, 2, false, false}, {nil, 0, -1, 2, false, false}, {nil, 0, 49, 2, false, false}, {nil, -1, 49, 2, true, true}} {
		g, c := test.capacity, test.capacity
		counts := make([]int32, max(test.frames, 0))
		out := make([]opuscc.OpusT_opus_extension_data, 12)
		words := extensionOutputWords(unsafe.SliceData(test.data), out)
		p := &out[0]
		n := &g
		if test.no {
			p = nil
		}
		if test.nc {
			n = nil
		}
		got := extensionCollectionResult(func() int32 {
			return opuscc.Opus_opus_packet_extensions_parse_ext(nil, unsafe.SliceData(test.data), test.length, p, n, unsafe.SliceData(counts), test.frames)
		})
		native := nativeExtensionParse(test.data, test.length, test.frames, words, &c, test.no, test.nc, 0, counts)
		if got != native || g != c || !slices.Equal(extensionOutputWords(unsafe.SliceData(test.data), out), words) {
			t.Fatal("assert/nil", test, got, native, g, c)
		}
	}
}

func TestExtensionFindAgainstC(t *testing.T) {
	fixtures := [][]byte{nil, {0}, {1}, {2}, {3, 0}, {3, 255}, {4}, {5}, {6}, {7}, {7, 44}, {65, 255}, {65, 2, 11, 12, 7, 99, 4, 2, 21, 22, 23, 31, 32, 33}, {6, 6, 6, 4}, {7, 11, 2, 7, 22}}
	rng := rand.New(rand.NewSource(51173))
	for i := 0; i < 500; i++ {
		data := make([]byte, rng.Intn(33))
		rng.Read(data)
		fixtures = append(fixtures, data)
	}
	for fi, data := range fixtures {
		for _, frames := range []int32{0, 1, 3, 48} {
			for _, id := range []int32{-1, 0, 2, 3, 31, 32, 127, 128} {
				for scenario := 0; scenario < 4; scenario++ {
					if scenario == 3 && id != 128 {
						continue
					} // NULL output is defined only when no match is found.
					base := unsafe.SliceData(data)
					var st opuscc.OpusT_OpusExtensionIterator
					opuscc.Opus_opus_extension_iterator_init(nil, &st, base, int32(len(data)), frames)
					if scenario == 1 {
						opuscc.Opus_opus_extension_iterator_set_frame_max(nil, &st, 1)
					} else if scenario == 2 {
						opuscc.Opus_opus_extension_iterator_set_frame_max(nil, &st, 0)
					}
					ext := opuscc.OpusT_opus_extension_data{Fid: 91, Fframe: 92, Fdata: base, Flen1: 93}
					want := int32(1)
					var out *opuscc.OpusT_opus_extension_data = &ext
					if scenario == 3 {
						out = nil
						want = 0
					}
					limit := (len(data)+1)*(int(frames)+1) + 1
					for step := 0; step < limit; step++ {
						v := extensionTestWithOutput(&st, &ext)
						result := opuscc.Opus_opus_extension_iterator_find(nil, &st, out, id)
						native := nativeExtensionIterator(data, 0, 0, &v, 3, want, id)
						if result != native || extensionTestWithOutput(&st, &ext) != v {
							t.Fatal(fi, frames, id, scenario, step, result, native, extensionTestWithOutput(&st, &ext), v)
						}
						if result <= 0 {
							break
						}
						if step == limit-1 {
							t.Fatal("find did not stop")
						}
					}
				}
			}
		}
	}
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
