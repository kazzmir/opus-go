//go:build compareopus

package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

// Full core still consumes legacy TLS scratch and output/pulse cursors.
//
//go:uintptrescapes
func CompareDecodeCore(tls *libc.TLS, decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, output uintptr, pulses *int16) {
	silk_decode_core(tls, decoder, control, output, pulses, 0)
}
func CompareDecodeCoreExcitation(decoder *OpusT_silk_decoder_state, pulses *int16, offset int32) int32 {
	return silkDecodeCoreExcitation(decoder, pulses, offset)
}
func CompareDecodeCoreTransition(decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, k int32) bool {
	return silkDecodeCoreTransition(decoder, control, k)
}
func CompareDecodeCoreHistory(decoder *OpusT_silk_decoder_state, frame []int16) {
	silkDecodeCoreHistory(decoder, unsafe.SliceData(frame))
}
