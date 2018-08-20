package main

import termbox "github.com/nsf/termbox-go"

type bufStack struct {
	bufs []*buffer
}

type buffer struct {
	cursor cursor
	lines  []*line
}

type cursor struct {
	x int
	y int
}

type line struct {
	text []rune
}

func (b *buffer) updateLines() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	for y, l := range b.lines {
		for x, r := range l.text {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (b *buffer) updateCursor() {
	termbox.SetCursor(b.cursor.x, b.cursor.y)
}

func (b *buffer) pushBufToUndoRedoBuffer() {
	tb := new(buffer)
	tb.cursor.x = b.cursor.x
	tb.cursor.y = b.cursor.y
	for i, l := range b.lines {
		tl := new(line)
		tb.lines = append(tb.lines, tl)
		tb.lines[i].text = l.text
	}
	undoBuf.bufs = append(undoBuf.bufs, tb)
}

func (l *line) insertChr(r rune, p int) {
	t := make([]rune, len(l.text), cap(l.text)+1)
	copy(t, l.text)
	l.text = append(t[:p+1], t[p:]...)
	l.text[p] = r
}
func (l *line) deleteChr(p int) {
	p = p - 1
	l.text = append(l.text[:p], l.text[p+1:]...)
}

func (l *line) split(pos int) ([]rune, []rune) {
	return l.text[:pos], l.text[pos:]
}

func (b *buffer) getTextOnCursorLine() []rune {
	return b.lines[b.cursor.y].text
}

func (b *buffer) numOfLines() int {
	return len(b.lines)
}

func (b *buffer) numOfColsOnCursor() int {
	return len(b.lines[b.cursor.y].text)
}

func (l *line) runenum() int {
	return len(l.text)
}
