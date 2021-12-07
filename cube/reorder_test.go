package cube

import "testing"

func TestReorder(t *testing.T) {

	// @todo we need to add the Y values
	// @todo we need to add the Z values
	// @todo we need to refactor to remove expansive integers
	var orderExpect = map[Move][7]int{
		Xy: [7]int{5, 1, 2, 3, 0, 6, 4},
		Xz: [7]int{4, 1, 2, 3, 6, 0, 5},
		Xn: [7]int{6, 1, 2, 3, 5, 4, 0},
	}

	for move, expectedOrder := range orderExpect {
		c := Cube{
			state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
			indices: [7]int{0, 1, 2, 3, 4, 5, 6},
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
