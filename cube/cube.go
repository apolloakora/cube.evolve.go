package cube

type Cube struct {
	state   *[28]int
	indices *[7]Cell
}

func (c *Cube) Revolve(move Move) {
	c.relocate(move)
	c.reorient(move)
	c.reorder(move)
}

func (c *Cube) relocate(move Move) {

	for d := range relocations[move] {
		translation := directions[move][d]
		result := c.state[c.indices[d]*4] & operands[translation]
		modifier := resultActions[translation][result]
		c.state[c.indices[d]*4] = c.state[c.indices[d]*4] + modifier
	}

}

func (c *Cube) reorient(move Move) {
	if !moves[move].isQuarter {
		return
	}
	orientation := xyz[moves[move].axis]
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
		c.reshuffle(move, relocations[move][toCell], c.indices[toCell], firstCell)
	}
	c.indices[toCell] = toValue
}

func (c *Cube) Solved() bool {
	for _, v := range c.indices {
		if c.state[v*4] != 0 || c.state[v*4+1] != 0 || c.state[v*4+2] != 1 || c.state[v*4+3] != 2 {
			return false
		}
	}
	return true
}

func (c *Cube) String() string {

	byteSeq := make([]byte, 21, 21)
	for i, v := range c.indices {
		byteSeq[i*3+0] = byte(ascii0 + c.state[v*4+0])
		byteSeq[i*3+1] = byte(asciiX + c.state[v*4+1])
		byteSeq[i*3+2] = byte(asciiX + c.state[v*4+2])
	}
	return string(byteSeq)
}
