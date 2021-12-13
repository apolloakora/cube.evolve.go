package cube

import "testing"

func TestRelocate(t *testing.T) {

	var relocateExpect = map[Move]string{
		Xy: "1xy0xy0xy0xy2xy2xy1xy",
		Xz: "2xy0xy0xy0xy1xy1xy2xy",
		Xn: "3xy0xy0xy0xy3xy3xy3xy",
		Yx: "0xy1xy0xy4xy0xy4xy1xy",
		Yz: "0xy4xy0xy1xy0xy1xy4xy",
		Yn: "0xy5xy0xy5xy0xy5xy5xy",
		Zx: "0xy0xy2xy4xy4xy0xy2xy",
		Zy: "0xy0xy4xy2xy2xy0xy4xy",
		Zn: "0xy0xy6xy6xy6xy0xy6xy",
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
