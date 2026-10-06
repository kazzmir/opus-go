//go:build compareopus && cgo

package main

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestDecoderDestroyAgainstC(t *testing.T) {
	for _, null := range []bool{false, true} {
		if !nativeOpusDestroy(null) {
			t.Fatal("decoder destruction must free only its base", null)
		}
	}
}

func TestModeTablesAgainstC(t *testing.T) {
	m, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	window := make([]float32, m.Foverlap)
	vectors := make([]byte, m.FnbAllocVectors*m.FnbEBands)
	caps := make([]byte, (m.FmaxLM+1)*2*m.FnbEBands)
	layout := nativeModeTables(window, vectors, caps)
	want := [5]uint64{uint64(unsafe.Sizeof(*m)), uint64(unsafe.Offsetof(m.FallocVectors)), uint64(unsafe.Offsetof(m.Fwindow)), uint64(unsafe.Sizeof(m.Fcache)), uint64(unsafe.Offsetof(m.Fcache.Fcaps))}
	if layout != want {
		t.Fatal(layout, want)
	}
	if !sameFloatBits(window, unsafe.Slice(m.Fwindow, len(window))) || !slices.Equal(vectors, unsafe.Slice(m.FallocVectors, len(vectors))) || !slices.Equal(caps, unsafe.Slice(m.Fcache.Fcaps, len(caps))) {
		t.Fatal("mode table payloads")
	}
}

func TestModeLogTableAgainstC(t *testing.T) {
	m, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]int16, m.FnbEBands)
	layout := nativeModeRemainingTable(0, unsafe.Pointer(&data[0]))
	want := [3]uint64{uint64(unsafe.Sizeof(*m)), uint64(unsafe.Offsetof(m.FlogN)), uint64(len(data) * 2)}
	if layout != want || !slices.Equal(data, unsafe.Slice(m.FlogN, len(data))) {
		t.Fatal("logN layout/payload", layout, want)
	}
}

func TestModePulseIndexTableAgainstC(t *testing.T) {
	m, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]int16, (m.FmaxLM+2)*m.FnbEBands)
	layout := nativeModeRemainingTable(1, unsafe.Pointer(&data[0]))
	want := [3]uint64{uint64(unsafe.Sizeof(m.Fcache)), uint64(unsafe.Offsetof(m.Fcache.Findex)), uint64(len(data) * 2)}
	if layout != want || !slices.Equal(data, unsafe.Slice(m.Fcache.Findex, len(data))) {
		t.Fatal("pulse index layout/payload", layout, want)
	}
}

func TestModePulseBitsTableAgainstC(t *testing.T) {
	m, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]byte, m.Fcache.Fsize)
	layout := nativeModeRemainingTable(2, unsafe.Pointer(&data[0]))
	want := [3]uint64{uint64(unsafe.Sizeof(m.Fcache)), uint64(unsafe.Offsetof(m.Fcache.Fbits)), uint64(len(data))}
	if layout != want || !slices.Equal(data, unsafe.Slice(m.Fcache.Fbits, len(data))) {
		t.Fatal("pulse bits layout/payload", layout, want)
	}
	index := unsafe.Slice(m.Fcache.Findex, int((m.FmaxLM+2)*m.FnbEBands))
	for LM := int32(-1); LM <= m.FmaxLM; LM++ {
		for band := int32(0); band < m.FnbEBands; band++ {
			offset := index[(LM+1)*m.FnbEBands+band]
			if offset < 0 {
				continue
			}
			max := int32(data[offset])
			for bits := int32(-2); bits <= 400; bits++ {
				pulse := bits % (max + 1)
				if pulse < 0 {
					pulse = 0
				}
				gq, gb := opuscc.CompareModePulseRate(m, band, LM, bits, pulse)
				cq, cb := nativeModePulseRate(band, LM, bits, pulse)
				if gq != cq || gb != cb {
					t.Fatal(band, LM, bits, pulse, gq, cq, gb, cb)
				}
			}
		}
	}
}

func TestModeBandTableAgainstC(t *testing.T) {
	m, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	data := make([]int16, m.FnbEBands+1)
	layout := nativeModeRemainingTable(3, unsafe.Pointer(&data[0]))
	want := [3]uint64{uint64(unsafe.Sizeof(*m)), uint64(unsafe.Offsetof(m.FeBands)), uint64(len(data) * 2)}
	if layout != want || !slices.Equal(data, unsafe.Slice(m.FeBands, len(data))) {
		t.Fatal("band layout/payload", layout, want)
	}
}

func TestCompositeProjectionSizeAgainstC(t *testing.T) {
	for _, shape := range [][3]int32{{1, 1, 0}, {2, 1, 1}, {4, 2, 2}, {0, 1, 0}, {-1, 1, 0}, {256, 1, 0}, {255, 255, 255}, {1, 0, 0}, {1, 2, -1}, {1, 1, 2}} {
		got := opuscc.Opus_opus_projection_decoder_get_size(nil, shape[0], shape[1], shape[2])
		if want := nativeProjectionSize(shape[0], shape[1], shape[2]); got != want {
			t.Fatal("projection size", shape, got, want)
		}
	}
}

func TestCompositeMatrixSizeAgainstC(t *testing.T) {
	for _, shape := range [][2]int32{{0, 0}, {1, 1}, {255, 127}, {255, 128}, {200, 162}, {256, 0}, {0, 256}, {-1, 3}, {-300, 2}} {
		got := opuscc.Opus_mapping_matrix_get_size(nil, shape[0], shape[1])
		if want := nativeMappingMatrixSize(shape[0], shape[1]); got != want {
			t.Fatal("matrix size", shape, got, want)
		}
	}
}

func TestCompositeMultistreamSizeAgainstC(t *testing.T) {
	for _, shape := range [][2]int32{{-1, -1}, {0, 0}, {1, 0}, {1, 1}, {2, 0}, {2, 1}, {2, 2}, {5, 2}, {1, 2}, {4, -1}, {255, 0}, {255, 255}} {
		got := opuscc.Opus_opus_multistream_decoder_get_size(nil, shape[0], shape[1])
		if want := nativeMSSize(shape[0], shape[1]); got != want {
			t.Fatal("multistream size", shape, got, want)
		}
	}
}

func TestCompositeDecoderSizeAgainstC(t *testing.T) {
	for _, ch := range []int32{-2147483648, -1, 0, 1, 2, 3, 2147483647} {
		got := opuscc.Opus_opus_decoder_get_size(nil, ch)
		if want := nativeOpusSize(ch); got != want {
			t.Fatal("composite decoder size", ch, got, want)
		}
	}
}

