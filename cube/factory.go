package cube

import (
	"math/rand"
	"time"
)

func SolvedCube() Cube {

	cube := Cube{
		state:   &[28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: &[7]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
	}

	return cube
}

const ascii0, asciiX = 48, 120

var hopper = [...]int{1, 2, 2, 1, 2, 1, 1, 2}

func FromString(in string) Cube {

	cube := Cube{
		state:   &[28]int{},
		indices: &[...]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
	}
	byteStr := []byte(in)
	for i := range cube.indices {
		cube.state[i*4+0] = int(byteStr[i*2+0]) - ascii0
		cube.state[i*4+1] = int(byteStr[i*2+1]) - asciiX
		cube.state[i*4+2] = (cube.state[i*4+1] + hopper[cube.state[i*4+0]]) % 3
		cube.state[i*4+3] = 3 - (cube.state[i*4+1] + cube.state[i*4+2])
	}
	return cube
}

const minimumMoves, extraMoveRange = 234567, 123456

func RandomCube() Cube {

	cube := SolvedCube()
	rand.Seed(time.Now().UnixNano())
	moveSequence, _ := RandomMoves(minimumMoves + rand.Intn(extraMoveRange))
	for _, v := range moveSequence {
		cube.Revolve(v)
	}
	return cube

}
