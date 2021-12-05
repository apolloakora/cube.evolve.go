package cube

import (
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
