package comp2d

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

//SpriteMover moves a sprite by a given velocity to a defined direction
type SpriteMover struct {
	container *sh2d.GameObject
	direction geometry2d.Vector2D
	velocity  float64
}

//Update moves a sprite
func (b *SpriteMover) Update(delta int64) error {
	vel := geometry2d.Vector2D{
		X: b.direction.X * b.velocity * float64(delta),
		Y: b.direction.Y * b.velocity * float64(delta),
	}

	pos := b.container.Position.Add(vel)

	b.container.Position = geometry2d.Vector2D{
		X: pos.X,
		Y: pos.Y,
	}

	return nil
}

//Draw Draw
func (b *SpriteMover) Draw(screen *ebiten.Image) error {

	return nil
}

//NewSpriteMover creates a new sprite mover
func NewSpriteMover(container *sh2d.GameObject, direction geometry2d.Vector2D, velocity float64) *SpriteMover {
	return &SpriteMover{
		container: container,
		direction: direction,
		velocity:  velocity,
	}
}
