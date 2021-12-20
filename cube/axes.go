package cube

type Axis int

const (
	X Axis = iota
	Y
	Z
)

var xyz = map[Axis][2]Axis{
	X: {Y, Z},
	Y: {X, Z},
	Z: {X, Y},
}
