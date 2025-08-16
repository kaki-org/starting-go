package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/kakikubo/starting-go/zoo/animals"
)

func expect(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Expected %d, got %d", b, a)
	}
}

func expectEqual(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %d, got %d", a, b)
	}
}

func ExampleMain() {
	main()
	// Output:
	// Zoo Application
	// os.Exit
	// !
}

func ExampleDoSomethingDo() {
	DoSomethingDo()
	// Output:
	// DoSomething Do
}

func ExamplePrivatedoSomethingDo() {
	doSomethingDo()
	// Output:
	// doSomething Do
}

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

func TestValidateOverflow(t *testing.T) {
	expect(t, validateOverflow(1, 2), true)
	expect(t, validateOverflow(math.MaxInt32, 1), false)
}

func ExampleWrapAround() {
	ui1 := uint32(400000000)
	ui2 := uint32(4000000000)
	if !validateOverflow(ui1, ui2) {
		fmt.Println(math.MaxUint32)
		fmt.Println("エラーが発生しました")
		return
	}

	sum := ui1 + ui2
	fmt.Printf("%d + %d = %d\n", ui1, ui2, sum)
	// Output:
	// 4294967295
	// エラーが発生しました
}

func ExampleWrapAround2() {
	ui1 := uint32(400000000)
	ui2 := uint32(4000000000)
	// ここでvalidateOverflowを呼び出さないとおかしな事になる
	sum := ui1 + ui2
	fmt.Printf("%d + %d = %d\n", ui1, ui2, sum)
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

const valueFormat = "value = %v\n"

func ExampleDouble() {
	fmt.Printf(valueFormat, 1.0000000000000000)
	fmt.Printf(valueFormat, 1.0000000000000001)
	fmt.Printf(valueFormat, 1.0000000000000002)
	fmt.Printf(valueFormat, 1.0000000000000003)
	fmt.Printf(valueFormat, 1.0000000000000004)
	fmt.Printf(valueFormat, 1.0000000000000005)
	fmt.Printf(valueFormat, 1.0000000000000006)
	fmt.Printf(valueFormat, 1.0000000000000007)
	fmt.Printf(valueFormat, 1.0000000000000008)
	fmt.Printf(valueFormat, 1.0000000000000009)
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
	fmt.Printf(valueFormat, float32(1.0000000000000000))
	fmt.Printf(valueFormat, float32(1.0000000000000001))
	fmt.Printf(valueFormat, float32(1.0000000000000002))
	fmt.Printf(valueFormat, float32(1.0000000000000003))
	fmt.Printf(valueFormat, float32(1.0000000000000004))
	fmt.Printf(valueFormat, float32(1.0000000000000005))
	fmt.Printf(valueFormat, float32(1.0000000000000006))
	fmt.Printf(valueFormat, float32(1.0000000000000007))
	fmt.Printf(valueFormat, float32(1.0000000000000008))
	fmt.Printf(valueFormat, float32(1.0000000000000009))
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

func ExampleHello() {
	hello()
	// Output:
	// Hello, World!
}

func TestPlus(t *testing.T) {
	expect(t, plus(1, 2), 3)
}

func ExampleDiv() {
	q, r := div(19, 7)
	fmt.Printf("商=%d 余り=%d\n", q, r) // => "商=2 余り=5" quotient, remainder
	q2, _ := div(19, 7)              // 余りは捨てる
	fmt.Printf("商=%d\n", q2)         // => "商=2"
	// Output:
	// 商=2 余り=5
	// 商=2
}

func TestDoSomethingA(t *testing.T) {
	expect(t, doSomethingA(), 0)
}

func TestDoSomethingXY(t *testing.T) {
	x, y := doSomethingXY()
	expect(t, x, 0)
	expect(t, y, 5)
}

func TestIgnoreArgs(t *testing.T) {
	expect(t, ignoreArgs(10, 20), 1)
}

func TestAnonymousFunc(t *testing.T) {
	fn := func(x, y int) int { return x + y }
	expect(t, fn(1, 2), 3)
	output := fmt.Sprintf("%T\n", fn) // => "func(int, int) int"
	expect(t, output, "func(int, int) int\n")
}

func TestClosure1(t *testing.T) {
	var f func(int, int) int
	f = func(x, y int) int { return x + y }
	output := fmt.Sprintln(f(1, 2))
	expect(t, output, "3\n")
}

func TestClosure2(t *testing.T) {
	output := fmt.Sprintf("ClosureSample %#v\n", func(x, y int) int { return x + y }(2, 3))
	expect(t, output, "ClosureSample 5\n")
}

func TestPlusAlias(t *testing.T) {
	// 定義済みの関数に別名をつけているかのような記述
	var plusAlias = plus
	expect(t, plusAlias(1, 2), 3)
}

func TestReturnFunc(t *testing.T) {
	// 関数を返す関数
	f := ReturnFunc()
	output := fmt.Sprintf("%T\n", f) // => "func()"
	expect(t, output, "func()\n")
}

func ExampleReturnFunc() {
	f := ReturnFunc()
	f()
	ReturnFunc()() // こうやっても実行できる
	// Output:
	// I'm a function
	// I'm a function
}

func ExampleCallFunction() {
	// 関数を引数にとる関数
	CallFunction(func() {
		fmt.Println("I'm a callFunction")
	})
	// Output:
	// I'm a callFunction
}

func ExampleLater() {
	// クロージャ(関数閉包)
	f := later()
	fmt.Println(f("Golang"))   // => ""
	fmt.Println(f("is"))       // => "Golang"
	fmt.Println(f("awesome!")) // => "is"
	fmt.Println(f("awesome!")) // => "awesome!"
	// Output:
	// Golang
	// is
	// awesome!
}

func TestClosureIntegers(t *testing.T) {
	// クロージャを利用してGeneratorを実装する
	ints := integers()
	expect(t, ints(), 1)
	expect(t, ints(), 2)
	expect(t, ints(), 3)
	otherInts := integers()
	expect(t, otherInts(), 1) // 1 (otherIntsの状態は別)
}

func TestOnetwo(t *testing.T) {
	// 関数内での定数宣言
	x, y := onetwo()
	expect(t, x, 1)
	expect(t, y, 2)
	expect(t, ONE, 1)
	// expect(t, TWO, 2) // これは参照できない
}

func ExampleConstant() {
	// 定数定義(値の省略)
	const (
		XX = 10
		YY
		ZZ
		S1 = "あ"
		S2
	)
	fmt.Println(XX, YY, ZZ, S1, S2) // => "10 10 10 あ あ"
	// Output:
	// 10 10 10 あ あ
}

func ExampleConstant1() {
	// 定数
	// const X = 1
	const (
		X = 1
		Y = 2
		Z = 3
	)
	fmt.Println(X, Y, Z) // => "1 2 3"
	// Output:
	// 1 2 3
}

func ExampleConstant2() {
	// 以前はこの書き方でもOKだったみたい(X1, X2が定義されていない状態でX12を定義している)
	// const (
	// 	X12 = X1 + X2
	// 	X1 = 1
	// 	X2 = 2
	// )

	const (
		X = 2
		Y = 7
		Z = X + Y // Z = 9

		S1 = "今日"
		S2 = "晴れ"
		S  = S1 + "は" + S2 // S = "今日は晴れ"
	)

	fmt.Println(Z, S)
	fmt.Printf("%T\n", Z) // => "int"
	// Output:
	// 9 今日は晴れ
	// int
}

func ExampleConstant3() {
	// 以下はオーバーフローする
	/*
		const (
			N = 999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
		)
		n999 := N
		fmt.Printf("%T %v\n", n999, n999) // => "int64 999999
	*/

	// 型が違う場合はコンパイルエラーにしてくれる
	/*
		const (
			UI64 = uint64(12345)
		)

		var i64 int64
		i64 = UI64 // cannot use UI64 (constant 12345 of type uint64) as int64 value in assignment
	*/
	const (
		// I64 int64   = -1
		// F64 float64 = 1.2
		I64 = int64(-1)
		F64 = float64(1.2)
	)
	fmt.Printf("%T %v\n%T %v\n", I64, I64, F64, F64) // => "int64 -1\nfloat64 1.2"
	// Output:
	// int64 -1
	// float64 1.2
}

func ExampleConstant4() {
	// コンパイル時に演算が処理される為、以下はコンパイルエラーにならない
	const (
		// uint64の最大値に1を足した値
		MAXUI64PLUS1 = math.MaxUint64 + 1
	)
	muint64 := uint64(MAXUI64PLUS1 - 1)     // コンパイル時に値が決定される為、18446744073709551615(uint64の最大値)になる
	fmt.Printf("%T %v\n", muint64, muint64) // => "uint64 18446744073709551615"
	// Output:
	// uint64 18446744073709551615
}

func ExampleConstant5() {
	// 浮動小数点数の定数
	const (
		Pi = 3.14
	)
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	fmt.Printf("%v\n", f32)
	fmt.Printf("%v\n", f64)
	const F = 1.0000000000001
	fmt.Println(float64(F) * 10000)
	fmt.Println(F * 10000)
	// Output:
	// 3.1415927
	// 3.141592653589793
	// 10000.000000000999
	// 10000.000000001
}

func ExampleConstantComplex() {
	// 複素数の定数
	const (
		C = 4.7 + 1.3i
	)
	fmt.Printf("%T %v\n", C, C) // => "complex128 (4.7+1.3i)"
	// Output:
	// complex128 (4.7+1.3i)
}

func ExampleConstantRune() {
	// ルーン、文字列の定数
	const (
		R  = 'あ'
		S  = "Go言語"
		RS = `秋の田のかりほの庵の苫をあらみ
わが衣手は露にぬれつつ`
	)
	fmt.Printf("%v\n", R)  // => "12354"
	fmt.Printf("%v\n", S)  // => "Go言語"
	fmt.Printf("%v\n", RS) // => "秋の田のかりほの庵の苫をあらみ\nわが衣手は露にぬれつつ"
	// Output:
	// 12354
	// Go言語
	// 秋の田のかりほの庵の苫をあらみ
	// わが衣手は露にぬれつつ
}

func ExampleConstantIota() {
	// 文字と認められていない〒(記号とされている)を使ってみる
	// const 〒 = "郵便番号" // 〒は記号として扱われる。エラー
	// iota
	const (
		A1 = iota + 1
		B1
		C1
		N = iota
	)
	// iotaは定数の宣言ごとにリセットされる
	const (
		A2 = iota
		B2
		C2
	)
	fmt.Println(A1, B1, C1, N) // => "1 2 3 3"
	fmt.Println(A2, B2, C2)    // => "0 1 2"
	// やらないけど言語仕様上は以下のようにもかける
	const (
		朝の挨拶 = "おはよう"
		昼の挨拶 = "こんにちは"
		夜の挨拶 = "こんばんは"
	)
	あいさつ(朝の挨拶) // => "おはよう"
	あいさつ(昼の挨拶) // => "こんにちは"
	あいさつ(夜の挨拶) // => "こんばんは"
	// Output:
	// 1 2 3 3
	// 0 1 2
	// おはよう
	// こんにちは
	// こんばんは
}

func TestAnimals(t *testing.T) {
	expect(t, animals.ElephantFeed(), "Grass")
	expect(t, animals.MonkeyFeed(), "Banana")
	expect(t, animals.RabbitFeed(), "Carrot")
	expect(t, animals.MAX, 100)
	expect(t, animals.FooFunc(5), 6)
	// fmt.Println(animals.internal_const) // コンパイルエラー
	// fmt.Println(animals.internalFunc(5)) // コンパイルエラー
}

func ExampleCondition() {
	for {
		fmt.Println("loop")
		break // このbreakがないと無限ループになる
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	x, y := 3, 5
	if n := x * y; n%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	/* 変数iが100になるまでループ */
	i := 0
	for i < 10 {
		// continue使うとdlvが停止してしまうのでコメントアウト
		// if i == 5 {
		// 	continue
		// }
		fmt.Println(i)
		i++
	}
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	/* rangeを伴うfor */
	for i, s := range fruits {
		// iはインデックス、sは要素(値)を取る
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
	/* 文字列の場合はrune型になる */
	for i, r := range "あいう" {
		fmt.Printf("[%d]=%d %v\n", i, r, string(r))
	}
	// Output:
	// loop
	// 0
	// 1
	// 2
	// 奇数
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// fruits[0]=Apple
	// fruits[1]=Banana
	// fruits[2]=Cherry
	// [0]=12354 あ
	// [3]=12356 い
	// [6]=12358 う
}

func ExampleSwitch() {
	n := 5
	// golangのswitchはbreakが不要(フォールスルーしない)
	switch n := 3; n { // ここで定義したnはswitch内でのみ有効
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4です")
		fallthrough // フォールスルーさせる場合は明示的に記述する
	// case "3": // 型が違うとコンパイルエラー
	// 	fmt.Println("3")
	case 5.0: // 整数5と互換性がある為エラーにならない
		fmt.Println("5.0です")
	case 6 + 0i: // 整数6と複素数6+0iは同じである為エラーにならない
		fmt.Println("6+0i")
	default:
		fmt.Println("unknown")
	}
	fmt.Printf("結局nは%d\n", n) // ここで定義したnはswitch外でも有効
	/* caseに定数と式が混在するswitch文はコンパイルエラー(正確には型が異なるとエラー) */
	// switch x := 1; x {
	// case 1, 2, 3:
	// 	fmt.Println(x)
	// case x > 3:
	// 	fmt.Println("x > 3")
	// }
	nn := 4
	switch {
	case nn > 0 && nn < 3:
		fmt.Println("0 < nn < 3")
		fallthrough
	case nn > 3 && nn < 6:
		fmt.Println("3 < nn < 6")
		fallthrough
	case nn > 0 && nn < 6:
		fmt.Println("0 < nn < 6")
	}
	// Output:
	// 3 or 4です
	// 5.0です
	// 結局nは5
	// 3 < nn < 6
	// 0 < nn < 6
}

func ExampleTypeAssertion() {
	anything(1)
	anything(3.14)
	anything(4 + 5i)
	anything('海')
	anything("日本語")
	anything([...]int{1, 2, 3, 4, 5})

	var x interface{} = 3
	i := x.(int)
	// f := x.(float64) // panic: interface conversion: interface {} is int, not float64
	fmt.Println(i)

	var y interface{} = 3.14

	ii, isInt := y.(int)
	ff, isFloat64 := y.(float64)
	s, isString := y.(string)
	fmt.Println(ii, isInt)     // 0 false
	fmt.Println(ff, isFloat64) // 3.14 true
	fmt.Println(s, isString)   //  false
	// Output:
	// 1
	// 3.14
	// (4+5i)
	// 28023
	// 日本語
	// [1 2 3 4 5]
	// 3
	// 0 false
	// 3.14 true
	//  false
}

func ExampleTypeAsertion2() {
	// 2つの値を返す型アサーションの結果で条件分岐
	var x interface{} = "abc"
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("int %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Printf("string %s\n", s)
	} else {
		fmt.Println("Unsupported type!")
	}
	// Output:
	// string abc
}

func ExampleTypeSwitch() {
	var x interface{} = 3

	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}
	// Output:
	// integer or unsigned integer
}

func ExampleTypeSwitch2() {
	var x interface{} = 3

	switch v := x.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println(v)
	default:
		fmt.Printf("%#v\n", v)
	}
	// Output:
	// 9
}

func ExampleGoto() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
L: /* ラベル */
	fmt.Println("C")
	// Output:
	// A
	// C
}

func ExampleLabel() {
LOOP:
	for {
		for {
			for {
				fmt.Println("START")
				break LOOP
			}
			fmt.Println("ここは通らない")
		}
		fmt.Println("ここは通らない")
	}
	fmt.Println("END")
	// Output:
	// START
	// END
}

func ExampleLabel2() {
L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("ここは処理されない")
	}
	// Output:
	// 1 * 1 = 1
	// 2 * 1 = 2
	// 3 * 1 = 3
}

