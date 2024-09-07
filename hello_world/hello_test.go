package main

import "testing"

func Test_Hello_Empty(t *testing.T) {
	got := Hello("", "")
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func Test_Hello_Romanian_Name(t *testing.T) {
	got := Hello("Ioan", "Romanian")
	want := "Salut, Ioan"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func Test_Hello_Empty_Name(t *testing.T) {
	got := Hello("John", "")
	want := "Hello, John"

	if got != want {
		t.Errorf("\ngot:  %q \nwant: %q", got, want)
	}
}
