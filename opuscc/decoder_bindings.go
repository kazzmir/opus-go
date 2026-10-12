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

type opusMSCodecBinding struct {
	Views            opusMSStorageViews
	Streams, Coupled int32
}

func opusBindMSCodecViews(tls *libc.TLS, views opusMSStorageViews) {
	for _, codec := range views.Codecs {
		libc.RegisterGoAllocationAlias(tls, views.Decoder, codec.Decoder)
		opusBindCodecViews(tls, codec)
	}
	libc.SetGoAllocationMetadata(tls, views.Decoder, opusMSCodecBinding{views, views.Decoder.Flayout.Fnb_streams, views.Decoder.Flayout.Fnb_coupled_streams})
}

func opusMSDecoderBoundAt(tls *libc.TLS, st *OpusT_OpusMSDecoder, offset uint) *OpusT_OpusDecoder {
	binding, ok := libc.GoAllocationMetadata[OpusT_OpusMSDecoder, opusMSCodecBinding](tls, st)
	if ok && binding.Views.Decoder == st && binding.Streams == st.Flayout.Fnb_streams && binding.Coupled == st.Flayout.Fnb_coupled_streams {
		for i, at := range binding.Views.Offsets {
			if at == offset {
				return binding.Views.Codecs[i].Decoder
			}
		}
	}
	return opusMSDecoderAt(st, offset)
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

type opusProjectionCodecBinding struct{ Views opusProjectionStorageViews }

func opusBindProjectionCodecViews(tls *libc.TLS, views opusProjectionStorageViews) {
	libc.RegisterGoAllocationAlias(tls, views.Decoder, views.Multistream.Decoder)
	opusBindMSCodecViews(tls, views.Multistream)
	libc.SetGoAllocationMetadata(tls, views.Decoder, opusProjectionCodecBinding{views})
}

func opusReleaseProjectionCodecViews(tls *libc.TLS, st *OpusT_OpusProjectionDecoder) {
	binding, ok := libc.GoAllocationMetadata[OpusT_OpusProjectionDecoder, opusProjectionCodecBinding](tls, st)
	if !ok {
		return
	}
	for _, codec := range binding.Views.Multistream.Codecs {
		libc.RemoveGoAllocation(tls, codec.Decoder)
	}
	libc.RemoveGoAllocation(tls, binding.Views.Multistream.Decoder)
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
