package codebox

type Direction []int

var (
	Right = Direction{1, 0}
	Down  = Direction{0, 1}
	Left  = Direction{-1, 0}
	Up    = Direction{0, -1}
)

// A little shorthand
func (d Direction) X() int {
	return d[0]
}

func (d Direction) Y() int {
	return d[0]
}
