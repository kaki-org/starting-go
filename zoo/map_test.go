package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	expect := map[int]string{1: "US", 81: "Japan", 86: "China"}
	actual := m
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(valueExpectFormat, expect, actual)
	}
}

func TestMapOverride(t *testing.T) {
	m := make(map[string]string)

	m["Yamada"] = "Taro"
	m["Sato"] = "Hanako"
	m["Yamada"] = "Jiro" // Yamadaの値が上書きされる

	expect := map[string]string{"Yamada": "Jiro", "Sato": "Hanako"}
	actual := m
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(valueExpectFormat, expect, actual)
	}
}

func TestMapFloat(t *testing.T) {
	m := make(map[float64]int)

	m[1.00000000000000001] = 1
	m[1.00000000000000002] = 2
	m[1.00000000000000003] = 3 // m[1.0]と同じ値になるので、全部上書き扱い

	expect := 3
	if expect != m[1.0] {
		t.Errorf(decimalExpectFormat, expect, m[1.0])
	}
}

func TestMapLiteral(t *testing.T) {
	// 以下、どちらも同じ意味
	m1 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	m2 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}

	if !reflect.DeepEqual(m1, m2) {
		t.Errorf(valueExpectFormat, m1, m2)
	}
}

func TestMapReference(t *testing.T) {
	// 基本的なMapの参照
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	s := m[1]
	expect := "A"
	if expect != s {
		t.Errorf(stringExpectFormat, expect, s)
	}
	s2 := m[9] // 存在しないキーを指定してもエラーにならない
	expect = ""
	if expect != s2 {
		t.Errorf(stringExpectFormat, expect, s2)
	}

	// 0が返るが、値がとれて0なのか、存在しないキーなのか区別できない
	m2 := map[int]int{1: 0}
	e := m2[1] // 値がとれている
	n := m2[7] // 存在しないキー

	if e != n {
		t.Errorf(decimalExpectFormat, e, n)
	}

	// キーが存在しなかったかどうかを調べる方法
	m3 := map[int]string{1: "A", 2: "B", 3: "C"}

	s, ok := m3[1]
	expect = "A"
	result := true
	if expect != s || result != ok {
		t.Errorf("%s != %s %t != %t", expect, s, result, ok)
	}

	s, ok = m3[9]
	expect = ""
	result = false
	if expect != s || result != ok {
		t.Errorf("%s != %s %t != %t", expect, s, result, ok)
	}

	_, ok = m3[3]
	result = true
	if result != ok {
		t.Errorf("%t != %t", result, ok)
	}

	// 頻出イディオム
	if _, ok := m[1]; ok {
		fmt.Println("m[1]は存在する")
	}
}

func TestMapFor(t *testing.T) {
	m := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	for k, v := range m {
		fmt.Printf("%d => %s\n", k, v)
		// result += fmt.Sprintf("%d => %s\n", k, v)
	}

	// 順序が保証されないので、結果が毎回異なる
	// expect := "1 => Apple\n2 => Banana\n3 => Cherry\n"

	// 個別に指定する
	expect_1 := "Apple"
	expect_2 := "Banana"
	expect_3 := "Cherry"

	if expect_1 != m[1] {
		t.Errorf(stringExpectFormat, expect_1, m[1])
	}
	if expect_2 != m[2] {
		t.Errorf(stringExpectFormat, expect_2, m[2])
	}
	if expect_3 != m[3] {
		t.Errorf(stringExpectFormat, expect_3, m[3])
	}
}

func TestMapLen(t *testing.T) {
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	expect := 3
	actual := len(m)
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}

	m[4] = "D"
	m[5] = "E"
	expect = 5
	actual = len(m)
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}
}

func TestMapDelete(t *testing.T) {
	m := map[int]string{1: "A", 2: "B", 3: "C"}
	delete(m, 2)

	expect := map[int]string{1: "A", 3: "C"}
	actual := m

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(valueExpectFormat, expect, actual)
	}
}

func TestMapMake(t *testing.T) {
	// 要素100のMapを作成(だからといってテストがかけるわけでもないので困った)
	m := make(map[int]string, 100)

	m[1] = "A"
	m[2] = "B"
	m[3] = "C"

	expect := map[int]string{1: "A", 2: "B", 3: "C"}
	actual := m

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(valueExpectFormat, expect, actual)
	}
}
