package window

import (
	"github.com/akif999/kennel/buffer"
	"github.com/akif999/kennel/cursor"
	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	lines  [][]rune
	Buf    *buffer.Buffer
	Cursor *cursor.Cursor
}

func NewWindow() *Window {
	b := buffer.NewBuffer()
	c := cursor.NewCursor()
	return &Window{Buf: b, Cursor: c}
}

func (w *Window) BufToLines() {
	var line []rune
	for _, r := range w.Buf.Runes {
		line = append(line, r)
		if r == '\n' {
			w.lines = append(w.lines, line)
		}
	}
}

func (w *Window) LineToBuf() {
	for _, line := range w.lines {
		for _, r := range line {
			w.Buf.Runes = append(w.Buf.Runes, r)
		}
	}
}

func (w *Window) Insert(chr rune) {
	w.Buf.Insert(chr)
	w.Cursor.MoveRight()
	if chr == '\n' {
		w.Cursor.MoveHead()
		w.Cursor.MoveDown()
	}
}

func (w *Window) Draw() {
	x, y := 0, 0
	for _, r := range w.Buf.Runes {
		termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		x++
		if r == '\n' {
			x = 0
			y++
		}
	}
}

func (w *Window) UpdateCursor() {
	w.Cursor.Update()
}
