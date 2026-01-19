package opuscc

import (
	"fmt"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// OpusError represents a libopus error code returned by the ccgo-transpiled API.
//
// Code matches one of the OPUS_* constants (e.g. OPUS_BAD_ARG).
// Message is derived from opus_strerror when a TLS is available.
type OpusError struct {
	Code    int32
	Message string
}

func (e *OpusError) Error() string {
	if e == nil {
		return "opuscc: (nil error)"
	}
	if e.Message != "" {
		return fmt.Sprintf("opuscc: %s (%d)", e.Message, e.Code)
	}
	return fmt.Sprintf("opuscc: error code %d", e.Code)
}

func opusErrorFromCode(code int32) error {
	if code == OPUS_OK {
		return nil
	}
	msg := ""
	p := Opus_opus_strerror(code)
	if p != 0 {
		msg = libc.GoString(p)
	}
	return &OpusError{Code: code, Message: msg}
}
