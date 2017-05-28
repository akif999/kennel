package control

import ()

type Cursor struct {
	X int
	Y int
}

func NewCursor() *Cursor {
	return &Cursor{X: 0, Y: 0}
}

func (c *Cursor) MoveCursorToUpper() {
}

func (c *Cursor) MoveCursorToLower() {
}

func (c *Cursor) MoveCursorToLeft() {
	if c.X != 0 {
		c.X--
	}
}

func (c *Cursor) MoveCursorToRight() {
	c.X++
}

func (c *Cursor) CopyLineLength(lens []int) {
}
