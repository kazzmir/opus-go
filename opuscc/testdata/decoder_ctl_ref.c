/* gcc -std=c99 -O2 -I ../opus -I ../opus/include -I ../opus/celt
   -I ../opus/silk -o /tmp/decoder_ctl_ref opuscc/testdata/decoder_ctl_ref.c
   ../opus/.libs/libopus.a -lm
   /tmp/decoder_ctl_ref > opuscc/testdata/decoder_ctl_ref.txt */
#include "config.h"
#include "src/opus_decoder.c"
#include <stdio.h>
#include <stdlib.h>
static void get(OpusDecoder *s, int req) {
 opus_int32 v=-999; int r=opus_decoder_ctl(s,req,&v);
 printf("get %d 0 %d %d\n",req,r,v);
 r=opus_decoder_ctl(s,req,(opus_int32*)NULL);
 printf("null %d 0 %d 0\n",req,r);
}
static void set(OpusDecoder *s,int req,int v) {
 int r=opus_decoder_ctl(s,req,v);
 printf("set %d %d %d 0\n",req,v,r);
}
int main(void) {
 int err; OpusDecoder *s=opus_decoder_create(48000,2,&err);
 int gets[]={OPUS_GET_BANDWIDTH_REQUEST,OPUS_GET_COMPLEXITY_REQUEST,
 OPUS_GET_FINAL_RANGE_REQUEST,OPUS_GET_SAMPLE_RATE_REQUEST,OPUS_GET_PITCH_REQUEST,
 OPUS_GET_GAIN_REQUEST,OPUS_GET_LAST_PACKET_DURATION_REQUEST,
 OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST,OPUS_GET_IGNORE_EXTENSIONS_REQUEST};
 for(int i=0;i<9;i++) get(s,gets[i]);
 int sets[]={OPUS_SET_COMPLEXITY_REQUEST,OPUS_SET_GAIN_REQUEST,
 OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST,OPUS_SET_IGNORE_EXTENSIONS_REQUEST};
 int queries[]={OPUS_GET_COMPLEXITY_REQUEST,OPUS_GET_GAIN_REQUEST,
 OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST,OPUS_GET_IGNORE_EXTENSIONS_REQUEST};
 int values[][5]={{10,-1,11,0,7},{-32768,-32769,32768,32767,-1234},{1,-1,2,0,1},{1,-1,2,0,1}};
 for(int i=0;i<4;i++) for(int j=0;j<5;j++) {set(s,sets[i],values[i][j]);get(s,queries[i]);}
 s->DecControl.prevPitchLag=77; s->prev_mode=MODE_SILK_ONLY;
 s->bandwidth=OPUS_BANDWIDTH_WIDEBAND; s->last_packet_duration=960; s->rangeFinal=123456789;
 printf("seed 0 0 0 0\n");
 get(s,OPUS_GET_PITCH_REQUEST);get(s,OPUS_GET_BANDWIDTH_REQUEST);
 get(s,OPUS_GET_LAST_PACKET_DURATION_REQUEST);get(s,OPUS_GET_FINAL_RANGE_REQUEST);
 s->prev_mode=MODE_CELT_ONLY; printf("celt 0 0 0 0\n");get(s,OPUS_GET_PITCH_REQUEST);
 int r=opus_decoder_ctl(s,OPUS_RESET_STATE);printf("reset %d 0 %d 0\n",OPUS_RESET_STATE,r);
 for(int i=0;i<9;i++) get(s,gets[i]);
 set(s,999999,0);
 opus_decoder_destroy(s);
}
