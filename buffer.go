package main

import (
	"os"

	termbox "github.com/nsf/termbox-go"
)

const (
	Up cursorDir = iota
	Down
	Left
	Right
)

type bufStack struct {
	bufs []*buffer
}

type buffer struct {
	cursor          cursor
	lines           []*line
	showStartHeight int
	showStartWidth  int
}

type cursor struct {
	x      int
	y      int
	offset int
}

type cursorDir uint8

type line struct {
	text []rune
}

func createBuffer(filename string) (*buffer, error) {
	b := new(buffer)

	if filename == "" {
		b.lines = []*line{&line{[]rune{}}}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		b.readFileToBuf(file)
	}
	b.pushBufToUndoRedoBuffer()

	return b, nil
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

func splitLine(l *line, pos int) ([]rune, []rune) {
	return l.text[:pos], l.text[pos:]
}

func joinLine(l *line, con []rune) {
	l.text = append(l.text, con...)
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

func (b *buffer) cursorUp() {
	// guard of top of "rows"
	if b.cursor.y > 0 {
		b.cursor.y--
		// guard of end of "row"
		if b.cursor.x > b.numOfColsOnCursor() {
			b.cursor.x = b.numOfColsOnCursor()
		}
		if b.isCursorOutOfWindowTop() {
			b.showStartHeight--
		}
	}
}

func (b *buffer) cursorDown() {
	// guard of end of "rows"
	if b.cursor.y < b.numOfLines()-1 {
		b.cursor.y++
		// guard of end of "row"
		if b.cursor.x > b.numOfColsOnCursor() {
			b.cursor.x = b.numOfColsOnCursor()
		}
		if b.isCursorOutOfWindowBottom() {
			b.showStartHeight++
		}
	}
}
func (b *buffer) cursorLeft() {
	winWidth, _ := termbox.Size()
	if b.cursor.x > 0 {
		b.cursor.x--
		if b.isCursorOutOfWindowLeft() {
			b.showStartWidth--
		}
	} else {
		// guard of top of "rows"
		if b.cursor.y > 0 {
			b.cursor.y--
			b.cursor.x = b.numOfColsOnCursor()
			b.showStartHeight--
			b.showStartWidth = winWidth - getDigit(b.numOfLines()) + 1 + b.numOfColsOnCursor()
		}
	}
}

func (b *buffer) cursorRight() {
	if b.cursor.x < b.lines[b.cursor.y].runenum() {
		b.cursor.x++
		if b.isCursorOutOfWindowRight() {
			b.showStartWidth++
		}
	} else {
		// guard of end of "rows"
		if b.cursor.y < b.numOfLines()-1 {
			b.cursor.x = 0
			b.cursor.y++
			b.showStartHeight++
			b.showStartWidth = 0
		}
	}
}

func (b *buffer) isCursorOutOfWindowTop() bool {
	return b.cursor.y < b.showStartHeight
}

func (b *buffer) isCursorOutOfWindowBottom() bool {
	_, winHeight := termbox.Size()
	return b.cursor.y+1 > b.showStartHeight+winHeight
}

func (b *buffer) isCursorOutOfWindowLeft() bool {
	return b.cursor.x < b.showStartWidth
}

func (b *buffer) isCursorOutOfWindowRight() bool {
	winWidth, _ := termbox.Size()
	offset := getDigit(b.numOfLines()) + 1
	return b.cursor.x+1+offset > b.showStartWidth+winWidth
}
