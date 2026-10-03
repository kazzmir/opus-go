//go:build compareopus

package opuscc

import "unsafe"

// ComparePLCEnergy exposes the internal helper only to the native-comparison build.
// The normal codec build does not export this test bridge.
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
