package main

import (
	"./buffer"
	"github.com/nsf/termbox-go"
	"log"
)

const ()

var (
	cp Cursol
)

type Cursol struct {
	x int
	y int
}

// TODO
// goroutine前提の設計とする
// Backspaceの実装
// Undo Redoの実装
// Window関連の機能をWindowパッケージへ切り出す

func main() {

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorBlue, termbox.ColorWhite)

	termbox.SetCursor(0, 0)
	termbox.Flush()

	sb := buffer.NewScenBuffer()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyEnter:
				sb.AddNewLine()
				cp.x = 0
				cp.y++
			case termbox.KeyArrowLeft:
				cp.x--
				sb.Pos--
			case termbox.KeyArrowRight:
				cp.x++
				sb.Pos++
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				BackSpace(sb)
				cp.x--
			case termbox.KeyCtrlS:
			default:
				if ev.Ch != 0 {
					sb.WriteChrToSBuf(ev.Ch)
					cp.x++
				}
			}
			termbox.SetCursor(cp.x, cp.y)
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
		CopyScrnBufToTermBoxBuf(sb)
		termbox.Flush()
	}
}

// 以下、package化する前の試作用関数

// CopyScrnBufToTermBoxBufは、引数で渡した内部バッファを、termboxのbackground bufferへコピーする
func CopyScrnBufToTermBoxBuf(buf *buffer.ScrnBuffer) {
	x := 0
	y := 0
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

func BackSpace(buf *buffer.ScrnBuffer) {
	buf.DelChrFromSBuf()
}
