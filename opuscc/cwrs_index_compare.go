//go:build compareopus

package opuscc

func CompareCWRSIndex(n int32, y *int32) uint32 { return icwrs(nil, n, y) }
