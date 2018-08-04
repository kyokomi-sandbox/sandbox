// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"

	"image"
	"image/color"
	gomath "math"
	"os"
	"log"
)

var data = flag.String("data", "", "path to data")

func buildStar(theme gxui.Theme, center math.Point, radius, rotation float32, points int) gxui.Image {
	p := make(gxui.Polygon, points*2)
	for i := 0; i < points*2; i++ {
		frac := float32(i) / float32(points*2)
		α := frac*math.TwoPi + rotation
		r := []float32{radius, radius / 2}[i&1]
		p[i] = gxui.PolygonVertex{
			Position: math.Point{
				X: center.X + int(r*math.Cosf(α)),
				Y: center.Y + int(r*math.Sinf(α)),
			},
			RoundedRadius: []float32{0, 50}[i&1],
		}
	}
	image := theme.CreateImage()
	image.SetPolygon(p, gxui.CreatePen(3, gxui.Red), gxui.CreateBrush(gxui.Yellow))
	return image
}

func buildMoon(theme gxui.Theme, center math.Point, radius float32) gxui.Image {
	c := 40
	p := make(gxui.Polygon, c*2)
	for i := 0; i < c; i++ {
		frac := float32(i) / float32(c)
		α := math.Lerpf(math.Pi*1.2, math.Pi*-0.2, frac)
		p[i] = gxui.PolygonVertex{
			Position: math.Point{
				X: center.X + int(radius*math.Sinf(α)),
				Y: center.Y + int(radius*math.Cosf(α)),
			},
			RoundedRadius: 0,
		}
	}
	for i := 0; i < c; i++ {
		frac := float32(i) / float32(c)
		α := math.Lerpf(math.Pi*-0.2, math.Pi*1.2, frac)
		r := math.Lerpf(radius, radius*0.5, math.Sinf(frac*math.Pi))
		p[i+c] = gxui.PolygonVertex{
			Position: math.Point{
				X: center.X + int(r*math.Sinf(α)),
				Y: center.Y + int(r*math.Cosf(α)),
			},
			RoundedRadius: 0,
		}
	}
	image := theme.CreateImage()
	image.SetPolygon(p, gxui.CreatePen(3, gxui.Gray80), gxui.CreateBrush(gxui.Gray40))
	return image
}

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(800, 600, "Polygon")

	container := theme.CreateLinearLayout()

//	texture := driver.CreateTexture(getImage(cacheDir, tweet.User.ProfileImageURL), 96)
	f, err := os.Open("data/icont19.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	im, _, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 100; i++ {
		pict := theme.CreateImage()
		texture := driver.CreateTexture(im, 96)
		texture.SetFlipY(true)
		pict.SetTexture(texture)
		pict.SetExplicitSize(math.Size{32, 32})
		pict.SetMargin(math.CreateSpacing(4))
		container.AddChild(pict)
	}

	label := theme.CreateLabel()
	label.SetText("hogehogehoge")
	label.SetMargin(math.CreateSpacing(200))
	container.AddChild(label)

	window.OnClose(driver.Terminate)

	window.AddChild(container)

	gxui.EventLoop(driver)
}

func depthToImage(img *image.RGBA, w int, h int, buffer []byte) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			bits := (uint32(buffer[3]) << 24) | (uint32(buffer[2]) << 16) | (uint32(buffer[1]) << 8) | (uint32(buffer[0]) << 0)
			depth := gomath.Float32frombits(bits)
			buffer = buffer[4:]

			d := 0.01 / (1.0 - depth)
			c := color.RGBA{
				R: byte(math.Cosf(d+math.TwoPi*0.000)*127.0 + 128.0),
				G: byte(math.Cosf(d+math.TwoPi*0.333)*127.0 + 128.0),
				B: byte(math.Cosf(d+math.TwoPi*0.666)*127.0 + 128.0),
				A: byte(0xFF),
			}
			img.Set(x, y, c)
		}
	}
}

func main() {
	flag.Parse()
	gl.StartDriver(*data, appMain)
}
