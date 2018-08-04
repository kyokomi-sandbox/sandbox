package main

import (
	"flag"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"image"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func main() {
	var imageURL string
	flag.StringVar(&imageURL, "image", "", "")
	flag.Parse()

	if imageURL == "" {
		imageURL = "http://pics.dmm.co.jp/digital/pcgame/guide/selen_0024/head.jpg"
	}

	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Header.Get("Content-Type"))

	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	{
		m := resize.Resize(300, 300, img, resize.Bicubic)

		out, err := os.Create("resize.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

	{
		m := resize.Resize(300, 0, img, resize.Bicubic)

		out, err := os.Create("resize_width.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

	{
		m := resize.Resize(0, 300, img, resize.Bicubic)

		out, err := os.Create("resize_height.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

	{
		m := resize.Thumbnail(300, 300, img, resize.Bicubic)
		out, err := os.Create("thumbnail.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

	{
		m, err := cutter.Crop(img, cutter.Config{
			Width:  300,
			Height: 300,
			Anchor: image.Point{0, 0},
			Mode:   cutter.Centered, // optional, default value
		})
		out, err := os.Create("crop_centered.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

	{
		x, y := uint(640), uint(640)
		if img.Bounds().Size().X < img.Bounds().Size().Y {
			y = 0
		} else {
			x = 0
		}
		m, err := cutter.Crop(resize.Resize(x, y, img, resize.Bicubic), cutter.Config{
			Width:  640,
			Height: 640,
			Anchor: image.Point{0, 0},
			Mode:   cutter.Centered, // optional, default value
		})
		out, err := os.Create("resize_crop_centered.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}
}
