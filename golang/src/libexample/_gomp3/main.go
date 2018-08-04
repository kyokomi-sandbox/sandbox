package main

import (
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/oto"
	"github.com/hajimehoshi/go-mp3"
)

func run() error {
	f, err := os.Open("The_long_journey_Free_Ver.mp3")
	if err != nil {
		return err
	}
	defer f.Close()
	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}
	defer d.Close()
	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 65536)
	if err != nil {
		return err
	}
	defer p.Close()
	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
