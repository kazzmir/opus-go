//go:build compareopus

package opuscc

func CompareFFTFactor(n int32, factors *[64]int32) int { return kf_factor(nil, n, factors) }
