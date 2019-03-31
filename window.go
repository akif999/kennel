package main

import (
	"strconv"

	termbox "github.com/nsf/termbox-go"
)

type window struct {
	cursor cursor
	lines  []*line
}

func createWindow(b *buffer) (*window, error) {
	w := new(window)
	w.copyBufToWindow(b, true)
	w.updateWindowLines(b)
	w.updateWindowCursor(b)
	return w, nil
}

func (w *window) updateWindowLines(b *buffer) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	offset := getDigit(b.numOfLines())
	w.cursor.offset = offset + 1
	for y, l := range w.lines {
		linenums := makeLineNum(y+1+b.showStartHeight, offset)
		t := append(linenums, l.text...)
		for x, r := range t {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (w *window) updateWindowCursor(b *buffer) {
	termbox.SetCursor(w.cursor.x+w.cursor.offset-b.showStartWidth, w.cursor.y-b.showStartHeight)
}

func (w *window) copyBufToWindow(b *buffer, addLinenum bool) {
	w.lines = []*line{}
	winWidth, winHeight := termbox.Size()
	for i := 0; i+b.showStartHeight < len(b.lines); i++ {
		if i > winHeight-1 {
			break
		}
		w.lines = append(w.lines, &line{})
		for j := 0; j+b.showStartWidth < len(b.lines[i+b.showStartHeight].text); j++ {
			debugPrint(j)
			if j+getDigit(b.numOfLines())+1 > winWidth-1 {
				break
			}
			w.lines[i].text = append(w.lines[i].text, b.lines[i+b.showStartHeight].text[j+b.showStartWidth])
		}
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
