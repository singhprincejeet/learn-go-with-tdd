package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Engligh when no language is supplied", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
