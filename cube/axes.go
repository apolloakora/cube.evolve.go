package cube

type Axis int

const (
	X Axis = iota
	Y
	Z
)

var AxisMirror = map[Axis][]Axis{
	X: {Y, Z},
	Y: {X, Z},
	Z: {X, Y},
}

var Axes = []Axis{X, Y, Z}
