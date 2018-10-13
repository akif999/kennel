package main

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

func TestGetDigits(t *testing.T) {
	got := getDigit(1)
	if got != 1 {
		t.Errorf("failed")
	}
	got = getDigit(12)
	if got != 2 {
		t.Errorf("failed")
	}
	got = getDigit(123)
	if got != 3 {
		t.Errorf("failed")
	}
	got = getDigit(1234)
	if got != 4 {
		t.Errorf("failed")
	}
	got = getDigit(12345)
	if got != 5 {
		t.Errorf("failed")
	}
}
