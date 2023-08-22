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
