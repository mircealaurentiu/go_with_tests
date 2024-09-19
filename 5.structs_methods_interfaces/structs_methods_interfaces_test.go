package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("\ngot:  %.2f \nwant: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(6.0, 7.0)
	want := 42.0

	if got != want {
		t.Errorf("\ngot:  %.2f \nwant: %.2f", got, want)
	}
}