func ExampleDefer() {
	/* deferに登録された式は関数の終了時に評価される。実行順序は最後に登録したものから先に実行される */
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}() // deferに登録する関数を即時実行する場合は()を付ける
	fmt.Println("done")
	// Output:
	// done
	// A
	// B
	// C
	// defer3
	// defer2
	// defer1
}

func TestRunPanic(t *testing.T) {
	defer func() {
		err := recover()
		if err != "runtime error" {
			t.Errorf("unexpected error: %v", err)
		}
	}()
	runPanic()
}

func ExampleRunRecover() {
	runRecover()
	// Output:
	// recover: runtime error
}

func ExampleTestRecover() {
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
	// Output:
	// panic: int=128
	// panic: string=hogehoge
	// panic: unknown=[1 2 3]
}

func TestSalutations(t *testing.T) {
	Salutations()
}

func TestArray0(t *testing.T) {
	Array0()
}

func TestTypeInference(t *testing.T) {
	TypeInference()
}

func TestDefineVariables(t *testing.T) {
	DefineVariables()
}

func TestPackageVariables(t *testing.T) {
	PackageVariables()
}

func TestVar1(t *testing.T) {
	Var1()
}

func TestVar2(t *testing.T) {
	Var2()
}

func TestVar3(t *testing.T) {
	Var3()
}

func TestFloat(t *testing.T) {
	Float()
}

func TestConstant(t *testing.T) {
	Constant()
}

func TestDefer(t *testing.T) {
	Defer()
}
