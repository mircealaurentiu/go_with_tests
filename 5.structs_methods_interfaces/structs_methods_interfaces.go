package structsmethodsinterfaces

import "math"

type Rectange struct {
	Width  float64
	Height float64
}

func (r Rectange) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectange) float64 {

	return_value := 2 * (rectangle.Width + rectangle.Height)

	return return_value
}

func Area(rectangle Rectange) float64 {

	return_value := rectangle.Width * rectangle.Height

	return return_value
}
