package cube

import "testing"

func TestDisplaceXy(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Xy)
	expectedState := "1xyz0xyz0xyz0xyz2xyz2xyz1xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceXz(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Xz)
	expectedState := "2xyz0xyz0xyz0xyz1xyz1xyz2xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceXn(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Xn)
	expectedState := "3xyz0xyz0xyz0xyz3xyz3xyz3xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceYx(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Yx)
	expectedState := "0xyz1xyz0xyz4xyz0xyz4xyz1xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceYz(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Yz)
	expectedState := "0xyz4xyz0xyz1xyz0xyz1xyz4xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceYn(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Yn)
	expectedState := "0xyz5xyz0xyz5xyz0xyz5xyz5xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceZx(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Zx)
	expectedState := "0xyz0xyz2xyz4xyz4xyz0xyz2xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceZy(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Zy)
	expectedState := "0xyz0xyz4xyz2xyz2xyz0xyz4xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}

func TestDisplaceZn(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.relocate(Zn)
	expectedState := "0xyz0xyz6xyz6xyz6xyz0xyz6xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}
