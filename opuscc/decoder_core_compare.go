//go:build compareopus

package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

// Full core accepts nil TLS and uses typed owners plus Go scratch.
func CompareDecodeCore(tls *libc.TLS, decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, output *int16, pulses *int16) {
	silk_decode_core(tls, decoder, control, output, pulses, 0)
}
func CompareDecodeCoreResidual(excitation, prediction int32) int32 {
	return silkDecodeCoreResidual(excitation, prediction)
}
func CompareDecodeCorePCM(sample, gain int32) int16 { return silkDecodeCorePCM(sample, gain) }
func CompareDecodeCoreLTPWhiten(history []int32, samples []int16, index, memory, lag, gain int32) {
	silkDecodeCoreLTPWhiten(history, samples, index, memory, lag, gain)
}
func CompareDecodeCoreLTPScale(history []int32, index, lag, gain int32) {
	silkDecodeCoreLTPScale(history, index, lag, gain)
}
func CompareDecodeCoreLTPPrediction(history []int32, index int32, b *[5]int16) int32 {
	return silkDecodeCoreLTPPrediction(history, index, b)
}
func CompareDecodeCoreWhiten(decoder *OpusT_silk_decoder_state, samples []int16, a *[16]int16, start, k int32) {
	silkDecodeCoreWhiten(nil, decoder, samples, a, start, k, 0)
}
func CompareDecodeCoreCoefficients(decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, k int32, snapshot *[16]int16) (*[16]int16, *[5]int16) {
	return silkDecodeCoreCoefficients(decoder, control, k, snapshot)
}
func CompareDecodeCoreExcitation(decoder *OpusT_silk_decoder_state, pulses *int16, offset int32) int32 {
	return silkDecodeCoreExcitation(decoder, pulses, offset)
}
func CompareDecodeCoreTransition(decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, k int32) bool {
	return silkDecodeCoreTransition(decoder, control, k)
}
func CompareDecodeAPICount(decoder *OpusT_silk_decoder, control *OpusT_silk_DecControlStruct, count *int32, internal int32) {
	silkDecodeAPICount(decoder, control, count, internal)
}
func CompareDecodeAPIOutput(out *float32, input *int16, count *int32, channel, stride int32) {
	silkDecodeAPIOutput(out, input, count, channel, stride)
}
func CompareDecodeAPIDuplicate(out *float32, count *int32) { silkDecodeAPIDuplicate(out, count) }
func CompareDecodeAPILBRR(decoder *OpusT_silk_decoder_state, ec *OpusT_ec_ctx) {
	silkDecodeAPILBRR(nil, decoder, ec)
}
func CompareDecodeAPIStartStereo(decoder *OpusT_silk_decoder, control *OpusT_silk_DecControlStruct) {
	silkDecodeAPIStartStereo(decoder, control)
}
func CompareDecodeAPIPacketStart(decoder *OpusT_silk_decoder, channels *int32, flag int32) {
	silkDecodeAPIPacketStart(decoder, channels, flag)
}
func CompareDecodeFrame(tls *libc.TLS, decoder *OpusT_silk_decoder_state, ec *OpusT_ec_ctx, output *int16, count *int32, lost, cond int32) int32 {
	return silk_decode_frame(tls, decoder, ec, output, count, lost, cond, 0)
}
func CompareDecodeFrameFinish(decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, count *int32, length int32) {
	silkDecodeFrameFinish(decoder, control, count, length)
}
func CompareDecodeFrameHistory(decoder *OpusT_silk_decoder_state, frame []int16) {
	silkDecodeFrameHistory(decoder, unsafe.SliceData(frame))
}
func CompareDecodeCoreHistory(decoder *OpusT_silk_decoder_state, frame []int16) {
	silkDecodeCoreHistory(decoder, unsafe.SliceData(frame))
}
