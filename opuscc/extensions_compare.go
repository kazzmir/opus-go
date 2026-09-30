//go:build compareopus

package opuscc

func CompareSkipExtension(data **byte, length int32, header *int32) int32 {
	return skip_extension(nil, data, length, header)
}

func CompareSkipPayload(data **byte, length int32, header *int32, id, trailing int32) int32 {
	return skip_extension_payload(nil, data, length, header, id, trailing)
}
