package main

import (
	"log"
	"os"

	termbox "github.com/nsf/termbox-go"
)

var (
	undoBuf = &bufStack{}
	redoBuf = &bufStack{}
)

func main() {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	err := initTermbox()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	buf, err := createBuffer(filename)
	if err != nil {
		log.Fatal(err)
	}
	win, err := createWindow(buf)
	if err != nil {
		log.Fatal(err)
	}

	termbox.Flush()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				buf.lineFeed()
			// mac delete-key is this
			case termbox.KeyCtrlH:
				fallthrough
			case termbox.KeyBackspace2:
				buf.backSpace()
			case termbox.KeyArrowUp:
				buf.moveCursor(Up)
			case termbox.KeyArrowDown:
				buf.moveCursor(Down)
			case termbox.KeyArrowLeft:
				buf.moveCursor(Left)
			case termbox.KeyArrowRight:
				buf.moveCursor(Right)
			case termbox.KeyCtrlZ:
				buf.undo()
			case termbox.KeyCtrlY:
				buf.redo()
			case termbox.KeyCtrlS:
				buf.saveAs()
			case termbox.KeyEsc:
				break mainloop
			default:
				// convert null charactor by space to space
				if ev.Ch == '\u0000' {
					buf.insertChr(' ')
				} else {
					buf.insertChr(ev.Ch)
				}
			}
		}
		win.copyBufToWindow(buf, true)
		win.updateWindowLines(buf)
		win.updateWindowCursor()
		buf.pushBufToUndoRedoBuffer()
		termbox.Flush()
	}
}

func initTermbox() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCursor(0, 0)
	return nil
}
