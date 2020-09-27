package bullet

import (
	"fmt"

	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/assets"
)

//NewBullet creates a new bullet that flies in a certain direction
func NewBullet(position sh2d.Vector2D, rotation float64, direction sh2d.Vector2D) *sh2d.GameObject {
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

	bullet.AddComponent(sh2dcomp.NewSpriteRenderer(bulletTexture, bullet, sh2dcomp.CenterCenter))
	bullet.AddComponent(sh2dcomp.NewSpriteMover(bullet, direction, bulletVelocity))

	return bullet
}
