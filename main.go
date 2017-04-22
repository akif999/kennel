package main

import (
	"./buffer"
	"github.com/nsf/termbox-go"
	"log"
)

const ()

type CurPos struct {
	x uint
	y uint
}

type WrtPos struct {
	x uint
	y uint
}

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

	termbox.SetCursor(0, 0)
	termbox.Flush()

	sb := buffer.NewScenBuffer()

	var x uint
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyEnter:
				sb.AddNewLine(WrtPos.x)
				WrtPos.x++
			case termbox.KeyBackspace, termbox.KeyBackspace2:
			case termbox.KeyCtrlS:
			default:
				if ev.Ch != 0 {
					sb.WriteChrToSBuf(WrtPos.x, ev.Ch)
					WrtPos.x++
				}
			}
			termbox.SetCursor(0, 0)
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
	for i, r := range buf.Chr {
		termbox.SetCell(i, 0, r, termbox.ColorBlue, termbox.ColorWhite)
	}
}
