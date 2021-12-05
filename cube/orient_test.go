package cube

import "testing"

func TestOrientXy(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.orient(Xy)
	expectedState := "0xzy0xyz0xyz0xyz0xzy0xzy0xzy"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}
