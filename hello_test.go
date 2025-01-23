package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to Yadnesh", func(t *testing.T) {
		got := Hello("Yadnesh", "")
		want := "Hello, Yadnesh"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Yadnesh", func(t *testing.T) {
		got := Hello("Yadnesh", "Spanish")
		want := "Hola, Yadnesh"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to empty string", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Yadnesh", func(t *testing.T) {
		got := Hello("Yadnesh", "French")
		want := "Bonjour, Yadnesh"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to empty string", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, the line number reported will be in our
	// function call rather than inside our test helper
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
