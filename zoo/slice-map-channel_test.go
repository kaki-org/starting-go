package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var a [10]int
	s := make([]int, 10)

	fmt.Printf("こっちは配列%T\n", a)
	fmt.Println(a)
	fmt.Printf("こっちはスライス%T\n", s)
	fmt.Println(s)
	expect := 10
	expect_value := 0
	actual := s
	if expect != len(actual) {
		t.Errorf("%d != %d", expect, actual)
	}
	if expect_value != actual[9] {
		t.Errorf("%d != %d", expect_value, actual[0])
	}
}
