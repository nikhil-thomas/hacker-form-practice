package array_slice_test

import (
	"testing"
	"github.com/nikhil-thomas/hacker-form-practice/03_golang/01_learn-go-with-tests/04_array_slice/array_slice"
)

func TestSumList(t *testing.T)  {
	numbers := [5]int{1,2,3,4,5}

	got := array_slice.SumList(numbers)
	want := 15

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
