package asteroids

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
)

type BulletMover struct {
	container *sh2d.GameObject
	velocity  sh2d.Vector2D
}

func (b *BulletMover) Update() error {
	b.container.Position.Add(b.velocity)
	return nil
}

func (b *BulletMover) Draw(screen *ebiten.Image) error {
	return nil
}
