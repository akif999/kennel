package main

import (
	"./buffer"
	"github.com/nsf/termbox-go"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"strconv"
)

const ()

var (
	debug = kingpin.Flag("debug", "Set debug mode").Default("false").Bool()

	cu Cursor
)

type Cursor struct {
	x           int
	y           int
	LineLengths []int
}

func Debug(buf *buffer.ScrnBuffer, row int) {
	xpos := []rune(strconv.Itoa(cu.x))
	ypos := []rune(strconv.Itoa(cu.y))
	llen := []rune(strconv.Itoa(cu.LineLengths[row]))
	bpos := []rune(strconv.Itoa(buf.Pos))
	for i, x := range xpos {
		termbox.SetCell(i, 25, x, termbox.ColorBlue, termbox.ColorWhite)
	}
	for i, y := range ypos {
		termbox.SetCell(i, 26, y, termbox.ColorBlue, termbox.ColorWhite)
	}
	for i, l := range llen {
		termbox.SetCell(i, 27, l, termbox.ColorBlue, termbox.ColorWhite)
	}
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
				cu.CopyLineLength(sb.LineLengths)
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
				cu.MoveCursorToLeft()
				cu.CopyLineLength(sb.LineLengths)
			case termbox.KeyCtrlS:
			default:
				if ev.Ch != 0 {
					sb.ConvertCursorToBufPos(cu.x, cu.y)
					sb.WriteChrToSBuf(ev.Ch, cu.y)
					cu.MoveCursorToRight()
					cu.CopyLineLength(sb.LineLengths)
				}
			}
			termbox.SetCursor(cu.x, cu.y)
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
		CopyScrnBufToTermBoxBuf(sb)
		if *debug == true {
			Debug(sb, cu.y)
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
