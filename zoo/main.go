package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */

	// fmt.Println(RequiredFunction(1))
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

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunction(f func()) {
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
