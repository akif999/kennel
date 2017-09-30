package main

import (
	"github.com/akif999/kennel/buffer"
	"github.com/akif999/kennel/command"
	"github.com/akif999/kennel/page"
	"github.com/akif999/kennel/window"
	"github.com/nsf/termbox-go"
)

type App struct {
	Pages []*page.Page
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() error {
	a.Pages = append(a.Pages, page.NewPage())
	a.Pages[0].Windows = append(a.Pages[0].Windows, window.NewWindow())
	a.Pages[0].Windows[0].Buf = buffer.NewBuffer()
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
			a.Chr(c, 0, 0)
		}
		a.Draw()
		termbox.Flush()
	}
	return nil
}

func (a *App) Chr(c *command.CommandSet, PageNum, WindowNum int) {
	a.Pages[PageNum].Windows[WindowNum].Buf.Insert(c.Chr)
}

func (a *App) Draw() {
	termbox.SetCell(0, 0, a.Pages[0].Windows[0].Buf.Runes[0], termbox.ColorWhite, termbox.ColorBlack)
}
