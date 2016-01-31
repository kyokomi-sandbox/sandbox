package main

import (
	"flag"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	var imageURL string
	flag.StringVar(&imageURL, "image", "", "")
	flag.Parse()

	if imageURL == "" {
		imageURL = "http://dengeki-hime.com/dengeki-hime/wp-content/uploads/2016/01/DH201601_khmg_06_01s.jpg"
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
		m := resize.Resize(100, 0, img, resize.Lanczos3)

		out, err := os.Create("resize_100_test_resized.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

	{
		m := resize.Thumbnail(800, 800, img, resize.Lanczos3)
		out, err := os.Create("thumb_400_test_resized.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}

}
