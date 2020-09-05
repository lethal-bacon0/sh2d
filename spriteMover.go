package sh2d

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
)

//Sp
type SpriteMover struct {
	container *sh2d.GameObject
	direction sh2d.Vector2D
	velocity  float64
}

func (b *SpriteMover) Update() error {
	vel := sh2d.Vector2D{
		X: b.direction.X * b.velocity,
		Y: b.direction.Y * b.velocity,
	}

	pos := b.container.Position.Add(vel)

	b.container.Position = sh2d.Vector2D{
		X: pos.X,
		Y: pos.Y,
	}

	return nil
}

func (b *SpriteMover) Draw(screen *ebiten.Image) error {

	return nil
}

func NewSpriteMover(container *sh2d.GameObject, direction sh2d.Vector2D, velocity float64) *SpriteMover {
	return &SpriteMover{
		container: container,
		direction: direction,
		velocity:  velocity,
	}
}
