package sh2d

import (
	"fmt"
	"image"
	"log"
	"os"

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
func (sp *SpriteRenderer) Update() error {
	return nil
}

//Draw draws the component.
func (sp *SpriteRenderer) Draw(screen *ebiten.Image) error {
	renderPos := ebiten.GeoM{}
	// renderPos.Scale(sp.Width, sp.Height)
	w, h := sp.Texture.Size()
	RotateAtCenter(&renderPos, sp.container.Rotation, float64(w), float64(h))
	switch *sp.Anchor {
	case LEFT:
		renderPos.Translate(sp.container.Position.X, sp.container.Position.Y)
	case CENTER:
		center := sp.GetAbsCenter()
		renderPos.Translate(center.X, center.Y)
	}
	screen.DrawImage(sp.Texture, &ebiten.DrawImageOptions{
		GeoM: renderPos,
	})
	return nil
}

//NewSpriteRenderer creates a new sprite renderer
func NewSpriteRenderer(texturePath string, container *GameObject, anchor RenderPositions) *SpriteRenderer {
	textureBytes, err := os.Open(texturePath)
	img, _, err := image.Decode(textureBytes)
	if err != nil {
		log.Fatal(err)
	}
	texture, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(fmt.Sprintf("Texture with path %v not found", texturePath))
	}
	return &SpriteRenderer{
		Texture:   texture,
		container: container,
		Anchor:    &anchor,
	}
}

func (sp *SpriteRenderer) GetAbsCenter() Vector2D {
	w, h := sp.Texture.Size()
	return Vector2D{
		X: sp.container.Position.X - float64(w/2),
		Y: sp.container.Position.Y - float64(h/2),
	}
}

func (sp *SpriteRenderer) GetRelCenter() Vector2D {
	w, h := sp.Texture.Size()
	return Vector2D{
		X: float64(w / 2),
		Y: float64(h / 2),
	}
}

type RenderPositions string

const (
	CENTER RenderPositions = "CENTER"
	LEFT   RenderPositions = "LEFT"
)