func TestCustomDecoderSizeAgainstC(t *testing.T) {
	for _, overlap := range []int32{0, 60, 120, 240} {
		for _, bands := range []int32{0, 1, 21, 25} {
			for _, channels := range []int32{0, 1, 2} {
				mode := opuscc.OpusT_OpusCustomMode{Foverlap: overlap, FnbEBands: bands}
				before := mode
				g := opuscc.CompareCustomDecoderSize(&mode, channels)
				c := nativeCustomDecoderSize(overlap, bands, channels)
				if g != c || mode != before {
					t.Fatal(overlap, bands, channels, g, c)
				}
			}
		}
	}
}

func TestMSDecoderDestroyAgainstC(t *testing.T) {
	for _, null := range []bool{false, true} {
		if !nativeMSDestroy(null) {
			t.Fatal("multistream must free only its base", null)
		}
	}
}

func TestMSValidationAgainstC(t *testing.T) {
	for _, channels := range []int32{0, 1, 2, 255} {
		for _, streams := range []int32{0, 1, 2, 127, 255, 256} {
			for _, coupled := range []int32{0, 1, 127} {
				for _, mapping := range []uint8{0, 1, 2, 254, 255} {
					st := opuscc.OpusT_OpusMSDecoder{}
					st.Flayout.Fnb_channels = channels
					st.Flayout.Fnb_streams = streams
					st.Flayout.Fnb_coupled_streams = coupled
					for i := range st.Flayout.Fmapping {
						st.Flayout.Fmapping[i] = mapping
					}
					before := st
					g := opuscc.CompareMSValidation(&st)
					c := nativeMSValidation(&st)
					if g != c || st != before {
						t.Fatal(channels, streams, coupled, mapping, g, c)
					}
				}
			}
		}
	}
}

func TestProjectionDecoderDestroyAgainstC(t *testing.T) {
	for _, null := range []bool{false, true} {
		if !nativeProjectionDestroy(null) {
			t.Fatal("projection must free only its base", null)
		}
	}
}

func TestProjectionDecoderInitAgainstC(t *testing.T) {
	for _, shape := range [][3]int32{{1, 1, 0}, {2, 1, 1}, {3, 2, 1}, {5, 3, 2}, {3, 1, 0}, {0, 1, 0}, {1, 0, 0}, {1, 0, 1}, {256, 1, 0}, {255, 255, 0}} {
		channels, streams, coupled := shape[0], shape[1], shape[2]
		size := int(opuscc.Opus_opus_projection_decoder_get_size(nil, channels, streams, coupled))
		if size != int(nativeProjectionSize(channels, streams, coupled)) {
			t.Fatal("size", shape)
		}
		validSize := size > 0
		if size == 0 {
			size = 4096
		}
		count := channels * (streams + coupled)
		matrix := make([]byte, count*2)
		values := []int16{0, 32767, -32768, -1, 12345, -23456, 1, -2}
		for i := int32(0); i < count; i++ {
			v := uint16(values[i%8])
			matrix[2*i] = byte(v)
			matrix[2*i+1] = byte(v >> 8)
		}
		msOffset := int((unsafe.Sizeof(opuscc.OpusT_OpusProjectionDecoder{})+7)&^uintptr(7)) + int(opuscc.Opus_mapping_matrix_get_size(nil, channels, streams+coupled))
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000, 44100} {
			for _, delta := range []int32{-1, 0, 1} {
				backing := make([]uint64, (size+7)/8+2)
				st := (*opuscc.OpusT_OpusProjectionDecoder)(unsafe.Pointer(&backing[0]))
				g := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
				for i := range g {
					g[i] = 0xa5
				}
				if validSize {
					normalizeMSModes(g[msOffset:], streams, coupled)
				}
				c := slices.Clone(g)
				code := opuscc.Opus_opus_projection_decoder_init(nil, st, rate, channels, streams, coupled, unsafe.SliceData(matrix), count*2+delta)
				native := nativeProjectionInitImage(c, rate, channels, streams, coupled, matrix, count*2+delta, -1)
				normalized := slices.Clone(g)
				if code == 0 {
					normalizeMSModes(normalized[msOffset:], streams, coupled)
				}
				if code != native || !slices.Equal(normalized, c) {
					t.Fatal(shape, rate, delta, "state image", code, native)
				}
			}
		}
	}
	// Snapshot coefficients before any state writes, including when input aliases the header.
	for _, offset := range []int32{0, 8, 24} {
		size := int(opuscc.Opus_opus_projection_decoder_get_size(nil, 3, 2, 1))
		b := make([]uint64, (size+7)/8+2)
		st := (*opuscc.OpusT_OpusProjectionDecoder)(unsafe.Pointer(&b[0]))
		g := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
		for i := range g {
			g[i] = byte(i * 17)
		}
		msOffset := 8 + int(opuscc.Opus_mapping_matrix_get_size(nil, 3, 3))
		normalizeMSModes(g[msOffset:], 2, 1)
		c := slices.Clone(g)
		code := opuscc.Opus_opus_projection_decoder_init(nil, st, 48000, 3, 2, 1, &g[offset], 18)
		native := nativeProjectionInitImage(c, 48000, 3, 2, 1, nil, 18, offset)
		if code == 0 {
			normalizeMSModes(g[msOffset:], 2, 1)
		}
		if code != native || !slices.Equal(g, c) {
			t.Fatal("alias", offset, code, native)
		}
	}
}

func normalizeMSModes(image []byte, streams, coupled int32) {
	ptr := int((unsafe.Sizeof(opuscc.OpusT_OpusMSDecoder{}) + 7) &^ uintptr(7))
	var silk int32
	opuscc.Opus_silk_Get_Decoder_Size(nil, &silk)
	celt := int((unsafe.Sizeof(opuscc.OpusT_OpusDecoder{})+7)&^uintptr(7)) + int((uint32(silk)+7)&^uint32(7))
	for i := int32(0); i < streams; i++ {
		clear(image[ptr+celt : ptr+celt+int(unsafe.Sizeof(uintptr(0)))])
		ch := int32(1)
		if i < coupled {
			ch = 2
		}
		ptr += int((uint32(opuscc.Opus_opus_decoder_get_size(nil, ch)) + 7) &^ uint32(7))
	}
}

