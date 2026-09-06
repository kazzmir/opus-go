/* Build with the include/link flags in decoder_ctl_ref.c and redirect stdout
   to extensions_generate_ref.txt. No floating-point/SIMD code is exercised. */
#include "config.h"
#include "src/opus_private.h"
#include <stdio.h>
#include <stdint.h>
#include <string.h>
typedef struct {int id,frame,len,seed;} Spec;
static void run(const char *name,int frames,const Spec *spec,int count) {
 opus_extension_data ext[16];unsigned char payload[16][600],output[1116];
 for(int i=0;i<count;i++) {
  for(int j=0;j<600;j++) payload[i][j]=(unsigned char)(spec[i].seed+j*17);
  ext[i].id=spec[i].id;ext[i].frame=spec[i].frame;ext[i].len=spec[i].len;ext[i].data=payload[i];
 }
 int caps[]={0,1,2,3,8,64,1100};
 for(int p=0;p<2;p++) for(int k=0;k<8;k++) {
  int cap=k<7?caps[k]:opus_packet_extensions_generate(NULL,1100,ext,count,frames,0);
  if(cap<0) continue;
  memset(output,0xa5,sizeof(output));
  int ret=opus_packet_extensions_generate(output,cap,ext,count,frames,p);
  int dry=opus_packet_extensions_generate(NULL,cap,ext,count,frames,p);
  uint64_t hash=UINT64_C(14695981039346656037);
  for(int i=0;i<cap+16;i++) {hash^=output[i];hash*=UINT64_C(1099511628211);}
  printf("%s %d %d %d %d %d %llu %d",name,frames,p,cap,ret,dry,(unsigned long long)hash,count);
  for(int i=0;i<count;i++) printf(" %d %d %d %d",spec[i].id,spec[i].frame,spec[i].len,spec[i].seed);
  puts("");
 }
}
#define RUN(name,frames,...) do {Spec s[]={__VA_ARGS__};run(name,frames,s,sizeof(s)/sizeof(*s));} while(0)
int main(void) {
 run("empty",1,NULL,0);run("zero_frames",0,NULL,0);run("too_many_frames",49,NULL,0);
 RUN("short",1,{3,0,0,0},{31,0,1,99});
 RUN("long_lengths",1,{32,0,0,0},{127,0,255,10},{33,0,256,20},{34,0,2,30});
 RUN("separators",48,{3,0,1,1},{4,1,1,2},{5,4,0,3},{32,47,2,4});
 RUN("unordered",4,{33,3,4,1},{3,0,1,2},{34,2,3,3},{32,0,2,4},{4,1,0,5});
 RUN("repeat_short",3,{3,0,1,1},{4,0,0,0},{3,1,1,2},{4,1,0,0},{3,2,1,3},{4,2,0,0});
 RUN("repeat_long",3,{32,0,2,1},{32,1,0,2},{32,2,260,3});
 RUN("repeat_remaining",3,{32,0,2,1},{3,0,1,8},{32,1,3,2},{4,1,0,0},{32,2,260,3});
 RUN("repeat_unordered",3,{32,2,2,3},{3,0,1,1},{3,1,1,2},{3,2,1,3},{32,0,4,1},{32,1,3,2});
 RUN("repeat_interleaved",3,{3,0,1,1},{3,1,1,2},{3,2,1,3},{32,1,3,2},{32,0,4,1},{32,2,2,3});
 RUN("different_short_lengths",2,{3,0,0,0},{3,1,1,1});
 RUN("different_ids",2,{3,0,1,1},{4,1,1,2});
 RUN("bad_frame",2,{3,-1,0,0});RUN("frame_past_end",2,{3,2,0,0});
 RUN("bad_id_low",1,{2,0,0,0});RUN("bad_id_high",1,{128,0,0,0});
 RUN("bad_short_length",1,{3,0,2,0});RUN("negative_short_length",1,{3,0,-1,0});
 RUN("negative_long_length",1,{32,0,-1,0});
}
