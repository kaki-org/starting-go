package main

import (
	"fmt"
	"math"

	"github.com/kakikubo/starting-go/zoo/animals"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
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

	// %vはさまざまな型のデータを埋め込む
	fmt.Printf("\n数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列=Golang 配列=[1 2 3]"
	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列="Golang" 配列=[3]int{1, 2, 3}"
	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=int 文字列=string 配列=[3]int"

	print("Hello, World!\n") // => "Hello, World!" を標準エラー出力へ出力
	println("Hello, World!") // => "Hello, World!\n" を標準エラー出力へ出力
	print(1, 2, 3, "\n")     // => "123\n" を標準エラー出力へ出力
	println(1, 2, 3)         // => "1 2 3\n" を標準エラー出力へ出力

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

	// 関数の戻り値を元に型推論を行って変数を定義、初期化
	n2 := one()
	fmt.Println(n2)

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

	// パッケージ変数n1の値を用いて計算
	n = n1 + 1
	fmt.Printf("n=%d\n", n)

	var_sample()
	var_sample3()
	var_sample4()
	var_sample5()
	var_sample6()
	wrap_around()

	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

	// 浮動小数点数
	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf, ninf, nan)

	fmt.Println(1.0e2)  // 1.0 * 10^2
	fmt.Println(1.0e+2) // 1.0 * 10^2
	fmt.Println(1.0e-2) // 1.0 / 10^2

	doubleValueSample()
	complexSomething()
	runeSomething()
	someString()
	someArray()
	someInterface()
	someMath()
	fmt.Println(plus(1, 2))
	hello()
	q, r := div(19, 7)
	fmt.Printf("商=%d 余り=%d\n", q, r) // => "商=2 余り=5" quotient, remainder
	q2, _ := div(19, 7)              // 余りは捨てる
	fmt.Printf("商=%d\n", q2)         // => "商=2"

	fmt.Println(doSomethingA())
	fmt.Println(doSomethingXY())
	fmt.Println(ignoreArgs(1, 2))

	// fmt.Println(RequiredFunction(1))

	// 無名関数
	fn := func(x, y int) int { return x + y }
	fmt.Println(fn(2, 3))
	fmt.Printf("%T\n", fn) // => "func(int, int) int"
	ClosureSample()

	// 定義済みの関数に別名をつけているかのような記述
	var plusAlias = plus
	fmt.Println(plusAlias(10, 20))

	// 関数を返す関数
	rfn := returnFunc()
	rfn()
	fmt.Printf("%T\n", rfn) // => "func()"
	returnFunc()()          // こうやっても実行できる

	// 関数を引数にとる関数
	callFunction(func() {
		fmt.Println("I'm a callFunction")
	})
}

func one() int {
	return 1
}

func var_sample() {
	a := 1
	// 以下を定義しているとエラーになってしまう
	// b := 2
	// c := 3
	// # command-line-arguments
	// ./main.go:93:2: b declared and not used
	// ./main.go:94:2: c declared and not used

	fmt.Println(a)
}

func var_sample2() {
	n4 := 9223372036854775807 // 符号付き64ビット整数で表現可能である最大値
	fmt.Println(n4)

	// var (
	// 	n1 int
	// 	n2 int64
	// )
	// n1 = 1
	// n2 = n1 // コンパイルエラー
	// fmt.Println(n1, n2)

}

func var_sample3() {
	// n := uint(17)
	n := 1
	b := byte(n)
	i64 := int64(n)
	u32 := uint32(n)
	fmt.Println(b, i64, u32)
	fmt.Printf("%T %T %T\n", b, i64, u32)
}

func var_sample4() {
	n := 256
	b := byte(n)
	fmt.Printf("b = %b\n", b)
}

func var_sample5() {
	b := byte(255)
	b = b + 1
	fmt.Println(b)
}

func var_sample6() {
	n := -1
	b := byte(n)
	fmt.Println(b)
}

func wrap_around() {
	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)
	if !doSomething(ui_1, ui_2) {
		fmt.Println("エラーが発生しました")
		return
	}

	sum := ui_1 + ui_2
	fmt.Printf("%d + %d = %d\n", ui_1, ui_2, sum)
	// 400000000 + 4000000000 = 105032704 オーバーフローして1億弱になってしまう
}

