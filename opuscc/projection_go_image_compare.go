//go:build compareopus

package opuscc

import "unsafe"

// CompareProjectionGoImage serializes logical ABI components, not the Go owner
// representation (which deliberately owns a separately allocated multistream).
func CompareProjectionGoImage(st *OpusProjectionDecoder) []byte {
	ms := CompareMSGoImage(st.multistream)
	offset := 8 + int(st.Fdemixing_matrix_size_in_bytes)
	image := make([]byte, offset+len(ms))
	copy(image, unsafe.Slice((*byte)(unsafe.Pointer(&st.opusProjectionGoHeader)), int(unsafe.Sizeof(st.opusProjectionGoHeader))))
	copy(image[8:offset], unsafe.Slice((*byte)(unsafe.Pointer(&st.matrix.Header)), int(st.Fdemixing_matrix_size_in_bytes)))
	copy(image[offset:], ms)
	return image
}
