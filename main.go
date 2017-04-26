package main

import (
	"./buffer"
	"github.com/nsf/termbox-go"
	"log"
)

const ()

type CurPos struct {
	x int
	y int
}

type WrtPos struct {
	x uint
	y uint
}

var (
	cp CurPos
	wp WrtPos
)

// TODO
// goroutinr前提の設計とする
// まずはbufferパッケージの機能から実装してリライトする
// 改行コードで改行する処理を実装する(CopyScrnBufToTermBoxBufの中でやる)

func main() {

	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorBlue, termbox.ColorWhite)

	termbox.SetCursor(cp.x, cp.y)
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
				sb.AddNewLine(wp.x)
				cp.x++
				wp.x++
				cp.x = 0
				cp.y++
			case termbox.KeyBackspace, termbox.KeyBackspace2:
			case termbox.KeyCtrlS:
			default:
				if ev.Ch != 0 {
					sb.WriteChrToSBuf(wp.x, ev.Ch)
					cp.x++
					wp.x++
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
