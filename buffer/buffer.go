package buffer

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	termbox "github.com/nsf/termbox-go"
)

const (
	Up CursorDir = iota
	Down
	Left
	Right
)

var (
	undoBuf = &bufStack{}
	redoBuf = &bufStack{}
)

type bufStack struct {
	bufs []*Buffer
}

type Buffer struct {
	Cursor          Cursor
	Lines           []*Line
	ShowStartHeight int
	ShowStartWidth  int
}

type Cursor struct {
	X      int
	Y      int
	Offset int
}

type CursorDir uint8

type Line struct {
	Text []rune
}

func New() (*Buffer, error) {
	b := new(Buffer)
	return b, nil
}

func (l *Line) insertChr(r rune, p int) {
	t := make([]rune, len(l.Text), cap(l.Text)+1)
	copy(t, l.Text)
	l.Text = append(t[:p+1], t[p:]...)
	l.Text[p] = r
}
func (l *Line) deleteChr(p int) {
	p = p - 1
	l.Text = append(l.Text[:p], l.Text[p+1:]...)
}

func splitLine(l *Line, pos int) ([]rune, []rune) {
	return l.Text[:pos], l.Text[pos:]
}

func joinLine(l *Line, con []rune) {
	l.Text = append(l.Text, con...)
}

func (b *Buffer) getTextOnCursorLine() []rune {
	return b.Lines[b.Cursor.Y].Text
}

func (b *Buffer) NumOfLines() int {
	return len(b.Lines)
}

func (b *Buffer) numOfColsOnCursor() int {
	return len(b.Lines[b.Cursor.Y].Text)
}

func (l *Line) runenum() int {
	return len(l.Text)
}

func (b *Buffer) PushBufToUndoRedoBuffer() {
	tb := new(Buffer)
	tb.Cursor.X = b.Cursor.X
	tb.Cursor.Y = b.Cursor.Y
	for i, l := range b.Lines {
		tl := new(Line)
		tb.Lines = append(tb.Lines, tl)
		tb.Lines[i].Text = l.Text
	}
	undoBuf.bufs = append(undoBuf.bufs, tb)
}

func (b *Buffer) cursorUp() {
	// guard of top of "rows"
	if b.Cursor.Y > 0 {
		b.Cursor.Y--
		// guard of end of "row"
		if b.Cursor.X > b.numOfColsOnCursor() {
			b.Cursor.X = b.numOfColsOnCursor()
		}
		if b.isCursorOutOfWindowTop() {
			b.ShowStartHeight--
		}
	}
}

func (b *Buffer) cursorDown() {
	// guard of end of "rows"
	if b.Cursor.Y < b.NumOfLines()-1 {
		b.Cursor.Y++
		// guard of end of "row"
		if b.Cursor.X > b.numOfColsOnCursor() {
			b.Cursor.X = b.numOfColsOnCursor()
		}
		if b.isCursorOutOfWindowBottom() {
			b.ShowStartHeight++
		}
	}
}
func (b *Buffer) cursorLeft() {
	winWidth, _ := termbox.Size()
	if b.Cursor.X > 0 {
		b.Cursor.X--
		if b.isCursorOutOfWindowLeft() {
			b.ShowStartWidth--
		}
	} else {
		// guard of top of "rows"
		if b.Cursor.Y > 0 {
			b.Cursor.Y--
			b.Cursor.X = b.numOfColsOnCursor()
			b.ShowStartHeight--
			b.ShowStartWidth = winWidth - GetDigit(b.NumOfLines()) + 1 + b.numOfColsOnCursor()
		}
	}
}

func (b *Buffer) cursorRight() {
	if b.Cursor.X < b.Lines[b.Cursor.Y].runenum() {
		b.Cursor.X++
		if b.isCursorOutOfWindowRight() {
			b.ShowStartWidth++
		}
	} else {
		// guard of end of "rows"
		if b.Cursor.Y < b.NumOfLines()-1 {
			b.Cursor.X = 0
			b.Cursor.Y++
			b.ShowStartHeight++
			b.ShowStartWidth = 0
		}
	}
}

func (b *Buffer) isCursorOutOfWindowTop() bool {
	return b.Cursor.Y < b.ShowStartHeight
}

func (b *Buffer) isCursorOutOfWindowBottom() bool {
	_, winHeight := termbox.Size()
	return b.Cursor.Y+1 > b.ShowStartHeight+winHeight
}

