package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	var a [10]int
	s := make([]int, 10)

	fmt.Printf("こっちは配列%T\n", a)
	fmt.Printf("こっちはスライス%T\n", s)
	fmt.Printf("%d\n%d\n", len(a), len(s))
	fmt.Printf("%v\n%v\n", a, s)
	expect := 10
	expectValue := 0
	actual := s
	if expect != len(actual) {
		t.Errorf(decimalExpectFormat, expect, actual)
	}
	if expectValue != actual[9] {
		t.Errorf(decimalExpectFormat, expectValue, actual[0])
	}
	// 要素数を超えた値を指定するとpanicになる
	// fmt.Println(actual[10]) // panic: runtime error: index out of range
}

func TestSliceCap(t *testing.T) {
	/* 要素数5、容量5のスライス */
	s5 := make([]int, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s5), cap(s5), s5)
	/* 要素数5、容量10のスライス */
	s10 := make([]int, 5, 10)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s10), cap(s10), s10)
	expect_cap5 := 5
	expect_cap10 := 10
	if (expect_cap5 != cap(s5)) || (expect_cap10 != cap(s10)) {
		t.Errorf(decimalExpectFormat, expect_cap5, cap(s5))
		t.Errorf(decimalExpectFormat, expect_cap10, cap(s10))
	}
	// fmt.Println(s10[9]) // これはpanicになる
	s10[4] = 100
	fmt.Println(s10[4])
	fmt.Printf("%#v\n", s10[0:10]) // 0から9までの要素を取得

	s5[0] = 1
	s5[1] = 2
	s5[2] = 3
	s5[3] = 4
	s5[4] = 5

	// makeを使用しないリテラルでのスライスの作成
	s := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(s5, s) {
		t.Errorf(valueExpectFormat, s5, s)
	}
}

func TestSimpleSliceExpressions(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	actual1 := a[0:2]
	actual2 := a[2:]
	actual3 := a[:4]
	actual4 := a[:]
	actual5 := a[len(a)-2:] // 式もかける

	expect1 := []int{1, 2}
	expect2 := []int{3, 4, 5}
	expect3 := []int{1, 2, 3, 4}
	expect4 := []int{1, 2, 3, 4, 5}
	expect5 := []int{4, 5}

	if !reflect.DeepEqual(expect1, actual1) {
		t.Errorf(valueExpectFormat, expect1, actual1)
	}
	if !reflect.DeepEqual(expect2, actual2) {
		t.Errorf(valueExpectFormat, expect2, actual2)
	}
	if !reflect.DeepEqual(expect3, actual3) {
		t.Errorf(valueExpectFormat, expect3, actual3)
	}
	if !reflect.DeepEqual(expect4, actual4) {
		t.Errorf(valueExpectFormat, expect4, actual4)
	}
	if !reflect.DeepEqual(expect5, actual5) {
		t.Errorf(valueExpectFormat, expect5, actual5)
	}

}

func TestSimpleSliceExpressionsString(t *testing.T) {
	s := "ABCDE"[1:3]
	multibytes := "あいうえお"[3:9] // バイト列([]byte)であるとみなされる

	expect := "BC"
	expectMultibytes := "いう"
	if !reflect.DeepEqual(expect, s) {
		t.Errorf(valueExpectFormat, expect, s)
	}
	if !reflect.DeepEqual(expectMultibytes, multibytes) {
		t.Errorf(valueExpectFormat, expectMultibytes, multibytes)
	}
}
func TestSliceAppend(t *testing.T) {
	// 配列と違って、スライスは拡張できる
	s := []int{1, 2, 3}
	s = append(s, 4)
	// append(s, 4) // 代入を伴わないappendはコンパイルエラー
	expect := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf(valueExpectFormat, expect, s)
	}

	// 5,6,7を続けて追加できる
	s = append(s, 5, 6, 7)
	expect = []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf(valueExpectFormat, expect, s)
	}

	// スライス同士を結合する。s1の追加の「...」に注意
	s0 := []int{1, 2, 3}
	s1 := []int{4, 5, 6}
	s2 := append(s0, s1...)

	expect = []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(expect, s2) {
		t.Errorf(valueExpectFormat, expect, s2)
	}

	// []byte型のスライスに文字列を追加する
	var b []byte
	b = append(b, "あいうえお"...)
	b = append(b, "かきくけこ"...)
	b = append(b, "さしすせそ"...)
	fmt.Println(b)         // [227 129 130 227 129 132 227 129 134 227 129 136 227 129 138 227 129 139 227 129 141 227 129 143 227 129 145 227 129 147 227 129 149 227 129 151 227 129 153 227 129 155 227 129 157]
	fmt.Println(string(b)) // あいうえおかきくけこさしすせそ
	expectStr := "あいうえおかきくけこさしすせそ"
	if expectStr != string(b) {
		t.Errorf(valueExpectFormat, expect, string(b))
	}

}

