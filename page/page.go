package page

import "github.com/akif999/kennel/window"

type Page struct {
	Windows []*window.Window
}

func NewPage() *Page {
	return &Page{}
}

func (p *Page) Insert(chr rune, WinNum int) {
	p.Windows[WinNum].Insert(chr)
}

func (p *Page) Draw(WinNum int) {
	p.Windows[WinNum].Draw()
}

func (p *Page) UpdateCursor(WinNum int) {
	p.Windows[WinNum].UpdateCursor()
}
