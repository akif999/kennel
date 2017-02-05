package main

import (
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"os"
	"strconv"
)

func FeedNewline(c *Cursor) {
	if c.y < TEXTBUFROWMAX {
		c.x = TEXTBUFCOLMIN
		c.y++
	}
}

func BackSpace(c *Cursor) {
	if c.x > TEXTBUFCOLMIN {
		c.x -= 1
		termbox.SetCell(c.x, c.y, []rune(" ")[0], termbox.ColorBlue, termbox.ColorWhite)
	} else {
	}
}

func DeleteChar(x, y int) {
	termbox.SetCell(x, y, []rune(" ")[0], termbox.ColorBlue, termbox.ColorWhite)
}

func DisplayPositon(c *Cursor) {
	xpos := []rune(strconv.Itoa(c.x - 1))
	ypos := []rune(strconv.Itoa(c.y))
	for i := 2; i < 7; i++ {
		DeleteChar(i, 30)
	}
	for i, x := range xpos {
		termbox.SetCell(i+2, TEXTBUFROWMAX+2, x, termbox.ColorBlue, termbox.ColorWhite)
	}
	for i, y := range ypos {
		termbox.SetCell(i+5, TEXTBUFROWMAX+2, y, termbox.ColorBlue, termbox.ColorWhite)
	}
}

func SaveBufToFile(filename string, cells []termbox.Cell) {
	var data []byte
	for _, c := range cells {
		data = append(data, byte(c.Ch))
	}
	ioutil.WriteFile(filename, data, os.ModePerm)
}

func InitBuffer() {
	title := []rune(TITLE)
	i := 0
	for ; i < 42; i++ {
		termbox.SetCell(i, 0, '=', termbox.ColorBlue, termbox.ColorWhite)
	}
	for _, t := range title {
		termbox.SetCell(i, 0, t, termbox.ColorBlue, termbox.ColorWhite)
		i++
	}
	for ; i < 90; i++ {
		termbox.SetCell(i, 0, '=', termbox.ColorBlue, termbox.ColorWhite)
	}

	for i = 1; i < 31; i++ {
		termbox.SetCell(TEXTBUFCOLMIN-2, i, '|', termbox.ColorBlue, termbox.ColorWhite)
	}
	for i = 1; i < 31; i++ {
		termbox.SetCell(TEXTBUFCOLMAX+1, i, '|', termbox.ColorBlue, termbox.ColorWhite)
	}
	for i = 0; i < 90; i++ {
		termbox.SetCell(i, TEXTBUFROWMAX+1, '=', termbox.ColorBlue, termbox.ColorWhite)
	}
	for i = 0; i < 90; i++ {
		termbox.SetCell(i, TEXTBUFROWMAX+3, '=', termbox.ColorBlue, termbox.ColorWhite)
	}
}
