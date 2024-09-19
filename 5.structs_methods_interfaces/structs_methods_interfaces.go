package structsmethodsinterfaces

type Rectange struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectange) float64 {

	return_value := 2 * (rectangle.Width + rectangle.Height)

	return return_value
}

func Area(rectangle Rectange) float64 {

	return_value := rectangle.Width * rectangle.Height

	return return_value
}