func TestSliceAppendCapacity(t *testing.T) {
	s := make([]int, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	expect := 0
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}

	s = append(s, 1)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	expect = 1
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}

	s = append(s, []int{2, 3, 4}...)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	expect = 4
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}

	s = append(s, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	expect = 8
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}

	s = append(s, 6, 7, 8, 9)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s), cap(s), s)
	expect = 16
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}

	s = make([]int, 1024, 1024)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	expect = 1024
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}

	// その時の実装においてどのくらい拡張されるかは不明
	s = append(s, 0)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	expect = 1536
	if expect != cap(s) {
		t.Errorf(decimalExpectFormat, expect, cap(s))
	}
}

func TestSliceCopy(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	n := copy(s1, s2) // s2の要素数分だけコピーする

	expect_len := 2
	expect_slice := []int{10, 11, 3, 4, 5}
	if expect_len != n {
		t.Errorf(decimalExpectFormat, expect_len, n)
	}
	if !reflect.DeepEqual(expect_slice, s1) {
		t.Errorf(valueExpectFormat, expect_slice, s1)
	}

	s3 := []int{1, 2, 3, 4, 5}
	s4 := []int{10, 11, 12, 13, 14, 15, 16}
	n2 := copy(s3, s4) // コピーできたs3の要素数がn2に入る
	expect_len = 5
	expect_slice = []int{10, 11, 12, 13, 14}
	if expect_len != n2 {
		t.Errorf(decimalExpectFormat, expect_len, n2)
	}
	if !reflect.DeepEqual(expect_slice, expect_slice) {
		t.Errorf(valueExpectFormat, expect_slice, expect_slice)
	}

	// あくまで[]byte単位でコピーする
	b := make([]byte, 9)
	n3 := copy(b, "あいうえお")

	expect_len = 9
	expect_string := "あいう"
	if expect_len != n3 {
		t.Errorf(decimalExpectFormat, expect_len, n3)
	}
	if expect_string != string(b) {
		t.Errorf(valueExpectFormat, expect_string, string(b))
	}
}

func TestFullSliceExpressions(t *testing.T) {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// a[low : high : max] という形式でスライスを作成する
	s1 := a[2:4]
	expect_len := 2
	expect_cap := 8 // (len(a) - low) → 10 - 2 = 8
	if expect_len != len(s1) {
		t.Errorf(decimalExpectFormat, expect_len, len(s1))
	}
	if expect_cap != cap(s1) {
		t.Errorf(decimalExpectFormat, expect_cap, cap(s1))
	}

	s2 := a[2:4:4]
	expect_len = 2
	expect_cap = 2 // max - low → 4 - 2 = 2
	if expect_len != len(s2) {
		t.Errorf(decimalExpectFormat, expect_len, len(s2))
	}
	if expect_cap != cap(s2) {
		t.Errorf(decimalExpectFormat, expect_cap, cap(s2))
	}

	s3 := a[2:4:6]
	expect_len = 2
	expect_cap = 4 // max - low → 6 - 2 = 4
	if expect_len != len(s3) {
		t.Errorf(decimalExpectFormat, expect_len, len(s3))
	}
	if expect_cap != cap(s3) {
		t.Errorf(decimalExpectFormat, expect_cap, cap(s3))
	}
}

