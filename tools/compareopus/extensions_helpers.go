//go:build compareopus && cgo

package main

/*
#define OPUS_BUILD 1
#define VAR_ARRAYS 1
#define opus_extension_iterator_init comparison_extensions_init
#define opus_extension_iterator_reset comparison_extensions_reset
#define opus_extension_iterator_set_frame_max comparison_extensions_frame_max
#define opus_extension_iterator_next comparison_extensions_next
#define opus_extension_iterator_find comparison_extensions_find
#define opus_packet_extensions_count comparison_extensions_count
#define opus_packet_extensions_count_ext comparison_extensions_count_ext
#define opus_packet_extensions_parse comparison_extensions_parse
#define opus_packet_extensions_parse_ext comparison_extensions_parse_ext
#define opus_packet_extensions_generate comparison_extensions_generate
#include "../../../opus/src/extensions.c"
static int native_write_extension(unsigned char *data,int capacity,int pos,int id,int length,const unsigned char *payload,int last) {
 opus_extension_data ext={0};ext.id=id;ext.len=length;ext.data=payload;return write_extension(data,capacity,pos,&ext,last);
}
static int native_write_payload(unsigned char *data,int capacity,int pos,int id,int length,const unsigned char *payload,int last) {
 opus_extension_data ext={0};ext.id=id;ext.len=length;ext.data=payload;return write_extension_payload(data,capacity,pos,&ext,last);
}
static int native_skip_payload(const unsigned char *base,int len,int id,int trailing,int *header,int *offset,int op) {
 const unsigned char *data=base;
 int result=op?skip_extension(&data,len,header):skip_extension_payload(&data,len,header,id,trailing);
 *offset=(int)(data-base);return result;
}
*/
import "C"
import "unsafe"

func nativeWriteExtension(data []byte, capacity, pos, id, length int32, payload []byte, last int32) int32 {
	return int32(C.native_write_extension((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(capacity), C.int(pos), C.int(id), C.int(length), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(payload))), C.int(last)))
}

func nativeWritePayload(data []byte, capacity, pos, id, length int32, payload []byte, last int32) int32 {
	return int32(C.native_write_payload((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(capacity), C.int(pos), C.int(id), C.int(length), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(payload))), C.int(last)))
}

func nativeSkipPayload(data []byte, length, id, trailing, header int32, op int32) (int32, int32, int32) {
	h := C.int(header)
	offset := C.int(0)
	result := C.native_skip_payload((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(length), C.int(id), C.int(trailing), &h, &offset, C.int(op))
	return int32(result), int32(offset), int32(h)
}
