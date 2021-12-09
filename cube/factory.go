package cube

func SolvedCube() Cube {

	cube := Cube{
		state:   [28]int{0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2, 0, 0, 1, 2},
		indices: [7]Cell{Xa, Ya, Za, Xd, Yd, Zd, Po},
	}

	return cube
}
