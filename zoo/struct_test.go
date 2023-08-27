package main

import (
	"fmt"
	"math"
	"reflect"
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

func TestStructInStruct(t *testing.T) {
	type Feed struct {
		Name   string
		Amount uint
	}
	type Animal struct {
		Name string
		Feed Feed // 構造体のフィールドに構造体を埋め込む
	}
	a := Animal{
		Name: "Monkey",
		Feed: Feed{ /* 複合リテラル内の複合リテラル */
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a)             // {Monkey {Banana 10}}
	fmt.Println(a.Name)        // Monkey
	fmt.Println(a.Feed.Name)   // Banana
	fmt.Println(a.Feed.Amount) // 10

	a.Feed.Amount = 15
	fmt.Println(a.Feed.Amount) // 15
}

func TestStructInStruct2(t *testing.T) {
	type Feed struct {
		Name   string
		Amount uint
	}
	type Animal struct {
		Name string
		Feed // 構造体のフィールドに構造体を埋め込む
	}
	a := Animal{
		Name: "Monkey",
		Feed: Feed{ /* 複合リテラル内の複合リテラル */
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a)           // {Monkey {Banana 10}}
	fmt.Println(a.Name)      // Monkey
	fmt.Println(a.Feed.Name) // Banana
	fmt.Println(a.Amount)    // 10(フィールド名を省略している a.Feed.Amountと同じ)

	a.Amount = 15
	fmt.Println(a.Amount) // 15
}

func TestStructInStruct3(t *testing.T) {
	type T1 struct {
		Name1 string
	}
	type T2 struct {
		T1
		Name2 string
	}
	type T3 struct {
		T2
		Name3 string
	}
	t123 := T3{T2: T2{T1: T1{Name1: "A"}, Name2: "B"}, Name3: "C"}
	if t123.Name1 != "A" {
		t.Errorf("t123.Name1 is not A")
	}
	if t123.Name2 != "B" {
		t.Errorf("t123.Name2 is not B")
	}
	if t123.Name3 != "C" {
		t.Errorf("t123.Name3 is not C")
	}
}

func TestStructInStruct4(t *testing.T) {
	type Base struct {
		Id    int
		Owner string
	}
	type A struct {
		Base /* 共通のフィールド */
		Name string
		Area string
	}
	type B struct {
		Base   /* 共通のフィールド */
		Title  string
		Bodies []string
	}

	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}
	b := B{
		Base:   Base{81, "Hanako"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}
	fmt.Println(a.Id)
	fmt.Println(a.Owner)
	fmt.Println(b.Id)
	fmt.Println(b.Owner)
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}

type Point struct {
	X, Y int
}
type FloatPoint struct {
	X, Y float64
}

func swap(p Point) {
	/* フィールドXとYを入れ替える */
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}
func swap2(p *Point) {
	/* フィールドXとYを入れ替える */
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func TestAnonymousStruct(t *testing.T) {
	// 無名構造体。あえてこの書き方にする必要はない
	s := struct{ X, Y int }{X: 1, Y: 2}
	showStruct(s)
	p := Point{X: 1, Y: 2}
	showStruct(p) // 互換性はあるのでPointを渡してもOK
	if reflect.DeepEqual(s, p) {
		t.Errorf("%v != %v", s, p)
	}
}

func TestStructValue(t *testing.T) {
	p := Point{X: 1, Y: 2}
	swap(p)
	// 構造体は値渡しなので、swap関数内での変更は反映されない
	fmt.Println(p) // {1 2}
	swap2(&p)
	// ポインタを渡すと、swap関数内での変更が反映される
	fmt.Println(p) // {2 1}
}

type Person struct {
	Id   int
	Name string
	Area string
}

func setPerson(p *Person, id int, name, area string) {
	p.Id = id
	p.Name = name
	p.Area = area
}

func TestStructNew(t *testing.T) {
	// newで構造体(のポインタ)を生成する
	p := new(Person)
	expectId := 0
	actualId := p.Id
	if expectId != actualId {
		t.Errorf("%d != %d", expectId, actualId)
	}
	expectName := ""
	actualName := p.Name
	if expectName != actualName {
		t.Errorf("%s != %s", expectName, actualName)
	}
	expectArea := ""
	actualArea := p.Area
	if expectArea != actualArea {
		t.Errorf("%s != %s", expectArea, actualArea)
	}

	setPerson(p, 1, "Gopher", "Tokyo")

	expectPerson := Person{Id: 1, Name: "Gopher", Area: "Tokyo"}
	actualPerson := *p
	if expectPerson != actualPerson {
		t.Errorf("%v != %v", expectPerson, actualPerson)
	}
}

/* *Point型のメソッドRender */
func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

/* IntPoint型の2点間の距離を求めるメソッドDistance */
func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

/* FloatPoint型の2点間の距離を求めるメソッドDistance */
func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(x*x + y*y)
}

func TestStructMethod(t *testing.T) {
	p := Point{X: 0, Y: 0}
	p.Render() // メソッド呼び出し<0, 0>

	distance := p.Distance(&Point{X: 1, Y: 1}) // メソッド呼び出し
	expect := 1.4142135623730951
	if expect != distance {
		t.Errorf("%f != %f", expect, distance)
	}

	fp := FloatPoint{X: 0.0, Y: 0.0}

	distanceFloat := fp.Distance(&FloatPoint{X: 1.0, Y: 1.0}) // メソッド呼び出し
	expectFloat := 1.4142135623730951
	if expectFloat != distanceFloat {
		t.Errorf("%f != %f", expectFloat, distanceFloat)
	}
}
