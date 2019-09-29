package user

import "github.com/akif999/kennel/buffer"

func LineFeed(b *buffer.Buffer) {
	b.LineFeed()
}

func BackSpace(b *buffer.Buffer) {
	b.BackSpace()
}

func InsertChr(b *buffer.Buffer, r rune) {
	b.InsertChr(r)
}

func MoveCursor(b *buffer.Buffer, d buffer.CursorDir) {
	b.MoveCursor(d)
}

func Undo(b *buffer.Buffer) {
	b.Undo()
}

func Redo(b *buffer.Buffer) {
	b.Redo()
}

func SaveAs(b *buffer.Buffer) {
	b.SaveAs()
}
