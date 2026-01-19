package opusccenc

import (
	"fmt"
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
		return "opusccenc: (nil error)"
	}
	if e.Message != "" {
		return fmt.Sprintf("opusccenc: %s (%d)", e.Message, e.Code)
	}
	return fmt.Sprintf("opusccenc: error code %d", e.Code)
}

func opusErrorFromCode(code int32) error {
	if code == OPUS_OK {
		return nil
	}
	return &OpusError{Code: code, Message: Opus_opus_strerror(code)}
}
