package cube

type Axis int

var xyz = map[Axis][2]Axis{
	X: {Y, Z},
	Y: {X, Z},
	Z: {X, Y},
}

var axes = [3]string{"x", "y", "z"}

const (
	X Axis = iota
	Y
	Z
)
