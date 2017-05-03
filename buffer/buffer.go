package buffer

import ()

type ScrnBuffer struct {
	Chr         []rune
	LineLengths []int
	Pos         int
}

// NewScenBufferは、新たなScrnBuffer型の参照を返す
func NewScenBuffer() *ScrnBuffer {
	b := new(ScrnBuffer)
	return b
}

// WriteChrToSBufは、内部バッファの位置(Pos)へrune型の文字を書き込む
func (b *ScrnBuffer) WriteChrToSBuf(chr rune, row int) {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[b.Pos+1:], b.Chr[b.Pos:])
	b.Chr[b.Pos] = chr
	b.LineLengths[row]++
}

// DelChrFromSBufは、内部バッファの位置(Pos)にある文字を削除する
func (b *ScrnBuffer) DelChrFromSBuf(row int) {
	if b.Pos == 0 {
		return
	}
	b.Chr = append(b.Chr[:b.Pos-1], b.Chr[b.Pos:]...)
	b.LineLengths[row]--
}

// AddNewLineは、内部バッファへ改行コードを追加する
func (b *ScrnBuffer) AddNewLine() {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[b.Pos+1:], b.Chr[b.Pos:])
	b.Chr[b.Pos] = rune('\n')
	b.Pos++
	b.LineLengths = append(b.LineLengths, 0)
}

func (b *ScrnBuffer) ConvertCursorToBufPos(row int, col int) {
	b.Pos = 0
	for i, l := range b.LineLengths {
		if col == 0 || col == l {
			b.Pos = (b.Pos + (l - col))
		} else {
			b.Pos = (b.Pos + (l - col)) - 1
		}
		if i == row {
			break
		}
	}
}
