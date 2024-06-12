package main

import "testing"

func TestHello(t *testing.T) {
	want := "hello world!"
	got := Hello()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
