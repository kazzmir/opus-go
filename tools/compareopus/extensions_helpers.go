//go:build compareopus && cgo

package main

/*
#include <setjmp.h>
#include <stdlib.h>
static _Thread_local jmp_buf extension_jump;
void comparison_extension_fatal(const char *str,const char *file,int line){longjmp(extension_jump,1);}
#define celt_fatal comparison_extension_fatal
#define ENABLE_ASSERTIONS 1
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
static const unsigned char *extension_pointer(const unsigned char *base,int offset){return offset<0?NULL:base+offset;}
static int extension_offset(const unsigned char *base,const unsigned char *pointer){return pointer==NULL?-1:(int)(pointer-base);}
static int native_extension_iterator(const unsigned char *base,int length,int frames,int *v,int op,int want,int id) {
 // Heap storage keeps partial writes defined after a captured assertion/longjmp.
 OpusExtensionIterator *st=calloc(1,sizeof(*st));opus_extension_data *ext=calloc(1,sizeof(*ext));
 st->data=extension_pointer(base,v[0]);st->curr_data=extension_pointer(base,v[1]);st->repeat_data=extension_pointer(base,v[2]);st->last_long=extension_pointer(base,v[3]);st->src_data=extension_pointer(base,v[4]);
 st->len=v[5];st->curr_len=v[6];st->repeat_len=v[7];st->src_len=v[8];st->trailing_short_len=v[9];st->nb_frames=v[10];st->frame_max=v[11];st->curr_frame=v[12];st->repeat_frame=v[13];st->repeat_l=v[14];
 if(op){ext->id=v[15];ext->frame=v[16];ext->data=extension_pointer(base,v[17]);ext->len=v[18];}
 int result=0;if(setjmp(extension_jump))result=-99;else if(op==0)opus_extension_iterator_init(st,base,length,frames);else if(op==1)result=opus_extension_iterator_next_repeat(st,want?ext:NULL);else if(op==2)result=opus_extension_iterator_next(st,want?ext:NULL);else result=opus_extension_iterator_find(st,want?ext:NULL,id);
 v[0]=extension_offset(base,st->data);v[1]=extension_offset(base,st->curr_data);v[2]=extension_offset(base,st->repeat_data);v[3]=extension_offset(base,st->last_long);v[4]=extension_offset(base,st->src_data);
 v[5]=st->len;v[6]=st->curr_len;v[7]=st->repeat_len;v[8]=st->src_len;v[9]=st->trailing_short_len;v[10]=st->nb_frames;v[11]=st->frame_max;v[12]=st->curr_frame;v[13]=st->repeat_frame;v[14]=st->repeat_l;
 if(op){v[15]=ext->id;v[16]=ext->frame;v[17]=extension_offset(base,ext->data);v[18]=ext->len;}free(ext);free(st);return result;
}
static int native_extension_count(const unsigned char *data,int length,int frames) {
 if(setjmp(extension_jump))return -99;return opus_packet_extensions_count(data,length,frames);
}
static int native_extension_count_ext(const unsigned char *data,int length,int *counts,int frames) {if(setjmp(extension_jump))return -99;return opus_packet_extensions_count_ext(data,length,counts,frames);}
static int native_extension_parse(const unsigned char *base,int length,int frames,int *words,int slots,int *nb,int null_output,int null_count,int capacity_alias,int *counts,int ordered) {
 opus_extension_data *out=calloc(slots,sizeof(*out));for(int i=0;i<slots;i++){out[i].id=words[4*i];out[i].frame=words[4*i+1];out[i].data=extension_pointer(base,words[4*i+2]);out[i].len=words[4*i+3];}
 int *capacity=null_count?NULL:capacity_alias==1?&out[0].frame:capacity_alias==2?counts:nb;
 int result;if(setjmp(extension_jump))result=-99;else if(ordered)result=opus_packet_extensions_parse_ext(base,length,null_output?NULL:out,capacity,counts,frames);else result=opus_packet_extensions_parse(base,length,null_output?NULL:out,capacity,frames);
 if(capacity&&nb)*nb=*capacity;
 for(int i=0;i<slots;i++){words[4*i]=out[i].id;words[4*i+1]=out[i].frame;words[4*i+2]=extension_offset(base,out[i].data);words[4*i+3]=out[i].len;}free(out);return result;
}
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

func nativeExtensionIterator(data []byte, length, frames int32, v *[19]int32, options ...int32) int32 {
	op, want, id := int32(0), int32(1), int32(0)
	if len(options) > 0 {
		op = options[0]
	}
	if len(options) > 1 {
		want = options[1]
	}
	if len(options) > 2 {
		id = options[2]
	}
	return int32(C.native_extension_iterator((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(length), C.int(frames), (*C.int)(unsafe.Pointer(&v[0])), C.int(op), C.int(want), C.int(id)))
}

func nativeExtensionCount(data []byte, length, frames int32) int32 {
	return int32(C.native_extension_count((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(length), C.int(frames)))
}

func nativeExtensionCountExt(data *byte, length int32, counts *int32, frames int32) int32 {
	return int32(C.native_extension_count_ext((*C.uchar)(unsafe.Pointer(data)), C.int(length), (*C.int)(unsafe.Pointer(counts)), C.int(frames)))
}

func nativeExtensionParse(data []byte, length, frames int32, words []int32, capacity *int32, nullOutput, nullCount bool, alias int32, frameCounts ...[]int32) int32 {
	no, nc := C.int(0), C.int(0)
	if nullOutput {
		no = 1
	}
	if nullCount {
		nc = 1
	}
	var counts *C.int
	ordered := C.int(0)
	if len(frameCounts) > 0 {
		ordered = 1
		counts = (*C.int)(unsafe.Pointer(unsafe.SliceData(frameCounts[0])))
	}
	return int32(C.native_extension_parse((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(length), C.int(frames), (*C.int)(unsafe.Pointer(unsafe.SliceData(words))), C.int(len(words)/4), (*C.int)(unsafe.Pointer(capacity)), no, nc, C.int(alias), counts, ordered))
}

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
