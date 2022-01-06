package cube

import (
	"reflect"
	"testing"
)

func TestSolved(t *testing.T) {
	c := SolvedCube()
	if !c.Solved() {
		t.Error("Expected solved cube", c, "to be solved but it was", c)
	}

	boomerang1 := [...]Move{Xy, Yx, Zx, Zy, Yz, Xz}
	for _, v := range boomerang1 {
		c.Revolve(v)
	}
	if !c.Solved() {
		t.Error("Expected a solved cube", c, "after this move sequence.", boomerang1)
	}

	boomerang2 := [...]Move{Xn, Yn, Zn, Zn, Yn, Xn}
	for _, v := range boomerang2 {
		c.Revolve(v)
	}
	if !c.Solved() {
		t.Error("Expected a solved cube", c, "after this move sequence.", boomerang2)
	}
}

func TestRevolutionsFromSolved(t *testing.T) {
	c := SolvedCube()
	forward, back := RandomMoves(234567)
	initial := c.String()
	for _, v := range forward {
		c.Revolve(v)
	}
	for _, v := range back {
		c.Revolve(v)
	}
	if c.String() != initial {
		t.Error("Expected solved state after thousands of Moves", initial, "but got", c.String())
	}
}

func TestRevolutionsFromRandom(t *testing.T) {

	c := RandomCube()
	forward, back := RandomMoves(123456)
	initial := c.String()
	for _, v := range forward {
		c.Revolve(v)
	}
	for _, v := range back {
		c.Revolve(v)
	}
	if c.String() != initial {
		t.Error("Expected starting state after thousands of Moves", initial, "but got", c.String())
	}
}

func TestRotations(t *testing.T) {

	for i := 0; i < 3; i++ {

		c1 := SolvedCube()
		c2 := SolvedCube()
		c3 := SolvedCube()
		c4 := SolvedCube()
		c5 := SolvedCube()

		c1.Revolve(Move(i * 3))
		c1.Revolve(Move(i * 3))
		c2.Revolve(Move(i*3 + 1))
		c2.Revolve(Move(i*3 + 1))
		c3.Revolve(Move(i*3 + 2))

		if c1.String() != c2.String() {
			t.Error("Expected", Move(i*3), "and", Move(i*3+1), "rotated twice to match", c1, c2)
		}
		if c1.String() != c3.String() {
			t.Error("Expected", Move(i*3), "rotated twice vs", Move(i*3+2), "to match", c1, c3)
		}
		c1.Revolve(Move(i * 3))
		c4.Revolve(Move(i*3 + 1))
		if c1.String() != c4.String() {
			t.Error("Expected", Move(i*3), "rotated 3 times vs", Move(i*3+1), "to match", c1, c4)
		}
		c2.Revolve(Move(i*3 + 1))
		c5.Revolve(Move(i * 3))
		if c2.String() != c5.String() {
			t.Error("Expected 3", Move(i*3+1), "rotations vs one", Move(i*3), "to match", c2, c5)
		}
	}
}

func TestString(t *testing.T) {
	c := SolvedCube()
	expectedState := "0x0x0x0x0x0x0x"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestOffsets(t *testing.T) {

	c := SolvedCube()
	for _, v := range c.Offsets() {
		if v != 0 {
			t.Error("Expected zero offeets but got", c.Offsets())
		}
	}

	c.Revolve(Xy)
	c.Revolve(Yz)
	c.Revolve(Zx)

	expect := []int{2, 1, 5, 2, 2, 4, 2}
	if !reflect.DeepEqual(c.Offsets(), expect) {
		t.Error("Expected cube offsets to be", expect, "but they were", c.Offsets())
	}
}

func TestMirror(t *testing.T) {
	src := SolvedCube()
	dst := src.Mirror()
	if !dst.Solved() {
		t.Error("Expected destination to be solved", dst, "like the source", src)
	}
	dst.Revolve(Xy)
	if dst.String() == src.String() {
		t.Error("Expected destination cube", dst, "to differ from source", src)
	}
	src.Revolve(Xn)
	dst.Revolve(Xy)
	if dst.String() != src.String() {
		t.Error("Expected destination cube", dst, "to mirror the source", src)
	}
}
