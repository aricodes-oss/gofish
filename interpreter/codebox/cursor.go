package codebox

import (
	"gofish/interpreter/parser"
)

type Cursor struct {
	x, y      int
	direction Direction
}

func (c *Cursor) Mirror(tok parser.Token) {
	switch tok {
	case parser.MIRFORWARD:
		c.x, c.y = -c.y, -c.x
	case parser.MIRBACKWARD:
		c.x, c.y = c.y, c.x
	case parser.MIRVERTICAL:
		c.x, c.y = -c.x, c.y
	case parser.MIRHORIZONTAL:
		c.x, c.y = c.x, -c.y
	case parser.MIRREVERSE:
		c.x, c.y = -c.x, -c.y
	}
}
