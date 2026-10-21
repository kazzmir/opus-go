//go:build compareopus

package opuscc

import "unsafe"

// CompareMSGoImage serializes logical ABI components for the native oracle.
// A normal Go multistream object itself is deliberately NOT a flat ABI image.
func CompareMSGoImage(st *OpusMSDecoder) []byte {
	image := make([]byte, Opus_opus_multistream_decoder_get_size(nil, st.Flayout.Fnb_streams, st.Flayout.Fnb_coupled_streams))
	copy(image, unsafe.Slice((*byte)(unsafe.Pointer(st.opusMSHeader())), int(unsafe.Sizeof(st.opusMSGoHeader))))
	for i, codec := range st.codecs {
		channels := int32(1)
		if int32(i) < st.Flayout.Fnb_coupled_streams {
			channels = 2
		}
		n := int(Opus_opus_decoder_get_size(nil, channels))
		copy(image[int(st.offsets[i]):int(st.offsets[i])+n], unsafe.Slice((*byte)(unsafe.Pointer(codec.Decoder)), n))
	}
	return image
}
