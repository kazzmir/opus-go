package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestCompositeSizeProjection(t *testing.T) {
	for _, shape := range [][3]int32{{1, 1, 0}, {2, 1, 1}, {4, 2, 2}, {0, 1, 0}, {-1, 1, 0}, {256, 1, 0}, {255, 255, 255}, {1, 0, 0}, {1, 2, -1}, {1, 1, 2}, {1, 2147483647, 0}} {
		channels, streams, coupled := shape[0], shape[1], shape[2]
		var want int32
		matrix := Opus_mapping_matrix_get_size(nil, streams+coupled, channels)
		if matrix != 0 {
			decoder := Opus_opus_multistream_decoder_get_size(nil, streams, coupled)
			if decoder != 0 {
				want = 8 + matrix + decoder
			}
		}
		if got := Opus_opus_projection_decoder_get_size(nil, channels, streams, coupled); got != want {
			t.Fatal("projection composition", shape, got, want)
		}
	}
}

func TestCompositeSizeMatrix(t *testing.T) {
	for _, shape := range [][2]int32{{0, 0}, {1, 1}, {255, 127}, {255, 128}, {200, 162}, {256, 0}, {0, 256}, {-1, 3}, {-300, 2}, {-2147483648, 2}, {-2147483648, -1}} {
		rows, cols := shape[0], shape[1]
		var want int32
		if rows <= 255 && cols <= 255 {
			size := int32(uint64(uint32(rows*cols)) * 2)
			if size <= 65004 {
				want = 16 + int32((uint32(size)+7)/8*8)
			}
		}
		if got := Opus_mapping_matrix_get_size(nil, rows, cols); got != want {
			t.Fatal("mapping size", shape, got, want)
		}
	}
}

func TestCompositeSizeMultistream(t *testing.T) {
	for _, shape := range [][2]int32{{-1, -1}, {0, 0}, {1, 0}, {1, 1}, {2, 0}, {2, 1}, {2, 2}, {5, 2}, {1, 2}, {4, -1}, {2147483647, 0}, {2147483647, 2147483647}} {
		streams, coupled := shape[0], shape[1]
		var want int32
		if streams >= 1 && coupled >= 0 && coupled <= streams {
			mono := Opus_opus_decoder_get_size(nil, 1)
			stereo := Opus_opus_decoder_get_size(nil, 2)
			want = 272 + coupled*int32((uint32(stereo)+7)/8*8) + (streams-coupled)*int32((uint32(mono)+7)/8*8)
		}
		if got := Opus_opus_multistream_decoder_get_size(nil, streams, coupled); got != want {
			t.Fatal("multistream composition", shape, got, want)
		}
	}
}

func TestCompositeSizeAlignment(t *testing.T) {
	for _, n := range []int32{-2147483648, -2147483647, -9, -8, -7, -1, 0, 1, 7, 8, 9, 2147483640, 2147483647} {
		want := int32((uint32(n) + uint32(8) - 1) / 8 * 8)
		if got := opusAlignSize8(n); got != want {
			t.Fatal("unsigned size alignment", n, got, want)
		}
	}
	for _, channels := range []int32{-2147483648, -1, 0, 3, 2147483647} {
		if Opus_opus_decoder_get_size(nil, channels) != 0 {
			t.Fatal("invalid channel size", channels)
		}
	}
	for _, channels := range []int32{1, 2} {
		var silk int32
		if Opus_silk_Get_Decoder_Size(nil, &silk) != 0 {
			t.Fatal("silk size")
		}
		want := int32(104) + int32((uint32(silk)+7)/8*8) + Opus_celt_decoder_get_size(nil, channels)
		if got := Opus_opus_decoder_get_size(nil, channels); got != want {
			t.Fatal("decoder composition", channels, got, want)
		}
	}
}

func TestOpusDecoderGetSizeLocalSilkSize(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	mono := Opus_opus_decoder_get_size(tls, 1)
	stereo := Opus_opus_decoder_get_size(tls, 2)
	if mono <= 0 || stereo <= mono {
		t.Fatalf("decoder sizes: mono=%d stereo=%d", mono, stereo)
	}
	if got := Opus_opus_decoder_get_size(tls, 0); got != 0 {
		t.Fatalf("zero-channel size: got %d, want 0", got)
	}
	if got := Opus_opus_decoder_get_size(tls, 3); got != 0 {
		t.Fatalf("three-channel size: got %d, want 0", got)
	}

	state := make([]byte, stereo)
	if got := Opus_opus_decoder_init(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(&state[0])), 48000, 2); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d, want %d", got, OPUS_OK)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(&state[0]))
	if decoder.Fsilk_dec_offset <= 0 || decoder.Fcelt_dec_offset <= decoder.Fsilk_dec_offset || decoder.Fframe_size != 120 {
		t.Fatalf("decoder layout: silk=%d celt=%d frame=%d", decoder.Fsilk_dec_offset, decoder.Fcelt_dec_offset, decoder.Fframe_size)
	}
}
