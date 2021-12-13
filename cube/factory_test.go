package cube

import "testing"

func TestFromString(t *testing.T) {

	endState := "0xy2zx"
	returnMoves := [...]Move{Xy, Zx}
	myCube := FromString(endState)
	for {

	}
	if !myCube.Solved() {
		t.Error("Starting from state", endState, "and revolving with", returnMoves, "we expected a solved cube but it was", myCube)
	}
}
