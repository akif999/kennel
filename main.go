package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

const (
	Up = iota
	Down
	Left
	Right
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
				text: []rune{},
			},
		},
	}

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				buf.lineFeed()
			case termbox.KeyArrowUp:
				buf.moveCursor(Up)
			case termbox.KeyArrowDown:
				buf.moveCursor(Down)
			case termbox.KeyArrowLeft:
				buf.moveCursor(Left)
			case termbox.KeyArrowRight:
				buf.moveCursor(Right)
			case termbox.KeyEsc:
				break mainloop
			default:
				buf.insertChr(ev.Ch)
			}
		}
		buf.updateLines()
		buf.updateCursor()
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

func (b *buffer) lineFeed() {
	p := b.cursor.y + 1
	t := make([]*line, len(b.lines), cap(b.lines)+1)
	copy(t, b.lines)
	b.lines = append(t[:p+1], t[p:]...)
	b.lines[p] = new(line)

	b.cursor.x = 0
	b.cursor.y++
}

func (b *buffer) insertChr(r rune) {
	b.lines[b.cursor.y].insertChr(r, b.cursor.x)
	b.cursor.x++
}

func (l *line) insertChr(r rune, p int) {
	t := make([]rune, len(l.text), cap(l.text)+1)
	copy(t, l.text)
	l.text = append(t[:p+1], t[p:]...)
	l.text[p] = r
}

func (b *buffer) updateLines() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	for y, l := range b.lines {
		for x, r := range l.text {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}

func (b *buffer) moveCursor(d int) {
	switch d {
	case Up:
		// guard of top of "rows"
		if b.cursor.y > 0 {
			b.cursor.y--
			// guard of end of "row"
			if b.cursor.x > len(b.lines[b.cursor.y].text) {
				b.cursor.x = len(b.lines[b.cursor.y].text)
			}
		}
		break
	case Down:
		// guard of end of "rows"
		if b.cursor.y < b.linenum()-1 {
			b.cursor.y++
			// guard of end of "row"
			if b.cursor.x > len(b.lines[b.cursor.y].text) {
				b.cursor.x = len(b.lines[b.cursor.y].text)
			}
		}
		break
	case Left:
		if b.cursor.x > 0 {
			b.cursor.x--
		}
		break
	case Right:
		if b.cursor.x < b.lines[b.cursor.y].runenum() {
			b.cursor.x++
		}
		break
	default:
	}
}

func (b *buffer) updateCursor() {
	termbox.SetCursor(b.cursor.x, b.cursor.y)
}

func (b *buffer) linenum() int {
	return len(b.lines)
}

func (l *line) runenum() int {
	return len(l.text)
}
