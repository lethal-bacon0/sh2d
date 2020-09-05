package player

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/assets"
	"github.com/lethal-bacon0/sh2d/example/bullet"
	"github.com/lethal-bacon0/sh2d/example/settings"
)

//NewPlayer creates a new player game object
func NewPlayer(position sh2d.Vector2D) *sh2d.GameObject {
	const (
		playerSize        = 1
		playerTexturePath = "resources/player/PlayerRed_Frame_01_png_processed.png"
	)

	var player = &sh2d.GameObject{
		Position: position,
		Active:   true,
	}

	playerTexture, err := assets.Asset(playerTexturePath)

	if err != nil {
		panic(fmt.Sprintf("Resource not found: %v", playerTexturePath))
	}

	player.AddComponent(newPlayerMouseController(player))
	player.AddComponent(sh2d.NewSpriteRenderer(playerTexture, player, sh2d.CenterCenter))
	player.AddComponent(newPlayerKeyboardController(player))

	return player
}

//KeyboardController handles player movement
type playerKeyboardController struct {
	isShooting bool
	container  *sh2d.GameObject
}

func (k *playerKeyboardController) Update() error {
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

func (k *playerKeyboardController) Draw(screen *ebiten.Image) error {
	return nil
}

func newPlayerKeyboardController(container *sh2d.GameObject) *playerKeyboardController {
	return &playerKeyboardController{
		container: container,
	}
}
