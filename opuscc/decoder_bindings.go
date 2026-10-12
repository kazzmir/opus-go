package opuscc

import libc "github.com/kazzmir/opus-go/libcshim"

// Bindings retain actual reflection-field views, not recovered byte images.
// Offset snapshots are guards only: mutated headers must retain layout behavior.
type opusCodecBinding struct {
	Views                  opusDecoderStorageViews
	SilkOffset, CeltOffset int32
}

func opusBindCodecViews(tls *libc.TLS, views opusDecoderStorageViews) {
	st := views.Decoder
	libc.SetGoAllocationMetadata(tls, st, opusCodecBinding{views, st.Fsilk_dec_offset, st.Fcelt_dec_offset})
}

type opusMSCodecBinding struct{ Views opusMSStorageViews }

func opusBindMSCodecViews(tls *libc.TLS, views opusMSStorageViews) {
	for _, codec := range views.Codecs {
		libc.RegisterGoAllocation(tls, codec.Decoder)
		opusBindCodecViews(tls, codec)
	}
	libc.SetGoAllocationMetadata(tls, views.Decoder, opusMSCodecBinding{views})
}

func opusReleaseMSCodecViews(tls *libc.TLS, st *OpusT_OpusMSDecoder) {
	binding, ok := libc.GoAllocationMetadata[OpusT_OpusMSDecoder, opusMSCodecBinding](tls, st)
	if !ok {
		return
	}
	for _, codec := range binding.Views.Codecs {
		libc.RemoveGoAllocation(tls, codec.Decoder)
	}
}

func opusCodecSilkState(tls *libc.TLS, st *OpusT_OpusDecoder) *OpusT_silk_decoder {
	binding, ok := libc.GoAllocationMetadata[OpusT_OpusDecoder, opusCodecBinding](tls, st)
	if ok && binding.Views.Decoder == st && binding.Views.Silk != nil && binding.SilkOffset == st.Fsilk_dec_offset {
		return binding.Views.Silk
	}
	return opusFrameSilkState(st)
}

func opusCodecCeltState(tls *libc.TLS, st *OpusT_OpusDecoder) *OpusT_OpusCustomDecoder {
	binding, ok := libc.GoAllocationMetadata[OpusT_OpusDecoder, opusCodecBinding](tls, st)
	if ok && binding.Views.Decoder == st && binding.Views.Celt != nil && binding.CeltOffset == st.Fcelt_dec_offset {
		return binding.Views.Celt
	}
	return opusFrameCeltState(st)
}
