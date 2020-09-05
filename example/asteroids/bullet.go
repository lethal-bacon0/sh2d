package asteroids

import (
	"fmt"

	"github.com/lethal-bacon0/Asteroid/assets"
	"github.com/lethal-bacon0/sh2d"
)

const bulletTexturePath = "resources/player/Laser_Large_png_processed.png"

//NewBullet creates a new bullet that flies in a certain direction
func NewBullet(position sh2d.Vector2D, rotation float64, velocity sh2d.Vector2D) *sh2d.GameObject {
	bullet := &sh2d.GameObject{
		Position: position,
		Active:   true,
		Rotation: rotation,
	}
	bulletTexture, err := assets.Asset(bulletTexturePath)
	if err != nil {
		panic(fmt.Sprintf("Resource not found: %v", bulletTexturePath))
	}

	bullet.AddComponent(sh2d.NewSpriteRenderer(bulletTexture, bullet, sh2d.CENTER))
	return bullet
}
