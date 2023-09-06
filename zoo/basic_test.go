package main

import (
	"fmt"
	"math"
)

func ExampleArray() {
	// コメントです。
	a := [3]string{
		"Yamada Taro",
		"Sato Hanako",
		"Suzuki Kenji",
	}
	fmt.Println(a[0], a[1], a[2])
	// 10進数の形式で数値5を%dの箇所へ埋め込む
	fmt.Printf("%d\n", 5) // => "数値=5"
	// 数値用の書式いろいろ
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17) // => "10進数=17 2進数=10001 8進数=21 16進数=11"
	// 埋め込むパラメータが足りない
	// fmt.Printf("%d年%d月%d日\n", 2015, 7) // => "2015年7月%!d(MISSING)日"
	// 埋め込むパラメータが過剰
	// fmt.Printf("%d年%d月%d日\n", 2015, 7, 15, 23)
	/* => "2015年7月15日
	%!(EXTRA int=23)" */

	// Output:
	// Yamada Taro Sato Hanako Suzuki Kenji
	// 5
	// 10進数=17 2進数=10001 8進数=21 16進数=11
}

func ExampleArray2() {
	// %vはさまざまな型のデータを埋め込む
	fmt.Printf("\n数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列=Golang 配列=[1 2 3]"
	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列="Golang" 配列=[3]int{1, 2, 3}"
	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=int 文字列=string 配列=[3]int"
	// Output:
	// 数値=5 文字列=Golang 配列=[1 2 3]
	// 数値=5 文字列="Golang" 配列=[3]int{1, 2, 3}
	// 数値=int 文字列=string 配列=[3]int
}

// TODO: 標準エラー出力をどうするのか
func ExamplePrint() {
	print("Hello, World!\n") // => "Hello, World!" を標準エラー出力へ出力
	println("Hello, World!") // => "Hello, World!\n" を標準エラー出力へ出力
	print(1, 2, 3, "\n")     // => "123\n" を標準エラー出力へ出力
	println(1, 2, 3)         // => "1 2 3\n" を標準エラー出力へ出力
}

func ExampleTypeInference() {
	var n int
	n = 5
	// n == 5
	// n = "string value" // コンパイルエラー
	fmt.Println(n)
	// int型の変数iを定義して1を代入(型推論)
	i := 1
	fmt.Println(i)
	// bool型の変数bを定義して真偽値trueを代入(型推論)
	b := true
	fmt.Println(b)
	// float64型の変数fを定義して浮動小数点数3.14を代入(型推論)
	f := 3.14
	fmt.Println(f)
	// string型の変数sを定義して文字列"abc"を代入(型推論)
	s := "abc"
	fmt.Println(s)
	// Output:
	// 5
	// 1
	// true
	// 3.14
	// abc
}

func ExampleTypeInference2() {
	// 関数の戻り値を元に型推論を行って変数を定義、初期化
	n2 := one()
	fmt.Println(n2)
	// Output:
	// 1
}

func ExampleDefineVariables() {
	// varで変数定義をまとめる書き方
	var (
		n3 = 1
		s3 = "string value"
		b3 = true
	)
	// 暗黙的な定義を並べる書き方だと以下
	// n3 := 1
	// s3 := "string value"
	// b3 := true
	fmt.Println(n3, s3, b3)
	// Output:
	// 1 string value true
}

func ExamplePackageVariables() {
	// パッケージ変数n1の値を用いて計算
	n := n1 + 1
	fmt.Printf("n=%d\n", n)
	// Output:
	// n=101
}

func ExampleVar1() {
	a := 1
	// 以下を定義しているとエラーになってしまう
	// b := 2
	// c := 3
	// # command-line-arguments
	// ./main.go:93:2: b declared and not used
	// ./main.go:94:2: c declared and not used
	fmt.Println(a)
	// Output:
	// 1
}

func ExampleVar2() {
	n4 := 9223372036854775807 // 符号付き64ビット整数で表現可能である最大値
	// var (
	// 	n1 int
	// 	n2 int64
	// )
	// n1 = 1
	// n2 = n1 // コンパイルエラー
	// fmt.Println(n1, n2)
	fmt.Println(n4)
	// Output:
	// 9223372036854775807
}

func ExampleVar3() {
	// n := uint(17)
	n := 1
	b := byte(n)
	i64 := int64(n)
	u32 := uint32(n)
	fmt.Println(b, i64, u32)
	fmt.Printf("%T %T %T\n", b, i64, u32)
	// Output:
	// 1 1 1
	// uint8 int64 uint32
}

func ExampleVar4() {
	// byte型に256を代入するとラップアラウンドして0になる
	n := 256
	b := byte(n)
	fmt.Printf("b = %b\n", b)
	// Output:
	// b = 0
}

func ExampleVar5() {
	// byte型255に1を足すとラップアラウンドして0になる
	b := byte(255)
	b = b + 1
	fmt.Println(b)
	// Output:
	// 0
}

func ExampleVar6() {
	// byte型-1は255になる
	n := -1
	b := byte(n)
	fmt.Println(b)
	// Output:
	// 255
}

func ExampleMaxUint32() {
	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
	// Output:
	// uint32 max value = 4294967295
}

func ExampleWrapAround() {
	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)
	if !validateOverflow(ui_1, ui_2) {
		fmt.Println(math.MaxUint32)
		fmt.Println("エラーが発生しました")
		return
	}

	sum := ui_1 + ui_2
	fmt.Printf("%d + %d = %d\n", ui_1, ui_2, sum)
	// Output:
	// 4294967295
	// エラーが発生しました
}

