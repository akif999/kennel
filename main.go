package main

import (
	"./buffer"
	"github.com/nsf/termbox-go"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

const ()

var (
	debug = kingpin.Flag("debug", "Set debug mode").Bool()

	cu Cursor
)

type Cursor struct {
	x           int
	y           int
	LineLengths []int
}

func Debug() {
}

func main() {

	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	sb := buffer.NewScenBuffer()
	sb.LineLengths = append(sb.LineLengths, 0)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyEnter:
				LineFeed(sb)
			case termbox.KeyArrowUp:
				cu.MoveCursorToUpper()
			case termbox.KeyArrowDown:
				cu.MoveCursorToLower()
			case termbox.KeyArrowLeft:
				cu.MoveCursorToLeft()
			case termbox.KeyArrowRight:
				cu.MoveCursorToRight()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				sb.ConvertCursorToBufPos(cu.x, cu.y)
				BackSpace(sb, cu.y)
				cu.CopyLineLength(sb.LineLengths)
			case termbox.KeyCtrlS:
			default:
				if ev.Ch != 0 {
					sb.ConvertCursorToBufPos(cu.x, cu.y)
					sb.WriteChrToSBuf(ev.Ch, cu.y)
					cu.CopyLineLength(sb.LineLengths)
					cu.x++
				}
			}
			termbox.SetCursor(cu.x, cu.y)
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
		CopyScrnBufToTermBoxBuf(sb)
		termbox.Flush()
		if *debug {
			Debug()
		}
	}
}

func Init() error {
	err := termbox.Init()
	termbox.Clear(termbox.ColorBlue, termbox.ColorWhite)
	termbox.SetCursor(0, 0)
	termbox.Flush()
	return err
}

// CopyScrnBufToTermBoxBufは、引数で渡した内部バッファを、termboxのbackground bufferへコピーする
func CopyScrnBufToTermBoxBuf(buf *buffer.ScrnBuffer) {
	x, y := 0, 0
	termbox.Clear(termbox.ColorBlue, termbox.ColorWhite)
	for _, r := range buf.Chr {
		if r == rune('\n') {
			x = 0
			y++
		} else {
			termbox.SetCell(x, y, r, termbox.ColorBlue, termbox.ColorWhite)
			x++
		}
	}
}

func LineFeed(buf *buffer.ScrnBuffer) {
	buf.AddNewLine()
	cu.x = 0
	cu.y++
}

func BackSpace(buf *buffer.ScrnBuffer, row int) {
	buf.DelChrFromSBuf(row)
	if cu.x != 0 {
		cu.x--
	}
}

func (c *Cursor) MoveCursorToUpper() {
}

func (c *Cursor) MoveCursorToLower() {
}

func (c *Cursor) MoveCursorToLeft() {
	if c.x != 0 {
		c.x--
	}
}

func (c *Cursor) MoveCursorToRight() {
	c.x++
}

func (c *Cursor) CopyLineLength(lens []int) {
	c.LineLengths = lens
}
