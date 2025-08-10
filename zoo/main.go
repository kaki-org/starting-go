package main

import (
	"fmt"
	"math"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */

	// fmt.Println(RequiredFunction(1))
	defer fmt.Println("!")

	fmt.Println("os.Exit")
	// os.Exit(0)
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
	fmt.Println(helloWorld)
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
	fmt.Println(helloWorld)
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
	fmt.Println(helloWorld)
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

func Salutations() {
	fmt.Println("Hello, and")
	fmt.Println("Goodbye.")
}

func Array0() {
	a := [3]string{
		"Yamada Taro",
		"Sato Hanako",
		"Suzuki Kenji",
	}
	fmt.Printf("%v\n", a[0])
}

func TypeInference() {
	n := 1
	fmt.Printf("%T\n", n)
}

func TypeInference2() {
	n := 1.0
	fmt.Printf("%T\n", n)
}

func DefineVariables() {
	var a, b, c int
	fmt.Printf("%d %d %d\n", a, b, c)
}

func PackageVariables() {
	fmt.Printf("%d %d\n", m, N)
}

func Var1() {
	var n1 int = 1
	var n2 = 2
	fmt.Printf("%d %d\n", n1, n2)
}

func Var2() {
	a := uint32(10)
	b := uint32(20)
	c := a + b
	fmt.Printf("%d\n", c)
}

func Var3() {
	a := int8(math.MaxInt8)
	fmt.Printf("%d\n", a)
}

func Var4() {
	a := int8(math.MaxInt8)
	a++
	fmt.Printf("%d\n", a)
}

func Var5() {
	a := uint32(math.MaxUint32)
	fmt.Printf("%d\n", a)
}

func Var6() {
	a := uint32(math.MaxUint32)
	a++
	fmt.Printf("%d\n", a)
}

func WrapAround() {
	a := uint32(1)
	a--
	fmt.Printf("%d\n", a)
}

func WrapAround2() {
	a := int32(math.MinInt32)
	a--
	fmt.Printf("%d\n", a)
}

func Float() {
	var a float64 = 1.0
	var b float64 = 1.1
	c := a + b
	fmt.Printf("%f\n", c)
}

func Double() {
	var a float64 = 0.1
	var b float64 = 0.2
	c := a + b
	fmt.Printf("%.1f\n", c)
	d := math.Round((a+b)*10) / 10
	fmt.Printf("%.1f\n", d)
}

func Double2() {
	a := 0.1 * 3
	b := 0.3
	fmt.Printf("%f\n", a)
	fmt.Printf("%t\n", a == b)
}

func DoubleCast() {
	var a int = 1
	var b int = 3
	c := float64(a) / float64(b)
	fmt.Printf("%.3f\n", c)
	c2 := a / b
	fmt.Printf("%d\n", c2)
}

func Fmain() {
	fmt.Printf("%f\n", math.Pi)
}

func Complex1() {
	var c complex64 = 1 + 2i
	fmt.Printf("%v\n", c)
}

func Complex2() {
	var r float32 = 1
	var i float32 = 2
	c := complex(r, i)
	fmt.Printf("%v\n", c)
}

func Complex3() {
	c := 1 + 2i
	fmt.Printf("%v\n", real(c))
	fmt.Printf("%v\n", imag(c))
}

func Rune() {
	r := '本'
	fmt.Printf("%d\n", r)
	fmt.Printf("%c\n", r)
	
	s := "本"
	fmt.Printf("%s\n", s)
	
	bs := []byte("本")
	fmt.Printf("%v\n", bs)
	
	rs := []rune("本")
	fmt.Printf("%v\n", rs)
}

func Array1() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Printf("%v\n", a)
}

func Array2() {
	a := [3]string{
		"Red",
		"Green", 
		"Blue",
	}
	fmt.Printf("%v\n", a)
}

func Array3() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
}

func Array4() {
	a := [5]int{1: 10, 3: 30}
	fmt.Printf("%v\n", a)
	
	b := [...]int{5: 50}
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)
}

func Math() {
	n1 := 9
	n2 := 2
	fmt.Printf("%d\n", n1/n2)
	fmt.Printf("%d\n", n1%n2)
	n3 := 9.0
	n4 := 2.0
	fmt.Printf("%f\n", math.Mod(n3, n4))
}

