package cube

import (
	"math/rand"
	"time"
)

func (m Move) String() string {
	return Moves[m].moveName
}

func RandomMoves(count int) ([]Move, []Move) {

	rand.Seed(time.Now().UnixNano())
	firstList := make([]Move, count)
	oppList := make([]Move, count)

	firstRand := Move(rand.Intn(len(Moves)))
	firstList[0], oppList[count-1] = firstRand, reciprocals[firstRand]
	moveAxis := Moves[firstRand].MoveAxis

	for i := 1; i < count; i++ {
		nextArray := next[moveAxis]
		listMove := nextArray[rand.Intn(len(Moves)-3)]
		firstList[i] = listMove
		oppList[count-1-i] = reciprocals[listMove]
		moveAxis = Moves[listMove].MoveAxis
	}

	return firstList, oppList

}

type Move int

type Movement struct {
	moveName  string
	MoveAxis  Axis
	isQuarter bool
}

var reciprocals = map[Move]Move{
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

var Moves = [9]Movement{
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

var Walkers = []struct {
	Routes [3]Move
	Mirror [3]Move
	Rewind Move
}{
	{[3]Move{Xn, Xy, Xn}, [3]Move{Xn, Xz, Xy}, Xz},
	{[3]Move{Yn, Yx, Yn}, [3]Move{Yn, Yz, Yx}, Yz},
	{[3]Move{Zn, Zx, Zn}, [3]Move{Zn, Zy, Zx}, Zy},
}

// The 3 Y and 3 Z moves are the only meaningful proceeding
// move possibilities after any of the 3 X axis rotations.
// Similar rules apply after a Y or Z axis rotation.
var next = map[Axis][6]Move{
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
