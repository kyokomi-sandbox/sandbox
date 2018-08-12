package main

import (
	"encoding/binary"
	"log"
	"math"
	"os"

	"github.com/r9y9/go-dsp/wav"
)

func ex5_4() {
	outFile := &wav.File{
		SampleRate:      44100,
		SignificantBits: 16,
		Channels:        1,
	}

	fs := float64(outFile.SampleRate)
	bits := uint16(outFile.SignificantBits)
	length := fs * 4
	data := make([]float64, uint64(length))

	for n := float64(0); n < length; n++ {
		i := uint64(n)
		data[i] += 0.5 * math.Sin(2.0*math.Pi*440.0*n/fs)
		data[i] += 1.0 * math.Sin(2.0*math.Pi*880*n/fs)
		data[i] += 0.7 * math.Sin(2.0*math.Pi*1320*n/fs)
		data[i] += 0.5 * math.Sin(2.0*math.Pi*1760*n/fs)
		data[i] += 0.3 * math.Sin(2.0*math.Pi*2200*n/fs)
	}

	gain := 0.1

	for n := float64(0); n < length; n++ {
		i := uint64(n)
		data[i] *= gain
	}

	for n := float64(0); n < (fs * 0.01); n++ {
		i := uint64(n)
		data[i] *= n / (fs * 0.01)
		data[uint64(length)-i-1] *= n / (fs * 0.01)
	}

	if err := WriteMono("ex5_4.wav", data, bits, 44100); err != nil {
		log.Fatal(err)
	}
}

func ex5_5() {
	fs := float64(44100)
	bits := uint16(16)
	length := fs * 4
	data := make([]float64, uint64(length))
	a0 := make([]float64, uint64(length))
	a1 := make([]float64, uint64(length))
	a2 := make([]float64, uint64(length))
	a3 := make([]float64, uint64(length))
	a4 := make([]float64, uint64(length))

	for n := float64(0); n < length; n++ {
		i := uint64(n)
		a0[i] = 0.5 * math.Exp(-5.0*n/(fs*4.0))
		a1[i] = 1.0 * math.Exp(-5.0*n/(fs*2.0))
		a2[i] = 0.7 * math.Exp(-5.0*n/(fs*1.0))
		a3[i] = 0.5 * math.Exp(-5.0*n/(fs*0.5))
		a4[i] = 0.3 * math.Exp(-5.0*n/(fs*0.2))
	}

	for n := float64(0); n < length; n++ {
		i := uint64(n)
		data[i] += a0[i] * math.Sin(2.0*math.Pi*440.0*n/fs)
		data[i] += a1[i] * math.Sin(2.0*math.Pi*880*n/fs)
		data[i] += a2[i] * math.Sin(2.0*math.Pi*1320*n/fs)
		data[i] += a3[i] * math.Sin(2.0*math.Pi*1760*n/fs)
		data[i] += a4[i] * math.Sin(2.0*math.Pi*2200*n/fs)
	}

	gain := 0.1

	for n := float64(0); n < length; n++ {
		i := uint64(n)
		data[i] *= gain
	}

	for n := float64(0); n < (fs * 0.01); n++ {
		i := uint64(n)
		data[i] *= n / (fs * 0.01)
		data[uint64(length)-i-1] *= n / (fs * 0.01)
	}

	if err := WriteMono("ex5_5.wav", data, bits, uint32(fs)); err != nil {
		log.Fatal(err)
	}
}

func mono16SinToByte(s float64) uint16 {
	s = (s + 1.0) / 2.0 * 65536.0

	if s > 65535.0 {
		s = 65535.0
	} else if s < 0.0 {
		s = 0.0
	}

	return uint16(s + 0.5 - 32768)
}

func WriteMono(filename string, data []float64, bits uint16, sampleRate uint32) error {
	channels := 1 // mono

	outFile := &wav.File{
		SampleRate:      sampleRate,
		SignificantBits: bits,
		Channels:        uint16(channels),
	}

	// []int to []bytes (assuming 16-bit samples)
	bytes := make([]byte, 2*len(data))
	for i, val := range data {
		start := i * 2
		binary.LittleEndian.PutUint16(bytes[start:start+2], mono16SinToByte(val))
	}

	ofile, oerr := os.Create(filename)
	if oerr != nil {
		return oerr
	}

	err := outFile.WriteData(ofile, bytes)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ex5_4()
	ex5_5()
}
