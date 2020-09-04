package sh2d

import "github.com/hajimehoshi/ebiten"

//RotateAtCenter totates a matrix on its center
//
//Parameters:
//	- `renderPos`: matrix to rotate
//	- `rotation`: rotation in radians
//	- `width`: width of the sprite
//	- `height`: height of the sprite
//	- `position`: position of the sprite
func RotateAtCenter(renderPos *ebiten.GeoM, rotation float64, width, height float64) {
	renderPos.Translate(-float64(width)/2, -float64(height)/2)
	renderPos.Rotate(rotation)
	renderPos.Translate(float64(width)/2, float64(height)/2)
}
