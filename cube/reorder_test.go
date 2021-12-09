package cube

import "testing"

func TestReorder(t *testing.T) {

	// @todo we need to add the Y values
	// @todo we need to add the Z values
	// @todo we need to refactor to remove expansive integers
	var orderExpect = map[Move][7]Cell{
		Xy: [7]Cell{Zd, Ya, Za, Xd, Xa, Po, Yd},
		Xz: [7]Cell{Yd, Ya, Za, Xd, Po, Xa, Zd},
		Xn: [7]Cell{Po, Ya, Za, Xd, Zd, Yd, Xa},
	}

	for move, expectedOrder := range orderExpect {
		c := Cube{
			state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
			indices: [7]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
		}
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
