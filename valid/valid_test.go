package valid

import (
	"cube.evolve.go/cube"
	"testing"
)

func TestCheckValid(t *testing.T) {

	c := cube.RandomCube()
	if !CheckValid(c) {
		t.Error("expected cube to be valid, but it wasn't", c)
	}

}

/*func TestCheckInvalid(t *testing.T) {

	var invalid = []string{"7xy7xy7xy7xy7xy7xy7xy", "0xy0xy0xy0xy0xy0xy1xy", "0xy7xy0xy0xy6xy0xy0xy"}

	for _, v := range invalid {
		invalidCube := cube.FromString(v)
		if CheckValid(invalidCube) {
			t.Error("cube expected to be invalid, but was valid. cube is", v)
		}
	}
}
*/
func TestCentreOfGravity(t *testing.T) {

	c := cube.RandomCube()
	if !CentreOfGravity(c) {
		t.Error("expected cube to be valid, but it wasn't", c.String())
	}

}

func TestCentreOfGravity2(t *testing.T) {

	var invalid = []string{"4x2z0x0x0x0x0x", "0x0x0x2y0x0x7z"}
	for _, v := range invalid {
		if CentreOfGravity(cube.FromString(v)) {
			t.Error("cube expected to be invalid, but was valid, cube is", v)
		}
	}
}

func TestSamePosition(t *testing.T) {

	c := cube.RandomCube()
	if !SamePosition(c) {
		t.Error("expected cube to be valid, but it wasn't", c.String())
	}

}

func TestSamePosition2(t *testing.T) {

	var invalid = []string{"2x4y0x0x0x0x0x", "0z0x7y0x0x0x1y"}
	for _, v := range invalid {
		if SamePosition(cube.FromString(v)) {
			t.Error("cube expected to be invalid, but it was valid, cube is", v)
		}
	}
}
