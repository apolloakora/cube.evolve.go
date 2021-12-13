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

	var intState [28]int
	charState := []rune(stringState)
	for i := 0; i < 7; i++ {
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

/*func FromString(State string) Cube {
	intCube := Cube{
		state:   [28]int{},
		indices: [7]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
	}

	stringSlice := make([]string, 0)
	intSlice := make([]int, 0)

	chars := []rune(State)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		stringSlice = append(stringSlice, char)
	}
	fmt.Println(stringSlice)

	for i, v := range stringSlice {
		if i%3 == 0 && i != 0 {
			fmt.Println("the index right now is", i)
			intSlice = append(intSlice, 3-(fromStringMap[stringSlice[i-2]]+fromStringMap[stringSlice[i-1]]))
		} else if i == 20 {
			intSlice = append(intSlice, 3-(fromStringMap[stringSlice[i-1]]+fromStringMap[stringSlice[i]]))
		}

		intSlice = append(intSlice, fromStringMap[v])
		fmt.Println(intSlice)

	}

	fmt.Println(intSlice)

	for i, v := range intSlice {
		intCube.state[i] = v
	}
	fmt.Println(intCube.state)
	return intCube
}
*/
func RandomCube() Cube {

	cube := SolvedCube()
	rand.Seed(time.Now().UnixNano())
	moveSequence, _ := RandomMoves(minimumMoves + rand.Intn(extraMoveRange))
	for _, v := range moveSequence {
		cube.revolve(v)
	}
	return cube

}
