package main

type bufStack struct {
	bufs []*buffer
}

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

func (b *buffer) lineFeed() {
	p := b.cursor.y + 1
	// split line by the cursor and store these
	fh, lh := b.lines[b.cursor.y].split(b.cursor.x)

	t := make([]*line, len(b.lines), cap(b.lines)+1)
	copy(t, b.lines)
	b.lines = append(t[:p+1], t[p:]...)
	b.lines[p] = new(line)

	// write back previous line and newline
	b.lines[p-1].text = fh
	b.lines[p].text = lh

	b.cursor.x = 0
	b.cursor.y++
}

func (b *buffer) backSpace() {
	if b.cursor.x == 0 && b.cursor.y == 0 {
		// nothing to do
	} else {
		if b.cursor.x == 0 {
			// store current line
			t := b.lines[b.cursor.y].text
			// delete current line
			b.lines = append(b.lines[:b.cursor.y], b.lines[b.cursor.y+1:]...)
			b.cursor.y--
			// // join stored lines to previous line-end
			plen := b.lines[b.cursor.y].text
			b.lines[b.cursor.y].text = append(b.lines[b.cursor.y].text, t...)
			b.cursor.x = len(plen)
		} else {
			b.lines[b.cursor.y].deleteChr(b.cursor.x)
			b.cursor.x--
		}
	}
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

func (l *line) deleteChr(p int) {
	p = p - 1
	l.text = append(l.text[:p], l.text[p+1:]...)
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
		} else {
			// guard of top of "rows"
			if b.cursor.y > 0 {
				b.cursor.y--
				b.cursor.x = len(b.lines[b.cursor.y].text)
			}
		}
		break
	case Right:
		if b.cursor.x < b.lines[b.cursor.y].runenum() {
			b.cursor.x++
		} else {
			// guard of end of "rows"
			if b.cursor.y < b.linenum()-1 {
				b.cursor.x = 0
				b.cursor.y++
			}
		}
		break
	default:
	}
}

func (b *buffer) undo() {
	if len(undoBuf.bufs) == 0 {
		return
	}
	if len(undoBuf.bufs) > 1 {
		redoBuf.bufs = append(redoBuf.bufs, undoBuf.bufs[len(undoBuf.bufs)-1])
		undoBuf.bufs = undoBuf.bufs[:len(undoBuf.bufs)-1]
	}
	tb := undoBuf.bufs[len(undoBuf.bufs)-1]
	undoBuf.bufs = undoBuf.bufs[:len(undoBuf.bufs)-1]
	b.cursor.x = tb.cursor.x
	b.cursor.y = tb.cursor.y
	for i, l := range tb.lines {
		tl := new(line)
		b.lines = append(b.lines, tl)
		b.lines[i].text = l.text
	}
}

func (b *buffer) redo() {
	if len(redoBuf.bufs) == 0 {
		return
	}
	tb := redoBuf.bufs[len(redoBuf.bufs)-1]
	redoBuf.bufs = redoBuf.bufs[:len(redoBuf.bufs)-1]
	b.cursor.x = tb.cursor.x
	b.cursor.y = tb.cursor.y
	for i, l := range tb.lines {
		tl := new(line)
		b.lines = append(b.lines, tl)
		b.lines[i].text = l.text
	}
}

func (l *line) split(pos int) ([]rune, []rune) {
	return l.text[:pos], l.text[pos:]
}

func (b *buffer) linenum() int {
	return len(b.lines)
}

func (l *line) runenum() int {
	return len(l.text)
}
