package main

import "github.com/hajimehoshi/ebiten"

type Vec2 struct {
	x, y float64
}

type Node interface {
	draw(scene *ebiten.Image) error
}
