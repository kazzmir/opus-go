//go:build compareopus

package opuscc

import "unsafe"
import libc "github.com/kazzmir/opus-go/libcshim"

// ComparePLCEnergy exposes the internal helper only to the native-comparison build.
// The normal codec build does not export this test bridge.
func CompareCeltDecodeSilence(ec *OpusT_ec_ctx, total int32) (int32, int32) {
	return celtDecodeSilence(nil, ec, total)
}
func CompareCeltDecodePacketError(state *OpusT_OpusCustomDecoder, ec *OpusT_ec_ctx, length int32) int32 {
	return celtDecodePacketError(state, ec, length)
}
func CompareCeltDecodeDeemphasis(state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, outputs **float32, pcm *float32, N, channels, accum int32) {
	celtDecodeDeemphasis(nil, state, mode, outputs, pcm, N, channels, accum)
}
func CompareCeltDecodePrefilterImage(data []byte, N int32) {
	state := (*OpusT_OpusCustomDecoder)(unsafe.Pointer(unsafe.SliceData(data)))
	state.Fmode = &mode48000_960_120
	celtDecodePrefilter(nil, state, N)
	state.Fmode = nil
}
func CompareCeltDecodeEnergyMergeMono(e *float32, bands int32) { celtDecodeEnergyMergeMono(e, bands) }
func CompareCeltDecodePostfilter(state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, outputs **float32, channels, N, LM, period int32, gain float32, tapset, overlap int32) {
	celtDecodePostfilter(nil, state, mode, outputs, channels, N, LM, period, gain, tapset, overlap)
}
func CompareCeltDecodePostfilterTail(state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, output *float32, N, period int32, gain float32, tapset, overlap int32) {
	celtDecodePostfilterTail(nil, state, mode, output, N, period, gain, tapset, overlap)
}
func CompareCeltDecodePostfilterFirst(state *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, output *float32, overlap int32) {
	celtDecodePostfilterFirst(nil, state, mode, output, overlap)
}
func CompareCeltDecodePostfilterClamp(state *OpusT_OpusCustomDecoder) {
	celtDecodePostfilterClamp(state)
}
func CompareCeltDecodePacketFinish(state *OpusT_OpusCustomDecoder) { celtDecodePacketFinish(state) }
func CompareCeltDecodeRecoverEnergy(state *OpusT_OpusCustomDecoder, e, l, p *float32, bands, start, end, LM, intra int32) {
	celtDecodeRecoverEnergy(state, e, l, p, bands, start, end, LM, intra)
}
func CompareCeltDecodeRecoveryBand(e, l, p *float32, missing int32, safety float32) {
	celtDecodeRecoveryBand(e, l, p, missing, safety)
}
func CompareCeltDecodeRecoverySafety(state *OpusT_OpusCustomDecoder, LM int32) (int32, float32) {
	return celtDecodeRecoverySafety(state, LM)
}
func CompareCeltDecodePostfilterFinish(state *OpusT_OpusCustomDecoder, period int32, gain float32, tapset, LM int32) {
	celtDecodePostfilterFinish(state, period, gain, tapset, LM)
}
func CompareCeltDecodeBoosts(bands *int16, cap, offsets *int32, start, end, C, LM, total int32, ec *OpusT_ec_ctx) (int32, int32) {
	return celtDecodeBoosts(nil, bands, cap, offsets, start, end, C, LM, total, ec)
}
func CompareCeltDecodeSilenceEnergy(energy *float32, bands, channels int32) {
	celtDecodeSilenceEnergy(energy, bands, channels)
}
func CompareCeltDecodeHistoryMove(history *float32, N, length int32) {
	celtDecodeHistoryMove(history, N, length)
}
func CompareCeltDecodeMaskStorage(bands, channels int32) []byte {
	return celtDecodeMaskStorage(bands, channels)
}
func CompareCeltDecodeSpectrumStorage(N, channels int32) []float32 {
	return celtDecodeSpectrumStorage(N, channels)
}
func CompareCeltDecodePriorityStorage(bands int32) []int32 { return celtDecodePriorityStorage(bands) }
func CompareCeltDecodePulseStorage(bands int32) []int32    { return celtDecodePulseStorage(bands) }
func CompareCeltDecodeFineStorage(bands int32) []int32     { return celtDecodeFineStorage(bands) }
func CompareCeltDecodeOffsetsStorage(bands int32) []int32  { return celtDecodeOffsetsStorage(bands) }
func CompareCeltDecodeOffsetsAllocation(mode *OpusT_OpusCustomMode, offsets, caps []int32, LM int32, ec *OpusT_ec_ctx) (coded int32, values [3]int32, outputs [3][21]int32) {
	pulses := celtDecodePulseStorage(21)
	priority := celtDecodePriorityStorage(21)
	coded = clt_compute_allocation(nil, mode, 0, 21, unsafe.SliceData(offsets), unsafe.SliceData(caps), 5, &values[1], &values[2], 512, &values[0], unsafe.SliceData(pulses), &outputs[1][0], unsafe.SliceData(priority), 2, LM, ec, 0, 0, 0)
	copy(outputs[0][:], pulses)
	copy(outputs[2][:], priority)
	return
}
func CompareCeltDecodeCapsStorage(mode *OpusT_OpusCustomMode, bands, LM, channels int32) []int32 {
	return celtDecodeCapsStorage(nil, mode, bands, LM, channels)
}
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
