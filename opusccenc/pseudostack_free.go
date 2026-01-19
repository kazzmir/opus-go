package opusccenc

import (
	"unsafe"

	libc "opusgo/libcshim"
)

// Matches the ccgo-only pseudostack storage in celt/stack_alloc.h.
const opusPseudostackTLSKey = uint32(0x6f707573) // 'opus'

// FreePseudostackTLS frees the per-thread pseudostack scratch buffer, if any.
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
	const offScratch = uintptr(0)
	offGlobal := uintptr(unsafe.Sizeof(uintptr(0)))

	scratch := *(*uintptr)(unsafe.Pointer(p + offScratch))
	if scratch != 0 {
		libc.Xfree(tls, scratch)
		*(*uintptr)(unsafe.Pointer(p + offScratch)) = 0
		*(*uintptr)(unsafe.Pointer(p + offGlobal)) = 0
	}
	libc.Xfree(tls, p)
	_ = libc.Xpthread_setspecific(tls, opusPseudostackTLSKey, 0)
}
