/* Build with the include/link flags in decoder_ctl_ref.c; redirect stdout
   to projection_init_ref.txt. */
#include "config.h"
#include "src/opus_projection_decoder.c"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
int main(void) {
 int cases[][5]={{1,1,0,48000,0},{2,1,1,24000,0},{4,3,1,16000,0},{2,3,0,12000,0},{8,4,4,8000,0},{4,2,0,48000,0},{2,1,1,44100,0},{2,1,1,48000,-1},{2,1,1,48000,1},{0,1,0,48000,0}};
 int16_t vals[]={0,32767,-32768,-1,12345,-23456,1,-2};
 for(int n=0;n<10;n++) {
  int ch=cases[n][0],streams=cases[n][1],coupled=cases[n][2],fs=cases[n][3],delta=cases[n][4];
  int count=ch*(streams+coupled);unsigned char matrix[128];
  for(int i=0;i<count;i++) {uint16_t v=vals[i%8];matrix[2*i]=v&255;matrix[2*i+1]=v>>8;}
  int size=opus_projection_decoder_get_size(ch,streams,coupled);
  OpusProjectionDecoder *d=calloc(1,size?size:4096);
  int ret=opus_projection_decoder_init(d,fs,ch,streams,coupled,matrix,count*2+delta);
  printf("%d %d %d %d %d %d %d",ch,streams,coupled,fs,delta,ret,size);
  if(!ret) {
   MappingMatrix *m=get_dec_demixing_matrix(d);int16_t *data=mapping_matrix_get_data(m);
   uint64_t hash=UINT64_C(14695981039346656037);
   for(int i=0;i<count;i++) for(int j=0;j<2;j++) {hash^=((uint16_t)data[i]>>(8*j))&255;hash*=UINT64_C(1099511628211);}
   OpusMSDecoder *ms=get_multistream_decoder(d);
   printf(" %d %d %d %d %llu",d->demixing_matrix_size_in_bytes,m->rows,m->cols,m->gain,(unsigned long long)hash);
   for(int i=0;i<ch;i++) printf(" %d",ms->layout.mapping[i]);
  }
  puts("");free(d);
 }
}
