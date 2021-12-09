package cube

import (
	"testing"
)

func TestMany(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}

	forward, back := RandomMoves(200000)
	initial := c.String()
	for _, v := range forward {
		c.revolve(v)
	}
	for _, v := range back {
		c.revolve(v)
	}

	if c.String() != initial {
		t.Error("expected", initial, "but received", c.String())
	}

}

func TestXRotations(t *testing.T) {

	c1 := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c2 := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c3 := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c4 := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c5 := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c1.revolve(Xy)
	c1.revolve(Xy)
	c2.revolve(Xz)
	c2.revolve(Xz)
	c3.revolve(Xn)

	if c1.String() != c2.String() {
		t.Error("Expected Xy and Xz rotated twice to match", c1, c2)
	}
	if c1.String() != c3.String() {
		t.Error("Expected Xy rotated twice vs Xn  to match", c1, c3)
	}
	c1.revolve(Xy)
	c4.revolve(Xz)
	if c1.String() != c4.String() {
		t.Error("Expected Xy rotated 3 times vs Xz  to match", c1, c4)
	}
	c2.revolve(Xz)
	c5.revolve(Xy)
	if c2.String() != c5.String() {
		t.Error("Expected 3 Xz rotations vs one Xy to match", c2, c5)
	}

}

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
