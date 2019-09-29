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

func TestGetDigits(t *testing.T) {
	got := GetDigit(1)
	if got != 1 {
		t.Errorf("failed")
	}
	got = GetDigit(12)
	if got != 2 {
		t.Errorf("failed")
	}
	got = GetDigit(123)
	if got != 3 {
		t.Errorf("failed")
	}
	got = GetDigit(1234)
	if got != 4 {
		t.Errorf("failed")
	}
	got = GetDigit(12345)
	if got != 5 {
		t.Errorf("failed")
	}
}
