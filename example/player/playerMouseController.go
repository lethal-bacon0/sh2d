package player

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/comp2d"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

//mouseController handles mouse input for the player
type mouseController struct {
	container *sh2d.GameObject
	position  *geometry2d.Vector2D
	sprite    *comp2d.SpriteRenderer
}

//Update updates the controller every frame
func (c *mouseController) Update(delta int64) error {
	x, y := ebiten.CursorPosition()
	c.position = &geometry2d.Vector2D{
		X: float64(x),
		Y: float64(y),
	}

	if c.sprite == nil {
		c.sprite = c.container.GetComponent(&comp2d.SpriteRenderer{}).(*comp2d.SpriteRenderer)
	}
	dx := c.container.Position.X - c.position.X
	dy := c.container.Position.Y - c.position.Y
	c.container.Rotation = math.Atan2(dy, dx) + float64(-90*math.Pi/180)
	return nil
}

//Draw draws the mouse controller
func (c *mouseController) Draw(screen *ebiten.Image) error {
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Rotation: %v", controller.container.Rotation))
	return nil
}

//mouseController creates a new PlayerMouseController
func newPlayerMouseController(container *sh2d.GameObject) *mouseController {
	return &mouseController{
		container: container,
	}
}
