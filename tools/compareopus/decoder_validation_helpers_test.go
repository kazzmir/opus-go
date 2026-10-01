//go:build compareopus && cgo

package main

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
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
