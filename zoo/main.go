package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"

	"github.com/kakikubo/starting-go/zoo/animals"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())

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

	// クロージャ(関数閉包)
	later := later()
	fmt.Println(later("Golang"))
	fmt.Println(later("is"))
	fmt.Println(later("awesome!"))
	fmt.Println(later("awesome!"))

	// クロージャを利用してGeneratorを実装する
	ints := integers()

	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3

	otherInts := integers()
	fmt.Println(otherInts()) // 1 (otherIntsの状態は別)

	// 定数
	// const X = 1
	const (
		X = 1
		Y = 2
		Z = 3
	)

	// 関数内での定数宣言
	x, y := onetwo()
	fmt.Printf("x=%d, y=%d\n", x, y) // => "x=1, y=2"

	// 定数定義(値の省略)
	const (
		XX = 10
		YY
		ZZ
		S1 = "あ"
		S2
	)
	fmt.Println(XX, YY, ZZ, S1, S2) // => "10 10 10 あ あ"

	const (
		XXX = 2
		YYY = 7
		ZZZ = XXX + YYY // ZZZ = 9

		S3  = "今日"
		S4  = "晴れ"
		S34 = S3 + "は" + S4 // S = "今日は晴れ"
	)

	fmt.Println(ZZZ, S34)
	fmt.Printf("%T\n", ZZZ) // => "int"
	// 以前はこの書き方でもOKだったみたい(X1, X2が定義されていない状態でX12を定義している)
	// const (
	// 	X12 = X1 + X2
	// 	X1 = 1
	// 	X2 = 2
	// )

	const (
		// I64 int64   = -1
		// F64 float64 = 1.2
		I64 = int64(-1)
		F64 = float64(1.2)
	)
	fmt.Printf("%T %v\n%T %v\n", I64, I64, F64, F64) // => "int64 -1\nfloat64 1.2"

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

	// コンパイル時に演算が処理される為、以下はコンパイルエラーにならない
	const (
		// uint64の最大値に1を足した値
		MAXUI64PLUS1 = math.MaxUint64 + 1
	)
	muint64 := uint64(MAXUI64PLUS1 - 1)     // コンパイル時に値が決定される為、18446744073709551615(uint64の最大値)になる
	fmt.Printf("%T %v\n", muint64, muint64) // => "uint64 18446744073709551615"

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

	// 複素数の定数
	const (
		C = 4.7 + 1.3i
	)
	fmt.Printf("%T %v\n", C, C) // => "complex128 (4.7+1.3i)"

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
	あいさつ(昼の挨拶) // => "こんにちは"

	// 文字と認められていない〒(記号とされている)を使ってみる
	// const 〒 = "郵便番号"

	fmt.Println(animals.MAX)
	// fmt.Println(animals.internal_const) // コンパイルエラー

	fmt.Println(animals.FooFunc(5)) // => "6"
	// fmt.Println(animals.internalFunc(5)) // コンパイルエラー
	fmain()

	someCondition()
	funcSwitch()
	typeAssertion()
	typeAssertion2()
	typeSwitch()
	typeSwitch2()
	gotoSample()
	labelSample()
	labelSample2()
	runDefer()
	// runPanic()
	runRecover()

	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
	defer fmt.Println("!")

	fmt.Println("os.Exit")
	os.Exit(0)
}

func one() int {
	return 1
}

func validateOverflow(a, b uint32) bool {
	if (math.MaxInt32 - a) < b {
		// オーバーフローするのでfalse
		return false
	} else {
		// チェック済みの為、問題なし
		sum := a + b
		fmt.Println(sum)
		return true
	}
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

	// fmt.Printf("ClosureSample %#v\n", func(x, y int) int { return x + y })
	fmt.Printf("ClosureSample %#v\n", func(x, y int) int { return x + y }(2, 3))
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	// 1つ前に与えられた文字列を保存する変数
	var store string
	// 引数に文字列を取り、文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

const ONE = 1

func onetwo() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func あいさつ(m string) {
	fmt.Println(m)
}

/* パッケージ変数 */
var (
	m = 256 // パッケージ内のみで参照出来る変数
	N = 512 // 公開される変数
)

/* 公開される関数 */
func DoSomethingDo() {
	fmt.Println("DoSomething Do")
}

/* パッケージ内のみで参照できる関数 */
func doSomethingDo() {
	fmt.Println("doSomething Do")
}

func someCondition() {
	for {
		fmt.Println("loop")
		break // このbreakがないと無限ループになる
	}
	for i := 0; i < 10; i++ {
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
	for i < 100 {
		// continue使うとdlvが停止してしまうのでコメントアウト
		// if i == 50 {
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

}

func funcSwitch() {
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

	nn := 4
	switch {
	case nn > 0 && nn < 3:
		fmt.Println("0 < nn < 3")
	case nn > 3 && nn < 6:
		fmt.Println("3 < nn < 6")
	}

	/* caseに定数と式が混在するswitch文はコンパイルエラー(正確には型が異なるとエラー) */
	// switch x := 1; x {
	// case 1, 2, 3:
	// 	fmt.Println(x)
	// case x > 3:
	// 	fmt.Println("x > 3")
	// }
}

func typeAssertion() {
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

}

func typeAssertion2() {
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
}

func anything(a interface{}) {
	fmt.Println(a)
}

func typeSwitch() {
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
}

func typeSwitch2() {
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
}

func gotoSample() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
L: /* ラベル */
	fmt.Println("C")
}

func labelSample() {
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
}

func labelSample2() {
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
}

func runDefer() {
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
}

func runPanic() {
	/* panic時でもdeferは実行される */
	defer fmt.Println("defer on runPanic")

	panic("runtime error") // ここでエラー終了
	fmt.Println("Hello, World!")
}

func runRecover() {
	/* recoverはpanic時に発生したエラー情報を取得する */
	defer func() {
		if x := recover(); x != nil {
			/* 変数xはpanicに渡されたinterface{}型の値 */
			fmt.Println("recover:", x)
		}
	}()
	panic("runtime error")
	/* これは実行されない */
	fmt.Println("Hello, World!")
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			/* panicによるinterface{}型の値に応じて処理を分岐 */
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Printf("panic: unknown=%v\n", v)
			}
		}
	}()
	panic(src)
	return
}

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
