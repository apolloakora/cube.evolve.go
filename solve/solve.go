package solve

import (
	"cube.evolve.go/cube"
	"cube.evolve.go/queue"
)

type Place struct {
	cube *cube.Cube
	path *[]cube.Move
}

func (p Place) Search() ([]*queue.Item, bool, *queue.Item) {

	var searchAxes = cube.Axes
	if len(*p.path) > 0 {
		lastMoveAxis := cube.Moves[(*p.path)[len(*p.path)-1]].MoveAxis
		searchAxes = cube.AxisMirror[lastMoveAxis]
	}

	var itemPointers []*queue.Item
	for _, thisAxis := range searchAxes {
		walker := cube.Walkers[thisAxis]
		for index, nextMove := range walker.Routes {

			p.cube.Revolve(nextMove)
			nextPath := make([]cube.Move, len(*p.path)+1)
			copy(nextPath, *p.path)
			nextPath[len(nextPath)-1] = walker.Mirror[index]
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
		p.cube.Revolve(walker.Rewind)
	}
	return itemPointers, false, nil
}

func (p Place) String() string {
	return p.cube.String()
}

// States will return all the valid cube states
func States() map[string]bool {
	startCube := cube.FromString("0y0y0y0y0y0y0y")
	firstPlace := Place{
		cube: &startCube,
		path: &[]cube.Move{},
	}
	var queueItem queue.Item = firstPlace
	ourQueue := queue.NewQueue(&queueItem)
	return ourQueue.Results()
}
