package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"strconv"
)

const (
	COLOFFSET = 2
	ROWOFFSET = 1
)

type Cursor struct {
	x int
	y int
}

func main() {
	c := new(Cursor)
	c.x = COLOFFSET
	c.y = ROWOFFSET

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorBlue, termbox.ColorWhite)
	termbox.SetCursor(c.x, c.y)
	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			termbox.SetCell(c.x, c.y, ev.Ch, termbox.ColorBlue, termbox.ColorWhite)
			c.x++
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyEnter:
				FeedNewline(c)
			}
			termbox.SetCursor(c.x, c.y)
			DisplayPositon(c)
		}
		termbox.Flush()
	}
}

func FeedNewline(c *Cursor) {
	c.x = 2
	c.y++
}

func BackSpace(c *Cursor) {
	termbox.SetCell(c.x, c.y, []rune(" ")[0], termbox.ColorBlue, termbox.ColorWhite)
}

func DeleteCell(x, y int) {
	termbox.SetCell(x, y, []rune(" ")[0], termbox.ColorBlue, termbox.ColorWhite)
}

func DisplayPositon(c *Cursor) {
	xpos := []rune(strconv.Itoa(c.x - 1))
	ypos := []rune(strconv.Itoa(c.y))
	for i := 2; i < 7; i++ {
		DeleteCell(i, 30)
	}
	for i, x := range xpos {
		termbox.SetCell(i+2, 30, x, termbox.ColorBlue, termbox.ColorWhite)
	}
	for i, y := range ypos {
		termbox.SetCell(i+5, 30, y, termbox.ColorBlue, termbox.ColorWhite)
	}
}
