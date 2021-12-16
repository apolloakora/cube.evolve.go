package cube

import "testing"

func TestFromStringSimple(t *testing.T) {
	cGreat := SolvedCube()
	cMaybe := FromString(cGreat.String())
	if !cMaybe.Solved() {
		t.Error("Expected our maybe cube to be solved but it was", cMaybe)
	}
}

func TestFromString(t *testing.T) {

	forwards, backwards := RandomMoves(123456)
	thisCube := SolvedCube()
	for _, move := range forwards {
		thisCube.Revolve(move)
	}
	thatCube := FromString(thisCube.String())
	for _, move := range backwards {
		thatCube.Revolve(move)
	}
	if !thatCube.Solved() {
		t.Error("Expected thatCube to be solved but it was", thatCube)
	}
}
