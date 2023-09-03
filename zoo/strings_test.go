package main

import (
	"strings"
	"testing"
)

func TestStringsJoin(t *testing.T) {
	expect(t, strings.Join([]string{"A", "B", "C"}, ","), "A,B,C")
	expect(t, strings.Join([]string{"Hello", ", ", "World!"}, ""), "Hello, World!")
}

func TestStringsIndex(t *testing.T) {
	expect(t, strings.Index("ABC", "A"), 0)
	expect(t, strings.Index("ABC", "B"), 1)
	expect(t, strings.Index("ABC", "X"), -1)
}

func TestStringsLastIndex(t *testing.T) {
	expect(t, strings.LastIndex("ABCABCABC", "ABC"), 6)
	expect(t, strings.LastIndex("ABCABCABC", "BCA"), 4)
	expect(t, strings.LastIndex("ABCABCABC", "X"), -1)
}
func TestStringsIndexAny(t *testing.T) {
	expect(t, strings.IndexAny("ABC", "ABC"), 0)
	expect(t, strings.IndexAny("ABC", "BCD"), 1)
	expect(t, strings.IndexAny("ABC", "CDE"), 2)
	expect(t, strings.IndexAny("ABC", "XYZ"), -1)
}

func TestStringsLastIndexAny(t *testing.T) {
	expect(t, strings.LastIndexAny("ABC", "ABC"), 2)  // Cが最後に出現する
	expect(t, strings.LastIndexAny("ABC", "BCD"), 2)  // Cが最後に出現する
	expect(t, strings.LastIndexAny("ABC", "CDE"), 2)  // Cが最後に出現する
	expect(t, strings.LastIndexAny("ABC", "XYZ"), -1) // 見つからない
}

func TestStringsHasPrefix(t *testing.T) {
	expect(t, strings.HasPrefix("Go言語", "Go"), true)
	expect(t, strings.HasPrefix("Go言語", "言語"), false)
}
func TestStringsHasSuffix(t *testing.T) {
	expect(t, strings.HasSuffix("Go言語", "Go"), false)
	expect(t, strings.HasSuffix("Go言語", "言語"), true)
}

func TestStringsContains(t *testing.T) {
	expect(t, strings.Contains("Go言語", "Go"), true)
	expect(t, strings.Contains("Go言語", "言語"), true)
	expect(t, strings.Contains("Go言語", "Java"), false)
	expect(t, strings.Contains("Go言語", ""), true)
	expect(t, strings.Contains("", ""), true)
}

func TestStringsContainsAny(t *testing.T) {
	expect(t, strings.ContainsAny("ABCDE", "AE"), true)
	expect(t, strings.ContainsAny("ABCDE", "Cookbook"), true)
	expect(t, strings.ContainsAny("ABCDE", "cookbook"), false)
	expect(t, strings.ContainsAny("ABCDE", "XYZ"), false)
}

func TestStringsCount(t *testing.T) {
	expect(t, strings.Count("ABC", "ABC"), 1)
	expect(t, strings.Count("ABCABCABC", "ABC"), 3)
	expect(t, strings.Count("ABCABCABC", "XYZ"), 0)
	expect(t, strings.Count("ABC", ""), 4)
}

func TestStringsRepeat(t *testing.T) {
	expect(t, strings.Repeat("ABC", 0), "")
	expect(t, strings.Repeat("ABC", 1), "ABC")
	expect(t, strings.Repeat("ABC", 2), "ABCABC")
	expect(t, strings.Repeat("ABC", 3), "ABCABCABC")
}

func TestStringsReplace(t *testing.T) {
	expect(t, strings.Replace("AAAAA", "A", "X", 1), "XAAAA")
	expect(t, strings.Replace("AAAAA", "A", "X", 2), "XXAAA")
	expect(t, strings.Replace("AAAAA", "A", "X", -1), "XXXXX")
	expect(t, strings.Replace("C言語", "C", "Go", 1), "Go言語")
}

func TestStringsSplit(t *testing.T) {
	expectEqual(t, strings.Split("A,B,C", ","), []string{"A", "B", "C"})
	expectEqual(t, strings.SplitAfter("A,B,C", ","), []string{"A,", "B,", "C"})
	expectEqual(t, strings.SplitN("A,B,C,D,E", ",", 3), []string{"A", "B", "C,D,E"})
	expectEqual(t, strings.SplitAfterN("A,B,C,D,E", ",", 3), []string{"A,", "B,", "C,D,E"})
}

func TestStringsLowerUpper(t *testing.T) {
	expect(t, strings.ToLower("ABC"), "abc")
	expect(t, strings.ToUpper("abc"), "ABC")
	expect(t, strings.ToUpper("ë"), "Ë") // ウムラウトも扱える
	expect(t, strings.ToLower("Ë"), "ë") // ウムラウトも扱える
}

func TestStringsTrimSpace(t *testing.T) {
	expect(t, strings.TrimSpace(" \t\n Hello, World! \n\t\r\n"), "Hello, World!")
	expect(t, strings.TrimSpace("　　←日本語の全角空白→　　　"), "←日本語の全角空白→")
}

func TestStringFields(t *testing.T) {
	expectEqual(t, strings.Fields("a b c"), []string{"a", "b", "c"})
	expectEqual(t, strings.Fields("A\tB\tC\nD"), []string{"A", "B", "C", "D"})
	expectEqual(t, strings.Fields("   X  Y   Z   "), []string{"X", "Y", "Z"})
	expectEqual(t, strings.Fields("　　い　　ろ　　は   "), []string{"い", "ろ", "は"})
}
