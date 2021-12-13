package cube

import "testing"

func TestFromString(t *testing.T) {

	endState := "1xz1zx7yx1xz2xz0xz4yx"
	returnMoves := [...]Move{Xz, Yx, Zy, Yz, Xn}
	myCube := FromString(endState)
	for _, v := range returnMoves {
		myCube.revolve(v)
	}

	if !myCube.Solved() {
		t.Error("Starting from state", endState, "and revolving with", returnMoves, "we expected a solved cube but it was", myCube)
	}
}
