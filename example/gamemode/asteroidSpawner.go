package gamemode

import (
	"fmt"
	"time"

	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/comp2d"
)

var (
	before = time.Now()
)

//NewAsteroidSpawner spawns asteroids
func NewAsteroidSpawner() *sh2d.GameObject {
	spawner := sh2d.GameObject{
		Active: true,
	}
	timerCallback := func() {
		delta := time.Now().Sub(before)
		fmt.Printf("Elapsed after %v\n", delta.Seconds())
		before = time.Now()
	}
	spawner.AddComponent(comp2d.NewTimer(timerCallback, true, 100*time.Millisecond))
	return &spawner
}