func TestMSDecoderInitAgainstC(t *testing.T) {
	for _, shape := range [][3]int32{{1, 1, 0}, {2, 1, 1}, {3, 2, 1}, {5, 3, 2}, {4, 4, 0}} {
		channels, streams, coupled := shape[0], shape[1], shape[2]
		size := int(opuscc.Opus_opus_multistream_decoder_get_size(nil, streams, coupled))
		if size != int(nativeMSSize(streams, coupled)) {
			t.Fatal("size", shape)
		}
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000, 44100} {
			for trial := 0; trial < 4; trial++ {
				backing := make([]uint64, (size+7)/8+2)
				st := (*opuscc.OpusT_OpusMSDecoder)(unsafe.Pointer(&backing[0]))
				g := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
				for i := range g {
					g[i] = 0xa5
				}
				normalizeMSModes(g, streams, coupled)
				c := slices.Clone(g)
				mapping := make([]byte, channels)
				for i := range mapping {
					mapping[i] = byte(i)
				}
				if trial == 1 {
					mapping[0] = 255
				}
				if trial == 2 {
					mapping[len(mapping)-1] = byte(streams + coupled)
				}
				if trial == 3 {
					clear(mapping)
				}
				code := opuscc.Opus_opus_multistream_decoder_init(nil, st, rate, channels, streams, coupled, &mapping[0])
				native := nativeMSInitImage(c, rate, channels, streams, coupled, mapping, -1)
				normalized := slices.Clone(g)
				if code == 0 {
					normalizeMSModes(normalized, streams, coupled)
				}
				if code != native || !slices.Equal(normalized, c) {
					t.Fatal(shape, rate, trial, "state image", code, native)
				}
			}
		}
	}
	// Forward alias stores, including a source one byte before the destination.
	for _, delta := range []int32{-1, 0, 1} {
		size := int(opuscc.Opus_opus_multistream_decoder_get_size(nil, 2, 1))
		b := make([]uint64, (size+7)/8+2)
		st := (*opuscc.OpusT_OpusMSDecoder)(unsafe.Pointer(&b[0]))
		g := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
		for i := range st.Flayout.Fmapping {
			st.Flayout.Fmapping[i] = byte(i % 3)
		}
		off := int32(unsafe.Offsetof(st.Flayout.Fmapping)) + delta
		c := slices.Clone(g)
		code := opuscc.Opus_opus_multistream_decoder_init(nil, st, 48000, 3, 2, 1, &g[off])
		native := nativeMSInitImage(c, 48000, 3, 2, 1, nil, off)
		if code == 0 {
			normalizeMSModes(g, 2, 1)
		}
		if code != native || !slices.Equal(g, c) {
			t.Fatal("alias", delta, code, native)
		}
	}
}

func factoryErrorCode(err error) int32 {
	if err == nil {
		return 0
	}
	return err.(*opuscc.OpusError).Code
}
func TestDecoderCreateAgainstC(t *testing.T) {
	for _, rate := range []int32{0, 8000, 12000, 16000, 24000, 48000, 44100} {
		for _, channels := range []int32{0, 1, 2, 3} {
			for _, fail := range []bool{false, true} {
				var tls *libc.TLS
				if !fail {
					tls = libc.NewTLS()
				}
				st, err := opuscc.Opus_opus_decoder_create_typed(tls, rate, channels)
				code := factoryErrorCode(err)
				var image []byte
				if st != nil {
					image = slices.Clone(unsafe.Slice((*byte)(unsafe.Pointer(st)), int(opuscc.Opus_opus_decoder_get_size(nil, channels))))
					off := int(st.Fcelt_dec_offset)
					clear(image[off : off+int(unsafe.Sizeof(uintptr(0)))])
				}
				native := make([]byte, len(image))
				want := nativeOpusCreateImage(native, rate, channels, fail)
				if code != want || !slices.Equal(image, native) {
					t.Fatal(rate, channels, fail, code, want, "factory image")
				}
				if tls != nil {
					libc.XfreePointer(tls, unsafe.Pointer(st))
					tls.Close()
				}
			}
		}
	}
}

func TestMSDecoderCreateAgainstC(t *testing.T) {
	for _, shape := range [][3]int32{{1, 1, 0}, {2, 1, 1}, {3, 2, 1}, {5, 3, 2}, {255, 1, 0}, {0, 1, 0}, {256, 1, 0}, {1, 0, 0}, {1, 1, -1}, {1, 1, 2}, {1, 256, 0}} {
		channels, streams, coupled := shape[0], shape[1], shape[2]
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000, 44100} {
			for _, fail := range []bool{false, true} {
				for variant := 0; variant < 3; variant++ {
					mapping := make([]byte, max(channels, 1))
					for i := range mapping {
						mapping[i] = byte(int32(i) % max(streams+coupled, 1))
					}
					if variant == 1 {
						mapping[0] = 255
					}
					if variant == 2 {
						mapping[0] = byte(streams + coupled)
					}
					var tls *libc.TLS
					if !fail {
						tls = libc.NewTLS()
					}
					st, err := opuscc.Opus_opus_multistream_decoder_create_typed(tls, rate, channels, streams, coupled, &mapping[0])
					code := factoryErrorCode(err)
					var image []byte
					if st != nil {
						image = slices.Clone(unsafe.Slice((*byte)(unsafe.Pointer(st)), int(opuscc.Opus_opus_multistream_decoder_get_size(nil, streams, coupled))))
						normalizeMSModes(image, streams, coupled)
					}
					native := make([]byte, len(image))
					want := nativeMSCreateImage(native, rate, channels, streams, coupled, mapping, fail)
					if code != want || !slices.Equal(image, native) {
						t.Fatal(shape, rate, fail, variant, code, want, "factory image")
					}
					if tls != nil {
						libc.XfreePointer(tls, unsafe.Pointer(st))
						tls.Close()
					}
				}
			}
		}
	}
}

func TestProjectionDecoderCreateAgainstC(t *testing.T) {
	for _, shape := range [][3]int32{{1, 1, 0}, {2, 1, 1}, {3, 2, 1}, {5, 3, 2}, {3, 1, 0}, {0, 1, 0}, {256, 1, 0}, {1, 0, 0}, {1, 1, 2}, {255, 255, 0}} {
		channels, streams, coupled := shape[0], shape[1], shape[2]
		count := channels * (streams + coupled)
		matrix := make([]byte, count*2)
		for i := range matrix {
			matrix[i] = byte(i * 37)
		}
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000, 44100} {
			for _, fail := range []bool{false, true} {
				for _, delta := range []int32{-1, 0, 1} {
					var tls *libc.TLS
					if !fail {
						tls = libc.NewTLS()
					}
					st, err := opuscc.Opus_opus_projection_decoder_create_typed(tls, rate, channels, streams, coupled, unsafe.SliceData(matrix), count*2+delta)
					code := factoryErrorCode(err)
					var image []byte
					if st != nil {
						image = slices.Clone(unsafe.Slice((*byte)(unsafe.Pointer(st)), int(opuscc.Opus_opus_projection_decoder_get_size(nil, channels, streams, coupled))))
						msOffset := int((unsafe.Sizeof(*st)+7)&^uintptr(7)) + int(st.Fdemixing_matrix_size_in_bytes)
						normalizeMSModes(image[msOffset:], streams, coupled)
					}
					native := make([]byte, len(image))
					want := nativeProjectionCreateImage(native, rate, channels, streams, coupled, matrix, count*2+delta, fail)
					if code != want || !slices.Equal(image, native) {
						t.Fatal(shape, rate, fail, delta, code, want, "factory image")
					}
					if tls != nil {
						libc.XfreePointer(tls, unsafe.Pointer(st))
						tls.Close()
					}
				}
			}
		}
	}
}

