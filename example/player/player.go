package player

import (
	"fmt"

	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/comp2d"
	"github.com/lethal-bacon0/sh2d/example/assets"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

//NewPlayer creates a new player game object
func NewPlayer(position geometry2d.Vector2D) *sh2d.GameObject {
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
	player.AddComponent(comp2d.NewSpriteRenderer(playerTexture, player, comp2d.CenterCenter))
	player.AddComponent(newKeyboardController(player))

	return player
}
