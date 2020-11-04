package iteration_test

import (
	"fmt"
	"testing"
	"github.com/nikhil-thomas/hacker-form-practice/03_golang/01_learn-go-with-tests/03_iteration/iteration"
)

func TestRepeat(t *testing.T) {
	n := 10
	repeated := iteration.Repeat("a", n)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B)  {
	for i:=0;i < b.N; i++{
		iteration.Repeat("a", 100)
	}
}

func ExampleRepeat() {
	rs := iteration.Repeat("a", 3)
	fmt.Println(rs)
	// Output: aaa
}

