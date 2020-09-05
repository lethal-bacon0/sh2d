package sh2d

//Vector2D represents a simple 2D vector
type Vector2D struct {
	X, Y float64
}

//Sub subtracts two vectors and return a new Vector2D
func (v1 Vector2D) Sub(v2 Vector2D) Vector2D {
	return Vector2D{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

//Add adds two vectors and return the result as a new vector
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

//GetDirVector gets the directional vector of two vectors
func (v1 Vector2D) GetDirVector(v2 Vector2D) Vector2D {
	return Vector2D{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
	}
}
