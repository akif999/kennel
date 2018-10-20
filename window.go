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

func (b *buffer) updateWindowLines() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	offset := getDigit(b.numOfLines())
	b.cursor.offset = offset + 1
	for y, l := range b.lines {
		linenums := makeLineNum(y+1, offset)
		t := append(linenums, l.text...)
		for x, r := range t {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (b *buffer) updateWindowCursor() {
	termbox.SetCursor(b.cursor.x+b.cursor.offset, b.cursor.y)
}

func (w *window) updateWindowSize() {
}

func (w *window) copyBufToWindow(b *buffer, addLinenum bool) {
	// offset := getDigit(b.numOfLines()) + 1
	/*
		for i, l := range b.lines {
			linenum := makeLineNum(i+1, offset)
			w.lines[i].text = append(w.lines[i].text)
		}
	*/
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
