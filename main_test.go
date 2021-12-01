package main

import (
	"testing"
)

var random = []int{3, 3, 3}

var originalCube = []int{
	5, 0, 2, 1,
	5, 1, 2, 0,
	5, 0, 2, 1,
	0, 0, 2, 1,
	0, 1, 2, 0,
	4, 1, 0, 2,
	1, 2, 1, 0,
}

var testMap = map[string]string{
	"xy": "6yzx5yzx5xzy0xzy4xyz0zxy2yxz",
	"xz": "1yxz5yzx5xzy0xzy3zxy7xyz5yzx",
	"xn": "2zyx5yzx5xzy0xzy7yxz3yzx6xzy",
	"yx": "5xzy0zxy5xzy4xzy0yzx0xyz4yzx",
	"yz": "5xzy1yzx5xzy5xyz0yzx1xzy5zxy",
	"yn": "5xzy4zyx5xzy1yxz0yzx5xzy0yzx",
	"zx": "5xzy5yzx4zyx7zxy3yzx4yxz4zxy",
	"zy": "5xzy5yzx2zxy5yzx1zxy4yxz2zyx",
	"zn": "5xzy5yzx7zyx6yzx6xzy4yxz3xzy",
}

func TestNextState(t *testing.T) {
	newMap := NextState(originalCube)

	if testMap["xy"] != newMap["xy"] {
		t.Error("testmap is", testMap["xy"])
		t.Error("newmap is", newMap["xy"])
	}

	//	for k := range testMap {
	//		if testMap[k] != newMap[k] {
	//			t.Error("test map is", testMap[k])
	//			t.Error("new map is", newMap[k])
	//			t.Error("the map given by the function is incorrect")
	//			for k, v := range newMap {
	//				t.Error("the move is", k, "the cube is", v)
}

//		}
//	}
