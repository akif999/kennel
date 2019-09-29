package window

import (
	"reflect"
	"testing"
)

func TestMakeLineNum(t *testing.T) {
	got := makeLineNum(1, 4)
	exp := []rune{' ', ' ', ' ', '1', ' '}
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("failed : got : %v, exp : %v", got, exp)
	}
	got = makeLineNum(12, 4)
	exp = []rune{' ', ' ', '1', '2', ' '}
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("failed : got : %v, exp : %v", got, exp)
	}
	got = makeLineNum(123, 4)
	exp = []rune{' ', '1', '2', '3', ' '}
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("failed : got : %v, exp : %v", got, exp)
	}
	got = makeLineNum(1234, 4)
	exp = []rune{'1', '2', '3', '4', ' '}
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("failed : got : %v, exp : %v", got, exp)
	}
}
