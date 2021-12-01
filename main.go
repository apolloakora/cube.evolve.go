package main

import (
	"fmt"
	"strconv"
	"strings"
)

var movemap = map[string]map[int]int{

	"xy": {
		0: 4,
		4: 6,
		6: 5,
		5: 0,
	},

	"xz": {
		0: 5,
		5: 6,
		6: 4,
		4: 0,
	},

	"xn": {
		0: 6,
		4: 5,
		6: 0,
		5: 4,
	},

	"yx": {
		1: 3,
		3: 6,
		6: 5,
		5: 1,
	},

	"yz": {
		1: 5,
		5: 6,
		6: 3,
		3: 1,
	},

	"yn": {
		1: 6,
		6: 1,
		5: 3,
		3: 5,
	},

	"zx": {
		2: 3,
		3: 6,
		6: 4,
		4: 2,
	},

	"zy": {
		2: 4,
		4: 6,
		6: 3,
		3: 2,
	},

	"zn": {
		2: 6,
		6: 2,
		4: 3,
		3: 4,
	},
}

var translations = map[string]map[int]string{

	"xy": {
		0: "zOnly",
		4: "yOnly",
		6: "zOnly",
		5: "yOnly",
	},

	"xz": {
		0: "yOnly",
		5: "zOnly",
		6: "yOnly",
		4: "zOnly",
	},

	"xn": {
		0: "y_z",
		4: "y_z",
		6: "y_z",
		5: "y_z",
	},

	"yx": {
		1: "zOnly",
		3: "xOnly",
		6: "zOnly",
		5: "xOnly",
	},

	"yz": {
		1: "xOnly",
		5: "zOnly",
		6: "xOnly",
		3: "zOnly",
	},

	"yn": {
		1: "x_z",
		6: "x_z",
		5: "x_z",
		3: "x_z",
	},

	"zx": {
		2: "yOnly",
		3: "xOnly",
		6: "yOnly",
		4: "xOnly",
	},

	"zy": {
		2: "xOnly",
		4: "yOnly",
		6: "xOnly",
		3: "yOnly",
	},

	"zn": {
		2: "x_y",
		6: "x_y",
		4: "x_y",
		3: "x_y",
	},
}

var operands = map[string]int{

	"zOnly": 1,
	"yOnly": 2,
	"xOnly": 4,
	"x_y":   6,
	"y_z":   3,
	"x_z":   5,
}

var resultActions = map[string]map[int]int{

	"zOnly": {
		0: 1,
		1: -1,
	},

	"yOnly": {
		0: 2,
		2: -2,
	},

	"xOnly": {
		0: 4,
		4: -4,
	},

	"x_y": {
		0: 6,
		2: 2,
		4: -2,
		6: -6,
	},

	"y_z": {
		0: 3,
		1: 1,
		2: -1,
		3: -3,
	},

	"x_z": {
		0: 5,
		1: 3,
		4: -3,
		5: -5,
	},
}

var rotations = map[string][2]int{

	"xy": {1, 2},
	"xz": {1, 2},

	"yx": {0, 2},
	"yz": {0, 2},

	"zx": {0, 1},
	"zy": {0, 1},
}

var faceMap = map[int]string{
	0: "x",
	1: "y",
	2: "z",
}

var checkMap = map[string]string{
	"xy": "xz",
	"xz": "xy",
	"xn": "xn",
	"yx": "yz",
	"yz": "yx",
	"yn": "yn",
	"zx": "zy",
	"zy": "zx",
	"zn": "zn",
}

var builders = []struct {
	move      string
	resembles string
}{
	{"xn", "xn"},
	{"xy", "xz"},
	{"xn", "xy"},
	{"yn", "yn"},
	{"yx", "yz"},
	{"yn", "yx"},
	{"zn", "zn"},
	{"zx", "zy"},
	{"zn", "zx"},
}

var locators = []int{0, 1, 2, 3, 4, 5, 6}

var cube = []int{
	5, 0, 2, 1,
	5, 1, 2, 0,
	5, 0, 2, 1,
	0, 0, 2, 1,
	0, 1, 2, 0,
	4, 1, 0, 2,
	1, 2, 1, 0,
}

var faces = make([][]int, 7)

var reverses = [3]string{
	"xz",
	"yz",
	"zy",
}

func main() {

	//	myMove := "xy"

	//	cubeToString(cube, locators)
	//	moveCube(cube, myMove)
	//	cubeToString(cube, locators)

	ourMap := NextState(cube)
	for k, v := range ourMap {
		fmt.Println("the move is", k, "the cube is", v)
	}

}

func moveCube(cube []int, move string) {

	translate(movemap, cube, move, locators)
	rotate(locators, faces, move)
	swap(locators, movemap[move])

}

func swap(dataList []int, swapIndices map[int]int) {

	dataCache := make(map[int]int)

	for k, v := range swapIndices {
		dataCache[v] = dataList[k]
	}

	for p, q := range dataCache {
		dataList[p] = q
	}

}

func translate(movemap map[string]map[int]int, cubeNumbers []int, move string, locate []int) {

	var cellMap = movemap[move]
	for k := range cellMap {
		translation := translations[move][k]
		operand := operands[translation]
		result := cubeNumbers[locate[k]*4] & operand
		modifier := resultActions[translation][result]

		cubeNumbers[locate[k]*4] = cubeNumbers[locate[k]*4] + modifier

	}

}

func rotate(location []int, faces [][]int, move string) {

	for i := range faces {
		theSlice := cube[(i*4 + 1):(i*4 + 4)]
		faces[i] = theSlice
	}

	rotation, ok := rotations[move]
	if !ok {
		return
	}

	for k := range movemap[move] {
		xyz := faces[location[k]]
		xyz[rotation[0]], xyz[rotation[1]] = xyz[rotation[1]], xyz[rotation[0]]
	}
}

func cubeToString(cube []int, location []int) string {
	state := make([]string, 28)
	for i, v := range location {
		state[i*4] = strconv.Itoa(cube[v*4])
		for j := 0; j < 3; j++ {
			state[i*4+1+j] = faceMap[cube[v*4+1+j]]
		}
	}
	return strings.Join(state, "")
}

func NextState(cube []int) map[string]string {

	fmt.Println("original cube is          ", cubeToString(cube, locators))

	nextMap := make(map[string]string)

	for index, builder := range builders {
		moveCube(cube, builder.move)
		stringCube := cubeToString(cube, locators)
		nextMap[builder.resembles] = stringCube

		if (index+1)%3 == 0 {
			moveCube(cube, reverses[((index+1)/3)-1])
		}

	}

	return nextMap

}
