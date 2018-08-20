package main

import (
	"fmt"
	"log"
	"os"

	termbox "github.com/nsf/termbox-go"
)

const (
	Up = iota
	Down
	Left
	Right
)

var (
	undoBuf = &bufStack{}
	redoBuf = &bufStack{}
)

func main() {
	filename := ""
	fmt.Print(len(os.Args))
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	err := startUp()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	buf := new(buffer)
	if filename == "" {
		buf.lines = []*line{&line{[]rune{}}}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		buf.readFileToBuf(file)
	}
	buf.updateLines()
	buf.updateCursor()
	buf.pushBufToUndoRedoBuffer()
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
				buf.writeBufToFile()
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
		buf.updateLines()
		buf.updateCursor()
		buf.pushBufToUndoRedoBuffer()
		termbox.Flush()
	}
}

func startUp() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCursor(0, 0)
	return nil
}