func TestOpusDecoderInitAgainstC(t *testing.T) {
	for _, ch := range []int32{1, 2} {
		size := int(opuscc.Opus_opus_decoder_get_size(nil, ch))
		if size != int(nativeOpusSize(ch)) {
			t.Fatal("size", ch, size, nativeOpusSize(ch))
		}
		var silkSize int32
		opuscc.Opus_silk_Get_Decoder_Size(nil, &silkSize)
		offset := int((unsafe.Sizeof(opuscc.OpusT_OpusDecoder{})+7)&^uintptr(7)) + int((uint32(silkSize)+7)&^uint32(7))
		for _, rate := range []int32{-1, 0, 8000, 12000, 16000, 24000, 44100, 48000, 96000} {
			for _, channels := range []int32{ch, 0, 3} {
				backing := make([]uint64, (size+7)/8+2)
				st := (*opuscc.OpusT_OpusDecoder)(unsafe.Pointer(&backing[0]))
				g := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
				for i := range g {
					g[i] = 0xa5
				}
				clear(g[offset : offset+int(unsafe.Sizeof(uintptr(0)))])
				c := slices.Clone(g)
				code := opuscc.Opus_opus_decoder_init(nil, st, rate, channels)
				native := nativeOpusInitImage(c, rate, channels)
				normalized := slices.Clone(g)
				if code == 0 {
					clear(normalized[offset : offset+int(unsafe.Sizeof(uintptr(0)))])
				}
				if code != native || !slices.Equal(normalized, c) {
					for i := range c {
						if c[i] != normalized[i] {
							t.Fatalf("ch=%d rate=%d channels=%d byte=%d Go=%02x C=%02x code=%d/%d", ch, rate, channels, i, normalized[i], c[i], code, native)
						}
					}
					t.Fatal("result", code, native)
				}
			}
		}
	}
}

func TestCeltDecoderInitAgainstC(t *testing.T) {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	for _, channels := range []int32{0, 1, 2} {
		for _, rate := range []int32{-1, 0, 8000, 12000, 16000, 24000, 44100, 48000, 96000} {
			size := int(opuscc.CompareCustomDecoderSize(mode, channels))
			storage := new(struct {
				State opuscc.OpusT_OpusCustomDecoder
				Tail  [6000]float32
			})
			g := unsafe.Slice((*byte)(unsafe.Pointer(&storage.State)), size+16)
			for i := int(unsafe.Sizeof(storage.State.Fmode)); i < len(g); i++ {
				g[i] = 0xa5
			}
			c := slices.Clone(g)
			code := int32(0)
			func() {
				defer func() {
					if recover() != nil {
						code = -99
					}
				}()
				code = opuscc.Opus_celt_decoder_init(nil, &storage.State, rate, channels)
			}()
			native := nativeCeltState(c, 2, channels, rate, 120, 21, 21)
			normalized := slices.Clone(g)
			clear(normalized[:unsafe.Sizeof(storage.State.Fmode)])
			if code != native || !slices.Equal(normalized, c) {
				t.Fatal(channels, rate, "initialization image", code, native)
			}
		}
	}
	for _, ch := range []int32{-1, 0, 1, 2, 3} {
		if g, c := opuscc.Opus_celt_decoder_init(nil, nil, 44100, ch), nativeCeltState(nil, 2, ch, 44100, 120, 21, 21); g != c {
			t.Fatal("nil/invalid", ch, g, c)
		}
	}
}

func TestCustomDecoderInitAgainstC(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, bands := range []int32{0, 1, 21, 32} {
			for _, overlap := range []int32{0, 12, 120} {
				mode := &opuscc.OpusT_OpusCustomMode{Foverlap: overlap, FnbEBands: bands, FeffEBands: bands - 1}
				size := int(opuscc.CompareCustomDecoderSize(mode, channels))
				storage := new(struct {
					State opuscc.OpusT_OpusCustomDecoder
					Tail  [6000]float32
				})
				g := unsafe.Slice((*byte)(unsafe.Pointer(&storage.State)), size+16)
				for i := int(unsafe.Sizeof(storage.State.Fmode)); i < len(g); i++ {
					g[i] = 0xa5
				}
				c := slices.Clone(g)
				code := opuscc.CompareCustomDecoderInit(&storage.State, mode, channels)
				native := nativeCeltState(c, 1, channels, 48000, overlap, bands, bands-1)
				normalized := slices.Clone(g)
				clear(normalized[:unsafe.Sizeof(storage.State.Fmode)])
				if code != native || !slices.Equal(normalized, c) {
					t.Fatal(channels, bands, overlap, "initialization image", code, native)
				}
			}
		}
	}
	for _, ch := range []int32{-1, 0, 1, 2, 3} {
		if g, c := opuscc.CompareCustomDecoderInit(nil, nil, ch), nativeCeltState(nil, 1, ch, 48000, 120, 21, 21); g != c {
			t.Fatal("nil/invalid", ch, g, c)
		}
	}
}

func TestAllocationDriverErrorAgainstC(t *testing.T) {
	m, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	for _, channels := range []int32{1, 2} {
		for _, start := range []int32{0, 3} {
			a, b := [7][23]int32{}, [7][23]int32{}
			s, cs := [3]int32{77, 88, 99}, [3]int32{77, 88, 99}
			cfg := [12]int32{start, start, 5, 0, 0, 0, 0, channels, 0, 0, 20, 20}
			g, c := opuscc.OpusT_ec_ctx{}, opuscc.OpusT_ec_ctx{}
			r := opuscc.CompareAllocationDriver(nil, m, &a, &s, &cfg, &g)
			cr := nativeAllocationDriver(&c, nil, &b, &cs, &cfg)
			if r != -99 || cr != -99 || a != b || s != cs || g != c {
				t.Fatal("driver assertion order", channels, start, r, cr, a, b, s, cs, g, c)
			}
		}
	}
}

