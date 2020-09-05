package mobile

import (
	"github.com/hajimehoshi/ebiten/mobile"
	"github.com/lethal-bacon0/Asteroid/asteroids"
	"github.com/lethal-bacon0/sh2d"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func init() {
	// yourgame.Game must implement mobile.Game (= ebiten.Game) interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten?tab=doc#Game
	game := sh2d.NewGame(screenWidth, screenHeight, "Asteroid")
	scene := sh2d.Scene{}
	scene.AddChild(asteroids.NewPlayer(sh2d.Vector2D{X: screenWidth / 2, Y: screenHeight / 2}))
	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
