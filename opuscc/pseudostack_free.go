package opuscc

import (
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
	// struct opusPseudostackState { uintptr scratchPtr; uintptr globalStack; }
	const offScratch = 0
	const offGlobal = uintptr(libc.PtrSize)

	scratch := libc.LoadUintptrAt(p, offScratch)
	if scratch != 0 {
		libc.Xfree(tls, scratch)
		libc.StoreUintptrAt(p, offScratch, 0)
		libc.StoreUintptrAt(p, offGlobal, 0)
	}
	libc.Xfree(tls, p)
	_ = libc.Xpthread_setspecific(tls, opusPseudostackTLSKey, 0)
}
