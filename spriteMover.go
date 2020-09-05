package sh2d

import (
	"github.com/hajimehoshi/ebiten"
)

//SpriteMover moves a sprite
type SpriteMover struct {
	container *GameObject
	direction Vector2D
	velocity  float64
}

func (b *SpriteMover) Update() error {
	vel := Vector2D{
		X: b.direction.X * b.velocity,
		Y: b.direction.Y * b.velocity,
	}

	pos := b.container.Position.Add(vel)

	b.container.Position = Vector2D{
		X: pos.X,
		Y: pos.Y,
	}

	return nil
}

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
