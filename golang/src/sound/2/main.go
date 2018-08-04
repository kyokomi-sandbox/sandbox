package main

import (
	"log"
	"os"

	"math"

	"encoding/binary"

	"fmt"

	"github.com/cryptix/wav"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "wave.h"

int ex2()
{
 MONO_PCM pcm;
 int n;
 double a, f0;

 pcm.fs = 44100;
 pcm.bits = 16;
 pcm.length = pcm.fs * 1;
 pcm.s = calloc(pcm.length, sizeof(double));

 a = 0.1;
 f0 = 500.0;


 for (n = 0; n < pcm.length; n++)
 {
   pcm.s[n] = a * sin(2.0 * M_PI * f0 * n / pcm.fs);
 }

 wave_write_16bit_mono(&pcm, "local_ex2_1.wav");

 free(pcm.s);

 return 0;
}
*/
import "C"

func mono16SinToByte(s float64) []byte {
	s = (s + 1.0) / 2.0 * 65536.0

	if s > 65535.0 {
		s = 65535.0
	} else if s < 0.0 {
		s = 0.0
	}

	data := int16(s + 0.5 - 32768)
	return Int16bytes(data)
}

func Int16bytes(data int16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(data))
	return buf
}

func main2() string {
	log.SetFlags(log.Llongfile)
	var wf = wav.File{
		SampleRate:      44100,
		Channels:        1,
		SignificantBits: 16,
		SoundSize:       44100 * 1,
	}

	f, err := os.Create(fmt.Sprintf("./%s.wav", "main2"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wr, err := wf.NewWriter(f)
	if err != nil {
		log.Fatal(err)
	}
	defer wr.Close()

	a := 0.1    // 振幅
	f0 := 500.0 // 周波数

	// サイン波
	for n := uint32(0); n < wf.SoundSize; n++ {
		// 丸め誤差を考慮してfloat64で音データを扱っている
		s := a * math.Sin(2.0*math.Pi*f0*float64(n)/float64(wf.SampleRate))
		// 16bitのモノラルの音を2byteのデータに変換している（32767 〜 -32768の間の値に変換）
		wr.Write(mono16SinToByte(s))
	}

	return f.Name()
}

func main() {
	//C.ex2()
	main2()
}
