//go:build compareopus

package opuscc

func CompareSmoothFade(a, b, out *float32, overlap, channels int32, window *float32, rate int32) {
	smooth_fade(nil, a, b, out, overlap, channels, window, rate)
}
