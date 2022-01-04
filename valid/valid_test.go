package valid

import (
	"cube.evolve.go/cube"
	"testing"
)

func TestCheckValid(t *testing.T) {

	c := cube.SolvedCube()
	c.Revolve(cube.Xy)
	c.Revolve(cube.Yz)
	c.Revolve(cube.Zx)

	validity := CheckValid(c)
	if !validity {
		t.Error("expected cube to be valid, but it wasn't", c)
	}

}

func TestCheckInvalid(t *testing.T) {

	var invalid = []string{"7xy7xy7xy7xy7xy7xy7xy", "0xy0xy0xy0xy0xy0xy1xy", "0xy7xy0xy0xy6xy0xy0xy"}

	for _, v := range invalid {
		invalidCube := cube.FromString(v)
		if CheckValid(invalidCube) {
			t.Error("cube expected to be invalid, but was valid. cube is", v)
		}
	}
}
