package sh2dcomp

import (
	"bytes"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

//SpriteRenderer is a component to render a sprite
type SpriteRenderer struct {
	Texture   *ebiten.Image
	container *GameObject
	Anchor    *RenderPositions
	Scale     *float64
}

//Update updates the component.
func (sp *SpriteRenderer) Update(delta int64) error {
	return nil
}

//Draw draws the component.
func (sp *SpriteRenderer) Draw(screen *ebiten.Image) error {
	renderPos := ebiten.GeoM{}
	// renderPos.Scale(sp.Width, sp.Height)
	w, h := sp.Texture.Size()
	RotateAtCenter(&renderPos, sp.container.Rotation, float64(w), float64(h))
	switch *sp.Anchor {
	case TopLeft:
		renderPos.Translate(sp.container.Position.X, sp.container.Position.Y)
	case CenterCenter:
		center := sp.GetAbsCenter()
		renderPos.Translate(center.X, center.Y)
	}
	screen.DrawImage(sp.Texture, &ebiten.DrawImageOptions{
		GeoM: renderPos,
	})
	return nil
}

//NewSpriteRenderer creates a new sprite renderer
func NewSpriteRenderer(texture []byte, container *GameObject, anchor RenderPositions) *SpriteRenderer {
	img, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		panic(fmt.Sprintf("Error while parsing sprite: %v", err.Error()))
	}
	screen, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(fmt.Sprintf("Texture with path %v not found", texture))
	}
	return &SpriteRenderer{
		Texture:   screen,
		container: container,
		Anchor:    &anchor,
	}
}

//GetAbsCenter gets centered absolute coordinates
func (sp *SpriteRenderer) GetAbsCenter() Vector2D {
	w, h := sp.Texture.Size()
	return Vector2D{
		X: sp.container.Position.X - float64(w/2),
		Y: sp.container.Position.Y - float64(h/2),
	}
}

//GetRelCenter gets centered relative coordinates
func (sp *SpriteRenderer) GetRelCenter() Vector2D {
	w, h := sp.Texture.Size()
	return Vector2D{
		X: float64(w / 2),
		Y: float64(h / 2),
	}
}

//GetAbsTopCenter gets absolute coordinates of top center
func (sp *SpriteRenderer) GetAbsTopCenter() Vector2D {
	w, _ := sp.Texture.Size()
	return Vector2D{
		X: sp.container.Position.X - float64(w/2),
		Y: sp.container.Position.Y,
	}
}

//RenderPositions defines anchor positions
type RenderPositions string

const (
	//CenterCenter center of sprite
	CenterCenter RenderPositions = "CENTER"
	//TopLeft topleft of sprite
	TopLeft RenderPositions = "LEFT"
)
