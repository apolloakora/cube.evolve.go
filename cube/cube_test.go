package cube

import (
	"testing"
)

func TestRevolutionsFromSolved(t *testing.T) {
	c := SolvedCube()
	forward, back := RandomMoves(234567)
	initial := c.String()
	for _, v := range forward {
		c.revolve(v)
	}
	for _, v := range back {
		c.revolve(v)
	}
	if c.String() != initial {
		t.Error("Expected solved state after thousands of moves", initial, "but got", c.String())
	}
}

func TestRevolutionsFromRandom(t *testing.T) {

	c := RandomCube()
	forward, back := RandomMoves(234567)
	initial := c.String()
	for _, v := range forward {
		c.revolve(v)
	}
	for _, v := range back {
		c.revolve(v)
	}
	if c.String() != initial {
		t.Error("Expected starting state after thousands of moves", initial, "but got", c.String())
	}
}

// @todo expand test to cover Y and Z scenarios
func TestXRotations(t *testing.T) {

	c1 := SolvedCube()
	c2 := SolvedCube()
	c3 := SolvedCube()
	c4 := SolvedCube()
	c5 := SolvedCube()
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
	c := SolvedCube()
	expectedState := "0xyz0xyz0xyz0xyz0xyz0xyz0xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}
