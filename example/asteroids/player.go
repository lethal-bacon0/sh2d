package asteroids

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/assets"
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
	player.AddComponent(newKeyboardController(player))

	return player
}

//KeyboardController handles player movement
type keyboardController struct {
	isShooting bool
	container  *sh2d.GameObject
}

func (k *keyboardController) Update() error {
	k.isShooting = false

	if ebiten.IsKeyPressed(shoot) {
		k.isShooting = true
		mouseX, mouseY := ebiten.CursorPosition()
		direction := k.container.Position.GetDirVector(sh2d.Vector2D{X: float64(mouseX), Y: float64(mouseY)})
		bullet := NewBullet(k.container.Position, k.container.Rotation, direction.GetUnitVector())
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

//playerMouseController handles mouse input for the player
type playerMouseController struct {
	container *sh2d.GameObject
	position  *sh2d.Vector2D
	sprite    *sh2d.SpriteRenderer
}

//Update updates the controller every frame
func (c *playerMouseController) Update() error {
	x, y := ebiten.CursorPosition()
	c.position = &sh2d.Vector2D{
		X: float64(x),
		Y: float64(y),
	}

	if c.sprite == nil {
		c.sprite = c.container.GetComponent(&sh2d.SpriteRenderer{}).(*sh2d.SpriteRenderer)
	}
	dx := c.container.Position.X - c.position.X
	dy := c.container.Position.Y - c.position.Y
	c.container.Rotation = math.Atan2(dy, dx) + float64(-90*math.Pi/180)
	return nil
}

//Draw draws the mouse controller
func (c *playerMouseController) Draw(screen *ebiten.Image) error {
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Rotation: %v", controller.container.Rotation))
	return nil
}

//newPlayerMouseController creates a new PlayerMouseController
func newPlayerMouseController(container *sh2d.GameObject) *playerMouseController {
	return &playerMouseController{
		container: container,
	}
}
