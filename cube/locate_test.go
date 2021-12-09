package cube

import "testing"

func TestRelocate(t *testing.T) {

	var relocateExpect = map[Move]string{
		Xy: "1xyz0xyz0xyz0xyz2xyz2xyz1xyz",
		Xz: "2xyz0xyz0xyz0xyz1xyz1xyz2xyz",
		Xn: "3xyz0xyz0xyz0xyz3xyz3xyz3xyz",
		Yx: "0xyz1xyz0xyz4xyz0xyz4xyz1xyz",
		Yz: "0xyz4xyz0xyz1xyz0xyz1xyz4xyz",
		Yn: "0xyz5xyz0xyz5xyz0xyz5xyz5xyz",
		Zx: "0xyz0xyz2xyz4xyz4xyz0xyz2xyz",
		Zy: "0xyz0xyz4xyz2xyz2xyz0xyz4xyz",
		Zn: "0xyz0xyz6xyz6xyz6xyz0xyz6xyz",
	}

	for move, expectation := range relocateExpect {
		c := SolvedCube()
		c.relocate(move)
		if c.String() == expectation {
			continue
		}
		t.Error("During relocation", move, "we expected state", expectation, "and not", c.String())
	}

}
