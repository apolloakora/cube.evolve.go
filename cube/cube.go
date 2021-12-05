package cube

import (
	"fmt"
	"strconv"
	"strings"
)

type Cube struct {
	state   [28]int
	indices [7]int
}

func (c *Cube) displace(move Move) {

	for d := range displacements[move] {
		translation := directions[move][d]
		operand := operands[translation]
		result := c.state[c.indices[d]*4] & operand
		modifier := resultActions[translation][result]
		c.state[c.indices[d]*4] = c.state[c.indices[d]*4] + modifier
	}

}

func (c *Cube) orient(move Move) {

	if !moves[move].isQuarter {
		return
	}
	orientation := orientations[moves[move].axis]
	for d := range displacements[move] {
		offset := c.indices[d]*4 + 1
		c.state[offset+int(orientation[0])], c.state[offset+int(orientation[1])] = c.state[offset+int(orientation[1])], c.state[offset+int(orientation[0])]
	}
}

func (c *Cube) relocate(move Move) {

	c.shuffle(move, Xa, 0)

	/*	dataCache := make(map[int]int)

		for k, v := range swapIndices {
			dataCache[v] = dataList[k]
		}

		for p, q := range dataCache {
			dataList[p] = q
		}
	*/
}

func (c *Cube) shuffle(move Move, fromCell Cell, count int) {
	toCell := displacements[move][fromCell]
	if count < 4 {
		count++
		c.shuffle(move, toCell, count)
	}
	fmt.Println("count", count, "copying from cell", fromCell, "to cell", toCell)
	c.indices[int(toCell)] = c.indices[int(fromCell)]
}

func (c Cube) String() string {

	var strState = make([]string, 28)
	for i, v := range c.indices {
		strState[i*4] = strconv.Itoa(c.state[v*4])
		for j := range faces {
			strState[i*4+1+j] = faces[c.state[v*4+1+j]].name
		}
	}
	return strings.Join(strState, "")

}
