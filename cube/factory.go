package cube

import (
	"math/rand"
	"time"
)

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

const minimumMoves, extraMoveRange = 234567, 123456
