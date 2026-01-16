package opuscc

import (
	"unsafe"

	"modernc.org/libc"
)

// Matches the ccgo-only pseudostack storage in celt/stack_alloc.h.
const opusPseudostackTLSKey libc.Tpthread_key_t = 0x6f707573 // 'opus'

type opusPseudostackState struct {
	scratchPtr  uintptr
	globalStack uintptr
}

// FreePseudostackTLS frees the per-decoder pseudostack scratch buffer, if any.
//
// Safe to call multiple times. Call before tls.Close().
func FreePseudostackTLS(tls *libc.TLS) {
	if tls == nil {
		return
	}
	p := libc.Xpthread_getspecific(tls, opusPseudostackTLSKey)
	if p == 0 {
		return
	}
	st := (*opusPseudostackState)(unsafe.Pointer(p))
	if st.scratchPtr != 0 {
		libc.Xfree(tls, st.scratchPtr)
		st.scratchPtr = 0
		st.globalStack = 0
	}
	libc.Xfree(tls, p)
	_ = libc.Xpthread_setspecific(tls, opusPseudostackTLSKey, 0)
}
