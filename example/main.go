package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/asteroids"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func main() {
	game := sh2d.NewGame(screenWidth, screenHeight, "Asteroid")
	scene := sh2d.Scene{}
	scene.AddChild(asteroids.NewPlayer(sh2d.Vector2D{X: screenWidth / 2, Y: screenHeight / 2}))
	sh2d.SetActiveScene(&scene)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
