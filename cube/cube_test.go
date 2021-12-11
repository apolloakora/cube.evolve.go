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

	var indexArray = [3]int{0, 1, 2}

	c1 := SolvedCube()
	c2 := SolvedCube()
	c3 := SolvedCube()
	c4 := SolvedCube()
	c5 := SolvedCube()

	for i := range indexArray {
		c1.revolve(Move(i * 3))
		c1.revolve(Move(i * 3))
		c2.revolve(Move(i*3 + 1))
		c2.revolve(Move(i*3 + 1))
		c3.revolve(Move(i*3 + 2))

		if c1.String() != c2.String() {
			t.Error("Expected", Move(i*3), "and", Move(i*3+1), "rotated twice to match", c1, c2)
		}
		if c1.String() != c3.String() {
			t.Error("Expected", Move(i*3), "rotated twice vs", Move(i*3+2), "to match", c1, c3)
		}
		c1.revolve(Move(i * 3))
		c4.revolve(Move(i*3 + 1))
		if c1.String() != c4.String() {
			t.Error("Expected", Move(i*3), "rotated 3 times vs", Move(i*3+1), "to match", c1, c4)
		}
		c2.revolve(Move(i*3 + 1))
		c5.revolve(Move(i * 3))
		if c2.String() != c5.String() {
			t.Error("Expected 3", Move(i*3+1), "rotations vs one", Move(i*3), "to match", c2, c5)
		}
	}
	/*

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
	*/

}

func TestString(t *testing.T) {
	c := SolvedCube()
	expectedState := "0xyz0xyz0xyz0xyz0xyz0xyz0xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}
