package main

import (
	"log"
	"os"

	"github.com/akif999/kennel/buffer"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	err := initTermbox()
	defer termbox.Close()
	if err != nil {
		log.Fatal(err)
	}
	buf, err := buffer.New()
	if err != nil {
		log.Fatal(err)
	}
	if filename == "" {
		buf.Lines = []*buffer.Line{&buffer.Line{[]rune{}}}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		buf.ReadFileToBuf(file)
	}
	buf.PushBufToUndoRedoBuffer()
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
				lineFeed(buf)
			// mac delete-key is this
			case termbox.KeyCtrlH:
				fallthrough
			case termbox.KeyBackspace2:
				backSpace(buf)
			case termbox.KeyArrowUp:
				moveCursor(buf, buffer.Up)
			case termbox.KeyArrowDown:
				moveCursor(buf, buffer.Down)
			case termbox.KeyArrowLeft:
				moveCursor(buf, buffer.Left)
			case termbox.KeyArrowRight:
				moveCursor(buf, buffer.Right)
			case termbox.KeyCtrlZ:
				undo(buf)
			case termbox.KeyCtrlY:
				redo(buf)
			case termbox.KeyCtrlS:
				saveAs(buf)
			case termbox.KeyEsc:
				break mainloop
			default:
				// convert null charactor by space to space
				if ev.Ch == '\u0000' {
					insertChr(buf, ' ')
				} else {
					insertChr(buf, ev.Ch)
				}
			}
		}
		win.copyBufToWindow(buf, true)
		win.updateWindowLines(buf)
		win.updateWindowCursor(buf)
		buf.PushBufToUndoRedoBuffer()
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
