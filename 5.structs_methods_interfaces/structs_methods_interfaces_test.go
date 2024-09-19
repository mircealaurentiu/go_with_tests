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
	rectange := Rectange{6.0, 7.0}
	got := Area(rectange)
	want := 42.0

	if got != want {
		t.Errorf("\ngot:  %.2f \nwant: %.2f", got, want)
	}
}
