package opus

import (
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// cBuf is scratch memory the transpiled libopus code can safely be handed
// a uintptr to. A caller's slice can't be: converting it to uintptr hides
// the pointer from escape analysis, so the slice may live on the goroutine
// stack, and the transpiled code's deep call chains can grow - move - that
// stack mid-call. Go fixes up real pointers when it moves a stack, but not
// uintptrs, so libopus would read from and write to the stack's old copy
// (observed as encoded packets left all-zero). cBuf memory comes from the
// shim's malloc, which keeps it reachable from the TLS and never moves it.
type cBuf struct {
	p uintptr
	n int
}

// ensure makes the buffer at least n bytes and returns its address.
func (b *cBuf) ensure(tls *libc.TLS, n int) uintptr {
	if n > b.n {
		b.free(tls)
		b.p = libc.Xmalloc(tls, uint64(n))
		b.n = n
	}
	return b.p
}

func (b *cBuf) free(tls *libc.TLS) {
	if b.p != 0 {
		libc.Xfree(tls, b.p)
	}
	b.p, b.n = 0, 0
}

// cSlice views n elements of T at p - memory from a cBuf.
func cSlice[T any](p uintptr, n int) []T {
	return unsafe.Slice((*T)(unsafe.Pointer(p)), n)
}

// copyIn copies src into b (growing it as needed) and returns its
// address, or 0 for an empty src (libopus's "no packet" / packet-loss
// convention).
func copyIn[T any](tls *libc.TLS, b *cBuf, src []T) uintptr {
	if len(src) == 0 {
		return 0
	}
	var zero T
	p := b.ensure(tls, len(src)*int(unsafe.Sizeof(zero)))
	copy(cSlice[T](p, len(src)), src)
	return p
}
