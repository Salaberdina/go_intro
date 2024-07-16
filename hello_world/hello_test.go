package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "hello, Sveta"
		got := Hello("Sveta", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world when parameter is empty", func(t *testing.T) {
		want := "hello, world"
		got := Hello("", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello spanish", func(t *testing.T) {
		want := "hola, Sveta"
		got := Hello("Sveta", spanish)

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello french", func(t *testing.T) {
		want := "bonjour, Sveta"
		got := Hello("Sveta", french)

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello havai", func(t *testing.T) {
		want := "aloha, Sveta"
		got := Hello("Sveta", havai)

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
