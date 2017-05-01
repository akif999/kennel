package buffer

import (
// "github.com/nsf/termbox-go"
// "io/ioutil"
//"os"
//"strconv"
)

// TODO
// bufferとして必要なメンバ、機能を追加する
// DelChrFromSBufのロジックを変更する(適切ではない可能性がある)
// Cuesorの行上げや、行末制限を実装する

type ScrnBuffer struct {
	Chr []rune
	Pos int
	// 2次元配列で保持しない(改行はあくまで改行コードで行う)
}

// NewScenBufferは、新たなScrnBuffer型の参照を返す
func NewScenBuffer() *ScrnBuffer {
	b := new(ScrnBuffer)
	return b
}

// WriteChrToSBufは、内部バッファの位置(Pos)へrune型の文字を書き込む
func (b *ScrnBuffer) WriteChrToSBuf(chr rune) {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[b.Pos+1:], b.Chr[b.Pos:])
	b.Chr[b.Pos] = chr
	b.Pos++
}

// DelChrFromSBufは、内部バッファの位置(Pos)にある文字を削除する
func (b *ScrnBuffer) DelChrFromSBuf() {
	if b.Pos == 0 {
		return
	}
	b.Chr = append(b.Chr[:b.Pos-1], b.Chr[b.Pos:]...)
	n := make([]rune, len(b.Chr))
	copy(n, b.Chr)
	b.Pos--
}

// AddNewLineは、内部バッファへ改行コードを追加する
func (b *ScrnBuffer) AddNewLine() {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[b.Pos+1:], b.Chr[b.Pos:])
	b.Chr[b.Pos] = rune('\n')
	b.Pos++
}
