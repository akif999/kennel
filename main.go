package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

type buffer struct {
	cursor cursor
	lines  []*line
}

type cursor struct {
	x int
	y int
}

type line struct {
	text []rune
}

func main() {
	err := startUp()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	buf := &buffer{
		cursor: cursor{
			x: 0,
			y: 0,
		},
		lines: []*line{
			&line{
				text: []rune{
					0,
				},
			},
		},
	}

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
			case termbox.KeyEsc:
				break mainloop
			default:
				buf.insertChr(ev.Ch)
				buf.updateLine()
				buf.updateCursor()
			}
		}
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

func (b *buffer) insertChr(r rune) {
	b.lines[b.cursor.y].insertChr(r, b.cursor.x)
	b.cursor.x++
}

func (l *line) insertChr(r rune, p int) {
	l.text = append(l.text[:p+1], l.text[p:]...)
	l.text[p] = r
}

func (b *buffer) updateLine() {
	for i, r := range b.lines[b.cursor.y].text {
		termbox.SetCell(i, b.cursor.y, r, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func (b *buffer) updateCursor() {
	termbox.SetCursor(b.cursor.x, b.cursor.y)
}
