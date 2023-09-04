package main

import (
	"regexp"
	"testing"
)

func TestRegexpMatchString(t *testing.T) {
	m, _ := regexp.MatchString("A", "ABC")
	expect(t, m, true)
	m, _ = regexp.MatchString("A", "XYZ")
	expect(t, m, false)
}

func TestRegexpCompile(t *testing.T) {
	// 正規表現のパターンをコンパイルする
	r, _ := regexp.Compile("A")
	expect(t, r.MatchString("ABC"), true)
	expect(t, r.MatchString("XYZ"), false)
}
