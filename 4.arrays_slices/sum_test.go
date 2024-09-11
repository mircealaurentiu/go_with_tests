package arraysslices

import (
	"reflect"
	"testing"
)

func Test_Sum1(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("\ngot:  %d \nwant: %d", got, want)
	}
}

func Test_SumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{9, 0})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot:  %d \nwant: %d", got, want)
	}
}
