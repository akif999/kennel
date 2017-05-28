package window

import (
	"../buffer/"
	"github.com/nsf/termbox-go"
)

type Window struct {
	LineLengths []uint32
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
