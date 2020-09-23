package sh2d

import (
	"github.com/hajimehoshi/ebiten"
)

//SpriteMover moves a sprite by a given velocity to a defined direction
type SpriteMover struct {
	container *GameObject
	direction Vector2D
	velocity  float64
}

//Update moves a sprite
func (b *SpriteMover) Update(delta int64) error {
	vel := Vector2D{
		X: b.direction.X * b.velocity * float64(delta),
		Y: b.direction.Y * b.velocity * float64(delta),
	}

	pos := b.container.Position.Add(vel)

	b.container.Position = Vector2D{
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
func NewSpriteMover(container *GameObject, direction Vector2D, velocity float64) *SpriteMover {
	return &SpriteMover{
		container: container,
		direction: direction,
		velocity:  velocity,
	}
}
