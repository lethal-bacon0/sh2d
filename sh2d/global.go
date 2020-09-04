package sh2d

var (
	windowWidth  int
	windowHeight int
	windowName   string
	activeScene  *Scene
)

func GetWindowWidth() int {
	return windowWidth
}

func GetWindowHeight() int {
	return windowHeight
}

func GetWindowName() string {
	return windowName
}

func GetActiveScene() *Scene {
	return activeScene
}

func SetActiveScene(newScene *Scene) {
	if newScene == nil {
		panic("Scene cannot be nil")
	}
	activeScene = newScene
}
