package player

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/bullet"
	"github.com/lethal-bacon0/sh2d/example/settings"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

//KeyboardController handles player movement
type keyboardController struct {
	isShooting bool
	container  *sh2d.GameObject
	lastShot   time.Time
	firingRate int64
}

func (k *keyboardController) Update(delta int64) error {
	const firingThreshhold = 9
	k.isShooting = false

	if ebiten.IsKeyPressed(settings.Shoot0) || ebiten.IsMouseButtonPressed(settings.Shoot1) {
		timeDelta := time.Now().Sub(k.lastShot)
		if timeDelta.Milliseconds() >= k.firingRate {
			k.shoot()
			k.lastShot = time.Now() //TODO: firing rate feels kinda hacky
		}
	}
	return nil
}

func (k *keyboardController) Draw(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %v", ebiten.CurrentFPS()))
	return nil
}

func (k *keyboardController) shoot() {
	var (
		mouseX, mouseY  int
		mousePos        geometry2d.Vector2D
		bulletDirection geometry2d.Vector2D
	)

	k.isShooting = true
	mouseX, mouseY = ebiten.CursorPosition()
	mousePos = geometry2d.Vector2D{
		X: float64(mouseX),
		Y: float64(mouseY),
	}
	//TODO: spawn bullet in front of sprite
	bulletDirection = k.container.Position.GetDirVector(mousePos)
	bulletDirUnit := bulletDirection.GetUnitVector()
	bullet := bullet.NewBullet(k.container.Position, k.container.Rotation, bulletDirUnit)
	sh2d.GetActiveScene().AddChild(bullet)
}

func newKeyboardController(container *sh2d.GameObject) *keyboardController {
	return &keyboardController{
		container:  container,
		firingRate: 99,
	}
}
