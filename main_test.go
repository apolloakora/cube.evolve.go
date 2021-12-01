package main

import (
	"testing"
)

var originalCube = []int{
	5, 1, 2, 0,
	5, 2, 0, 1,
	5, 1, 2, 0,
	0, 1, 2, 0,
	0, 2, 0, 1,
	4, 1, 0, 2,
	6, 2, 1, 0,
}

var testMap = map[string]string{
	"xy": "6yxz5zyx5yxz0yxz0zyx4yzx6zxy",
}

func TestNextState(t *testing.T) {

	for i := range faces {
		theSlice := originalCube[(i*4 + 1):(i*4 + 4)]
		faces[i] = theSlice
	}

	newMap := NextState(originalCube)

	if testMap["xy"] != newMap["xy"] {
		t.Error("testmap is", testMap["xy"])
		t.Error("newmap is", newMap["xy"])
	}

	//	for k := range testMap {
	//		if testMap[k] != newMap[k] {
	//			t.Error("test map is", testMap[k])
	//			t.Error("new map is", newMap[k])
	//		}

	//	}

}
