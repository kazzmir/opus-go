//go:build compareopus

package opuscc

import "unsafe"
import libc "github.com/kazzmir/opus-go/libcshim"

// ComparePLCEnergy exposes the internal helper only to the native-comparison build.
// The normal codec build does not export this test bridge.
func CompareCeltDecodeTFStorage(bands, start, end, transient, LM int32, ec *OpusT_ec_ctx) []int32 {
	return celtDecodeTFStorage(nil, bands, start, end, transient, LM, ec)
}
func CompareCeltDecodeEnergyClear(energy, log, previous *float32, bands, start, end int32) {
	celtDecodeEnergyClear(energy, log, previous, bands, start, end)
}
func CompareCeltDecodeEnergyBackground(state *OpusT_OpusCustomDecoder, background, energy *float32, bands, M int32) {
	celtDecodeEnergyBackground(state, background, energy, bands, M)
}
func CompareCeltDecodeEnergyLogs(energy, log, previous *float32, bands, transient int32) {
	celtDecodeEnergyLogs(energy, log, previous, bands, transient)
}
func CompareCeltDecodeEnergyMono(energy *float32, bands int32) { celtDecodeEnergyMono(energy, bands) }
func CompareCeltPLCLost(tls *libc.TLS, state *OpusT_OpusCustomDecoder, N, LM int32) {
	celt_decode_lost(tls, state, N, LM)
}
func CompareCeltPLCHistoryViews(state *OpusT_OpusCustomDecoder, overlap, bands, channels, N int32) ([2][]float32, [2]*float32, *float32, *float32, *float32) {
	return celtPLCHistoryViews(state, overlap, bands, channels, N)
}
func CompareCeltPLCMode(state *OpusT_OpusCustomDecoder) (*OpusT_OpusCustomMode, int32, int32, *int16) {
	return celtPLCMode(state)
}
func CompareCeltPLCDispatch(state *OpusT_OpusCustomDecoder) (int32, int32, int32) {
	return celtPLCDispatch(state)
}
func CompareCeltPLCNoise(state *OpusT_OpusCustomDecoder, bands *int16, spectrum *float32, N, start, end, LM, channels int32) {
	celtPLCNoise(nil, state, bands, spectrum, N, start, end, LM, channels)
}
func CompareCeltPLCFinish(state *OpusT_OpusCustomDecoder, loss, LM, frameType int32) {
	celtPLCFinish(state, loss, LM, frameType)
}
func CompareCeltPLCLPCHistory(memory *[24]float32, history *float32, size, N int32) {
	celtPLCLPCHistory(memory, history, size, N)
}
func CompareCeltPLCExtrapolate(history, exc *float32, size, period, N, overlap, pitch int32, fade, decay float32) float32 {
	return celtPLCExtrapolate(history, exc, size, period, N, overlap, pitch, fade, decay)
}
func CompareCeltPLCExcitationHistory(exc, history *float32, size, period int32) {
	celtPLCExcitationHistory(exc, history, size, period)
}
func CompareCeltPLCSynthesisAttenuate(output, window *float32, length, overlap int32, s1 float32) {
	celtPLCSynthesisAttenuate(output, window, length, overlap, s1)
}
func CompareCeltPLCExcitationDecay(exc *float32, period, length int32) float32 {
	return celtPLCExcitationDecay(exc, period, length)
}
func CompareCeltPLCLagWindow(ac *[25]float32) { celtPLCLagWindow(ac) }
func CompareCeltPLCDecay(a, b *float32, bands, start, end, channels, loss int32) {
	celtPLCDecay(a, b, bands, start, end, channels, loss)
}
func ComparePLCEnergy(excitation []int32, gains [2]int32, subframeLength, subframes int32) [4]int32 {
	var result [4]int32
	silk_PLC_energy(nil, &result[0], &result[1], &result[2], &result[3], unsafe.SliceData(excitation), &gains, subframeLength, subframes)
	return result
}
