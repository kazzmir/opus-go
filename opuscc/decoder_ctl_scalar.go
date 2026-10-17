package opuscc

// A real Go scalar cell may back either C int32/uint32 output convention.
// Conversions preserve its word bits without reinterpreting pointer types.
// This private alternative is opt-in; existing public I32/U32 slots and nil
// validation retain their old meaning and take precedence.
// Keep the public argument record's original fields/layout/source shape.
// Only internal dispatch carries the extra concrete scalar endpoint.
type opusCtlArgs struct {
	OpusDecoderCtlArgs
	scalar *opusCtlScalar
}

type opusCtlScalar struct {
	signed   *int32
	unsigned *uint32
}

func (a opusCtlArgs) hasI32() bool {
	return a.I32 != nil || a.scalar != nil && (a.scalar.signed != nil || a.scalar.unsigned != nil)
}
func (a opusCtlArgs) hasU32() bool {
	return a.U32 != nil || a.scalar != nil && (a.scalar.signed != nil || a.scalar.unsigned != nil)
}
func (a opusCtlArgs) storeI32(value int32) {
	if a.I32 != nil {
		*a.I32 = value
	} else if a.scalar.signed != nil {
		*a.scalar.signed = value
	} else {
		*a.scalar.unsigned = uint32(value)
	}
}
func (a opusCtlArgs) storeU32(value uint32) {
	if a.U32 != nil {
		*a.U32 = value
	} else if a.scalar.unsigned != nil {
		*a.scalar.unsigned = value
	} else {
		*a.scalar.signed = int32(value)
	}
}
func (a opusCtlArgs) loadU32() uint32 {
	if a.U32 != nil {
		return *a.U32
	}
	if a.scalar.unsigned != nil {
		return *a.scalar.unsigned
	}
	return uint32(*a.scalar.signed)
}
