package main

import (
	"github.com/nsf/termbox-go"
	"log"
)

const (
	COLOFFSET     = 2
	ROWOFFSET     = 1
	TEXTBUFROWMAX = 28
	TEXTBUFCOLMAX = 88
	TEXTBUFCOLMIN = 2
	TITLE         = "KENNEL"
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
	InitBuffer()
	DisplayPositon(c)
	termbox.SetCursor(c.x, c.y)
	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyEnter:
				FeedNewline(c)
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				BackSpace(c)
			case termbox.KeyCtrlS:
				SaveBufToFile("hoge.txt", termbox.CellBuffer())
			default:
				if ev.Ch != 0 {
					if c.x < TEXTBUFCOLMAX {
						termbox.SetCell(c.x, c.y, ev.Ch, termbox.ColorBlue, termbox.ColorWhite)
						c.x++
					}
				}
			}
			termbox.SetCursor(c.x, c.y)
			DisplayPositon(c)
		}
		termbox.Flush()
	}
}
