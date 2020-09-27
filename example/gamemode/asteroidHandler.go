package gamemode

import (
	"math/rand"
	"time"

	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/comp2d"
	"github.com/lethal-bacon0/sh2d/example/enemy"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

var (
	before          = time.Now()
	activeAsteroids []*sh2d.GameObject
	windowWidth     int
	windowHeight    int
)

func init() {
	go destroyAsteroids()
}

//NewAsteroidSpawner spawns asteroids
func NewAsteroidSpawner() *sh2d.GameObject {
	spawner := sh2d.GameObject{
		Active: true,
	}
	windowWidth = sh2d.GetWindowWidth()
	windowHeight = sh2d.GetWindowHeight()
	timerCallback := func() {
		position := geometry2d.Vector2D{
			X: float64(rand.Intn(windowWidth)),
			Y: float64(rand.Intn(windowHeight)),
		}
		asteroid := enemy.NewAsteroid(position)
		sh2d.GetActiveScene().AddChild(asteroid)
		activeAsteroids = append(activeAsteroids, asteroid)
	}
	spawner.AddComponent(comp2d.NewTimer(timerCallback, true, 500*time.Millisecond))
	return &spawner
}

func destroyAsteroids() {
	for true {
		for i, asteroid := range activeAsteroids {
			if !asteroid.Active {
				sh2d.GetActiveScene().RemoveChild(asteroid)
				activeAsteroids = append(activeAsteroids[:i], activeAsteroids[i+1:]...)
			}
		}
	}
}
