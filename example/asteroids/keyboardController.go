package asteroids

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
)

//KeyboardController handles player movement
type KeyboardController struct {
	isShooting bool
	container  *sh2d.GameObject
}

func (k *KeyboardController) Update() error {
	k.isShooting = false

	if ebiten.IsKeyPressed(shoot) {
		k.isShooting = true
		bullet := NewBullet(k.container.Position, k.container.Rotation, sh2d.Vector2D{})
		sh2d.GetActiveScene().AddChild(bullet)
	}
	return nil
}

func (k *KeyboardController) Draw(screen *ebiten.Image) error {
	return nil
}

func NewKeyboardController(container *sh2d.GameObject) *KeyboardController {
	return &KeyboardController{
		container: container,
	}
}
