package cube

type Cell int
type Direction int

// These displacements for each move are maps with
// the keys that are the source index of the moving cell
// and values that are the destination index.
var displacements = map[Move]map[Cell]Cell{

	Xy: {Xa: Yd, Yd: Po, Po: Zd, Zd: Xa},
	Xz: {Xa: Zd, Zd: Po, Po: Yd, Yd: Xa},
	Xn: {Xa: Po, Yd: Zd, Po: Xa, Zd: Yd},
	Yx: {Ya: Xd, Xd: Po, Po: Zd, Zd: Ya},
	Yz: {Ya: Zd, Zd: Po, Po: Xd, Xd: Ya},
	Yn: {Ya: Po, Po: Ya, Zd: Xd, Xd: Zd},
	Zx: {Za: Xd, Xd: Po, Po: Yd, Yd: Za},
	Zy: {Za: Yd, Yd: Po, Po: Xd, Xd: Za},
	Zn: {Za: Po, Po: Za, Yd: Xd, Xd: Yd},
}

var directions = map[Move]map[Cell]Direction{

	Xy: {Xa: ZOnly, Yd: YOnly, Po: ZOnly, Zd: YOnly},
	Xz: {Xa: YOnly, Zd: ZOnly, Po: YOnly, Yd: ZOnly},
	Xn: {Xa: YAndZ, Yd: YAndZ, Po: YAndZ, Zd: YAndZ},
	Yx: {Ya: ZOnly, Xd: XOnly, Po: ZOnly, Zd: XOnly},
	Yz: {Ya: XOnly, Zd: ZOnly, Po: XOnly, Xd: ZOnly},
	Yn: {Ya: XAndZ, Po: XAndZ, Zd: XAndZ, Xd: XAndZ},
	Zx: {Za: YOnly, Xd: XOnly, Po: YOnly, Yd: XOnly},
	Zy: {Za: XOnly, Yd: YOnly, Po: XOnly, Xd: YOnly},
	Zn: {Za: XAndY, Po: XAndY, Yd: XAndY, Xd: XAndY},
}

var operands = map[Direction]int{
	ZOnly: 1,
	YOnly: 2,
	XOnly: 4,
	XAndY: 6,
	YAndZ: 3,
	XAndZ: 5,
}

var resultActions = map[Direction]map[int]int{
	ZOnly: {0: 1, 1: -1},
	YOnly: {0: 2, 2: -2},
	XOnly: {0: 4, 4: -4},
	XAndY: {0: 6, 2: 2, 4: -2, 6: -6},
	YAndZ: {0: 3, 1: 1, 2: -1, 3: -3},
	XAndZ: {0: 5, 1: 3, 4: -3, 5: -5},
}

const (
	Xa Cell = iota
	Ya
	Za
	Xd
	Yd
	Zd
	Po
)

const (
	XAndY Direction = iota
	XAndZ
	XOnly
	YAndX
	YAndZ
	YOnly
	ZAndX
	ZAndY
	ZOnly
)
