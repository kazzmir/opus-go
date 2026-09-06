/* Build with the include/link flags in decoder_ctl_ref.c and redirect stdout
   to multistream_ctl_ref.txt. Includes decoder internals to seed range state. */
#include "config.h"
#include "src/opus_decoder.c"
#include "opus_multistream.h"
#include <stdio.h>
static OpusDecoder *dec[3];
static void get(OpusMSDecoder *s,int stream,int req) {
 opus_int32 v=-999;int r=stream<0?opus_multistream_decoder_ctl(s,req,&v):opus_decoder_ctl(dec[stream],req,&v);
 printf("get %d %d 0 %d %d\n",stream,req,r,v);
}
static void set(OpusMSDecoder *s,int req,int v) {
 int r=opus_multistream_decoder_ctl(s,req,v);printf("set -1 %d %d %d 0\n",req,v,r);
}
int main(void) {
 unsigned char map[]={0,1,2,3};int err;
 OpusMSDecoder *s=opus_multistream_decoder_create(48000,4,3,1,map,&err);
 for(int i=-1;i<=3;i++) {
  OpusDecoder *p=NULL;int r=opus_multistream_decoder_ctl(s,OPUS_MULTISTREAM_GET_DECODER_STATE(i,&p));
  printf("state -1 %d %d %d %ld\n",OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST,i,r,p?(long)((char*)p-(char*)s):0);
  if(!r) dec[i]=p;
 }
 int r=opus_multistream_decoder_ctl(s,OPUS_MULTISTREAM_GET_DECODER_STATE(0,(OpusDecoder**)NULL));
 printf("state_null -1 %d 0 %d 0\n",OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST,r);
 int gets[]={OPUS_GET_BANDWIDTH_REQUEST,OPUS_GET_SAMPLE_RATE_REQUEST,OPUS_GET_GAIN_REQUEST,OPUS_GET_LAST_PACKET_DURATION_REQUEST,OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST,OPUS_GET_COMPLEXITY_REQUEST,OPUS_GET_FINAL_RANGE_REQUEST};
 for(int i=0;i<7;i++) {
  get(s,-1,gets[i]);r=opus_multistream_decoder_ctl(s,gets[i],(opus_int32*)NULL);
  printf("null -1 %d 0 %d 0\n",gets[i],r);
 }
 for(int i=0;i<3;i++) {
  dec[i]->rangeFinal=0x87654321u+i*0x12345678u;
  dec[i]->bandwidth=1101+i;dec[i]->last_packet_duration=120*(i+1);
  opus_decoder_ctl(dec[i],OPUS_SET_GAIN(100+i));
 }
 puts("seed -1 0 0 0 0");
 for(int i=0;i<7;i++) get(s,-1,gets[i]);
 int sets[]={OPUS_SET_GAIN_REQUEST,OPUS_SET_COMPLEXITY_REQUEST,OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST};
 int queries[]={OPUS_GET_GAIN_REQUEST,OPUS_GET_COMPLEXITY_REQUEST,OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST};
 int vals[][4]={{-32768,32768,-32769,1234},{10,11,-1,3},{1,2,-1,0}};
 for(int i=0;i<3;i++) for(int j=0;j<4;j++) {
  set(s,sets[i],vals[i][j]);get(s,-1,queries[i]);
  for(int k=0;k<3;k++) get(s,k,queries[i]);
 }
 r=opus_multistream_decoder_ctl(s,OPUS_RESET_STATE);printf("reset -1 %d 0 %d 0\n",OPUS_RESET_STATE,r);
 for(int k=-1;k<3;k++) for(int i=0;i<7;i++) get(s,k,gets[i]);
 set(s,999999,0);set(s,OPUS_SET_IGNORE_EXTENSIONS_REQUEST,1);
 opus_multistream_decoder_destroy(s);
}
