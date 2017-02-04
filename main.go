package main

import (
	"github.com/nsf/termbox-go"
	"log"
)

func main() {
	var x int = 1
	var y int = 1

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
			termbox.SetCell(x, y, ev.Ch, termbox.ColorWhite, termbox.ColorBlue)
			x++
		case termbox.EventResize:
		}
	}
}
