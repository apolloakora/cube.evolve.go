package valid

import (
	"cube.evolve.go/cube"
	"fmt"
)

var CellOffset = map[cube.Cell]map[cube.Axis]int{
	cube.Xa: {cube.X: -1, cube.Y: 1, cube.Z: 1},
	cube.Ya: {cube.X: 1, cube.Y: -1, cube.Z: 1},
	cube.Za: {cube.X: 1, cube.Y: 1, cube.Z: -1},
	cube.Xd: {cube.X: 1, cube.Y: -1, cube.Z: -1},
	cube.Yd: {cube.X: -1, cube.Y: 1, cube.Z: -1},
	cube.Zd: {cube.X: -1, cube.Y: -1, cube.Z: 1},
	cube.Po: {cube.X: -1, cube.Y: -1, cube.Z: -1},
}

func CheckValid(c cube.Cube) bool {
	offsets := c.Offsets()
	fmt.Println(offsets)
	counters := [3]int{}
	for t, v := range offsets {
		for i, a := range cube.Axes {
			power := 1 << (2 - i)
			signal := power & v
			if signal == 0 {
				continue
			}
			counters[i] += CellOffset[cube.Cell(t)][a]
		}
	}

	for _, v := range counters {
		if v == 0 {
			continue
		}
		return false
	}
	return true
}

var mySlice = []int{4, 2, 1, 3, 5, 6, 7}

func CentreOfGravity(c cube.Cube) bool {
	offsets := c.Offsets()
	for i, v := range offsets {
		if mySlice[i] == v {
			return false
		}
	}
	return true
}

func SamePosition(c cube.Cube) bool {
	offsets := c.Offsets()
	tookPlaces := make([]bool, 7)
	for i, v := range offsets {
		xorValue := (mySlice[i] ^ v) - 1
		if tookPlaces[xorValue] {
			return false
		}
		tookPlaces[xorValue] = true
	}
	return true
}
