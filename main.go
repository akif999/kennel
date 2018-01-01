package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

type buffer struct {
	lines []line
}

type line struct {
	text []rune
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCursor(0, 0)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			default:
				termbox.SetCell(0, 0, ev.Ch, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
		termbox.Flush()
	}
}
