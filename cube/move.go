package cube

type Move int

type movement struct {
	name      string
	axis      Axis
	isQuarter bool
}

// If this move rotated the X axis there are 6 possible
// next moves that spin the Y and Z layers. Similarly if
// this move spins the Y or indeed Z axes.
func (m Move) NextMoves() [6]Move {
	return nextMoveMap[moves[m].axis]
}

// Return the string representation of a Move
func (m Move) String() string {
	return moves[m].name
}

var moves = [9]movement{
	{"Xy", X, true},
	{"Xz", X, true},
	{"Xn", X, false},
	{"Yx", Y, true},
	{"Yz", Y, true},
	{"Yn", Y, false},
	{"Zx", Z, true},
	{"Zy", Z, true},
	{"Zn", Z, false},
}

var nextMoveMap = map[Axis][6]Move{
	X: [6]Move{Yx, Yz, Yn, Zx, Zy, Zn},
	Y: [6]Move{Xy, Xz, Xn, Zx, Zy, Zn},
	Z: [6]Move{Xy, Xz, Xn, Yx, Yz, Yn},
}

const (
	Xy Move = iota
	Xz
	Xn
	Yx
	Yz
	Yn
	Zx
	Zy
	Zn
)
