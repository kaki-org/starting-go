package main

import (
	"testing"
)

func TestSlice(t *testing.T) {
	s := make([]int, 10)

	expect := 10
	actual := s
	if expect != len(actual) {
		t.Errorf("%d != %d", expect, actual)
	}
}
