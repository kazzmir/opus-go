//go:build compareopus

package opuscc

import "unsafe"

func CompareExtensionGenerate(data *byte, length int32, exts *OpusT_opus_extension_data, count, frames, pad int32) (result int32) {
	defer func() {
		if recover() != nil {
			result = -99
		}
	}()
	return Opus_opus_packet_extensions_generate(nil, uintptr(unsafe.Pointer(data)), length, exts, count, frames, pad)
}

func CompareExtensionRepeat(iter *OpusT_OpusExtensionIterator, ext *OpusT_opus_extension_data) int32 {
	return opus_extension_iterator_next_repeat(nil, iter, ext)
}

func CompareWriteExtension(data *byte, capacity, pos, id, length int32, payload *byte, last int32) int32 {
	return write_extension(nil, data, capacity, pos, id, length, payload, last)
}

func CompareWritePayload(data *byte, capacity, pos, id, length int32, payload *byte, last int32) int32 {
	return write_extension_payload(nil, data, capacity, pos, id, length, payload, last)
}

func CompareSkipExtension(data **byte, length int32, header *int32) int32 {
	return skip_extension(nil, data, length, header)
}

func CompareSkipPayload(data **byte, length int32, header *int32, id, trailing int32) int32 {
	return skip_extension_payload(nil, data, length, header, id, trailing)
}
