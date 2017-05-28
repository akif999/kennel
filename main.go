package main

import (
	"./buffer"
	"./control"
	"./window"
	"github.com/nsf/termbox-go"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"strconv"
)

const ()

var (
	debug = kingpin.Flag("debug", "Set debug mode").Short('d').Default("false").Bool()
)

func Debug(buf *buffer.ScrnBuffer, cu *control.Cursor, row int) {
	xpos := []rune(strconv.Itoa(cu.X))
	ypos := []rune(strconv.Itoa(cu.Y))
	// llen := []rune(strconv.Itoa(cu.LineLengths[row]))
	bpos := []rune(strconv.Itoa(buf.Pos))
	for i, x := range xpos {
		termbox.SetCell(i, 25, x, termbox.ColorBlue, termbox.ColorWhite)
	}
	for i, y := range ypos {
		termbox.SetCell(i, 26, y, termbox.ColorBlue, termbox.ColorWhite)
	}
	/*
		for i, l := range llen {
			termbox.SetCell(i, 27, l, termbox.ColorBlue, termbox.ColorWhite)
		}
	*/
	for i, p := range bpos {
		termbox.SetCell(i, 28, p, termbox.ColorBlue, termbox.ColorWhite)
	}
}

func main() {

	kingpin.Parse()

	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	sb := buffer.NewScenBuffer()
	cu := control.NewCursor()
	sb.LineLengths = append(sb.LineLengths, 0)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyEnter:
				LineFeed(sb, cu)
			case termbox.KeyArrowUp:
				cu.MoveCursorToUpper()
			case termbox.KeyArrowDown:
				cu.MoveCursorToLower()
			case termbox.KeyArrowLeft:
				cu.MoveCursorToLeft()
			case termbox.KeyArrowRight:
				cu.MoveCursorToRight()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				sb.ConvertCursorToBufPos(cu.Y, cu.X)
				BackSpace(sb, cu, cu.Y)
				cu.MoveCursorToLeft()
			case termbox.KeyCtrlS:
			default:
				if ev.Ch != 0 {
					sb.ConvertCursorToBufPos(cu.Y, cu.X)
					sb.WriteChrToSBuf(ev.Ch, cu.Y)
					cu.MoveCursorToRight()
				}
			}
			termbox.SetCursor(cu.X, cu.Y)
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
		window.CopyScrnBufToTermBoxBuf(sb)
		if *debug == true {
			Debug(sb, cu, cu.Y)
		}
		termbox.Flush()
	}
}

func Init() error {
	err := termbox.Init()
	termbox.Clear(termbox.ColorBlue, termbox.ColorWhite)
	termbox.SetCursor(0, 0)
	termbox.Flush()
	return err
}

func LineFeed(buf *buffer.ScrnBuffer, cu *control.Cursor) {
	buf.AddNewLine()
	cu.X = 0
	cu.Y++
}

func BackSpace(buf *buffer.ScrnBuffer, cu *control.Cursor, row int) {
	buf.DelChrFromSBuf(row)
	if cu.X != 0 {
		cu.X--
	}
}
