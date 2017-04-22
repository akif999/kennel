package buffer

import (
// "github.com/nsf/termbox-go"
// "io/ioutil"
//"os"
//"strconv"
)

// TODO
// bufferパッケージとしてリライトする
// とりあえずは、内部バッファへライトして、termboxで出力する実装をする

type ScrnBuffer struct {
	Chr []rune
	// 2次元配列で保持しない(改行はあくまで改行コードで行う)
	/* 別個のサイズの情報は必要ない気もする(Bufから直接計算できるからメソッドで実装できそう)
	vSize uint
	hSize uint
	*/
}

// NewScenBufferは、新たなScrnBuffer型の参照を返す
func NewScenBuffer() *ScrnBuffer {
	b := new(ScrnBuffer)
	return b
}

// WriteChrToSBufは、内部バッファの任意の座標へrune型の文字を書き込む
func (b *ScrnBuffer) WriteChrToSBuf(l uint, chr rune) {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[l+1:], b.Chr[l:])
	b.Chr[l] = chr
}

func (b *ScrnBuffer) AddNewLine(l uint) {
	b.Chr = append(b.Chr, 0)
	copy(b.Chr[l+1:], b.Chr[l:])
	b.Chr[l] = rune('\n')
}

// 以下旧ソースのfunction
// run & buildのため削除
