package player

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/bullet"
	"github.com/lethal-bacon0/sh2d/example/settings"
)

//KeyboardController handles player movement
type keyboardController struct {
	isShooting bool
	container  *sh2d.GameObject
}

func (k *keyboardController) Update() error {
	k.isShooting = false

	if ebiten.IsKeyPressed(settings.Shoot0) || ebiten.IsMouseButtonPressed(settings.Shoot1) {
		k.isShooting = true
		mouseX, mouseY := ebiten.CursorPosition()
		direction := k.container.Position.GetDirVector(sh2d.Vector2D{X: float64(mouseX), Y: float64(mouseY)})
		bullet := bullet.NewBullet(k.container.Position, k.container.Rotation, direction.GetUnitVector())
		sh2d.GetActiveScene().AddChild(bullet)
	}
	return nil
}

func (k *keyboardController) Draw(screen *ebiten.Image) error {
	return nil
}

func newKeyboardController(container *sh2d.GameObject) *keyboardController {
	return &keyboardController{
		container: container,
	}
}
