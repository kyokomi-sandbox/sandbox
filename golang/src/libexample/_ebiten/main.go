package main

import (
	"github.com/hajimehoshi/ebiten"
	_ "image/jpeg"
	"log"
	"fmt"
	"image/color"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	count     int
	sprites   []*Sprite
)

func Attach(s *Sprite) error {
	
	if sprites == nil {
		sprites = make([]*Sprite, 0)
	}
	
	sprites = append(sprites, s)
	
	return nil
}

func update(screen *ebiten.Image) error {
	count++

	for _, s := range sprites {
		if err := s.draw(screen); err != nil {
			fmt.Println(err)
		}
	}

	message := "Hello World!"
	if err := ArcadeFont.DrawTextWithShadow(
		screen, message, 40, screenHeight - 40, 1,
		color.NRGBA{0xff, 0xff, 0xff, 0xff}); err != nil {
		return err
	}

	return nil
}

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {

	s, err := NewSprite("images/actor37_0.png")
	if err != nil {
		log.Fatal(err)
	}
	
	w, h := s.Size()
	
	s.SetScale(0.5, 0.5)
	
	x := (screenWidth / 2) - (float64(w) / 2 * s.ScaleX()) + (screenWidth / 3)
	y := 0 + (float64(h) / 3 * s.ScaleY())
	s.SetPosition(x, y)

	if err := Attach(s); err != nil {
		log.Fatal(err)
	}

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Ebiten Demo"); err != nil {
		log.Fatal(err)
	}
}
