package main

import "github.com/akif999/kennel/buffer"

func lineFeed(b *buffer.Buffer) {
	b.LineFeed()
}

func backSpace(b *buffer.Buffer) {
	b.BackSpace()
}

func insertChr(b *buffer.Buffer, r rune) {
	b.InsertChr(r)
}

func moveCursor(b *buffer.Buffer, d buffer.CursorDir) {
	b.MoveCursor(d)
}

func undo(b *buffer.Buffer) {
	b.Undo()
}

func redo(b *buffer.Buffer) {
	b.Redo()
}

func saveAs(b *buffer.Buffer) {
	b.SaveAs()
}
