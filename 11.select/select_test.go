package selectt

import "testing"

func TestRace(t *testing.T) {
	slowURL := "http://facebook.com"
	fastURL := "http://quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("\ngot %q \nwant %q", got, want)
	}
}
