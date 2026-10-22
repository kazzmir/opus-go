package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

// A matrix has at most 65004 coefficient bytes, rounded to an eight-byte extent.
// This bounded int16 array is actual Go coefficient storage, not an opaque image.
type opusGoProjectionMatrix struct {
	Header       OpusT_MappingMatrix
	Padding      [4]byte
	Coefficients [65008 / 2]int16
}
type opusProjectionGoHeader OpusT_OpusProjectionDecoder

// OpusProjectionDecoderState distinguishes ordinary Go owners from flat ABI state.
type OpusProjectionDecoderState interface {
	projectionMS(*libc.TLS) OpusMSDecoderState
	projectionMatrix(*libc.TLS) *OpusT_MappingMatrix
	projectionDestroy(*libc.TLS)
}

func (st *OpusT_OpusProjectionDecoder) projectionMS(tls *libc.TLS) OpusMSDecoderState {
	return get_multistream_decoder(tls, st)
}
func (st *OpusT_OpusProjectionDecoder) projectionMatrix(tls *libc.TLS) *OpusT_MappingMatrix {
	return get_dec_demixing_matrix(tls, st)
}
func (st *OpusT_OpusProjectionDecoder) projectionDestroy(tls *libc.TLS) {
	opusReleaseProjectionCodecViews(tls, st)
	libc.XfreePointer(tls, unsafe.Pointer(st))
}
func (st *OpusProjectionDecoder) projectionMS(_ *libc.TLS) OpusMSDecoderState { return st.multistream }
func (st *OpusProjectionDecoder) projectionMatrix(_ *libc.TLS) *OpusT_MappingMatrix {
	return &st.matrix.Header
}
func (st *OpusProjectionDecoder) projectionDestroy(_ *libc.TLS) {}

func projectionGoCreate(rate, channels, streams, coupled int32, matrix *byte, matrixBytes int32) (*OpusProjectionDecoder, error) {
	// Keep size-before-init ordering; nil TLS is no longer a simulated malloc failure.
	if Opus_opus_projection_decoder_get_size(nil, channels, streams, coupled) == 0 {
		return nil, opusErrorFromCode(-7)
	}
	inputs := streams + coupled
	count := inputs * channels
	if count*2 != matrixBytes {
		return nil, opusErrorFromCode(-1)
	}
	src := unsafe.Slice(matrix, matrixBytes)
	coefficients := make([]int16, count)
	for i := int32(0); i < count; i++ {
		s := int32(src[2*i+1])<<8 | int32(src[2*i])
		coefficients[i] = int16(((s & 0xffff) ^ 0x8000) - 0x8000)
	}
	st := new(OpusProjectionDecoder)
	st.Fdemixing_matrix_size_in_bytes = Opus_mapping_matrix_get_size(channels, inputs)
	if st.Fdemixing_matrix_size_in_bytes == 0 {
		return nil, opusErrorFromCode(-1)
	}
	Opus_mapping_matrix_init(nil, &st.matrix.Header, channels, inputs, 0, unsafe.SliceData(coefficients), matrixBytes)
	var mapping [255]byte
	for i := int32(0); i < channels; i++ {
		mapping[i] = byte(i)
	}
	if channels < 1 || channels > 255 || streams < 1 || coupled < 0 || coupled > streams || streams > 255-coupled {
		return nil, opusErrorFromCode(-1)
	}
	st.multistream = opusMSDecoderOwnedStorage(streams, coupled)
	st.multistream.projection = st
	if ret := opusMSDecoderInitWithCodecs(nil, st.multistream.opusMSHeader(), rate, channels, streams, coupled, &mapping[0], nil, st.multistream.codecs); ret != 0 {
		return nil, opusErrorFromCode(ret)
	}
	return st, nil
}
