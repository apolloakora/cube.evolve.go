package cube

import (
	"math/rand"
	"time"
)

const minimumMoves, extraMoveRange = 234567, 123456

// @todo write a FromString() function which takes a 21 long string and returns a correctly initialized cube.

// @todo write a test function that hands over a string (of a cube that has been moved forward (from solved) by 10 moves)
// @todo you then run the received cube through a hardcoded 10 reverse move and assert that the resultant string is equal to the one produced by a solved cube

func SolvedCube() Cube {

	cube := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
	}

	return cube
}

func RandomCube() Cube {

	cube := SolvedCube()
	rand.Seed(time.Now().UnixNano())
	moveSequence, _ := RandomMoves(minimumMoves + rand.Intn(extraMoveRange))
	for _, v := range moveSequence {
		cube.revolve(v)
	}
	return cube

}
