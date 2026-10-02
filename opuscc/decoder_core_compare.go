//go:build compareopus

package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

// Full core still consumes legacy TLS scratch and output/pulse cursors.
//
//go:uintptrescapes
func CompareDecodeCore(tls *libc.TLS, decoder *OpusT_silk_decoder_state, control uintptr, output, pulses uintptr) {
	silk_decode_core(tls, decoder, control, output, pulses, 0)
}
func CompareDecodeCoreHistory(decoder *OpusT_silk_decoder_state, frame []int16) {
	silkDecodeCoreHistory(decoder, unsafe.SliceData(frame))
}
