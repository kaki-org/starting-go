package main

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)

	fmt.Printf("%v %T\n", pair, pair) // [1 2] main.IntPair
	fmt.Printf("%v %T\n", strs, strs) // [Apple Banana Cherry] main.Strings
	fmt.Printf("%v %T\n", amap, amap) // map[Tokyo:[35.689488 139.691706]] main.AreaMap
	fmt.Printf("%v %T\n", ich, ich)   // 0xc0000a4000 chan []int
}

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func TestTypeCallback(t *testing.T) {
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)

	expect := 30
	actual := n
	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestStruct(t *testing.T) {
	type Point struct {
		X, Y int
	}

	var pt Point
	if (pt.X != pt.Y) || (pt.X != 0) {
		t.Errorf("pt.X or pt.Y is not 0")
	}

	pt.X = 10
	pt.Y = 8

	expect := Point{10, 8} // 省略してかけるが、後述のようにフィールド名を指定する書き方が推奨
	actual := pt
	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}
	expect = Point{X: 10, Y: 8} // フィールド名を指定して初期化
	if expect != actual {
		t.Errorf("%v != %v", expect, actual)
	}

	pt2 := Point{Y: 5}

	expectX := 0
	actualX := pt2.X
	if expectX != actualX {
		t.Errorf("%d != %d", expectX, actualX)
	}
	expectY := 5
	actualY := pt2.Y
	if expectY != actualY {
		t.Errorf("%d != %d", expectY, actualY)
	}
}

func TestStructField(ts *testing.T) {
	// 構造体のフィールド名指定を省略すると、宣言順に初期化される
	type T struct {
		int
		float64
		string
	}
	t := T{5, 3.14, "文字列"}
	expectInt := 5
	actualInt := t.int
	if expectInt != actualInt {
		ts.Errorf("%d != %d", expectInt, actualInt)
	}
	expectFloat64 := 3.14
	actualFloat64 := t.float64
	if expectFloat64 != actualFloat64 {
		ts.Errorf("%f != %f", expectFloat64, actualFloat64)
	}
	expectString := "文字列"
	actualString := t.string
	if expectString != actualString {
		ts.Errorf("%s != %s", expectString, actualString)
	}

	// 無名フィールドを用いた構造体の埋め込み
	type T2 struct {
		N uint
		_ int16
		S string
	}

	t2 := T2{
		N: 1,
		S: "文字列",
	}
	expect2 := T2{1, 0, "文字列"} // 0はint16のゼロ値。無名フィールドにも値は存在する
	if expect2 != t2 {
		ts.Errorf("%v != %v", expect2, t2)
	}
	fmt.Println(t2)
}
