package main

import (
	"github.com/akif999/kennel/command"
	"github.com/akif999/kennel/page"
	"github.com/nsf/termbox-go"
)

type App struct {
	pages []page.Page
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() error {
	err := termbox.Init()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCursor(0, 0)
	termbox.Flush()
	return err
}

func (a *App) End() {
	defer termbox.Close()
}

func (a *App) Run() error {
	c := command.NewCommandSet()
mainloop:
	for {
		err := c.Parse(termbox.PollEvent())
		if err != nil {
			return err
		}
		switch c.Cmd {
		case command.QuitApp:
			break mainloop
		case command.Chr:
		}
		termbox.Flush()
	}
	return nil
}

func (a *App) Chr() {

}
