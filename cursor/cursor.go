package cursor

import termbox "github.com/nsf/termbox-go"

type Cursor struct {
	X int
	Y int
}

func NewCursor() *Cursor {
	return &Cursor{}
}

func (c *Cursor) Update() {
	termbox.SetCursor(c.X, c.Y)
}

func (c *Cursor) MoveRight() {
	c.X++
}

func (c *Cursor) MoveLeft() {
	if c.X > 0 {
		c.X--
	}
}

func (c *Cursor) MoveUp() {
	if c.Y > 0 {
		c.Y--
	}
}

func (c *Cursor) MoveDown() {
	c.Y++
}

func (c *Cursor) MoveHead() {
	c.X = 0
}

func (c *Cursor) MoveTop() {
	c.Y = 0
}

func (c *Cursor) MoveHorizontalByNum(num int) {
	c.X = num
}

func (c *Cursor) MoveVerticalByNum(num int) {
	c.Y = num
}
