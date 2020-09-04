package sh2d

type Scene struct {
	SceneElements []GameObject
}

func (s *Scene) AddChild(child *GameObject) {
	s.SceneElements = append(s.SceneElements, *child)
}
