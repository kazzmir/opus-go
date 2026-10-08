package opus

import (
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// cBuf owns registered, stable shim scratch through a GC-visible byte pointer.
// Typed decoder calls use the pointer directly. Legacy encoder calls still use
// explicit integer adapters: stack slices must not cross those raw interfaces.
// The allocation is bytes, not a generally GC-scanned C struct image.
type cBuf struct {
	p *byte
	n int
}

func (b *cBuf) ensurePointer(tls *libc.TLS, n int) *byte {
	if n > b.n {
		b.free(tls)
		b.p = (*byte)(libc.XmallocPointer(tls, uint64(n)))
		b.n = n
	}
	return b.p
}

// Legacy encoder address boundary.
func (b *cBuf) ensure(tls *libc.TLS, n int) uintptr {
	return uintptr(unsafe.Pointer(b.ensurePointer(tls, n)))
}

func (b *cBuf) free(tls *libc.TLS) {
	if b.p != nil {
		libc.XfreePointer(tls, unsafe.Pointer(b.p))
	}
	b.p, b.n = nil, 0
}

// Reinterpret registered scratch at the concrete element-type boundary.
func cPointer[T any](p *byte) *T {
	return (*T)(unsafe.Pointer(p))
}

func cBufferSlice[T any](p *byte, n int) []T {
	return unsafe.Slice(cPointer[T](p), n)
}

// Legacy encoder view boundary.
func cSlice[T any](p uintptr, n int) []T {
	return cBufferSlice[T]((*byte)(unsafe.Pointer(p)), n)
}

func copyInPointer[T any](tls *libc.TLS, b *cBuf, src []T) *byte {
	if len(src) == 0 {
		return nil // packet-loss convention; retain any existing scratch owner
	}
	var zero T
	p := b.ensurePointer(tls, len(src)*int(unsafe.Sizeof(zero)))
	copy(cBufferSlice[T](p, len(src)), src)
	return p
}

// Legacy encoder input boundary.
func copyIn[T any](tls *libc.TLS, b *cBuf, src []T) uintptr {
	return uintptr(unsafe.Pointer(copyInPointer(tls, b, src)))
}
