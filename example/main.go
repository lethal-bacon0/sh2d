package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/example/gamemode"
	"github.com/lethal-bacon0/sh2d/example/player"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	game := sh2d.NewGame(screenWidth, screenHeight, "Asteroid")
	scene := sh2d.Scene{}
	scene.AddChild(player.NewPlayer(geometry2d.Vector2D{X: screenWidth / 2, Y: screenHeight / 2}))
	scene.AddChild(gamemode.NewAsteroidSpawner())
	// scene.AddChild(gamemode.NewAsteroidDestroyer())
	sh2d.SetActiveScene(&scene)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
