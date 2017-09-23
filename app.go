package main

import (
	"github.com/akif999/kennel/command"
	"github.com/nsf/termbox-go"
)

func Init() error {
	err := termbox.Init()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCursor(0, 0)
	termbox.Flush()
	return err
}

func End() {
	defer termbox.Close()
}

func Run() error {
mainloop:
	for {
		c, err := command.ParseKeyToCommand(termbox.PollEvent())
		if err != nil {
			return err
		}
		switch c {
		case command.QuitApp:
			break mainloop
		}
		termbox.Flush()
	}
	return nil
}
