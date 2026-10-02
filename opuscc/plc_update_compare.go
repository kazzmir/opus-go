//go:build compareopus

package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
)

func ComparePLCDispatch(tls *libc.TLS, dec *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, frame *int16, lost, arch int32) {
	silk_PLC(tls, dec, control, frame, lost, arch)
}

func ComparePLCLPC(decoder *OpusT_silk_decoder_state, history []int32, A *[MAX_LPC_ORDER]int16, frame *int16, gain int32) {
	silkPLCLPC(nil, decoder, history, A, frame, gain)
}

func ComparePLCPCM(sample, gain int32) int16 { return silkPLCPCM(sample, gain) }

func ComparePLCNoise(prediction int32, random []int32, index int32, scale int16) int32 {
	return silkPLCNoise(prediction, random, index, scale)
}

func ComparePLCDecay(coefficients *[LTP_ORDER]int16, gain int32) { silkPLCDecayLTP(coefficients, gain) }

func ComparePLCUpdate(dec *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control) {
	silk_PLC_update(nil, dec, control)
}
