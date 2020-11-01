package hello_test

import (
	"fmt"
	"testing"
	"github.com/nikhil-thomas/hacker-form-practice/03_golang/01_learn-go-with-tests/01_hello-world/hello"
)

func TestHello(t *testing.T) {
	name := "Chris"
	want := fmt.Sprintf("Hello, %s", name)
	got := hello.Hello(name)

	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}