func TestAllocationDriverAgainstC(t *testing.T) {
	m, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	bands := unsafe.Slice(m.FeBands, 22)
	for _, channels := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, start := range []int32{0, 3, 18} {
				for _, budget := range []int32{0, 8, 100, 512, 4096, 12000} {
					for _, trim := range []int32{0, 5, 10} {
						for _, encode := range []int32{0, 1} {
							for _, boost := range []bool{false, true} {
								if boost && budget < 512 {
									continue
								}
								for alias := int32(0); alias <= 7; alias++ {
									if alias != 0 && (channels != 2 || LM != 1 || start != 0 || budget != 4096 || boost) {
										continue
									}
									var a [7][23]int32
									for k := range a {
										a[k][0], a[k][22] = 77, 88
									}
									for j := 0; j < 21; j++ {
										a[3][j+1] = int32(bands[j+1]-bands[j]) * channels << LM << 6
									}
									if boost {
										a[0][min(start+3, 20)+1] = 64
									}
									b := a
									cfg := [12]int32{start, 21, trim, budget, 0, 0, 0, channels, LM, encode, 20, 20}
									s, cs := [3]int32{77, 21, 1}, [3]int32{77, 21, 1}
									data := make([]byte, 256)
									for i := range data {
										data[i] = byte(i*17 + 31)
									}
									cb := slices.Clone(data)
									var g opuscc.OpusT_ec_ctx
									if encode != 0 {
										opuscc.Opus_ec_enc_init(nil, &g, &data[0], 256)
									} else {
										opuscc.Opus_ec_dec_init(nil, &g, &data[0], 256)
									}
									c := g
									c.Fbuf = &cb[0]
									r := opuscc.CompareAllocationDriverAlias(nil, m, &a, &s, &cfg, &g, alias)
									cr := nativeAllocationCall(&c, cb, &b, &cs, &cfg, alias, 1)
									g.Fbuf = nil
									c.Fbuf = nil
									if r != cr || a != b || s != cs || g != c || !slices.Equal(data, cb) {
										t.Fatal("allocation driver", channels, LM, start, budget, trim, encode, boost, alias, r, cr, s, cs, a, b, g, c)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func TestAllocationInterpAliasAgainstC(t *testing.T) {
	m, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	bands := unsafe.Slice(m.FeBands, 22)
	for _, alias := range []int32{1, 2, 3, 4, 5} {
		for _, encode := range []int32{0, 1} {
			var a [7][23]int32
			for k := range a {
				a[k][0], a[k][22] = 77, 88
			}
			for j := 0; j < 21; j++ {
				width := int32(bands[j+1] - bands[j])
				a[1][j+1] = width * 32
				a[2][j+1] = max(16, width*8)
				a[3][j+1] = width * 256
			}
			b := a
			cfg := [12]int32{0, 21, 0, 4096, 8, 40, 8, 2, 1, encode, 20, 20}
			s, cs := [3]int32{77, 21, 1}, [3]int32{77, 21, 1}
			data := make([]byte, 256)
			for i := range data {
				data[i] = byte(i*17 + 31)
			}
			cb := slices.Clone(data)
			var g opuscc.OpusT_ec_ctx
			if encode != 0 {
				opuscc.Opus_ec_enc_init(nil, &g, &data[0], 256)
			} else {
				opuscc.Opus_ec_dec_init(nil, &g, &data[0], 256)
			}
			c := g
			c.Fbuf = &cb[0]
			r := opuscc.CompareAllocationInterpAlias(m, &a, &s, &cfg, &g, alias)
			cr := nativeAllocationInterpAlias(&c, cb, &b, &cs, &cfg, alias)
			g.Fbuf = nil
			c.Fbuf = nil
			if r != cr || a != b || s != cs || g != c || !slices.Equal(data, cb) {
				t.Fatal("interp alias", alias, encode, r, cr, s, cs, a, b)
			}
		}
	}
}

func TestAllocationInterpAgainstC(t *testing.T) {
	m, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	bands := unsafe.Slice(m.FeBands, 22)
	for _, C := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, start := range []int32{0, 3, 18} {
				for _, budget := range []int32{0, 8, 100, 512, 4096, 12000} {
					for _, encode := range []int32{0, 1} {
						for _, reserved := range []bool{false, true} {
							if reserved && (budget < 100 || C == 1) {
								continue
							}
							var a [7][23]int32
							for k := range a {
								a[k][0], a[k][22] = 77, 88
							}
							for j := 0; j < 21; j++ {
								width := int32(bands[j+1] - bands[j])
								a[1][j+1] = width * C << LM << 3
								a[2][j+1] = max(C<<3, width<<LM<<2)
								a[3][j+1] = width * C << LM << 6
							}
							b := a
							cfg := [12]int32{start, 21, start, budget, 0, 0, 0, C, LM, encode, 20, 20}
							if budget >= 8 {
								cfg[4] = 8
							}
							if reserved {
								cfg[5] = 40
								cfg[6] = 8
							}
							s, cs := [3]int32{77, 21, 1}, [3]int32{77, 21, 1}
							data := make([]byte, 256)
							for i := range data {
								data[i] = byte(i*17 + 31)
							}
							cb := slices.Clone(data)
							var g opuscc.OpusT_ec_ctx
							if encode != 0 {
								opuscc.Opus_ec_enc_init(nil, &g, &data[0], 256)
							} else {
								opuscc.Opus_ec_dec_init(nil, &g, &data[0], 256)
							}
							c := g
							c.Fbuf = &cb[0]
							r := opuscc.CompareAllocationInterp(m, &a, &s, &cfg, &g)
							cr := nativeAllocationInterp(&c, cb, &b, &cs, &cfg)
							g.Fbuf = nil
							c.Fbuf = nil
							if r != cr || a != b || s != cs || g != c || !slices.Equal(data, cb) {
								t.Fatal("allocation interp", C, LM, start, budget, encode, reserved, r, cr, s, cs, a, b, g, c)
							}
						}
					}
				}
			}
		}
	}
}

func TestPrefilterFoldRoundingAgainstC(t *testing.T) {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	data := make([]byte, opuscc.CompareCustomDecoderSize(mode, 1))
	st := (*opuscc.OpusT_OpusCustomDecoder)(unsafe.Pointer(&data[0]))
	st.Fchannels = 1
	st.Foverlap = 3
	history := unsafe.Slice(&st.F_decode_mem[0], 2048+3)
	copy(history[2048-120:], []float32{-2422, -2249, -2076})
	nativePrefilterFold(data, 120)
	if bits := math.Float32bits(history[2048-120]); bits != 0xc086cced {
		t.Fatalf("C fold product rounding: %08x want c086cced", bits)
	}
}

func TestPrefilterFoldAgainstC(t *testing.T) {
	mode, err := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	if err != nil {
		t.Fatal(err)
	}
	for _, channels := range []int32{0, 1, 2} {
		for _, N := range []int32{0, 120, 240, 480, 960} {
			for _, overlap := range []int32{0, 1, 2, 3, 119, 120} {
				for _, periods := range [][2]int32{{0, 0}, {15, 15}, {31, 128}, {1024, 1024}} {
					for _, gains := range [][2]float32{{0, 0}, {.25, .5}, {-.25, .75}} {
						for _, taps := range [][2]int32{{0, 0}, {1, 2}, {2, 1}} {
							size := int(opuscc.CompareCustomDecoderSize(mode, max(channels, 1)))
							data := make([]byte, size+16)
							st := (*opuscc.OpusT_OpusCustomDecoder)(unsafe.Pointer(&data[0]))
							st.Fchannels = channels
							st.Foverlap = overlap
							st.Fpostfilter_period_old, st.Fpostfilter_period = periods[0], periods[1]
							st.Fpostfilter_gain_old, st.Fpostfilter_gain = gains[0], gains[1]
							st.Fpostfilter_tapset_old, st.Fpostfilter_tapset = taps[0], taps[1]
							history := unsafe.Slice(&st.F_decode_mem[0], (2048+overlap)*max(channels, 1))
							for i := range history {
								history[i] = float32(i%29-14) * 173
							}
							for i := size; i < len(data); i++ {
								data[i] = 165
							}
							c := slices.Clone(data)
							opuscc.ComparePrefilterFold(nil, data, N)
							nativePrefilterFold(c, N)
							if !slices.Equal(data, c) {
								for i := range data {
									if data[i] != c[i] {
										t.Fatal("prefilter", channels, N, overlap, periods, gains, taps, i, data[i], c[i])
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func TestCeltSynthesisOutputAliasAgainstC(t *testing.T) {
	// Normal stereo calls access the two (possibly overlapping) outputs in
	// separate MDCT invocations; neither output aliases a live spectral input.
	for LM := int32(0); LM <= 3; LM++ {
		N := int32(120) << LM
		for _, transient := range []int32{0, 1} {
			for _, offset := range []int32{0, 17, 60} {
				x, e := make([]float32, 2*N), make([]float32, 42)
				for i := range x {
					x[i] = float32(i%19-9) / 32
				}
				for i := range e {
					e[i] = float32(i%5 - 3)
				}
				g := make([]float32, N+60+offset+2)
				for i := range g {
					g[i] = float32(i%11-5) / 31
				}
				g[0], g[len(g)-1] = 77, 88
				c := slices.Clone(g)
				opuscc.CompareCeltSynthesis(nil, &x[0], &e[0], &g[1], &g[1+offset], 0, 21, 2, 2, transient, LM, 1, 0)
				nativeCeltSynthesis(x, e, c[1:], c[1+offset:], 0, 21, 2, 2, transient, LM, 1, 0)
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
						t.Fatal("synthesis output alias", LM, transient, offset, i, g[i], c[i])
					}
				}
			}
		}
	}
}

func TestCeltSynthesisAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		N := int32(120) << LM
		for _, channels := range [][2]int32{{1, 1}, {2, 2}, {1, 2}, {2, 1}, {0, 0}, {0, 1}} {
			for _, transient := range []int32{0, 1} {
				for _, downsample := range []int32{1, 2, 3, 4, 6} {
					for _, silence := range []int32{0, 1} {
						for _, rangeBands := range [][2]int32{{0, 21}, {0, 20}, {1, 19}, {5, 7}, {21, 21}} {
							x := make([]float32, N*max(channels[0], channels[1], 1)+2)
							energy := make([]float32, 21*max(channels[0], channels[1], 1)+2)
							for i := range x {
								x[i] = float32(i%19-9) / 32
							}
							for i := range energy {
								energy[i] = []float32{-28, -9, .25, 8, 32}[i%5]
							}
							cx, ce := slices.Clone(x), slices.Clone(energy)
							gl, gr := make([]float32, N+62), make([]float32, N+62)
							for i := range gl {
								gl[i] = float32(i%11-5) / 31
								gr[i] = float32(i%13-6) / 37
							}
							gl[0], gl[len(gl)-1], gr[0], gr[len(gr)-1] = 77, 88, 99, 111
							cl, cr := slices.Clone(gl), slices.Clone(gr)
							opuscc.CompareCeltSynthesis(nil, &x[1], &energy[1], &gl[1], &gr[1], rangeBands[0], rangeBands[1], channels[0], channels[1], transient, LM, downsample, silence)
							nativeCeltSynthesis(cx[1:], ce[1:], cl[1:], cr[1:], rangeBands[0], rangeBands[1], channels[0], channels[1], transient, LM, downsample, silence)
							for _, pair := range [][2][]float32{{x, cx}, {energy, ce}, {gl, cl}, {gr, cr}} {
								for i := range pair[0] {
									if math.Float32bits(pair[0][i]) != math.Float32bits(pair[1][i]) {
										t.Fatal("synthesis", LM, channels, transient, downsample, silence, rangeBands, i, pair[0][i], pair[1][i])
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func TestDeemphasisDriverAgainstC(t *testing.T) {
	for _, N := range []int32{0, 1, 2, 8, 120, 960} {
		for _, channels := range []int32{0, 1, 2} {
			for _, downsample := range []int32{1, 2, 3, 4, 6} {
				for _, accum := range []int32{0, 1} {
					for _, coef := range []float32{0, .85, -.25, 1} {
						left, right := make([]float32, N), make([]float32, N)
						for i := range left {
							left[i] = float32(i%17-8) * 1700
							right[i] = float32(i%13-6) * 2300
						}
						g := make([]float32, (N/downsample)*max(channels, 1)+2)
						for i := range g {
							g[i] = float32(i+77) / 31
						}
						c := slices.Clone(g)
						gm, cm := [2]float32{.25, -.5}, [2]float32{.25, -.5}
						opuscc.CompareDeemphasis(nil, unsafe.SliceData(left), unsafe.SliceData(right), &g[1], N, channels, downsample, &coef, &gm[0], accum)
						nativeDeemphasis(left, right, c[1:], N, channels, downsample, coef, &cm, accum)
						for i := range g {
							if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
								t.Fatal("deemphasis", N, channels, downsample, accum, coef, i, g[i], c[i])
							}
						}
						if gm != cm {
							t.Fatal("deemphasis memory", N, channels, downsample, accum, coef, gm, cm)
						}
					}
				}
			}
		}
	}
}

func TestProjectionCtlAgainstC(t *testing.T) {
	requests := []int32{opuscc.OPUS_GET_BANDWIDTH_REQUEST, opuscc.OPUS_GET_SAMPLE_RATE_REQUEST, opuscc.OPUS_GET_GAIN_REQUEST, opuscc.OPUS_GET_LAST_PACKET_DURATION_REQUEST, opuscc.OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_GET_COMPLEXITY_REQUEST, opuscc.OPUS_GET_FINAL_RANGE_REQUEST, opuscc.OPUS_RESET_STATE, opuscc.OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, opuscc.OPUS_SET_GAIN_REQUEST, opuscc.OPUS_SET_COMPLEXITY_REQUEST, opuscc.OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_GET_PITCH_REQUEST, 123456}
	for _, streams := range []int32{1, 3} {
		for _, coupled := range []int32{0, streams} {
			for _, request := range requests {
				for _, value := range []int32{-32769, -1, 0, 1, 2, 3, 10, 11, 32767, 32768} {
					for _, alias := range []int32{-2, -1} {
						channels := streams + coupled
						matrix := make([]byte, 2*channels*channels)
						for i := int32(0); i < channels; i++ {
							matrix[2*(i*channels+i)] = 255
							matrix[2*(i*channels+i)+1] = 127
						}
						g := make([]byte, int(opuscc.Opus_opus_projection_decoder_get_size(nil, channels, streams, coupled))+16)
						st := (*opuscc.OpusT_OpusProjectionDecoder)(unsafe.Pointer(&g[0]))
						if opuscc.Opus_opus_projection_decoder_init(nil, st, 48000, channels, streams, coupled, unsafe.SliceData(matrix), int32(len(matrix))) != 0 {
							t.Fatal("projection fixture")
						}
						msOff := nativeProjectionMultistream(unsafe.Pointer(st))
						normalizeMSModes(g[msOff:], streams, coupled)
						c := slices.Clone(g)
						ret, out := opuscc.CompareProjectionCtl(g, request, value, alias)
						cr, co := nativeProjectionCtl(c, request, value, alias)
						if ret != cr || out != co || !slices.Equal(g, c) {
							t.Fatal("projection CTL", streams, coupled, request, value, alias, ret, cr, out, co)
						}
					}
				}
			}
		}
	}
}

func TestMSCtlAgainstC(t *testing.T) {
	requests := []int32{opuscc.OPUS_GET_BANDWIDTH_REQUEST, opuscc.OPUS_GET_SAMPLE_RATE_REQUEST, opuscc.OPUS_GET_GAIN_REQUEST, opuscc.OPUS_GET_LAST_PACKET_DURATION_REQUEST, opuscc.OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_GET_COMPLEXITY_REQUEST, opuscc.OPUS_GET_FINAL_RANGE_REQUEST, opuscc.OPUS_RESET_STATE, opuscc.OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, opuscc.OPUS_SET_GAIN_REQUEST, opuscc.OPUS_SET_COMPLEXITY_REQUEST, opuscc.OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_GET_PITCH_REQUEST, opuscc.OPUS_GET_IGNORE_EXTENSIONS_REQUEST, 123456}
	for _, streams := range []int32{1, 2, 3} {
		for coupled := int32(0); coupled <= streams; coupled++ {
			for _, request := range requests {
				for _, value := range []int32{-32769, -32768, -1, 0, 1, 2, 3, 10, 11, 32767, 32768} {
					for _, alias := range []int32{-2, -1, int32(unsafe.Offsetof(opuscc.OpusT_OpusMSDecoder{}.Flayout) + unsafe.Offsetof(opuscc.OpusT_ChannelLayout{}.Fnb_streams))} {
						channels := streams + coupled
						mapping := make([]byte, channels)
						for i := range mapping {
							mapping[i] = byte(i)
						}
						g := make([]byte, int(opuscc.Opus_opus_multistream_decoder_get_size(nil, streams, coupled))+16)
						st := (*opuscc.OpusT_OpusMSDecoder)(unsafe.Pointer(&g[0]))
						opuscc.Opus_opus_multistream_decoder_init(nil, st, 48000, channels, streams, coupled, unsafe.SliceData(mapping))
						offset := int((unsafe.Sizeof(*st) + 7) &^ uintptr(7))
						for i := int32(0); i < streams; i++ {
							dec := (*opuscc.OpusT_OpusDecoder)(unsafe.Add(unsafe.Pointer(st), offset))
							dec.FrangeFinal = uint32(0x12345678 + i*37)
							dec.Fdecode_gain = i * 17
							dec.Fbandwidth = 1105
							(*opuscc.OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(dec), dec.Fcelt_dec_offset)).Fmode = nil
							ch := int32(1)
							if i < coupled {
								ch = 2
							}
							offset += int((uint32(opuscc.Opus_opus_decoder_get_size(nil, ch)) + 7) &^ uint32(7))
						}
						c := slices.Clone(g)
						ret, out := opuscc.CompareMSCtl(g, request, value, alias)
						cr, co := nativeMSCtl(c, request, value, alias)
						if ret != cr || out != co || !slices.Equal(g, c) {
							t.Fatal("MS CTL", streams, coupled, request, value, alias, ret, cr, out, co)
						}
					}
				}
			}
		}
	}
}

func TestOpusCtlAgainstC(t *testing.T) {
	requests := []int32{opuscc.OPUS_GET_BANDWIDTH_REQUEST, opuscc.OPUS_SET_COMPLEXITY_REQUEST, opuscc.OPUS_GET_COMPLEXITY_REQUEST, opuscc.OPUS_GET_FINAL_RANGE_REQUEST, opuscc.OPUS_RESET_STATE, opuscc.OPUS_GET_SAMPLE_RATE_REQUEST, opuscc.OPUS_GET_PITCH_REQUEST, opuscc.OPUS_GET_GAIN_REQUEST, opuscc.OPUS_SET_GAIN_REQUEST, opuscc.OPUS_GET_LAST_PACKET_DURATION_REQUEST, opuscc.OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_SET_IGNORE_EXTENSIONS_REQUEST, opuscc.OPUS_GET_IGNORE_EXTENSIONS_REQUEST, 123456}
	for _, ch := range []int32{1, 2} {
		for _, prev := range []int32{opuscc.MODE_SILK_ONLY, opuscc.MODE_CELT_ONLY} {
			for _, request := range requests {
				for _, value := range []int32{-32769, -32768, -1, 0, 1, 10, 11, 32767, 32768} {
					for _, alias := range []int32{-2, -1, int32(unsafe.Offsetof(opuscc.OpusT_OpusDecoder{}.Fdecode_gain))} {
						size := int(opuscc.Opus_opus_decoder_get_size(nil, ch))
						g := make([]byte, size+16)
						st := (*opuscc.OpusT_OpusDecoder)(unsafe.Pointer(&g[0]))
						opuscc.Opus_opus_decoder_init(nil, st, 48000, ch)
						st.Fprev_mode = prev
						st.Fdecode_gain = -99
						st.FDecControl.FprevPitchLag = 55
						st.Fbandwidth = 1105
						st.FrangeFinal = 0xfedcba98
						st.Flast_packet_duration = 960
						celt := (*opuscc.OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(st), st.Fcelt_dec_offset))
						celt.Fpostfilter_period = 37
						celt.Fmode = nil
						c := slices.Clone(g)
						ret, out := opuscc.CompareOpusCtl(g, request, value, alias)
						cr, co := nativeOpusCtl(c, request, value, alias)
						if ret != cr || out != co || !slices.Equal(g, c) {
							t.Fatal("Opus CTL", ch, prev, request, value, alias, ret, cr, out, co)
						}
					}
				}
			}
		}
	}
}

func TestCustomCtlAgainstC(t *testing.T) {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	requests := []int32{opuscc.OPUS_SET_COMPLEXITY_REQUEST, opuscc.OPUS_GET_COMPLEXITY_REQUEST, opuscc.CELT_SET_START_BAND_REQUEST, opuscc.CELT_SET_END_BAND_REQUEST, opuscc.CELT_SET_CHANNELS_REQUEST, opuscc.CELT_GET_AND_CLEAR_ERROR_REQUEST, opuscc.OPUS_GET_LOOKAHEAD_REQUEST, opuscc.OPUS_RESET_STATE, opuscc.OPUS_GET_PITCH_REQUEST, opuscc.CELT_GET_MODE_REQUEST, opuscc.CELT_SET_SIGNALLING_REQUEST, opuscc.OPUS_GET_FINAL_RANGE_REQUEST, opuscc.OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, opuscc.OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, 123456}
	for _, request := range requests {
		for _, value := range []int32{-32769, -1, 0, 1, 2, 10, 11, 20, 21, 32768} {
			// CELT's C state parameter is restrict-qualified; state/output aliases
			// are checked in Go, not submitted to this C oracle.
			for _, alias := range []int32{-2, -1} {
				// Pointer-output mode aliasing uses the mode slot, never a numeric slot.
				actualAlias := alias
				if request == opuscc.CELT_GET_MODE_REQUEST && alias >= 0 {
					actualAlias = int32(unsafe.Offsetof(opuscc.OpusT_OpusCustomDecoder{}.Fmode))
				}
				g := make([]byte, opuscc.CompareCustomDecoderSize(mode, 1)+16)
				st := (*opuscc.OpusT_OpusCustomDecoder)(unsafe.Pointer(&g[0]))
				opuscc.CompareCustomDecoderInit(st, mode, 1)
				st.Ferror1 = 123
				st.Fpostfilter_period = 37
				st.Frng = 0x89abcdef
				st.Fmode = nil
				c := slices.Clone(g)
				ret, out := opuscc.CompareCustomCtl(g, request, value, actualAlias)
				cr, co := nativeCustomCtl(c, request, value, actualAlias)
				if ret != cr || out != co || !slices.Equal(g, c) {
					t.Fatal("custom CTL", request, value, actualAlias, ret, cr, out, co)
				}
			}
		}
	}
}

func TestCeltResetAgainstC(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, bands := range []int32{0, 1, 21, 32} {
			for _, overlap := range []int32{0, 12, 120} {
				mode := &opuscc.OpusT_OpusCustomMode{Foverlap: overlap, FnbEBands: bands}
				size := int(opuscc.CompareCustomDecoderSize(mode, channels))
				storage := new(struct {
					State opuscc.OpusT_OpusCustomDecoder
					Tail  [6000]float32
				})
				g := unsafe.Slice((*byte)(unsafe.Pointer(&storage.State)), size+16)
				for i := int(unsafe.Sizeof(storage.State.Fmode)); i < len(g); i++ {
					g[i] = 0xa5
				}
				storage.State.Fmode = mode
				storage.State.Foverlap = overlap
				storage.State.Fchannels = channels
				c := slices.Clone(g)
				opuscc.CompareCeltReset(&storage.State)
				nativeCeltState(c, 0, channels, 48000, overlap, bands, bands)
				normalized := slices.Clone(g)
				clear(normalized[:unsafe.Sizeof(storage.State.Fmode)])
				if !slices.Equal(normalized, c) {
					t.Fatal(channels, bands, overlap, "reset image")
				}
			}
		}
	}
}

func TestCustomModeAgainstC(t *testing.T) {
	for _, rate := range []int32{-1, 0, 8000, 16000, 24000, 44100, 48000, 96000} {
		for frame := int32(0); frame <= 2000; frame++ {
			mode, err := opuscc.Opus_opus_custom_mode_create(nil, rate, frame)
			code, values := nativeCustomMode(rate, frame)
			if (mode != nil) != (code == 0) || (err != nil) != (code != 0) {
				t.Fatal(rate, frame, mode, err, code)
			}
			if mode != nil && [7]int32{mode.FFs, mode.Foverlap, mode.FnbEBands, mode.FeffEBands, mode.FshortMdctSize, mode.FnbShortMdcts, mode.FmaxLM} != values {
				t.Fatal("mode fields", rate, frame)
			}
		}
	}
}

func TestCeltValidationAgainstC(t *testing.T) {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	base := opuscc.OpusT_OpusCustomDecoder{Fmode: mode, Foverlap: 120, Fchannels: 2, Fstream_channels: 1, Fdownsample: 1, Fend: 21}
	for field := 0; field < 13; field++ {
		for _, v := range []int32{-1, 0, 1, 2, 3, 15, 16, 17, 18, 20, 21, 22, 100, 120, 1023, 1024} {
			st := base
			switch field {
			case 0:
				st.Foverlap = v
			case 1:
				st.Fend = v
			case 2:
				st.Fchannels = v
			case 3:
				st.Fstream_channels = v
			case 4:
				st.Fdownsample = v
			case 5:
				st.Fstart = v
			case 6:
				st.Farch = v
			case 7:
				st.Flast_pitch_index = v
			case 8:
				st.Fpostfilter_period = v
			case 9:
				st.Fpostfilter_period_old = v
			case 10:
				st.Fpostfilter_tapset = v
			case 11:
				st.Fpostfilter_tapset_old = v
			case 12:
				st.Fmode = nil
				if v != 0 {
					st.Fmode = &opuscc.OpusT_OpusCustomMode{}
				}
			}
			before := st
			g := opuscc.CompareCeltValidation(&st)
			c := nativeCeltValidation(&st)
			if g != c || st != before {
				t.Fatal(field, v, g, c)
			}
		}
	}
}

func TestOpusValidationAgainstC(t *testing.T) {
	base := opuscc.OpusT_OpusDecoder{Fchannels: 2, FFs: 48000, Fstream_channels: 1}
	base.FDecControl.FAPI_sampleRate = 48000
	base.FDecControl.FnChannelsAPI = 2
	for field := 0; field < 9; field++ {
		for _, value := range []int32{-1, 0, 1, 2, 3, 10, 20, 40, 60, 8000, 12000, 16000, 24000, 48000, 96000} {
			st := base
			switch field {
			case 0:
				st.Fchannels = value
			case 1:
				st.FFs = value
			case 2:
				st.FDecControl.FAPI_sampleRate = value
			case 3:
				st.FDecControl.FinternalSampleRate = value
			case 4:
				st.FDecControl.FnChannelsAPI = value
			case 5:
				st.FDecControl.FnChannelsInternal = value
			case 6:
				st.FDecControl.FpayloadSize_ms = value
			case 7:
				st.Farch = value
			case 8:
				st.Fstream_channels = value
			}
			before := st
			g := opuscc.CompareOpusValidation(&st)
			c := nativeOpusValidation(&st)
			if g != c || st != before {
				t.Fatal(field, value, g, c)
			}
		}
	}
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		for _, ch := range []int32{1, 2} {
			st := base
			st.FFs = rate
			st.FDecControl.FAPI_sampleRate = rate
			st.Fchannels = ch
			st.FDecControl.FnChannelsAPI = ch
			if opuscc.CompareOpusValidation(&st) || nativeOpusValidation(&st) {
				t.Fatal(rate, ch)
			}
		}
	}
}
