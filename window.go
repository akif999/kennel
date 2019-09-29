package main

import (
	"strconv"

	"github.com/akif999/kennel/buffer"
	termbox "github.com/nsf/termbox-go"
)

type window struct {
	cursor buffer.Cursor
	lines  []*buffer.Line
}

func createWindow(b *buffer.Buffer) (*window, error) {
	w := new(window)
	w.copyBufToWindow(b, true)
	w.updateWindowLines(b)
	w.updateWindowCursor(b)
	return w, nil
}

func (w *window) updateWindowLines(b *buffer.Buffer) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	offset := buffer.GetDigit(b.NumOfLines())
	w.cursor.Offset = offset + 1
	for y, l := range w.lines {
		linenums := makeLineNum(y+1+b.ShowStartHeight, offset)
		t := append(linenums, l.Text...)
		for x, r := range t {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (w *window) updateWindowCursor(b *buffer.Buffer) {
	termbox.SetCursor(w.cursor.X+w.cursor.Offset-b.ShowStartWidth, w.cursor.Y-b.ShowStartHeight)
}

func (w *window) copyBufToWindow(b *buffer.Buffer, addLinenum bool) {
	w.lines = []*buffer.Line{}
	winWidth, winHeight := termbox.Size()
	for i := 0; i+b.ShowStartHeight < len(b.Lines); i++ {
		if i > winHeight-1 {
			break
		}
		w.lines = append(w.lines, &buffer.Line{})
		for j := 0; j+b.ShowStartWidth < len(b.Lines[i+b.ShowStartHeight].Text); j++ {
			if j+buffer.GetDigit(b.NumOfLines())+1 > winWidth-1 {
				break
			}
			w.lines[i].Text = append(w.lines[i].Text, b.Lines[i+b.ShowStartHeight].Text[j+b.ShowStartWidth])
		}
	}
	w.cursor.X = b.Cursor.X
	w.cursor.Y = b.Cursor.Y
	w.cursor.Offset = b.Cursor.Offset
}

func makeLineNum(num int, digit int) []rune {
	numstr := strconv.Itoa(num)
	lineNum := make([]rune, digit+1)
	for i := 0; i < len(lineNum); i++ {
		lineNum[i] = ' '
	}
	// get digit of argument num
	cdigit := buffer.GetDigit(num)
	for i, c := range numstr {
		lineNum[i+(digit-cdigit)] = c
	}
	return lineNum
}
