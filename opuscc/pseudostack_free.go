package opuscc

import (
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// Matches the ccgo-only pseudostack storage in celt/stack_alloc.h.
const opusPseudostackTLSKey libc.Tpthread_key_t = 0x6f707573 // 'opus'

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
	state := libc.RegisteredGoAllocation[OpusT_opus_ccgo_pseudostack_state](tls, p)
	if state == nil {
		// Explicit compatibility boundary for externally supplied ABI headers.
		state = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(p))
	}
	scratch := state.Fscratch_ptr
	if scratch != nil {
		libc.XfreePointer(tls, unsafe.Pointer(scratch))
		state.Fscratch_ptr = nil
		state.Fglobal_stack = nil
	}
	libc.Xfree(tls, p)
	_ = libc.Xpthread_setspecific(tls, opusPseudostackTLSKey, 0)
}
