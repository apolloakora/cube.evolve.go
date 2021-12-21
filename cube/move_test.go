package cube

import (
	"testing"
)

func TestRandomMovesLength(t *testing.T) {
	const quantity = 12345
	outMoves, inMoves := RandomMoves(quantity)
	if len(outMoves) != quantity || len(inMoves) != quantity {
		t.Error("Expected lengths to be", quantity, "but outMoves is", len(outMoves), "and inMoves is", len(inMoves))
	}
}

const quantity, minimum, maximum = 1234567, 76543, 234567

func TestRandomOutboundMoves(t *testing.T) {
	outboundMoves, _ := RandomMoves(quantity)
	outcomes := map[Move]int{
		Xy: 0, Xz: 0, Xn: 0,
		Yx: 0, Yz: 0, Yn: 0,
		Zx: 0, Zy: 0, Zn: 0,
	}
	for _, move := range outboundMoves {
		outcomes[move] = outcomes[move] + 1
	}
	for k, v := range outcomes {
		if v < minimum {
			t.Error("Expected count of", k, "outbound Moves to be at least", minimum, "but it was", v)
		}
		if v > maximum {
			t.Error("Expected count of", k, "outbound Moves not to exceed", maximum, "but it was", v)
		}
	}
	if len(outboundMoves) != quantity {
		t.Error("Expected outbound move counts to be", quantity, "not", len(outboundMoves))
	}
}

func TestRandomInboundMoves(t *testing.T) {
	_, inboundMoves := RandomMoves(quantity)
	outcomes := map[Move]int{
		Xy: 0, Xz: 0, Xn: 0,
		Yx: 0, Yz: 0, Yn: 0,
		Zx: 0, Zy: 0, Zn: 0,
	}
	for _, move := range inboundMoves {
		outcomes[move] = outcomes[move] + 1
	}
	for k, v := range outcomes {
		if v < minimum {
			t.Error("Expected count of", k, "inbound Moves to be at least", minimum, "but it was", v)
		}
		if v > maximum {
			t.Error("Expected count of", k, "inbound Moves not to exceed", maximum, "but it was", v)
		}
	}
	if len(inboundMoves) != quantity {
		t.Error("Expected inbound move counts to be", quantity, "not", len(inboundMoves))
	}
}
