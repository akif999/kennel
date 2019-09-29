package buffer

import (
	"reflect"
	"testing"
)

func TestInsertChr(t *testing.T) {
	type test struct {
		r    rune
		p    int
		text []rune
		exp  []rune
		got  []rune
	}
	ts := []test{
		test{
			r:    'c',
			p:    2,
			text: []rune{'a', 'b', 'd'},
			exp:  []rune{'a', 'b', 'c', 'd'},
		},
		test{
			r:    '3',
			p:    2,
			text: []rune{'a', 'b', 'd'},
			exp:  []rune{'a', 'b', '3', 'd'},
		},
		test{
			r:    'e',
			p:    3,
			text: []rune{'a', 'b', 'd'},
			exp:  []rune{'a', 'b', 'd', 'e'},
		},
	}

	for _, te := range ts {
		l := new(Line)
		l.Text = te.text
		l.insertChr(te.r, te.p)
		te.got = l.Text
		if !reflect.DeepEqual(te.got, te.exp) {
			t.Errorf("failed got : %v, exp : %v", te.got, te.exp)
		}
	}
}

func TestDeleteChr(t *testing.T) {
	type test struct {
		p    int
		text []rune
		exp  []rune
		got  []rune
	}
	ts := []test{
		test{
			p:    1,
			text: []rune{'a', 'b', 'c'},
			exp:  []rune{'b', 'c'},
		},
		test{
			p:    2,
			text: []rune{'a', 'b', 'c'},
			exp:  []rune{'a', 'c'},
		},
		test{
			p:    3,
			text: []rune{'a', 'b', 'c'},
			exp:  []rune{'a', 'b'},
		},
	}

	for _, te := range ts {
		l := new(Line)
		l.Text = te.text
		l.deleteChr(te.p)
		te.got = l.Text
		if !reflect.DeepEqual(te.got, te.exp) {
			t.Errorf("failed got : %v, exp : %v", te.got, te.exp)
		}
	}
}
