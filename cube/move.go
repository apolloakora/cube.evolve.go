package cube

import (
	"math/rand"
	"time"
)

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

func RandomMoves(count int) ([]Move, []Move) {

	rand.Seed(time.Now().UnixNano())
	firstList := make([]Move, count)
	oppList := make([]Move, count)

	firstRand := Move(rand.Intn(len(opposites)))
	firstList[0], oppList[count-1] = firstRand, opposites[firstRand]
	moveAxis := moves[firstRand].axis

	for i := 1; i < count; i++ {
		nextArray := nextMoveMap[moveAxis]
		listMove := nextArray[rand.Intn(len(opposites)-3)]
		firstList[i] = listMove
		oppList[count-1-i] = opposites[listMove]
		moveAxis = moves[listMove].axis
	}

	return firstList, oppList

}

var opposites = map[Move]Move{
	Xy: Xz,
	Xz: Xy,
	Xn: Xn,
	Yx: Yz,
	Yz: Yx,
	Yn: Yn,
	Zx: Zy,
	Zy: Zx,
	Zn: Zn,
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
	X: {Yx, Yz, Yn, Zx, Zy, Zn},
	Y: {Xy, Xz, Xn, Zx, Zy, Zn},
	Z: {Xy, Xz, Xn, Yx, Yz, Yn},
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
