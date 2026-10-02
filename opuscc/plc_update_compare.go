//go:build compareopus

package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

func ComparePLCDispatch(tls *libc.TLS, dec *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, frame *int16, lost, arch int32) {
	silk_PLC(tls, dec, control, uintptr(unsafe.Pointer(frame)), lost, arch)
}

func ComparePLCUpdate(dec *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control) {
	silk_PLC_update(nil, dec, control)
}
