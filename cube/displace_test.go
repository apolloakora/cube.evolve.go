package cube

import "testing"

func TestDisplaceXy(t *testing.T) {
	c := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	c.displace(Xy)
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
	c.displace(Xz)
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
	c.displace(Xn)
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
	c.displace(Yx)
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
	c.displace(Yz)
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
	c.displace(Yn)
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
	c.displace(Zx)
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
	c.displace(Zy)
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
	c.displace(Zn)
	expectedState := "0xyz0xyz6xyz6xyz6xyz0xyz6xyz"
	if c.String() == expectedState {
		return
	}
	t.Error("Expected state", expectedState, "but got", c.String())
}
