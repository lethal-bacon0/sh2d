package sh2d

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Gameloop implements ebiten.Gameloop interface.
type Gameloop struct {
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Gameloop) Update(screen *ebiten.Image) error {
	for _, element := range activeScene.SceneElements {
		err := element.Update()
		if err != nil {
			ebitenutil.DebugPrint(screen, err.Error())
		}
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Gameloop) Draw(screen *ebiten.Image) {
	for _, element := range activeScene.SceneElements {
		err := element.Draw(screen)
		if err != nil {
			ebitenutil.DebugPrint(screen, err.Error())
		}
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Gameloop) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

//NewGame creates a new game
func NewGame(screenWidth, screenHeight int, name string) *Gameloop {
	game := &Gameloop{}
	windowHeight = screenHeight
	windowWidth = screenWidth
	windowName = name
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(name)

	return game
}
