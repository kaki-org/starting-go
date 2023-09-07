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

func ExampleArray0() {
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

func ExampleComplex1() {
	// 複素数型のサンプル
	c := 1.0 + 3i         // complex128型の変数cを定義して、1.0 + 3iを代入
	fmt.Println(c)        // 出力: (1+3i)
	c2 := complex(1.0, 3) // 別の定義方法
	fmt.Println(c2 == c)  // 出力: true
	// Output:
	// (1+3i)
	// true
}
func ExampleComplex2() {
	// 複素数リテラル
	fmt.Println(0i)
	fmt.Println(11i)
	fmt.Println(0.i)
	fmt.Println(2.71828i)
	fmt.Println(6.67428e-11i)
	fmt.Println(1e6i)
	fmt.Println(.25i)
	fmt.Println(.12345e+5i)
	// Output:
	// (0+0i)
	// (0+11i)
	// (0+0i)
	// (0+2.71828i)
	// (0+6.67428e-11i)
	// (0+1e+06i)
	// (0+0.25i)
	// (0+12345i)
}

func ExampleComplex3() {
	// 複素数の実部と虚部
	c3 := 1.3 + 4.2i
	fmt.Println(real(c3)) // real number (実数)
	fmt.Println(imag(c3)) // imaginary number (虚数)
	// Output:
	// 1.3
	// 4.2
}

func ExampleRune() {
	r := '松'
	fmt.Printf("%v\n", r) // 出力: 26494
	// Output:
	// 26494
}

func ExampleString() {
	s := "Goの文字列"
	fmt.Printf("%v\n", s) // 出力: Goの文字列

	s2 := `
GOの
RAW文字列リテラルによる
複数行に渡る
文字列
`
	fmt.Printf("%v\n", s2)
	s3 := `abc`
	fmt.Printf("%v\n", s3)
	s4 := `\n
\n`
	fmt.Printf("%v\n", s4)
	// Output:
	// Goの文字列
	//
	// GOの
	// RAW文字列リテラルによる
	// 複数行に渡る
	// 文字列
	//
	// abc
	// \n
	// \n
}
func ExampleArray1() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", a[0]) // 出力: 1
	a1 := [5]int{}
	fmt.Printf("%v\n", a1)
	a2 := [5]int{1, 2, 3}
	fmt.Printf("%v\n", a2)
	// a3 := [5]int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("%v\n", a3) // ./main.go:281:30: index 5 is out of bounds (>= 5)
	var a3 [5]int
	a4 := [5]int{}
	fmt.Printf("%v\n", a3 == a4) // 出力: true
	// Output:
	// [1 2 3 4 5]
	// 1
	// [0 0 0 0 0]
	// [1 2 3 0 0]
	// true
}

func ExampleArray2() {
	ia := [3]int{}
	fmt.Printf("int array %v\n", ia) // 出力: [0 0 0]
	ua := [3]uint{}
	fmt.Printf("unsigned int array %v\n", ua) // 出力: [0 0 0]
	ba := [3]bool{}
	fmt.Printf("bool array %v\n", ba) // 出力: [false false false]
	fa := [3]float64{}
	fmt.Printf("float64 array %v\n", fa) // 出力: [0 0 0]
	ca := [3]complex128{}
	fmt.Printf("complex128 array %v\n", ca) // 出力: [(0+0i) (0+0i) (0+0i)]
	ra := [3]rune{}
	fmt.Printf("rune array %v\n", ra) // 出力: [0 0 0]
	sa := [3]string{}
	fmt.Printf("string array %v\n", sa) // 出力: [  ]
	zeroa := [0]int{}
	fmt.Printf("zero array %v\n", zeroa) // 出力: []
	// Output:
	// int array [0 0 0]
	// unsigned int array [0 0 0]
	// bool array [false false false]
	// float64 array [0 0 0]
	// complex128 array [(0+0i) (0+0i) (0+0i)]
	// rune array [0 0 0]
	// string array [  ]
	// zero array []
}

func ExampleArray3() {
	// 以下はエラーになる
	// var (
	// 	a6 [3]int
	// 	a7 [5]int
	// )
	// a6 = a7 // ./main.go:317:7: cannot use a7 (variable of type [5]int) as type [3]int value in assignment(exit status 1)
	a10 := [...]int{1, 2, 3}
	a11 := [...]int{1, 2, 3, 4, 5}
	a12 := [...]int{}
	fmt.Printf("%v\n%v\n%v\n", a10, a11, a12)
	a5 := [...]int{1, 2, 3}
	a5[0] = 0
	a5[2] = 0
	fmt.Printf("%v\n", a5) // 出力: [0 2 0]
	// Output:
	// [1 2 3]
	// [1 2 3 4 5]
	// []
	// [0 2 0]
}

func ExampleArray4() {
	// 異なる型の配列は代入できません
	// var (
	// 	a8 [5]int
	// 	a9 [5]uint
	// )
	// a8 = a9 // ./main.go:325:7: cannot use a9 (variable of type [5]uint) as [5]int value in assignment (exit status 1)
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}
	a1 = a2
	a1[0] = 0
	a1[2] = 0
	fmt.Printf("a1 = %v\n", a1) // a13の値は[0 5 0]
	fmt.Printf("a2 = %v\n", a2) // a14の値は[4 5 6]※a13とa14は別の配列
	// Output:
	// a1 = [0 5 0]
	// a2 = [4 5 6]
}

func ExampleInterface() {
	// interfaceはすべての型の値を汎用的に表す手段である為、演算の対象としては利用できない
	// var xx, yy interface{}
	// xx, yy = 1, 2 // 代入
	// z := xx + yy  // 演算できない
	var x interface{}
	fmt.Printf("%#v\n", x) // 出力: <nil>
	x = 1
	x = 3.14
	x = '山'
	x = "文字列"
	x = [...]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", x) // 出力: [5]uint8{0x1, 0x2, 0x3, 0x4, 0x5}
	// Output:
	// <nil>
	// [5]uint8{0x1, 0x2, 0x3, 0x4, 0x5}
}

func ExampleMath() {
	var n int
	s := "Go言語"
	x := 10
	n += 5
	fmt.Println(n)
	s += "の解説"
	fmt.Println(s)
	n *= 10
	fmt.Println(n)
	n &= x
	fmt.Println(n)
	// Output:
	// 5
	// Go言語の解説
	// 50
	// 2
}
