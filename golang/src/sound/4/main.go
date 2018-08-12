package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "wave.h"
#include "fft.h"

int hoge()
{
  MONO_PCM pcm;
  int n, k, N;
  double *x_real, *x_imag;

  wave_read_16bit_mono(&pcm, "sine_500hz.wav");

  N = 64;

  x_real = calloc(N, sizeof(double));
  x_imag = calloc(N, sizeof(double));

  for (n = 0; n < N; n++)
  {
    x_real[n] = pcm.s[n];
    x_imag[n] = 0.0;
  }

  FFT(x_real, x_imag, N);

  for (k = 0; k < N; k++)
  {
    printf("X(%d) = %f+j%f\n", k, x_real[k], x_imag[k]);
  }

  free(pcm.s);
  free(x_real);
  free(x_imag);

  return 0;
}
*/
import "C"
import (
	"fmt"
	"log"
	"os"

	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/wav"
)

func mono16SinToInt16(s float64) int16 {
	s = (s + 1.0) / 2.0 * 65536.0

	if s > 65535.0 {
		s = 65535.0
	} else if s < 0.0 {
		s = 0.0
	}

	return int16(s + 0.5 - 32768)
}

func monoInt16ToFloat64(s int16) float64 {
	a := float64(s) + 32768
	a = a - 0.5
	a = a - 1.0*2.0/65536.0
	return a
}

func main() {
	file, err := os.Open("./sine_500hz.wav")
	if err != nil {
		log.Fatal(err)
	}
	w, err := wav.New(file)
	if err != nil {
		log.Fatal(err)
	}

	sample, err := w.ReadSamples(w.Samples)
	if err != nil {
		log.Fatal(err)
	}

	samples := sample.([]int16)

	N := 64
	xReal := make([]float64, N)
	for i := 0; i < N; i++ {
		xReal[i] = monoInt16ToFloat64(samples[i])
	}

	C.hoge()
	fftReals := fft.FFTReal(xReal)
	ifftReals := fft.IFFTReal(xReal)

	for k := 0; k < N; k++ {
		fmt.Printf("X(%d) = %f+j%f\n", k, fftReals[k], ifftReals[k])
	}
}
