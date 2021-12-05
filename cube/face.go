package cube

type face struct {
	name string
	axis Axis
}

var faces = [3]face{
	{"x", X},
	{"y", Y},
	{"z", Z},
}
