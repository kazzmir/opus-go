/* Build like decoder_ctl_ref.c, then redirect stdout to multistream_validate_ref.txt. */
#include "config.h"
#include "src/opus_multistream_decoder.c"
#include <stdio.h>
#include <string.h>
int main(void) {
 const char *hex[]={"", "00", "000000", "000008", "010001", "0303000303", "0000000000", "000100", "03", "0300", "0341", "ff", "0000", "0303000302", "0000ff", "0200"};
 int rates[]={48000,8000};
 for(int i=0;i<16;i++) for(int n=1;n<=3;n++) for(int f=0;f<2;f++) {
  unsigned char data[64]={0}; int len=strlen(hex[i])/2;
  for(int j=0;j<len;j++) {unsigned x;sscanf(hex[i]+2*j,"%2x",&x);data[j]=x;}
  printf("%s %d %d %d\n",len?hex[i]:"-",rates[f],n,opus_multistream_packet_validate(data,len,n,rates[f]));
 }
}
