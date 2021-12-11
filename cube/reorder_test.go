package cube

import "testing"

func TestReorder(t *testing.T) {

	// @todo we need to add the Y values
	// @todo we need to add the Z values
	var orderExpect = map[Move][7]Cell{
		Xy: {Zd, Ya, Za, Xd, Xa, Po, Yd},
		Xz: {Yd, Ya, Za, Xd, Po, Xa, Zd},
		Xn: {Po, Ya, Za, Xd, Zd, Yd, Xa},
	}

	for move, expectedOrder := range orderExpect {
		c := SolvedCube()
		c.reorder(move)
		for i, v := range expectedOrder {
			if c.indices[i] == v {
				continue
			}
			t.Error(
				"Reorder move", move,
				"expected", v,
				"at index", i,
				"not", c.indices[i],
				"Expected", orderExpect[move],
				"not", c.indices,
			)
		}
	}
}
