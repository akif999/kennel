package buffer

import (
// "github.com/nsf/termbox-go"
// "io/ioutil"
//"os"
//"strconv"
)

// TODO
// bufferパッケージとしてリライトする
// bufferとして必要なメンバ、機能を追加する

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

// WriteChrToSBufは、内部バッファの任意の座標へrune型の文字を書き込む
func (b *ScrnBuffer) WriteChrToSBuf(chr rune) {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[b.Pos+1:], b.Chr[b.Pos:])
	b.Chr[b.Pos] = chr
	b.Pos++
}

// AddNewLineは、内部バッファへ改行コードを追加する
func (b *ScrnBuffer) AddNewLine() {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[b.Pos+1:], b.Chr[b.Pos:])
	b.Chr[b.Pos] = rune('\n')
	b.Pos++
}

// 以下旧ソースのfunction
// run & buildのため削除
