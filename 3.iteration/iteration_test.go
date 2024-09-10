package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("9", 5)
	want := "99999"

	if got != want {
		t.Errorf("\ngot:  %q \nwant: %q", got, want)
	}
}
