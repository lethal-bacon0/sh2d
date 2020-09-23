package sh2d

import "github.com/hajimehoshi/ebiten"

//Component asd
type Component interface {
	Update(delta int64) error
	Draw(screen *ebiten.Image) error
}
