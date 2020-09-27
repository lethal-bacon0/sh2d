package enemy

import (
	"math"
	"math/rand"
	"time"

	"github.com/lethal-bacon0/sh2d"
	"github.com/lethal-bacon0/sh2d/comp2d"
	"github.com/lethal-bacon0/sh2d/example/assets"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

var (
	textures [][]byte
)

func init() {
	tex1, err := assets.Asset("resources/enemy/Asteroid 01_png_processed.png")
	if err != nil {
		panic(err.Error())
	}
	tex2, err := assets.Asset("resources/enemy/Asteroid 02_png_processed.png")
	if err != nil {
		panic(err.Error())
	}
	tex3, err := assets.Asset("resources/enemy/Asteroid 03_png_processed.png")
	if err != nil {
		panic(err.Error())
	}
	tex4, err := assets.Asset("resources/enemy/Asteroid 04_png_processed.png")
	if err != nil {
		panic(err.Error())
	}
	textures = [][]byte{
		tex1, tex2, tex3, tex4,
	}
}

//NewAsteroid creates a new asteroid
func NewAsteroid(position geometry2d.Vector2D) *sh2d.GameObject {
	var (
		asteroid         *sh2d.GameObject
		lifetimeCallback func()
		texIndex         int
		timer            *comp2d.Timer
		direction        geometry2d.Vector2D
		windowWidth      = sh2d.GetWindowWidth()
		windowHeight     = sh2d.GetWindowHeight()
	)
	asteroid = &sh2d.GameObject{
		Position: position,
		Active:   true,
		Rotation: float64(rand.Intn(360)) * math.Pi / 180,
	}
	texIndex = rand.Intn(len(textures))
	lifetimeCallback = func() {
		asteroid.Active = false
	}
	target := geometry2d.Vector2D{
		X: float64(rand.Intn(windowWidth)),
		Y: float64(rand.Intn(windowHeight)),
	}
	direction = asteroid.Position.GetDirVector(target)
	timer = comp2d.NewTimer(lifetimeCallback, false, time.Minute)
	timer.Start()
	asteroid.AddComponent(timer)
	asteroid.AddComponent(comp2d.NewSpriteRenderer(textures[texIndex], asteroid, comp2d.CenterCenter))
	asteroid.AddComponent(comp2d.NewSpriteMover(asteroid, direction.GetUnitVector(), 0.2))

	return asteroid
}
