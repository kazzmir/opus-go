package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

// OpusMSDecoderState supports normal Go owners and explicit flat ABI states.
// The two storage representations never get reinterpreted as one another.
type OpusMSDecoderState interface {
	opusMSHeader() *OpusT_OpusMSDecoder
	opusMSChild(*libc.TLS, uint) *OpusT_OpusDecoder
}

// Keep the normal owner's header private: exposing it as a flat decoder would
// invite callers to use C child-address arithmetic on a noncontiguous Go object.
type opusMSGoHeader OpusT_OpusMSDecoder

func (st *OpusT_OpusMSDecoder) opusMSHeader() *OpusT_OpusMSDecoder { return st }
func (st *OpusT_OpusMSDecoder) opusMSChild(tls *libc.TLS, offset uint) *OpusT_OpusDecoder {
	return opusMSDecoderBoundAt(tls, st, offset)
}
func (st *OpusMSDecoder) opusMSHeader() *OpusT_OpusMSDecoder {
	if st == nil {
		return nil
	}
	return (*OpusT_OpusMSDecoder)(&st.opusMSGoHeader)
}
func (st *OpusMSDecoder) opusMSChild(_ *libc.TLS, offset uint) *OpusT_OpusDecoder {
	for i, at := range st.offsets {
		if at == offset {
			return st.codecs[i].Decoder
		}
	}
	// A Go child is an object, not arbitrary bytes at a mutated C address.
	panic("opus: invalid Go multistream child coordinate")
}
func opusMSHeaderOf(st OpusMSDecoderState) *OpusT_OpusMSDecoder {
	if st == nil {
		return nil
	}
	return st.opusMSHeader()
}
func opusMSChildOf(tls *libc.TLS, st OpusMSDecoderState, offset uint) *OpusT_OpusDecoder {
	if st == nil {
		return opusMSDecoderBoundAt(tls, nil, offset)
	}
	return st.opusMSChild(tls, offset)
}

// Shared single/multistream Go structs retain the canonical decoder ABI offsets. Only the
// standard 48k/960 CELT mode is used by the Opus decoder (21 bands, overlap120).
// Tail words match opus_custom_decoder_get_size, not a cast byte allocation.
type opusGoDecoderPrefix struct {
	Decoder       OpusT_OpusDecoder
	HeaderPadding [((unsafe.Sizeof(OpusT_OpusDecoder{}) + 7) &^ 7) - unsafe.Sizeof(OpusT_OpusDecoder{})]byte
	Silk          OpusT_silk_decoder
	SilkPadding   [((unsafe.Sizeof(OpusT_silk_decoder{}) + 7) &^ 7) - unsafe.Sizeof(OpusT_silk_decoder{})]byte
	Celt          OpusT_OpusCustomDecoder
}

const opusGoMonoTailWords = (DEC_PITCH_BUF_SIZE + 120) - 1 + 21*8 + CELT_LPC_ORDER
const opusGoStereoTailWords = 2*(DEC_PITCH_BUF_SIZE+120) - 1 + 21*8 + 2*CELT_LPC_ORDER

type opusGoMonoBody struct {
	opusGoDecoderPrefix
	Tail [opusGoMonoTailWords]float32
}
type opusGoStereoBody struct {
	opusGoDecoderPrefix
	Tail [opusGoStereoTailWords]float32
}

// Back references are outside the ABI image and outside Clear. A returned child
// therefore owns its storage AND its multistream parent/siblings; GC handles cycles.
type opusGoMonoCell struct {
	Body  opusGoMonoBody
	Owner *OpusMSDecoder
}
type opusGoStereoCell struct {
	Body  opusGoStereoBody
	Owner *OpusMSDecoder
}

// A fixed 32KiB scanned allocation avoids the small-object malloc header that
// can shift ordinary Go cells by eight bytes. Supported Go allocators give this
// allocation page-aligned storage; this is not a language-level alignment rule.
// Padding is an explicit memory cost, outside the codec ABI and Clear extent.
type opusGoMonoAllocation struct {
	Cell    opusGoMonoCell
	Padding [32768 - unsafe.Sizeof(opusGoMonoCell{})]byte
}

type opusGoStereoAllocation struct {
	Cell    opusGoStereoCell
	Padding [32768 - unsafe.Sizeof(opusGoStereoCell{})]byte
}

// Keep this allocation on the heap, not an inlined caller's stack.
//
//go:noinline
func opusGoMonoStorage(parent *OpusMSDecoder) opusDecoderStorageViews {
	cell := &new(opusGoMonoAllocation).Cell
	cell.Owner = parent
	return opusDecoderStorageViews{Decoder: &cell.Body.Decoder, Silk: &cell.Body.Silk, Celt: &cell.Body.Celt, Clear: func() {
		cell.Body.opusGoDecoderPrefix = opusGoDecoderPrefix{}
		clear(cell.Body.Tail[:])
	}}
}

// Keep this allocation on the heap, not an inlined caller's stack.
//
//go:noinline
func opusGoStereoStorage(parent *OpusMSDecoder) opusDecoderStorageViews {
	cell := &new(opusGoStereoAllocation).Cell
	cell.Owner = parent
	return opusDecoderStorageViews{Decoder: &cell.Body.Decoder, Silk: &cell.Body.Silk, Celt: &cell.Body.Celt, Clear: func() {
		cell.Body.opusGoDecoderPrefix = opusGoDecoderPrefix{}
		clear(cell.Body.Tail[:])
	}}
}

func opusDecoderGoStorageCheck(channels int32) {
	if channels != 1 && channels != 2 {
		panic("opus Go decoder channel geometry")
	}
	if mode48000_960_120.Foverlap != 120 || mode48000_960_120.FnbEBands != 21 {
		panic("opus Go decoder standard mode geometry")
	}
	prefix := opusGoDecoderPrefix{}
	if uint(unsafe.Offsetof(prefix.Silk)) != uint(opusAlignSize8(int32(unsafe.Sizeof(prefix.Decoder)))) || uint(unsafe.Offsetof(prefix.Celt)) != uint(unsafe.Offsetof(prefix.Silk))+uint(opusAlignSize8(int32(unsafe.Sizeof(prefix.Silk)))) {
		panic("opus Go child field offsets")
	}
	words := opusGoMonoTailWords
	if channels == 2 {
		words = opusGoStereoTailWords
	}
	if int(unsafe.Sizeof(prefix))+words*4 != int(Opus_opus_decoder_get_size(nil, channels)) {
		panic("opus Go child extent")
	}
}

func opusMSDecoderOwnedStorage(streams, coupled int32) *OpusMSDecoder {
	if mode48000_960_120.Foverlap != 120 || mode48000_960_120.FnbEBands != 21 {
		panic("opus Go decoder standard mode geometry")
	}
	st := &OpusMSDecoder{codecs: make([]opusDecoderStorageViews, streams), offsets: make([]uint, streams)}
	offset := uint(opusAlignSize8(int32(unsafe.Sizeof(OpusT_OpusMSDecoder{}))))
	for i := range st.codecs {
		channels := int32(1)
		if int32(i) < coupled {
			channels = 2
			st.codecs[i] = opusGoStereoStorage(st)
		} else {
			st.codecs[i] = opusGoMonoStorage(st)
		}
		// Verify the real static Go fields against the unchanged codec ABI.
		opusDecoderGoStorageCheck(channels)
		st.offsets[i] = offset
		offset += uint(opusAlignSize8(Opus_opus_decoder_get_size(nil, channels)))
	}
	return st
}
