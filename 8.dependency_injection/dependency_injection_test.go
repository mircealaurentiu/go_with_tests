package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John")

	got := buffer.String()
	want := "Hello, John"

	if got != want {
		t.Errorf("\ngot %q \nwant %q", got, want)
	}
}
