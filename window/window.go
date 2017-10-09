package window

import (
	"github.com/akif999/kennel/buffer"
	"github.com/akif999/kennel/cursor"
	termbox "github.com/nsf/termbox-go"
)

const (
	InitialBufferSize = 4095
)

type Window struct {
	lines  [][]rune
	Buf    *buffer.Buffer
	Cursor *cursor.Cursor
}

type lines struct {
}

func NewWindow() *Window {
	buf := make([][]rune, InitialBufferSize)
	b := buffer.NewBuffer()
	c := cursor.NewCursor()
	return &Window{lines: buf, Buf: b, Cursor: c}
}

func (w *Window) LineToBuf() {
copyloop:
	for _, line := range w.lines {
		for _, r := range line {
			w.Buf.Runes = append(w.Buf.Runes, r)
			if r == 0 {
				break copyloop
			}
		}
		w.Buf.Runes = append(w.Buf.Runes, '\n')
	}
}

func (w *Window) Insert(chr rune) {
	line := w.lines[w.Cursor.Y]
	if len(line) > 1 {
		line = append(line[:w.Cursor.X+1], line[w.Cursor.X:]...)
		line[w.Cursor.X] = chr
	} else {
		line = append(line, chr)
	}
	w.Cursor.MoveRight()
	if chr == '\n' {
		w.Cursor.MoveHead()
		w.Cursor.MoveDown()
	}
}

func (w *Window) Draw() {
	w.LineToBuf()
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
