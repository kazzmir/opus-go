package opuscc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestProjectionDecoderInitCReference(t *testing.T) {
	f, err := os.Open("testdata/projection_init_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		t.Run(fmt.Sprint(line), func(t *testing.T) {
			input := strings.NewReader(scanner.Text())
			var ch, streams, coupled, fs, delta, wantRet, wantSize int32
			if _, err := fmt.Fscan(input, &ch, &streams, &coupled, &fs, &delta, &wantRet, &wantSize); err != nil {
				t.Fatal(err)
			}
			tls := libc.NewTLS()
			defer tls.Close()
			size := Opus_opus_projection_decoder_get_size(tls, ch, streams, coupled)
			// The fixture sizes include one amd64 decoder allocation per stream.
			if wantSize > 0 {
				wantSize += streams * decoderReferenceLayoutDelta()
			}
			if size != wantSize {
				t.Fatalf("size %d, want %d", size, wantSize)
			}
			if size == 0 {
				size = 4096
			}
			st := libc.Xmalloc(tls, uint64(size))
			defer libc.Xfree(tls, st)
			libc.Xmemset(tls, st, 0, uint64(size))
			matrix := libc.Xmalloc(tls, 128)
			defer libc.Xfree(tls, matrix)
			values := []int16{0, 32767, -32768, -1, 12345, -23456, 1, -2}
			count := ch * (streams + coupled)
			bytes := unsafe.Slice((*byte)(unsafe.Pointer(matrix)), 128)
			for i := int32(0); i < count; i++ {
				v := uint16(values[i%8])
				bytes[2*i] = byte(v)
				bytes[2*i+1] = byte(v >> 8)
			}
			if ret := Opus_opus_projection_decoder_init(tls, st, fs, ch, streams, coupled, matrix, 2*count+delta); ret != wantRet {
				t.Fatalf("return %d, want %d", ret, wantRet)
			}
			if wantRet != 0 {
				return
			}
			var matrixSize, rows, cols, gain int32
			var wantHash uint64
			if _, err := fmt.Fscan(input, &matrixSize, &rows, &cols, &gain, &wantHash); err != nil {
				t.Fatal(err)
			}
			d := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st))
			m := (*OpusT_MappingMatrix)(unsafe.Pointer(get_dec_demixing_matrix(tls, st)))
			if d.Fdemixing_matrix_size_in_bytes != matrixSize || m.Frows != rows || m.Fcols != cols || m.Fgain != gain {
				t.Fatalf("matrix header mismatch: %+v, bytes %d", m, d.Fdemixing_matrix_size_in_bytes)
			}
			hash := uint64(14695981039346656037)
			data := Opus_mapping_matrix_get_data(tls, get_dec_demixing_matrix(tls, st))
			for _, b := range unsafe.Slice((*byte)(unsafe.Pointer(data)), int(count)*2) {
				hash ^= uint64(b)
				hash *= 1099511628211
			}
			if hash != wantHash {
				t.Fatalf("matrix hash %d, want %d", hash, wantHash)
			}
			ms := (*OpusT_OpusMSDecoder)(unsafe.Pointer(get_multistream_decoder(tls, st)))
			if ms.Flayout.Fnb_channels != ch || ms.Flayout.Fnb_streams != streams || ms.Flayout.Fnb_coupled_streams != coupled {
				t.Fatal("multistream layout mismatch")
			}
			for i := int32(0); i < ch; i++ {
				var want uint8
				if _, err := fmt.Fscan(input, &want); err != nil {
					t.Fatal(err)
				}
				if ms.Flayout.Fmapping[i] != want {
					t.Fatalf("mapping[%d] = %d, want %d", i, ms.Flayout.Fmapping[i], want)
				}
			}
		})
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
