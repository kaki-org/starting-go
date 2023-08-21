package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := make([]int, 10)

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
