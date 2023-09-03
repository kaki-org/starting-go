package main

import (
	"testing"
	"unicode"
)

func TestIsDigit(t *testing.T) {
	expect(t, unicode.IsDigit('X'), false)
	expect(t, unicode.IsDigit('1'), true)
	expect(t, unicode.IsDigit('２'), true)
	expect(t, unicode.IsDigit('三'), false)
}

func TestIsLetter(t *testing.T) {
	expect(t, unicode.IsLetter('X'), true)
	expect(t, unicode.IsLetter('Ｘ'), true)
	expect(t, unicode.IsLetter('3'), false)
	expect(t, unicode.IsLetter('３'), false)
	expect(t, unicode.IsLetter('三'), true)
	expect(t, unicode.IsLetter('〒'), false) // 〒は郵便記号
}

func TestIsSpace(t *testing.T) {
	expect(t, unicode.IsSpace(' '), true)
	expect(t, unicode.IsSpace('　'), true)
	expect(t, unicode.IsSpace('A'), false)
	expect(t, unicode.IsSpace('１'), false)
	expect(t, unicode.IsSpace('あ'), false)
}

func TestIsControl(t *testing.T) {
	expect(t, unicode.IsControl(' '), false)
	expect(t, unicode.IsControl('　'), false)
	expect(t, unicode.IsControl('A'), false)
	expect(t, unicode.IsControl('１'), false)
	expect(t, unicode.IsControl('あ'), false)
	expect(t, unicode.IsControl('\n'), true)
	expect(t, unicode.IsControl('\t'), true)
}

func expect(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Expected %d, got %d", a, b)
	}
}
