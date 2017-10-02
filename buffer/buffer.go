package buffer

const (
	InitialBufferSize = 4095
)

type Buffer struct {
	Runes     []rune
	InsertPos uint64
}

func NewBuffer() *Buffer {
	buf := make([]rune, InitialBufferSize)
	return &Buffer{Runes: buf, InsertPos: 0}
}

func (b *Buffer) Insert(chr rune) {
	b.Runes[b.InsertPos] = chr
	b.InsertPos++
}
