package solve

import (
	"cube.evolve.go/cube"
	"cube.evolve.go/queue"
	"fmt"
	"reflect"
	"testing"
)

func TestSolveFromRandom(t *testing.T) {
	//toSolveCube := cube.FromString("0zx5zx2yx1xz0yz0yz6zx" )
	// solution is Yx Xz Yn Zy Yx Zx Yx Xn Zx
	//	toSolveCube := cube.FromString("7zy3yz5zx4yx7xz0xy2zy")
	//	toSolveCube := cube.FromString("0yz7yx0yz5xy6xy4xz0zx")
	//	toSolveCube := cube.FromString("0zx3yz2yx6zx7zy0xy0yz")
	//	toSolveCube := cube.FromString("3xy7zy2zy5xy4xz4zy3zx")
	//	toSolveCube := cube.FromString("0yz5zx0yz5yz0zx5yz5xy")
	//	startCube := cube.FromString("0yz0xy0xy0yz0zx0yz0zx")

	//	startCube := cube.FromString("0xy0xy0yz0xy0xy0yz0xy")
	//	startCube := cube.FromString("0xy0xy0xy0xy0xy0xy0xy")
	//	startCube := cube.FromString("2zy6yz0xy0xy0xy4zy0xy")
	//	startCube := cube.FromString("0xy0xy0xy4xz0zx0xy4yx")

	startCube := cube.FromString("0xy1yx0xy4xz0xy4xz1zy")
	firstPlace := Place{
		cube: &startCube,
		path: &[]cube.Move{},
	}
	fmt.Println("The starting cube id", startCube.String())
	var queueItem queue.Item = firstPlace
	ourQueue := queue.NewQueue(&queueItem)
	foundItemPtr := ourQueue.Seek()
	var solution = (*foundItemPtr).(Place)
	fmt.Println("Solution Path has", len(*solution.path), "moves", *solution.path)
}

// @todo switch this test back on
func offTestSolveKnownCube(t *testing.T) {
	knownCube := cube.FromString("0zx5zx2yx1xz0yz0yz6zx")
	knownSolution := []cube.Move{cube.Yx, cube.Xz, cube.Yn, cube.Zy, cube.Yx, cube.Zx, cube.Yx, cube.Xn, cube.Zx}
	firstPlace := Place{
		cube: &knownCube,
		path: &[]cube.Move{},
	}
	var queueItem queue.Item = firstPlace
	ourQueue := queue.NewQueue(&queueItem)
	var solution = (*ourQueue.Seek()).(Place)
	if !reflect.DeepEqual(knownSolution, *solution.path) {
		t.Error("Expected solution for", knownCube.String(), "to be", knownSolution, "but it was", *solution.path)
	}
}
