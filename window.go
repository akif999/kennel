package main

import termbox "github.com/nsf/termbox-go"

func (b *buffer) updateWindowLines() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	for y, l := range b.lines {
		for x, r := range l.text {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (b *buffer) updateWindowCursor() {
	termbox.SetCursor(b.cursor.x, b.cursor.y)
}
