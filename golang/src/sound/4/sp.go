package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/r9y9/gossp"
	"github.com/r9y9/gossp/io"
	"github.com/r9y9/gossp/stft"
	"github.com/r9y9/gossp/window"
)

func main() {
	filename := flag.String("i", "./sine_500hz.wav", "Input filename")
	flag.Parse()

	w, werr := io.ReadWav(*filename)
	if werr != nil {
		log.Fatal(werr)
	}
	data := w.GetMonoData()

	s := &stft.STFT{
		FrameShift: int(float64(w.SampleRate) / 100.0), // 0.01 sec,
		FrameLen:   2048,
		Window:     window.CreateHanning(2048),
	}

	spectrogram := s.STFT(data)
	amplitudeSpectrogram, _ := gossp.SplitSpectrogram(spectrogram)
	PrintMatrixAsGnuplotFormat(amplitudeSpectrogram)

	//if err := wav.WriteMono("copy.wav", s.ISTFT(spectrogram), w.SampleRate); err != nil {
	//	log.Fatal(err)
	//}
}

func PrintMatrixAsGnuplotFormat(matrix [][]float64) {
	fmt.Println("#", len(matrix[0]), len(matrix))
	for i, vec := range matrix {
		for j, val := range vec {
			fmt.Println(i, j, math.Log(val))
		}
		fmt.Println("")
	}
}
