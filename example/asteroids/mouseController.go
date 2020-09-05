package asteroids

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
)

//PlayerMouseController handles mouse input for the player
type PlayerMouseController struct {
	container *sh2d.GameObject
	position  *sh2d.Vector2D
	sprite    *sh2d.SpriteRenderer
}

//Update updates the controller every frame
func (c *PlayerMouseController) Update() error {
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

	//TODO: check if player is shooting
	return nil
}

//Draw draws the mouse controller
func (controller *PlayerMouseController) Draw(screen *ebiten.Image) error {
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Rotation: %v", controller.container.Rotation))
	return nil
}

//NewPlayerMouseController creates a new PlayerMouseController
func NewPlayerMouseController(container *sh2d.GameObject) *PlayerMouseController {
	return &PlayerMouseController{
		container: container,
	}
}
