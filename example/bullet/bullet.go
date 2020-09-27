package bullet

import (
	"fmt"

	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/comp2d"
	"github.com/lethal-bacon0/sh2d/example/assets"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

//NewBullet creates a new bullet that flies in a certain direction
func NewBullet(position geometry2d.Vector2D, rotation float64, direction geometry2d.Vector2D) *sh2d.GameObject {
	const (
		bulletTexturePath = "resources/player/Minigun_Large_png_processed.png"
		bulletVelocity    = 0.9
	)

	bullet := &sh2d.GameObject{
		Position: position,
		Active:   true,
		Rotation: rotation,
	}

	bulletTexture, err := assets.Asset(bulletTexturePath)
	if err != nil {
		panic(fmt.Sprintf("Resource not found: %v", bulletTexturePath))
	}

	bullet.AddComponent(comp2d.NewSpriteRenderer(bulletTexture, bullet, comp2d.CenterCenter))
	bullet.AddComponent(comp2d.NewSpriteMover(bullet, direction, bulletVelocity))

	return bullet
}
