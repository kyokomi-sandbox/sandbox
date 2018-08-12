#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "wave.h"
#include "fft.h"

int main(void)
{
  MONO_PCM pcm;
  int n, k, N;
  double *x_real, *x_imag;
  
  wave_read_16bit_mono(&pcm, "sine_500hz.wav");
  
  N = 64; /* DFTÇÃÉTÉCÉY */
  
  x_real = calloc(N, sizeof(double));
  x_imag = calloc(N, sizeof(double));
  
  /* îgå` */
  for (n = 0; n < N; n++)
  {
    x_real[n] = pcm.s[n]; /* x(n)ÇÃé¿êîïî */
    x_imag[n] = 0.0; /* x(n)ÇÃãïêîïî */
  }
  
  FFT(x_real, x_imag, N); /* FFTÇÃåvéZåãâ ÇÕx_realÇ∆x_imagÇ…è„èëÇ´Ç≥ÇÍÇÈ */
  
  /* é¸îgêîì¡ê´ */
  for (k = 0; k < N; k++)
  {
    printf("X(%d) = %f+j%f\n", k, x_real[k], x_imag[k]);
  }
  
  free(pcm.s);
  free(x_real);
  free(x_imag);
  
  return 0;
}
