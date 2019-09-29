package main

import (
	"log"
	"os"

	"github.com/akif999/kennel/buffer"
	"github.com/akif999/kennel/user"
	"github.com/akif999/kennel/window"
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
	win, err := window.New(buf)
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
				user.LineFeed(buf)
			// mac delete-key is this
			case termbox.KeyCtrlH:
				fallthrough
			case termbox.KeyBackspace2:
				user.BackSpace(buf)
			case termbox.KeyArrowUp:
				user.MoveCursor(buf, buffer.Up)
			case termbox.KeyArrowDown:
				user.MoveCursor(buf, buffer.Down)
			case termbox.KeyArrowLeft:
				user.MoveCursor(buf, buffer.Left)
			case termbox.KeyArrowRight:
				user.MoveCursor(buf, buffer.Right)
			case termbox.KeyCtrlZ:
				user.Undo(buf)
			case termbox.KeyCtrlY:
				user.Redo(buf)
			case termbox.KeyCtrlS:
				user.SaveAs(buf)
			case termbox.KeyEsc:
				break mainloop
			default:
				// convert null charactor by space to space
				if ev.Ch == '\u0000' {
					user.InsertChr(buf, ' ')
				} else {
					user.InsertChr(buf, ev.Ch)
				}
			}
		}
		win.CopyBufToWindow(buf, true)
		win.UpdateWindowLines(buf)
		win.UpdateWindowCursor(buf)
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
