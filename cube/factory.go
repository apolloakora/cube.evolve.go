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

var fromStringMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"x": 0,
	"y": 1,
	"z": 2,
}

func FromString(stringState string) Cube {

	// @todo create the cube first to avoid array copy operation when assigning state to intState
	var intState [28]int
	charState := []rune(stringState)
	for i := 0; i < 7; i++ {
		// @todo remove ranging over axes because it is misleading
		// @todo use ASCII subtraction to achieve digit and 0-2 from x-z
		// @todo delete fromStringMap
		for index := range axes {
			intState[i*4+index] = fromStringMap[string(charState[i*3+index])]
		}
		intState[i*4+3] = 3 - (intState[i*4+1] + intState[i*4+2])
	}

	cube := Cube{
		state:   intState,
		indices: [7]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
	}

	return cube
}

const minimumMoves, extraMoveRange = 234567, 123456

func RandomCube() Cube {

	cube := SolvedCube()
	rand.Seed(time.Now().UnixNano())
	moveSequence, _ := RandomMoves(minimumMoves + rand.Intn(extraMoveRange))
	for _, v := range moveSequence {
		cube.revolve(v)
	}
	return cube

}
