//go:build compareopus

package opuscc

func ComparePLCUpdate(dec *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control) {
	silk_PLC_update(nil, dec, control)
}
