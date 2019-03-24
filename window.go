package main

import (
	"strconv"

	termbox "github.com/nsf/termbox-go"
)

type window struct {
	cursor cursor
	lines  []*line
	size   size
}

type size struct {
	x int
	y int
}

func createWindow(b *buffer) (*window, error) {
	w := new(window)
	w.copyBufToWindow(b, true)
	w.updateWindowLines(b)
	w.updateWindowCursor()
	return w, nil
}

func (w *window) updateWindowLines(b *buffer) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	offset := getDigit(b.numOfLines())
	w.cursor.offset = offset + 1
	for y, l := range w.lines {
		linenums := makeLineNum(y+1, offset)
		t := append(linenums, l.text...)
		for x, r := range t {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (w *window) updateWindowCursor() {
	termbox.SetCursor(w.cursor.x+w.cursor.offset, w.cursor.y)
}

func (w *window) updateWindowSize() {
}

func (w *window) copyBufToWindow(b *buffer, addLinenum bool) {
	w.lines = []*line{}
	for _, l := range b.lines {
		w.lines = append(w.lines, l)
	}
	w.cursor.x = b.cursor.x
	w.cursor.y = b.cursor.y
	w.cursor.offset = b.cursor.offset
}

func makeLineNum(num int, digit int) []rune {
	numstr := strconv.Itoa(num)
	lineNum := make([]rune, digit+1)
	for i := 0; i < len(lineNum); i++ {
		lineNum[i] = ' '
	}
	// get digit of argument num
	cdigit := getDigit(num)
	for i, c := range numstr {
		lineNum[i+(digit-cdigit)] = c
	}
	return lineNum
}

func getDigit(linenum int) int {
	d := 0
	for linenum != 0 {
		linenum = linenum / 10
		d++
	}
	return d
}
