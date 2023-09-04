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

func TestRegexpMustCompile(t *testing.T) {
	// 正規表現のパターンをコンパイルする(エラー時はランタイムパニックが発生する)
	// 返り値は*Regexp型を一つだけ返す
	r := regexp.MustCompile("AB")
	expect(t, r.MatchString("ABC"), true)
	expect(t, r.MatchString("XYZ"), false)

	r = regexp.MustCompile("\\d") // バックスラッシュをエスケープしないといけない

	expect(t, r.MatchString("123"), true)
	expect(t, r.MatchString("XYZ"), false)

	r = regexp.MustCompile(`\d`) // RAW文字列ならバックスラッシュのエスケープが不要
	expect(t, r.MatchString("123"), true)
	expect(t, r.MatchString("XYZ"), false)

	/* (?i)は大文字小文字を区別しない。 Perlなどでいう、/str/iのようなオプション */
	r = regexp.MustCompile(`(?i)abc`)
	expect(t, r.MatchString("ABC"), true)
	expect(t, r.MatchString("aBc"), true)

	// 幅を持たない正規表現のパターン
	r = regexp.MustCompile(`^XYZ$`)
	expect(t, r.MatchString("XYZ"), true)
	expect(t, r.MatchString(" XYZ "), false)
}
