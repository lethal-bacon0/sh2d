package sh2d

//Scene contains all elements on screen
type Scene struct {
	SceneElements []GameObject
}

//TODO: add layer functionality

//AddChild adds new child to scene
func (s *Scene) AddChild(child *GameObject) {
	s.SceneElements = append(s.SceneElements, *child)
}
