package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/oto"
	"github.com/k0kubun/pp"
	"github.com/mjibson/go-dsp/wav"
)

func main() {
	// wavファイルを読み込み
	file, err := os.Open("./ji00101.wav") // http://otosozai.com/?sozai=ji00101
	if err != nil {
		log.Fatal(err)
	}
	w, err := wav.New(file)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(w)

	// playerを生成（bitsPerSample/8でbytesPerSampleにしている）
	p, err := oto.NewPlayer(int(w.SampleRate), int(w.NumChannels), int(w.BitsPerSample/8), 65536)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	// TODO: 8bitモノラルにしか対応してない?
	// wavの音dataをplayerにcopyして再生
	sample, err := w.ReadSamples(w.Samples)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(p, bytes.NewReader(sample.([]byte))); err != nil {
		log.Fatal(err)
	}
}
