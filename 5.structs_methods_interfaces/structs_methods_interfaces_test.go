package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectange := Rectange{10.0, 10.0}
	got := Perimeter(rectange)
	want := 40.0

	if got != want {
		t.Errorf("\ngot:  %.2f \nwant: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectange{6.0, 7.0}
		got := rectangle.Area()
		want := 42.0

		if got != want {
			t.Errorf("\ngot:  %.2f \nwant: %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("\ngot:  %g \nwant: %g", got, want)
		}

	})
}
