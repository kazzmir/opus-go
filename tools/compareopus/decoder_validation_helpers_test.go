//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
)

func TestCustomDecoderSizeAgainstC(t *testing.T) {
	for _, overlap := range []int32{0, 60, 120, 240} {
		for _, bands := range []int32{0, 1, 21, 25} {
			for _, channels := range []int32{0, 1, 2} {
				mode := opuscc.OpusT_OpusCustomMode{Foverlap: overlap, FnbEBands: bands}
				before := mode
				g := opuscc.CompareCustomDecoderSize(&mode, channels)
				c := nativeCustomDecoderSize(overlap, bands, channels)
				if g != c || mode != before {
					t.Fatal(overlap, bands, channels, g, c)
				}
			}
		}
	}
}

func TestMSValidationAgainstC(t *testing.T) {
	for _, channels := range []int32{0, 1, 2, 255} {
		for _, streams := range []int32{0, 1, 2, 127, 255, 256} {
			for _, coupled := range []int32{0, 1, 127} {
				for _, mapping := range []uint8{0, 1, 2, 254, 255} {
					st := opuscc.OpusT_OpusMSDecoder{}
					st.Flayout.Fnb_channels = channels
					st.Flayout.Fnb_streams = streams
					st.Flayout.Fnb_coupled_streams = coupled
					for i := range st.Flayout.Fmapping {
						st.Flayout.Fmapping[i] = mapping
					}
					before := st
					g := opuscc.CompareMSValidation(&st)
					c := nativeMSValidation(&st)
					if g != c || st != before {
						t.Fatal(channels, streams, coupled, mapping, g, c)
					}
				}
			}
		}
	}
}

func TestCeltValidationAgainstC(t *testing.T) {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	base := opuscc.OpusT_OpusCustomDecoder{Fmode: mode, Foverlap: 120, Fchannels: 2, Fstream_channels: 1, Fdownsample: 1, Fend: 21}
	for field := 0; field < 13; field++ {
		for _, v := range []int32{-1, 0, 1, 2, 3, 15, 16, 17, 18, 20, 21, 22, 100, 120, 1023, 1024} {
			st := base
			switch field {
			case 0:
				st.Foverlap = v
			case 1:
				st.Fend = v
			case 2:
				st.Fchannels = v
			case 3:
				st.Fstream_channels = v
			case 4:
				st.Fdownsample = v
			case 5:
				st.Fstart = v
			case 6:
				st.Farch = v
			case 7:
				st.Flast_pitch_index = v
			case 8:
				st.Fpostfilter_period = v
			case 9:
				st.Fpostfilter_period_old = v
			case 10:
				st.Fpostfilter_tapset = v
			case 11:
				st.Fpostfilter_tapset_old = v
			case 12:
				st.Fmode = uintptr(v)
			}
			before := st
			g := opuscc.CompareCeltValidation(&st)
			c := nativeCeltValidation(&st)
			if g != c || st != before {
				t.Fatal(field, v, g, c)
			}
		}
	}
}

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
