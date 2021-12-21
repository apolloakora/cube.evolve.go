package solve

import (
	"cube.evolve.go/cube"
	"cube.evolve.go/queue"
	"fmt"
	"testing"
)

func TestSolveFromSolved(t *testing.T) {
	alreadySolvedCube := cube.SolvedCube()
	firstPlace := Place{
		cube: &alreadySolvedCube,
		path: &[]cube.Move{},
	}
	var queueItem queue.Item = firstPlace
	ourQueue := queue.NewQueue(&queueItem)
	foundItemPtr := ourQueue.Seek()
	fmt.Println("The queue seek operation returned", (*foundItemPtr).String())
	var solution = (*foundItemPtr).(Place)
	fmt.Println("Solution Path", solution.path)
}

func TestSolveFromRandom(t *testing.T) {
	//toSolveCube := cube.FromString("0zx5zx2yx1xz0yz0yz6zx" )
	// solution is Yx Xz Yn Zy Yx Zx Yx Xn Zx
	toSolveCube := cube.FromString("7zy3yz5zx4yx7xz0xy2zy")
	firstPlace := Place{
		cube: &toSolveCube,
		path: &[]cube.Move{},
	}
	var queueItem queue.Item = firstPlace
	ourQueue := queue.NewQueue(&queueItem)
	foundItemPtr := ourQueue.Seek()
	fmt.Println("The queue seek operation returned", (*foundItemPtr).String())
	var solution = (*foundItemPtr).(Place)
	fmt.Println("Solution Path", *solution.path)
}
