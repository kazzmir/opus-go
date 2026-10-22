package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"
	"weak"
)

// Only explicit flat ABI/registry fixtures use this old representation.
func legacyRegisteredProjectionTestDecoder(tls *libc.TLS, rate, channels, streams, coupled int32, matrix *byte, matrixBytes int32) (*OpusT_OpusProjectionDecoder, error) {
	owned := opusProjectionDecoderLegacyOwnedStorage(channels, streams, coupled)
	st := libc.RegisterGoAllocation(tls, owned.Decoder)
	if st == nil {
		return nil, opusErrorFromCode(-7)
	}
	if ret := opusProjectionDecoderInitWithStorage(tls, st, rate, channels, streams, coupled, matrix, matrixBytes, &owned); ret != 0 {
		libc.RemoveGoAllocation(tls, st)
		return nil, opusErrorFromCode(ret)
	}
	opusBindProjectionCodecViews(tls, owned)
	return st, nil
}

//go:noinline
func droppedProjectionGoOwner(t *testing.T, tls *libc.TLS) weak.Pointer[OpusProjectionDecoder] {
	matrix := [18]byte{}
	st, err := Opus_opus_projection_decoder_create_typed(48000, 3, 2, 1, &matrix[0], 18)
	if err != nil {
		t.Fatal(err)
	}
	if libc.RegisteredGoAllocation[OpusProjectionDecoder](tls, uintptr(unsafe.Pointer(st))) != nil {
		t.Fatal("normal projection registered")
	}
	return weak.Make(st)
}
func TestProjectionGoCollectionPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	w := droppedProjectionGoOwner(t, tls)
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	if w.Value() != nil {
		t.Fatal("Go projection retained without destructor")
	}
	runtime.KeepAlive(tls)
}

func TestProjectionGoDecodePointers(t *testing.T) {
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	for method := 0; method < 3; method++ {
		tls := libc.NewTLS()
		matrix := [2]byte{0, 0x80}
		st, err := Opus_opus_projection_decoder_create_typed(48000, 1, 1, 0, &matrix[0], 2)
		if err != nil {
			t.Fatal(err)
		}
		flat, err := legacyRegisteredProjectionTestDecoder(tls, 48000, 1, 1, 0, &matrix[0], 2)
		if err != nil {
			t.Fatal(err)
		}
		for _, data := range [][]byte{packet, nil} {
			entropyInitGrowStack(12)
			runtime.GC()
			var a, b int32
			switch method {
			case 0:
				x, y := make([]float32, 2882), make([]float32, 2882)
				x[0], y[0], x[2881], y[2881] = 77, 77, 88, 88
				a = Opus_opus_projection_decode_float_typed(nil, st, unsafe.SliceData(data), int32(len(data)), &x[1], 2880, 0)
				b = Opus_opus_projection_decode_float_typed(tls, flat, unsafe.SliceData(data), int32(len(data)), &y[1], 2880, 0)
				for i := range x {
					if math.Float32bits(x[i]) != math.Float32bits(y[i]) {
						t.Fatal("float bits", i)
					}
				}
				if x[0] != 77 || x[2881] != 88 {
					t.Fatal("float guard")
				}
			case 1:
				x, y := make([]int16, 2882), make([]int16, 2882)
				x[0], y[0], x[2881], y[2881] = 77, 77, 88, 88
				a = Opus_opus_projection_decode(nil, st, unsafe.SliceData(data), int32(len(data)), &x[1], 2880, 0)
				b = Opus_opus_projection_decode_typed(tls, flat, unsafe.SliceData(data), int32(len(data)), &y[1], 2880, 0)
				if !slices.Equal(x, y) || x[0] != 77 || x[2881] != 88 {
					t.Fatal("int16/guards")
				}
			case 2:
				x, y := make([]int32, 2882), make([]int32, 2882)
				x[0], y[0], x[2881], y[2881] = 77, 77, 88, 88
				a = Opus_opus_projection_decode24(nil, st, unsafe.SliceData(data), int32(len(data)), &x[1], 2880, 0)
				b = Opus_opus_projection_decode24_typed(tls, flat, unsafe.SliceData(data), int32(len(data)), &y[1], 2880, 0)
				if !slices.Equal(x, y) || x[0] != 77 || x[2881] != 88 {
					t.Fatal("int24/guards")
				}
			}
			if a != 2880 || a != b {
				t.Fatal("decode status", a, b)
			}
			var ra, rb uint32
			if Opus_opus_projection_decoder_ctl_typed(nil, st, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: &ra}) != 0 || Opus_opus_projection_decoder_ctl_typed(tls, flat, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: &rb}) != 0 || ra != rb {
				t.Fatal("range")
			}
		}
		Opus_opus_projection_decoder_destroy_typed(tls, flat)
		tls.Close()
	}
}

func TestProjectionGoChildrenPointers(t *testing.T) {
	matrix := [18]byte{0, 0x80, 0xff, 0x7f, 1, 0}
	st, err := Opus_opus_projection_decoder_create_typed(48000, 3, 2, 1, &matrix[0], 18)
	if err != nil {
		t.Fatal(err)
	}
	if unsafe.Offsetof(st.matrix.Coefficients)-unsafe.Offsetof(st.matrix.Header) != 16 {
		t.Fatal("declared coefficient offset")
	}
	w := weak.Make(st)
	var child *OpusT_OpusDecoder
	if Opus_opus_projection_decoder_ctl_typed(nil, st, OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, OpusDecoderCtlArgs{Value: 1, Decoder: &child}) != 0 {
		t.Fatal("child")
	}
	st = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if w.Value() == nil || w.Value().matrix.Coefficients[0] != -32768 || w.Value().matrix.Coefficients[1] != 32767 {
		t.Fatal("child lost projection matrix")
	}
	runtime.KeepAlive(child)
}
