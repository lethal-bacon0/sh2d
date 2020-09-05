package asteroids

import (
	"fmt"

	"github.com/lethal-bacon0/Asteroid/assets"
	"github.com/lethal-bacon0/sh2d"
)

const (
	playerSize        = 1
	playerTexturePath = "resources/player/PlayerRed_Frame_01_png_processed.png"
)

//NewPlayer creates a new player game object
func NewPlayer(position sh2d.Vector2D) *sh2d.GameObject {
	var player = &sh2d.GameObject{
		Position: position,
		Active:   true,
	}
	playerTexture, err := assets.Asset(playerTexturePath)
	if err != nil {
		panic(fmt.Sprintf("Resource not found: %v", playerTexturePath))
	}

	player.AddComponent(NewPlayerMouseController(player))
	player.AddComponent(sh2d.NewSpriteRenderer(playerTexture, player, sh2d.CENTER))
	player.AddComponent(NewKeyboardController(player))
	return player
}