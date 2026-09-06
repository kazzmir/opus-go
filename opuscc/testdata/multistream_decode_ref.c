/* Build with the same include/link flags as decoder_ctl_ref.c.
   Redirect stdout to multistream_decode_ref.txt.
   Default: silent coupled stream, nonzero mono stream (both SILK NB).
   Pass any argument to reproduce the separate, pre-existing noisy coupled
   stream mismatch in the first two decoded packets. */
#include "config.h"
#include "src/opus_multistream_decoder.c"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
static void run(OpusMSDecoder *d,const char *name,const unsigned char *p,int len,int frame,int fec,int short_out) {
 if(frame<=640) frame*=3;
 float pcm[6400*5]; memset(pcm,0xa5,sizeof(pcm));
 int ret=short_out?opus_multistream_decode(d,p,len,(opus_int16*)pcm,frame,fec):opus_multistream_decode_float(d,p,len,pcm,frame,fec);
 uint64_t hash=UINT64_C(14695981039346656037);
 if(ret>0) for(int i=0;i<ret*5;i++) {
  uint32_t bits;int bytes=short_out?2:4;
  if(short_out) bits=(uint16_t)((opus_int16*)pcm)[i]; else memcpy(&bits,pcm+i,4);
  for(int j=0;j<bytes;j++) {hash^=(bits>>(8*j))&255;hash*=UINT64_C(1099511628211);}
 }
 opus_uint32 range=0;opus_multistream_decoder_ctl(d,OPUS_GET_FINAL_RANGE(&range));
 printf("%s %d %d %d %d %d %llu %u ",name,len,frame,fec,short_out,ret,(unsigned long long)hash,range);
 if(len>0) for(int i=0;i<len;i++) printf("%02x",p[i]);else printf("-");puts("");
}
int main(int argc, char **argv) {
 (void)argv;
 unsigned char map[]={2,0,1,2,255};int err;
 OpusMSDecoder *d=opus_multistream_decoder_create(48000,5,2,1,map,&err);
 OpusEncoder *enc[2];
 for(int s=0;s<2;s++) {
  enc[s]=opus_encoder_create(16000,s?1:2,OPUS_APPLICATION_VOIP,&err);
  opus_encoder_ctl(enc[s],OPUS_SET_BITRATE(s?12000:24000));
  opus_encoder_ctl(enc[s],OPUS_SET_FORCE_CHANNELS(1));
  opus_encoder_ctl(enc[s],OPUS_SET_MAX_BANDWIDTH(OPUS_BANDWIDTH_NARROWBAND));
 }
 uint32_t seed=1;unsigned char combined[4000];int total=0;
 for(int f=0;f<3;f++) {
  total=0;
  for(int s=0;s<2;s++) {
   opus_int16 pcm[640];int channels=s?1:2;
   for(int i=0;i<320*channels;i++) {seed=seed*1664525u+1013904223u;pcm[i]=(s || argc>1)?((int)(seed>>18)-8192):0;}
   unsigned char packet[2000];int n=opus_encode(enc[s],pcm,320,packet,sizeof(packet));
   if(n<0) return 1;
   if(!s) {
    OpusRepacketizer *rp=opus_repacketizer_create();opus_repacketizer_cat(rp,packet,n);
    n=opus_repacketizer_out_range_impl(rp,0,1,combined,sizeof(combined),1,0,NULL,0);opus_repacketizer_destroy(rp);
   } else memcpy(combined+total,packet,n);
   total+=n;
  }
  run(d,"packet",combined,total,320,0,0);
 }
 run(d,"plc",NULL,0,320,0,0);
 run(d,"fec",combined,total,320,1,0);
 run(d,"short",combined,total,320,0,1);
 run(d,"too_small",combined,total,160,0,0);
 run(d,"bad_frame",combined,total,0,0,0);
 run(d,"negative_len",combined,-1,320,0,0);
 run(d,"truncated",combined,2,320,0,0);
 unsigned char mismatch[]={0,0,8};run(d,"mismatch",mismatch,3,640,0,0);
 run(d,"bad_fec",combined,total,320,2,0);
 run(d,"clamped_plc",NULL,0,6400,0,0);
 opus_multistream_decoder_destroy(d);for(int i=0;i<2;i++) opus_encoder_destroy(enc[i]);
}
