package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("\ngot:  %d \nwant: %d", got, want)
	}
}

func ExampleAdd() {
	got := Add(3, 4)
	fmt.Println(got)
	// Output: 7
}
