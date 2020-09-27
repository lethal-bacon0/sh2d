package sh2d

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten"
	"github.com/lethal-bacon0/sh2d/geometry2d"
)

//GameObject is a generic type all elements in the game should inherit
type GameObject struct {
	Position   geometry2d.Vector2D
	Rotation   float64
	Active     bool
	Components []Component
	// Width, Height float64
}

//AddComponent adds a new component to the game object
func (g *GameObject) AddComponent(newComponent Component) {
	for _, existing := range g.Components {
		if reflect.TypeOf(newComponent) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("Duplicate component of type %v", reflect.TypeOf(newComponent)))
		}
	}
	g.Components = append(g.Components, newComponent)
}

//GetComponent gets a component of the speciied type
func (g *GameObject) GetComponent(withType Component) Component {
	for _, component := range g.Components {
		if reflect.TypeOf(withType) == reflect.TypeOf(component) {
			return component
		}
	}

	panic(fmt.Sprintf("Component of type %v was not found", reflect.TypeOf(withType)))
}

//Update updates all conmponents
func (g *GameObject) Update(delta int64) error {
	for _, component := range g.Components {
		err := component.Update(delta)
		if err != nil {
			return err
		}
	}
	return nil
}

//Draw draws all components to screen
func (g *GameObject) Draw(screen *ebiten.Image) error {
	for _, component := range g.Components {
		err := component.Draw(screen)
		if err != nil {
			return err
		}
	}

	return nil
}

// //GetCenter gets the center vector of the current game object
// func (g *GameObject) GetCenter() Vector2D {
// 	return Vector2D{
// 		X: g.Position.X - (g.Width / 2),
// 		Y: g.Position.Y - (g.Height / 2),
// 	}
// }
