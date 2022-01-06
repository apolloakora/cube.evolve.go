package cube

import "testing"

func TestRelocate(t *testing.T) {

	var relocateExpect = map[Move]string{
		Xy: "1x0x0x0x2x2x1x",
		Xz: "2x0x0x0x1x1x2x",
		Xn: "3x0x0x0x3x3x3x",
		Yx: "0x1x0x4x0x4x1x",
		Yz: "0x4x0x1x0x1x4x",
		Yn: "0x5x0x5x0x5x5x",
		Zx: "0x0x2x4x4x0x2x",
		Zy: "0x0x4x2x2x0x4x",
		Zn: "0x0x6x6x6x0x6x",
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
