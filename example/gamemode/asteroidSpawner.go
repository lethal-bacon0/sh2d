package gamemode

import "github.com/lethal-bacon0/sh2d"

//AsteroidSpawner spawns asteroids
func AsteroidSpawner() *sh2d.GameObject {
	spawner := sh2d.GameObject{
		Active: true,
	}
	return &spawner
}
