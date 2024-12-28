package integer

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("expected %d but got %d", expect, sum)
	}
}