func Hello() {
	fmt.Println(helloWorld)
}

func Div() {
	q, r := div(19, 7)
	fmt.Printf("19/7=%d...%d\n", q, r)
	
	q2, _ := div(19, 7)
	fmt.Printf("19/7=%d\n", q2)
	
	_, r2 := div(19, 7)
	fmt.Printf("19%%7=%d\n", r2)
}

func Later() {
	f := later()
	fmt.Printf("%s\n", f("Golang"))
	fmt.Printf("%s\n", f("is"))
	fmt.Printf("%s\n", f("awesome!"))
}

func Constant() {
	const X = 1
	fmt.Printf("%d\n", X)
}

func Constant1() {
	const (
		X = 1
		Y = 2
		Z = 3
	)
	fmt.Printf("%d %d %d\n", X, Y, Z)
}

func Constant2() {
	const (
		X = 1 + 2
		Y = X + 3
		Z
	)
	fmt.Printf("%d %d %d\n", X, Y, Z)
}

func Constant3() {
	const (
		X = 1
		Y = "ABC"
	)
	fmt.Printf("%d\n", X)
	fmt.Printf("%s\n", Y)
}

func Constant4() {
	const (
		X = 1
		Y
		Z
	)
	fmt.Printf("%d %d %d\n", X, Y, Z)
}

func Constant5() {
	const (
		X = 1 << 1
		Y = 1 << 2
		Z = 1 << 3
	)
	fmt.Printf("%d %d %d\n", X, Y, Z)
}

func ConstantComplex() {
	const C = 1 + 2i
	fmt.Printf("%v\n", C)
}

func ConstantRune() {
	const (
		A = 'A'
		B = 'B'
	)
	fmt.Printf("%c %c\n", A, B)
}

func ConstantIota() {
	const (
		A = iota
		B
		C
	)
	fmt.Printf("%d %d %d\n", A, B, C)
	
	const (
		KB = 1 << (10 * (1 + iota))
		MB
		GB
	)
	fmt.Printf("%d %d %d\n", KB, MB, GB)
}

func Condition() {
	a := 1
	if a == 1 {
		fmt.Printf("a is 1\n")
	} else if a == 2 {
		fmt.Printf("a is 2\n")
	} else {
		fmt.Printf("I don't know\n")
	}
	
	if b := 2; b == 2 {
		fmt.Printf("b is 2\n")
	}
}

func Switch() {
	n := 3
	switch n {
	case 1:
		fmt.Printf("n is 1\n")
	case 2, 3:
		fmt.Printf("n is 2 or 3\n")
	default:
		fmt.Printf("I don't know\n")
	}
	
	switch n2 := n + 1; n2 {
	case 1:
		fmt.Printf("n2 is 1\n")
	case 2, 3:
		fmt.Printf("n2 is 2 or 3\n")
	default:
		fmt.Printf("n2 is %d\n", n2)
	}
	
	switch {
	case n > 0:
		fmt.Printf("n is positive\n")
	case n < 0:
		fmt.Printf("n is negative\n")
	default:
		fmt.Printf("n is 0\n")
	}
}

func TypeAssertion() {
	var i interface{} = 1
	n := i.(int)
	fmt.Printf("%d\n", n)
}

func TypeAsertion2() {
	var i interface{} = 1
	if n, ok := i.(int); ok {
		fmt.Printf("%d\n", n)
	} else {
		fmt.Printf("型が違う\n")
	}
}

func TypeSwitch() {
	var i interface{} = 1
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	default:
		fmt.Printf("I don't know\n")
	}
}

func TypeSwitch2() {
	var i interface{} = "Hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	default:
		fmt.Printf("I don't know\n")
	}
}

func Goto() {
	n := 1
	if n == 1 {
		goto A
	}
	fmt.Printf("ここは通らない\n")
A:
	fmt.Printf("ここを実行\n")
}

func Label() {
LOOP:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j == 2 {
				break LOOP
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}

func Label2() {
LOOP:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j == 2 {
				continue LOOP
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}

func Defer() {
	defer fmt.Printf("defer\n")
	fmt.Printf("1\n")
	fmt.Printf("2\n")
}

func RunRecover() {
	runRecover()
}

func TestRecover() {
	testRecover(10)
	testRecover("エラーだ")
}

func PrivatedoSomethingDo() {
	doSomethingDo()
}