func TestSliceFor(t *testing.T) {
	s := []string{"Apple", "Banana", "Cherry"}

	output := ""
	for i, v := range s {
		output += fmt.Sprintf("i=%d v=%s\n", i, v)
		fmt.Printf("i=%d v=%s\n", i, v)
		s = append(s, "Melon") // 要素の追加をしても元のスライスsには影響しない
	}
	expect := "i=0 v=Apple\ni=1 v=Banana\ni=2 v=Cherry\n"
	if expect != output {
		t.Errorf(valueExpectFormat, expect, output)
	}
	expectSlice := []string{"Apple", "Banana", "Cherry", "Melon", "Melon", "Melon"}
	if !reflect.DeepEqual(expectSlice, s) {
		t.Errorf(valueExpectFormat, expectSlice, s)
	}

}

// 可変長引数を取る関数
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// func sum(s ...int) int { // OK
// func sum(s ...int, s...int) int { // コンパイルエラー
// func sum(s ...int, b bool) int { // コンパイルエラー

func TestVariableLengthArguments(t *testing.T) {
	expect := 6
	actual := sum(1, 2, 3)
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}

	expect = 15
	actual = sum(1, 2, 3, 4, 5)
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}

	expect = 0
	actual = sum()
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}

	s := []int{1, 2, 3}
	expect = 6
	actual = sum(s...)
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}
}

func powArray(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
func powSlice(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
func TestArrayAndSlice(t *testing.T) {
	a := [3]int{1, 2, 3}
	powArray(a)
	expect := [3]int{1, 2, 3} // 配列は値渡し
	if !reflect.DeepEqual(expect, a) {
		t.Errorf(valueExpectFormat, expect, a)
	}

	s := []int{1, 2, 3}
	powSlice(s)
	expectSlice := []int{1, 4, 9} // スライスは参照渡し
	if !reflect.DeepEqual(expectSlice, s) {
		t.Errorf(valueExpectFormat, expectSlice, s)
	}

	var (
		a2 [3]int
		s2 []int
	)
	expectArray := [3]int{0, 0, 0}
	expectSlice2 := []int(nil)
	if !reflect.DeepEqual(expectArray, a2) {
		t.Errorf(valueExpectFormat, expectArray, a2)
	}
	if !reflect.DeepEqual(expectSlice2, s2) {
		t.Errorf(valueExpectFormat, expectSlice2, s2)
	}

}

func TestSlicePitfall(t *testing.T) {
	a := [3]int{1, 2, 3}
	s := a[:]

	expectLen := 3
	expectCap := 3
	if expectLen != len(s) {
		t.Errorf(decimalExpectFormat, expectLen, len(s))
	}
	if expectCap != cap(s) {
		t.Errorf(decimalExpectFormat, expectCap, cap(s))
	}
	// この時点では配列とスライスの要素は同じ
	a[0] = 9
	// 配列に代入した値はスライスにも反映される
	if a[0] != s[0] {
		t.Errorf(decimalExpectFormat, a[0], s[0])
	}

	// スライスの自動拡張が起こるように要素を追加する
	s = append(s, 4, 5, 6)
	// 再度配列の要素を変更する
	a[0] = 99
	// 配列に代入した値はもちろん反映される
	if a[0] != 99 {
		t.Errorf("%d == %d", a[0], 99)
	}
	// スライスは別の参照領域を指すようになった為、9のまま
	if s[0] != 9 {
		t.Errorf(decimalExpectFormat, s[0], 9)
	}
	// この時点で配列とスライスは値が異なる
	if a[0] == s[0] {
		t.Errorf("%d == %d", a[0], s[0])
	}

}
