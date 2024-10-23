package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectange := Rectange{10.0, 10.0}
	got := Perimeter(rectange)
	want := 40.0

	if got != want {
		t.Errorf("\ngot:  %.2f \nwant: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("\ngot:  %g \nwant: %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectange{6.0, 7.0}
		checkArea(t, rectangle, 42)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)

	})
}

func TestArea_TableDrivenTests(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectange{Width: 12, Height: 6}, want: 72.0},
		{name: "circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		//intentional failing test
		{name: "triangle", shape: Triangle{8, 5}, want: 21.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("\ngot %g \nwant %g", got, tt.want)
			}
		})
	}
}
