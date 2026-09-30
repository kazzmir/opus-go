//go:build compareopus

package opuscc

func CompareWritePayload(data *byte, capacity, pos, id, length int32, payload *byte, last int32) int32 {
	return write_extension_payload(nil, data, capacity, pos, id, length, payload, last)
}

func CompareSkipExtension(data **byte, length int32, header *int32) int32 {
	return skip_extension(nil, data, length, header)
}

func CompareSkipPayload(data **byte, length int32, header *int32, id, trailing int32) int32 {
	return skip_extension_payload(nil, data, length, header, id, trailing)
}
