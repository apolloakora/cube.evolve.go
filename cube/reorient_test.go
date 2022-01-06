package cube

import "testing"

func TestOrient(t *testing.T) {

	var orientExpect = map[Move]string{
		Xy: "0x0x0x0x0x0x0x",
		Xz: "0x0x0x0x0x0x0x",
		Xn: "0x0x0x0x0x0x0x",
		Yx: "0x0z0x0z0x0z0z",
		Yz: "0x0z0x0z0x0z0z",
		Yn: "0x0x0x0x0x0x0x",
		Zx: "0x0x0y0y0y0x0y",
		Zy: "0x0x0y0y0y0x0y",
		Zn: "0x0x0x0x0x0x0x",
	}

	for move, expectation := range orientExpect {
		c := SolvedCube()
		c.reorient(move)
		if c.String() == expectation {
			continue
		}
		t.Error("During move", move, "we expected state", expectation, "and not", c.String())
	}

}
