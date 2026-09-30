//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
)

func TestOpusValidationAgainstC(t *testing.T) {
	base := opuscc.OpusT_OpusDecoder{Fchannels: 2, FFs: 48000, Fstream_channels: 1}
	base.FDecControl.FAPI_sampleRate = 48000
	base.FDecControl.FnChannelsAPI = 2
	for field := 0; field < 9; field++ {
		for _, value := range []int32{-1, 0, 1, 2, 3, 10, 20, 40, 60, 8000, 12000, 16000, 24000, 48000, 96000} {
			st := base
			switch field {
			case 0:
				st.Fchannels = value
			case 1:
				st.FFs = value
			case 2:
				st.FDecControl.FAPI_sampleRate = value
			case 3:
				st.FDecControl.FinternalSampleRate = value
			case 4:
				st.FDecControl.FnChannelsAPI = value
			case 5:
				st.FDecControl.FnChannelsInternal = value
			case 6:
				st.FDecControl.FpayloadSize_ms = value
			case 7:
				st.Farch = value
			case 8:
				st.Fstream_channels = value
			}
			before := st
			g := opuscc.CompareOpusValidation(&st)
			c := nativeOpusValidation(&st)
			if g != c || st != before {
				t.Fatal(field, value, g, c)
			}
		}
	}
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		for _, ch := range []int32{1, 2} {
			st := base
			st.FFs = rate
			st.FDecControl.FAPI_sampleRate = rate
			st.Fchannels = ch
			st.FDecControl.FnChannelsAPI = ch
			if opuscc.CompareOpusValidation(&st) || nativeOpusValidation(&st) {
				t.Fatal(rate, ch)
			}
		}
	}
}
