package command

import termbox "github.com/nsf/termbox-go"

type Command int

const (
	QuitApp Command = 0
	DmyCmd  Command = 1
	NilCmd  Command = 2
)

func ParseKeyToCommand(event termbox.Event) (cmd Command, err error) {
	switch event.Type {
	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyEsc:
			return QuitApp, nil
		case termbox.KeyEnter:
		case termbox.KeyArrowUp:
		case termbox.KeyArrowDown:
		case termbox.KeyArrowLeft:
		case termbox.KeyArrowRight:
		case termbox.KeyBackspace, termbox.KeyBackspace2:
		case termbox.KeyCtrlS:
		default:
			if event.Ch != 0 {
			}
		}
	case termbox.EventError:
		return NilCmd, event.Err
	}
	return NilCmd, nil
}
