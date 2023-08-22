package main

import (
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
		t.Errorf("%v != %v", expect, actual)
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
		t.Errorf("%v != %v", expect, actual)
	}
}

func TestMapFloat(t *testing.T) {
	m := make(map[float64]int)

	m[1.00000000000000001] = 1
	m[1.00000000000000002] = 2
	m[1.00000000000000003] = 3 // m[1.0]と同じ値になるので、全部上書き扱い

	expect := 3
	if expect != m[1.0] {
		t.Errorf("%d != %d", expect, m[1.0])
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
		t.Errorf("%v != %v", m1, m2)
	}
}
