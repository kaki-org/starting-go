package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var a [10]int
	s := make([]int, 10)

	fmt.Printf("こっちは配列%T\n", a)
	fmt.Printf("こっちはスライス%T\n", s)
	fmt.Printf("%d\n%d\n", len(a), len(s))
	fmt.Printf("%v\n%v\n", a, s)
	expect := 10
	expect_value := 0
	actual := s
	if expect != len(actual) {
		t.Errorf("%d != %d", expect, actual)
	}
	if expect_value != actual[9] {
		t.Errorf("%d != %d", expect_value, actual[0])
	}
	// 要素数を超えた値を指定するとpanicになる
	// fmt.Println(actual[10]) // panic: runtime error: index out of range
}
