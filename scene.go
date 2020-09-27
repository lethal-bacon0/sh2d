package sh2d

//Scene contains all elements on screen
type Scene struct {
	SceneElements []*GameObject
}

//TODO: add layer functionality

//AddChild adds new child to scene
func (s *Scene) AddChild(child *GameObject) {
	s.SceneElements = append(s.SceneElements, child)
}

//RemoveChild removes a given child from the scene
func (s *Scene) RemoveChild(child *GameObject) {
	for i, element := range s.SceneElements {
		if child == element {
			s.SceneElements = append(s.SceneElements[:i], s.SceneElements[i+1:]...)
		}
	}
}