func (b *Buffer) isCursorOutOfWindowLeft() bool {
	return b.Cursor.X < b.ShowStartWidth
}

func (b *Buffer) isCursorOutOfWindowRight() bool {
	winWidth, _ := termbox.Size()
	offset := GetDigit(b.NumOfLines()) + 1
	return b.Cursor.X+1+offset > b.ShowStartWidth+winWidth
}

func GetDigit(linenum int) int {
	d := 0
	for linenum != 0 {
		linenum = linenum / 10
		d++
	}
	return d
}

func (b *Buffer) ReadFileToBuf(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		l := new(Line)
		l.Text = []rune(scanner.Text())
		b.Lines = append(b.Lines, l)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (b *Buffer) writeBufToFile(path string) {
	content := []byte{}
	for _, l := range b.Lines {
		l.Text = append(l.Text, '\n')
		content = append(content, string(l.Text)...)
	}
	ioutil.WriteFile(path, content, os.ModePerm)
}

func (b *Buffer) LineFeed() {
	p := b.Cursor.Y + 1
	// split line by the Cursor and store these
	fh, lh := splitLine(b.Lines[b.Cursor.Y], b.Cursor.X)

	t := make([]*Line, len(b.Lines), cap(b.Lines)+1)
	copy(t, b.Lines)
	b.Lines = append(t[:p+1], t[p:]...)
	b.Lines[p] = new(Line)

	// write back previous line and newline
	b.Lines[p-1].Text = fh
	b.Lines[p].Text = lh

	b.Cursor.X = 0
	b.Cursor.Y++
}

func (b *Buffer) BackSpace() {
	if b.Cursor.X == 0 && b.Cursor.Y == 0 {
		// nothing to do
	} else {
		if b.Cursor.X == 0 {
			// store current line
			current := b.getTextOnCursorLine()
			// delete current line
			b.Lines = append(b.Lines[:b.Cursor.Y], b.Lines[b.Cursor.Y+1:]...)
			b.Cursor.Y--
			// // join stored Lines to previous line-end
			prev := b.getTextOnCursorLine()
			joinLine(b.Lines[b.Cursor.Y], current)
			b.Cursor.X = len(prev)
		} else {
			b.Lines[b.Cursor.Y].deleteChr(b.Cursor.X)
			b.Cursor.X--
		}
	}
}

func (b *Buffer) InsertChr(r rune) {
	b.Lines[b.Cursor.Y].insertChr(r, b.Cursor.X)
	b.Cursor.X++
}

func (b *Buffer) MoveCursor(d CursorDir) {
	switch d {
	case Up:
		b.cursorUp()
	case Down:
		b.cursorDown()
	case Left:
		b.cursorLeft()
	case Right:
		b.cursorRight()
	default:
	}
}

func (b *Buffer) Undo() {
	if len(undoBuf.bufs) == 0 {
		return
	}
	if len(undoBuf.bufs) > 1 {
		redoBuf.bufs = append(redoBuf.bufs, undoBuf.bufs[len(undoBuf.bufs)-1])
		undoBuf.bufs = undoBuf.bufs[:len(undoBuf.bufs)-1]
	}
	tb := undoBuf.bufs[len(undoBuf.bufs)-1]
	undoBuf.bufs = undoBuf.bufs[:len(undoBuf.bufs)-1]
	b.Cursor.X = tb.Cursor.X
	b.Cursor.Y = tb.Cursor.Y
	for i, l := range tb.Lines {
		tl := new(Line)
		b.Lines = append(b.Lines, tl)
		b.Lines[i].Text = l.Text
	}
}

func (b *Buffer) Redo() {
	if len(redoBuf.bufs) == 0 {
		return
	}
	tb := redoBuf.bufs[len(redoBuf.bufs)-1]
	redoBuf.bufs = redoBuf.bufs[:len(redoBuf.bufs)-1]
	b.Cursor.X = tb.Cursor.X
	b.Cursor.Y = tb.Cursor.Y
	for i, l := range tb.Lines {
		tl := new(Line)
		b.Lines = append(b.Lines, tl)
		b.Lines[i].Text = l.Text
	}
}

func (b *Buffer) SaveAs() {
	b.writeBufToFile("./output.txt")
}
