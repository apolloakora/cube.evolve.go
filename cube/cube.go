package cube

import (
	"strconv"
	"strings"
)

type Cube struct {
	state   [28]int
	indices [7]Cell
}

func (c *Cube) revolve(move Move) {
	c.relocate(move)
	c.reorient(move)
	c.reorder(move)
}

func (c *Cube) relocate(move Move) {

	for d := range relocations[move] {
		translation := directions[move][d]
		operand := operands[translation]
		result := c.state[c.indices[d]*4] & operand
		modifier := resultActions[translation][result]
		c.state[c.indices[d]*4] = c.state[c.indices[d]*4] + modifier
	}

}

func (c *Cube) reorient(move Move) {

	if !moves[move].isQuarter {
		return
	}
	orientation := orientations[moves[move].axis]
	for d := range relocations[move] {
		offset := int(c.indices[d])*4 + 1
		c.state[offset+int(orientation[0])], c.state[offset+int(orientation[1])] =
			c.state[offset+int(orientation[1])], c.state[offset+int(orientation[0])]
	}
}

func (c *Cube) reorder(move Move) {
	if moves[move].isQuarter {
		c.replace(move)
		return
	}
	quarterTurn := Move(moves[move].axis * 3)
	c.replace(quarterTurn)
	c.replace(quarterTurn)
}

func (c *Cube) replace(move Move) {
	startIndex := int(moves[move].axis)
	c.reshuffle(move, relocations[move][Cell(startIndex)], c.indices[startIndex], Cell(startIndex))
}

func (c *Cube) reshuffle(move Move, toCell Cell, toValue Cell, firstCell Cell) {
	if toCell != firstCell {
		c.reshuffle(move, relocations[move][toCell], c.indices[int(toCell)], firstCell)
	}
	// @todo will the below work - surely it should be c.indices[int(toCell)]
	c.indices[toCell] = toValue
}

func (c Cube) String() string {

	var strState = make([]string, 28)
	for i, v := range c.indices {
		strState[i*4] = strconv.Itoa(c.state[v*4])
		for j := range faces {
			strState[i*4+1+j] = faces[c.state[int(v)*4+1+j]].name
		}
	}
	return strings.Join(strState, "")

}
