package sh2d

import "github.com/hajimehoshi/ebiten"

//Component asd
type Component interface {
	Update() error
	Draw(screen *ebiten.Image) error
}
