package solve

import (
	"cube.evolve.go/cube"
	"cube.evolve.go/queue"
)

type Place struct {
	cube *cube.Cube
	path *[]cube.Move
}

var walkers = []struct {
	routes [3]cube.Move
	mirror [3]cube.Move
	rewind cube.Move
}{
	{[3]cube.Move{cube.Xn, cube.Xy, cube.Xn}, [3]cube.Move{cube.Xn, cube.Xz, cube.Xy}, cube.Xz},
	{[3]cube.Move{cube.Yn, cube.Yx, cube.Yn}, [3]cube.Move{cube.Yn, cube.Yz, cube.Yx}, cube.Yz},
	{[3]cube.Move{cube.Zn, cube.Zx, cube.Zn}, [3]cube.Move{cube.Zn, cube.Zy, cube.Zx}, cube.Zy},
}

func (p Place) Search() ([]*queue.Item, bool, *queue.Item) {

	var searchAxes = cube.Axes
	if len(*p.path) > 0 {
		lastMoveAxis := cube.Moves[(*p.path)[len(*p.path)-1]].MoveAxis
		searchAxes = cube.AxisMirror[lastMoveAxis]
	}

	var itemPointers []*queue.Item
	for _, thisAxis := range searchAxes {
		walker := walkers[thisAxis]
		for index, nextMove := range walker.routes {

			p.cube.Revolve(nextMove)
			nextPath := make([]cube.Move, len(*p.path)+1)
			copy(nextPath, *p.path)
			nextPath[len(nextPath)-1] = walker.mirror[index]
			nextCube := *p.cube.Mirror()
			nextPlace := Place{
				cube: &nextCube,
				path: &nextPath,
			}
			var item queue.Item = nextPlace
			if nextCube.Solved() {
				return nil, true, &item
			}
			itemPointers = append(itemPointers, &item)
		}
		p.cube.Revolve(walker.rewind)
	}
	return itemPointers, false, nil
}

func (p Place) String() string {
	return p.cube.String()
}
