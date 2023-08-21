package main

import (
	"fmt"
	"reflect"
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

func TestSliceCap(t *testing.T) {
	/* 要素数5、容量5のスライス */
	s5 := make([]int, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s5), cap(s5), s5)
	/* 要素数5、容量10のスライス */
	s10 := make([]int, 5, 10)
	fmt.Printf("len=%d cap=%d value=%v\n", len(s10), cap(s10), s10)
	expect_cap5 := 5
	expect_cap10 := 10
	if (expect_cap5 != cap(s5)) || (expect_cap10 != cap(s10)) {
		t.Errorf("%d != %d", expect_cap5, cap(s5))
		t.Errorf("%d != %d", expect_cap10, cap(s10))
	}
	// fmt.Println(s10[9]) // これはpanicになる
	s10[4] = 100
	fmt.Println(s10[4])
	fmt.Printf("%#v\n", s10[0:10]) // 0から9までの要素を取得

	s5[0] = 1
	s5[1] = 2
	s5[2] = 3
	s5[3] = 4
	s5[4] = 5

	// makeを使用しないリテラルでのスライスの作成
	s := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(s5, s) {
		t.Errorf("%v != %v", s5, s)
	}
}
