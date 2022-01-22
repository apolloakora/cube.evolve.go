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

// If the previous move was X based the next move can only
// be Y or Z based. A similar symmetrical rule applies after
// a Y or Z rotation.
var next = map[Axis][6]Move{
	X: {Yx, Yz, Yn, Zx, Zy, Zn},
	Y: {Xy, Xz, Xn, Zx, Zy, Zn},
	Z: {Xy, Xz, Xn, Yx, Yz, Yn},
}

var reorient = map[Axis]map[Axis]map[uint8]Axis{
	Y: {
		X: {1: Y, 2: Y, 4: Y, 7: Y, 0: Z, 3: Z, 5: Z, 6: Z},
		Y: {1: Z, 2: Z, 4: Z, 7: Z, 0: X, 3: X, 5: X, 6: X},
		Z: {1: X, 2: X, 4: X, 7: X, 0: Y, 3: Y, 5: Y, 6: Y},
	},

	Z: {
		X: {1: Z, 2: Z, 4: Z, 7: Z, 0: Y, 3: Y, 5: Y, 6: Y},
		Y: {1: X, 2: X, 4: X, 7: X, 0: Z, 3: Z, 5: Z, 6: Z},
		Z: {1: Y, 2: Y, 4: Y, 7: Y, 0: X, 3: X, 5: X, 6: X},
	},
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
