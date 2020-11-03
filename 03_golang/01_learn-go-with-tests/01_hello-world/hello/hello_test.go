package hello_test

import (
	"fmt"
	"testing"

	"github.com/nikhil-thomas/hacker-form-practice/03_golang/01_learn-go-with-tests/01_hello-world/hello"
)

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		name := "Chris"
		language := "english"
		want := fmt.Sprintf("Hello, %s", name)
		got := hello.Hello(name, language)

		assertCorrectMessage(t, want, got)
	})
	t.Run("saying Hello, World when empty string is supplied", func(t *testing.T) {
		name := ""
		language := "english"
		want := fmt.Sprintf("Hello, World")
		got := hello.Hello(name, language)

		assertCorrectMessage(t, want, got)
	})
	t.Run("greeting in Spanish", func(t *testing.T) {
		name := "Chris"
		want := fmt.Sprintf("Hola, %s", name)
		got := hello.Hello(name, "spanish")

		assertCorrectMessage(t, want, got)
	})

	t.Run("greeting in French", func(t *testing.T) {
		name := "Chris"
		want := fmt.Sprintf("Bonjour, %s", name)
		got := hello.Hello(name, "french")

		assertCorrectMessage(t, want, got)
	})
}

func assertCorrectMessage(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

