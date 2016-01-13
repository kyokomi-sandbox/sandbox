package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Sprite struct {
	*ebiten.Image
	scaleX, scaleY float64
	position Vec2
}

func NewSprite(fileName string) (*Sprite, error) {
	i, _, err := ebitenutil.NewImageFromFile(fileName, ebiten.FilterNearest)
	if err != nil {
		return nil, err
	}

	s := Sprite{}
	s.Image = i
	s.scaleX = 1
	s.position = Vec2{0, 0}
	
	return &s, nil
}

func (s *Sprite) draw(scene *ebiten.Image) error {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.scaleX, s.scaleY)
	op.GeoM.Translate(s.position.x, s.position.y)
	if err := scene.DrawImage(s.Image, op); err != nil {
		return err
	}
	return nil
}

func (s Sprite) Scale() (x, y float64) {
	return s.scaleX, s.scaleY
}

func (s Sprite) ScaleX() (float64) {
	return s.scaleX
}

func (s Sprite) ScaleY() (float64) {
	return s.scaleY
}

func (s *Sprite) SetScale(x, y float64) {
	s.scaleX = x
	s.scaleY = y
}

func (s *Sprite) SetPositionVec2(p Vec2) {
	s.position = p
}

func (s *Sprite) SetPosition(x, y float64) {
	s.SetPositionVec2(Vec2{x: x, y: y})
}

var _ Node = (*Sprite)(nil)