func doSomething(a, b uint32) bool {
	if (math.MaxInt32 - a) < b {
		return false
	} else {
		// チェック済みの為、問題なし
		sum := a + b
		fmt.Println(sum)
		return true
	}
}

func doubleValueSample() {
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

	f := 3.14
	n := int(f)
	fmt.Println(n)

	nf := -3.14
	nn := int(nf)
	fmt.Println(nn)
}

func complexSomething() {
	// 複素数型のサンプル
	c := 1.0 + 3i         // complex128型の変数cを定義して、1.0 + 3iを代入
	fmt.Println(c)        // 出力: (1+3i)
	c2 := complex(1.0, 3) // 別の定義方法
	fmt.Println(c2 == c)  // 出力: true
	// 複素数リテラル
	fmt.Println(0i)
	fmt.Println(11i)
	fmt.Println(0.i)
	fmt.Println(2.71828i)
	fmt.Println(6.67428e-11i)
	fmt.Println(1e6i)
	fmt.Println(.25i)
	fmt.Println(.12345e+5i)

	// 複素数の実部と虚部
	c3 := 1.3 + 4.2i
	fmt.Println(real(c3)) // real number (実数)
	fmt.Println(imag(c3)) // imaginary number (虚数)
}

func runeSomething() {
	r := '松'
	fmt.Printf("%v\n", r) // 出力: 26494
}

func someString() {
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
}

func someArray() {
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

	a10 := [...]int{1, 2, 3}
	a11 := [...]int{1, 2, 3, 4, 5}
	a12 := [...]int{}
	fmt.Printf("%v\n%v\n%v\n", a10, a11, a12)

	a5 := [...]int{1, 2, 3}
	a5[0] = 0
	a5[2] = 0
	fmt.Printf("%v\n", a5) // 出力: [0 2 0]
	// 以下はエラーになる
	// var (
	// 	a6 [3]int
	// 	a7 [5]int
	// )
	// a6 = a7 // ./main.go:317:7: cannot use a7 (variable of type [5]int) as type [3]int value in assignment(exit status 1)

	// 異なる型の配列は代入できません
	// var (
	// 	a8 [5]int
	// 	a9 [5]uint
	// )
	// a8 = a9 // ./main.go:325:7: cannot use a9 (variable of type [5]uint) as [5]int value in assignment (exit status 1)

	a13 := [3]int{1, 2, 3}
	a14 := [3]int{4, 5, 6}
	a13 = a14
	a13[0] = 0
	a13[2] = 0
	fmt.Printf("a13 = %v\n", a13) // a13の値は[0 5 0]
	fmt.Printf("a14 = %v\n", a14) // a14の値は[4 5 6]※a13とa14は別の配列
}

func someInterface() {
	var x interface{}
	fmt.Printf("%#v\n", x) // 出力: <nil>
	x = 1
	x = 3.14
	x = '山'
	x = "文字列"
	x = [...]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", x) // 出力: [5]uint8{0x1, 0x2, 0x3, 0x4, 0x5}

	// interfaceはすべての型の値を汎用的に表す手段である為、演算の対象としては利用できない
	// var xx, yy interface{}
	// xx, yy = 1, 2 // 代入
	// z := xx + yy  // 演算できない
}

func someMath() {
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
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello, World!")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

// 戻り値を表す変数
func doSomethingA() (a int) {
	return
	// 以下のコードと同じ
	// var a int
	// return a
}

func doSomethingXY() (x, y int) {
	y = 5
	return
	// 以下のコードと同じ
	// var x, y int
	// y = 5
	// return x, y
}

func ignoreArgs(_, _ int) int {
	return 1
}

// 型Tの定義
type T struct {
	value int
}

// インターフェース型I
type I interface {
	// 引数が2つ必要であると定義
	RequiredFunction(a, b int) int
}

// T型のインターフェースIを満たす関数(メソッド)
func (*T) RequiredFunction(a, _ int) int {
	// 実装に2番目の引数は不要
	return a
}

func ClosureSample() {
	var f func(int, int) int
	f = func(x, y int) int { return x + y }
	fmt.Println(f(1, 2))

	fmt.Printf("%#v\n", func(x, y int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}
