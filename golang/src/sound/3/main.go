package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"

	"math/rand"
	"time"

	"github.com/cryptix/wav"
)

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

func waveWrite16bitMono(fileName string, gain float64, sinFunc func(wf wav.File) []float64) string {
	var wf = wav.File{
		SampleRate:      44100,
		Channels:        1,
		SignificantBits: 16,
		SoundSize:       44100 * 1,
	}

	f, err := os.Create(fmt.Sprintf("./%s.wav", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wr, err := wf.NewWriter(f)
	if err != nil {
		log.Fatal(err)
	}
	defer wr.Close()

	s := sinFunc(wf)
	/*
		wavファイルの音量を調整する時は、一般的に波形のゲインと呼ばれる値を調整します。
		ゲインを上げれば、波形の縦の振幅は大きくなり、下げれば小さくなります。
	*/
	for i := range s {
		s := mono16SinToByte(s[i] * gain)
		wr.Write(s)
	}

	return f.Name()
}

// ノコギリ波
func main3to1() string {
	f0 := 500.0 // 周波数
	return waveWrite16bitMono("main3to1", 0.1, func(wf wav.File) []float64 {
		s := make([]float64, wf.SoundSize)
		for i := float64(1); i <= 44; i++ {
			for n := uint32(0); n < wf.SoundSize; n++ {
				// 丸め誤差を考慮してfloat64で音データを扱っている
				s[n] += 1.0 / i * math.Sin(2.0*math.Pi*i*f0*float64(n)/float64(wf.SampleRate))
			}
		}
		return s
	})
}

// 矩形波
func main3to2() string {
	f0 := 500.0 // 周波数
	return waveWrite16bitMono("main3to2", 0.1, func(wf wav.File) []float64 {
		s := make([]float64, wf.SoundSize)
		for i := float64(1); i <= 44; i += 2 {
			for n := uint32(0); n < wf.SoundSize; n++ {
				// 丸め誤差を考慮してfloat64で音データを扱っている
				s[n] += 1.0 / i * math.Sin(2.0*math.Pi*i*f0*float64(n)/float64(wf.SampleRate))
			}
		}
		return s
	})
}

// 三角波
func main3to3() string {
	f0 := 500.0 // 周波数
	return waveWrite16bitMono("main3to3", 0.1, func(wf wav.File) []float64 {
		s := make([]float64, wf.SoundSize)
		for i := float64(1); i <= 44; i += 2 {
			for n := uint32(0); n < wf.SoundSize; n++ {
				// 丸め誤差を考慮してfloat64で音データを扱っている
				s[n] += 1.0 / i / i * math.Sin(math.Pi*i/2.0) * math.Sin(2.0*math.Pi*i*f0*float64(n)/float64(wf.SampleRate))
			}
		}
		return s
	})
}

func randomFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

// 白色雑音
func main3to5() string {
	f0 := 1.0 // 周波数
	return waveWrite16bitMono("main3to5", 0.001, func(wf wav.File) []float64 {
		s := make([]float64, wf.SoundSize)
		for i := float64(1); i <= 22050; i++ {
			theta := randomFloat64(0.0, 1.0) * 2.0 * math.Pi
			for n := uint32(0); n < wf.SoundSize; n++ {
				// 丸め誤差を考慮してfloat64で音データを扱っている
				s[n] += math.Sin((2.0 * math.Pi * i * f0 * float64(n) / float64(wf.SampleRate)) + theta)
			}
		}
		return s
	})
}

func main() {
	log.SetFlags(log.Llongfile)

	//main3to1()
	//main3to2()
	//main3to3()
	main3to5()
}
