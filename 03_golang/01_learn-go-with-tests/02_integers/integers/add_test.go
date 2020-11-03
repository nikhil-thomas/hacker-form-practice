package integers_test

import (
	"fmt"
	"testing"
"github.com/nikhil-thomas/hacker-form-practice/03_golang/01_learn-go-with-tests/02_integers/integers"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2,2)
	expected := 4

	if sum != expected {
		t.Errorf("expected: %d, got: %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := integers.Add(1,5)
	fmt.Println(sum)
	// Output: 6
}