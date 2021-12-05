package cube

import (
	"testing"
)

func TestString(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 1, 0, 2, 1, 2, 1, 0, 2, 3, 1, 2, 0, 4, 2, 0, 1, 5, 2, 1, 0, 6, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	expectedState := "0xyz1xzy2yxz3yzx4zxy5zyx6xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}
