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

func anything(a interface{}) {
	fmt.Println(a)
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
