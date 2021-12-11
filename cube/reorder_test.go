package cube

import "testing"

func TestReorder(t *testing.T) {

	var orderExpect = map[Move][7]Cell{
		Xy: {Zd, Ya, Za, Xd, Xa, Po, Yd},
		Xz: {Yd, Ya, Za, Xd, Po, Xa, Zd},
		Xn: {Po, Ya, Za, Xd, Zd, Yd, Xa},
		Yx: {Xa, Zd, Za, Ya, Yd, Po, Xd},
		Yz: {Xa, Xd, Za, Po, Yd, Ya, Zd},
		Yn: {Xa, Po, Za, Zd, Yd, Xd, Ya},
		Zx: {Xa, Ya, Yd, Za, Po, Zd, Xd},
		Zy: {Xa, Ya, Xd, Po, Za, Zd, Yd},
		Zn: {Xa, Ya, Po, Yd, Xd, Zd, Za},
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
