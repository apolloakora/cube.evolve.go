package cube

type Axis int

type face struct {
	name string
	axis Axis
}

var faces = [3]face{
	{"x", X},
	{"y", Y},
	{"z", Z},
}

var orientations = map[Axis][2]Axis{
	X: {Y, Z},
	Y: {X, Z},
	Z: {X, Y},
}

const (
	X Axis = iota
	Y
	Z
)
