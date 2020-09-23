package sh2d

var (
	windowWidth  int
	windowHeight int
	windowName   string
	activeScene  *Scene
)

//GetWindowWidth GetWindowWidth
func GetWindowWidth() int {
	return windowWidth
}

//GetWindowHeight GetWindowHeight
func GetWindowHeight() int {
	return windowHeight
}

//GetWindowName GetWindowName
func GetWindowName() string {
	return windowName
}

//GetActiveScene GetActiveScene
func GetActiveScene() *Scene {
	return activeScene
}

//SetActiveScene SetActiveScene
func SetActiveScene(newScene *Scene) {
	if newScene == nil {
		panic("Scene cannot be nil")
	}
	activeScene = newScene
}
