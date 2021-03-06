// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin linux

package main

import (
	"image"
	"log"

	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
)

const (
	tileWidth, tileHeight = 180, 180 // width and height of each tile
	tilesX, tilesY        = 16, 16   // number of horizontal tiles

	gopherTile = 1 // which tile the gopher is standing on (0-indexed)
)

type Game struct {
}

func NewGame() *Game {
	var g Game
	return &g
}

func (g *Game) Scene(eng sprite.Engine) *sprite.Node {
	texs := loadTextures(eng)

	scene := &sprite.Node{}
	eng.Register(scene)
	eng.SetTransform(scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})

	newNode := func(fn arrangerFunc) {
		n := &sprite.Node{Arranger: arrangerFunc(fn)}
		eng.Register(n)
		scene.AppendChild(n)
	}

	// The gopher.
	newNode(func(eng sprite.Engine, n *sprite.Node, t clock.Time) {
		eng.SetSubTex(n, texs[texGopher])
		eng.SetTransform(n, f32.Affine{
			{tileWidth, 0, 0},
			{0, tileHeight, 0},
		})
	})
	return scene
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }

const (
	texGopher = iota
)

func loadTextures(eng sprite.Engine) []sprite.SubTex {
	a, err := asset.Open("koha1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	m, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}

	t, err := eng.LoadTexture(m)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", m.Bounds())

	x, y := t.Bounds()
	return []sprite.SubTex{
		texGopher: sprite.SubTex{t, image.Rect(0, 0, x, y)},
	}
}