func ExampleWrapAround2() {
	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)
	// ここでvalidateOverflowを呼び出さないとおかしな事になる
	sum := ui_1 + ui_2
	fmt.Printf("%d + %d = %d\n", ui_1, ui_2, sum)
	// 400000000 + 4000000000 = 105032704 オーバーフローして1億弱になってしまう
	// Output:
	// 400000000 + 4000000000 = 105032704
}

func ExampleFloat() {
	// 浮動小数点数
	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf, ninf, nan)

	fmt.Println(1.0e2)  // 1.0 * 10^2
	fmt.Println(1.0e+2) // 1.0 * 10^2
	fmt.Println(1.0e-2) // 1.0 / 10^2
	// Output:
	// +Inf -Inf NaN
	// 100
	// 100
	// 0.01
}

func ExampleDouble() {
	fmt.Printf("value = %v\n", 1.0000000000000000)
	fmt.Printf("value = %v\n", 1.0000000000000001)
	fmt.Printf("value = %v\n", 1.0000000000000002)
	fmt.Printf("value = %v\n", 1.0000000000000003)
	fmt.Printf("value = %v\n", 1.0000000000000004)
	fmt.Printf("value = %v\n", 1.0000000000000005)
	fmt.Printf("value = %v\n", 1.0000000000000006)
	fmt.Printf("value = %v\n", 1.0000000000000007)
	fmt.Printf("value = %v\n", 1.0000000000000008)
	fmt.Printf("value = %v\n", 1.0000000000000009)
	// Output:
	// value = 1
	// value = 1
	// value = 1.0000000000000002
	// value = 1.0000000000000002
	// value = 1.0000000000000004
	// value = 1.0000000000000004
	// value = 1.0000000000000007
	// value = 1.0000000000000007
	// value = 1.0000000000000009
	// value = 1.0000000000000009
}

func ExampleDouble2() {
	f := 3.14
	n := int(f)
	fmt.Println(n)
	nf := -3.14
	nn := int(nf)
	fmt.Println(nn)
	// Output:
	// 3
	// -3
}

func ExampleDoubleCast() {
	fmt.Printf("value = %v\n", float32(1.0000000000000000))
	fmt.Printf("value = %v\n", float32(1.0000000000000001))
	fmt.Printf("value = %v\n", float32(1.0000000000000002))
	fmt.Printf("value = %v\n", float32(1.0000000000000003))
	fmt.Printf("value = %v\n", float32(1.0000000000000004))
	fmt.Printf("value = %v\n", float32(1.0000000000000005))
	fmt.Printf("value = %v\n", float32(1.0000000000000006))
	fmt.Printf("value = %v\n", float32(1.0000000000000007))
	fmt.Printf("value = %v\n", float32(1.0000000000000008))
	fmt.Printf("value = %v\n", float32(1.0000000000000009))
	fmt.Println(float32(1.0) / float32(3.0))
	fmt.Println(float64(1.0) / float64(3.0))
	// Output:
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// value = 1
	// 0.33333334
	// 0.3333333333333333
}

func ExampleFmain() {
	fmain()
	// Output:
	// Hello, World! fmain
}
