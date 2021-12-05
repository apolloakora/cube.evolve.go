package cube

import "testing"

var orientExpect = map[Move]string{
	Xy: "0xzy0xyz0xyz0xyz0xzy0xzy0xzy",
	Xz: "0xzy0xyz0xyz0xyz0xzy0xzy0xzy",
	Xn: "0xyz0xyz0xyz0xyz0xyz0xyz0xyz",
	Yx: "0xyz0zyx0xyz0zyx0xyz0zyx0zyx",
	Yz: "0xyz0zyx0xyz0zyx0xyz0zyx0zyx",
	Yn: "0xyz0xyz0xyz0xyz0xyz0xyz0xyz",
	Zx: "0xyz0xyz0yxz0yxz0yxz0xyz0yxz",
	Zy: "0xyz0xyz0yxz0yxz0yxz0xyz0yxz",
	Zn: "0xyz0xyz0xyz0xyz0xyz0xyz0xyz",
}

func TestOrientXy(t *testing.T) {

	for move, expectation := range orientExpect {
		c := Cube{
			state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
			indices: [7]int{0, 1, 2, 3, 4, 5, 6},
		}
		c.orient(move)
		if c.String() == expectation {
			continue
		}
		t.Error("During move", move, "we expected state", expectation, "and not", c.String())
	}

}

func TestRelocate(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Xy)
}
