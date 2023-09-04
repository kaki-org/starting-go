package main

import (
	"fmt"
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

	// 基本的な正規表現のパターン(並びがそのまま適用される)
	r = regexp.MustCompile(`bc`)
	expect(t, r.MatchString("abcd"), true)
	expect(t, r.MatchString("acbd"), false)

	// .は任意の一文字にマッチする
	r = regexp.MustCompile(`.`)
	expect(t, r.MatchString("ABC"), true)
	expect(t, r.MatchString("日本語"), true)
	expect(t, r.MatchString("\n"), false)

	// |は正規表現パターンのORを表す
	r = regexp.MustCompile(`abc|xyz`)
	expect(t, r.MatchString("abc"), true)
	expect(t, r.MatchString("xyz"), true)
}

func TestRegexpRepeat(t *testing.T) {
	re := regexp.MustCompile(`a+b*`)
	expect(t, re.MatchString("ab"), true)
	expect(t, re.MatchString("a"), true)
	expect(t, re.MatchString("aaaaaaabbb"), true)
	expect(t, re.MatchString("b"), false)

	re = regexp.MustCompile(`A+?A+?X`)
	expect(t, re.MatchString("AX"), false)
	expect(t, re.MatchString("AAX"), true)
	expect(t, re.MatchString("AAAX"), true)
	expect(t, re.MatchString("AAAAAX"), true)
}

func TestRegexpCharacterClass(t *testing.T) {
	re := regexp.MustCompile(`[XYZ]`)
	expect(t, re.MatchString("X"), true)
	expect(t, re.MatchString("Y"), true)
	expect(t, re.MatchString("Z"), true)
	expect(t, re.MatchString("A"), false)

	// 英数字とアンダースコア3文字を表す正規表現
	re = regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	expect(t, re.MatchString("ABC"), true)
	expect(t, re.MatchString("x01"), true)
	expect(t, re.MatchString("abcdefg"), false)
	expect(t, re.MatchString("あいう"), false)

	// 英数字とアンダースコア「以外」を表す正規表現
	re = regexp.MustCompile(`[^0-9A-Za-z_]`)
	expect(t, re.MatchString("ABC"), false)
	expect(t, re.MatchString("x01"), false)
	expect(t, re.MatchString("あいう"), true)
}

func TestRegexpPerl(t *testing.T) {
	// Perlの正規表現のような記法
	re := regexp.MustCompile(`\d+`)
	expect(t, re.MatchString("123456"), true)
	expect(t, re.MatchString("X=1"), true)
	expect(t, re.MatchString("abcde"), false)
}

func TestRegexpUnicode(t *testing.T) {
	// Unicodeにおける「カタカナ」にマッチする正規表現
	re := regexp.MustCompile(`\p{Katakana}`)
	expect(t, re.MatchString("ア"), true)
	expect(t, re.MatchString("ｱ"), true)
	expect(t, re.MatchString("あ"), false)
}

func TestRegexpGroup(t *testing.T) {
	/* グルーピングを利用した正規表現 */
	re := regexp.MustCompile(`(佐藤|鈴木)(太郎|花子)`)
	expect(t, re.MatchString("佐藤太郎"), true)
	expect(t, re.MatchString("鈴木花子"), true)
	expect(t, re.MatchString("佐藤一郎"), false)
}

func TestRegexpFindString(t *testing.T) {
	/* 連続した英数字かアンダースコアの繰り返し */
	re := regexp.MustCompile(`\w+`)
	// マッチした最初の文字列を取得
	expect(t, re.FindString("abc xyz 999"), "abc")
	// マッチした文字列を指定した数だけ取得
	expectEqual(t, re.FindAllString("abc xyz 999", 2), []string{"abc", "xyz"})
	expectEqual(t, re.FindAllString("abc xyz 999", -1), []string{"abc", "xyz", "999"})
}

func TestRegexpSplit(t *testing.T) {
	// タブやスペースなどの空白にマッチ
	re := regexp.MustCompile(`\s+`)

	expectEqual(t, re.Split("A B  C   D\tE", 3), []string{"A", "B", "C   D\tE"})
	expectEqual(t, re.Split("A B  C   D\tE", -1), []string{"A", "B", "C", "D", "E"}) // \tはなくなる
}

func TestRegexpReplaceAllString(t *testing.T) {
	// 「佐藤」にマッチする正規表現
	re := regexp.MustCompile(`佐藤`)

	expect(t, re.ReplaceAllString("佐藤太郎", "鈴木"), "鈴木太郎")
	expect(t, re.ReplaceAllString("XYZ", "鈴木"), "XYZ")
}

func TestRegexpFindAllStringSubmatch(t *testing.T) {
	// 「000-1111-2222」のような形式にマッチする正規表現
	re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)

	s := `
	00-1111-2222
	3333-44-55
	666-777-888
	9-9-9
	`

	// 正規表現のグルーピングも含めて取得
	ms := re.FindAllStringSubmatch(s, -1)

	for _, v := range ms {
		// fmt.Println(i) // インデックス
		fmt.Println(v)
	}

	expect(t,
		re.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"),
		"Tel: 222-111-000")
}
