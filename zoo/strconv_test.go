package main

import (
	"strconv"
	"testing"
)

func TestFormatBool(t *testing.T) {
	b := true

	expect := "true"
	boolStr := strconv.FormatBool(b)
	if boolStr != expect {
		t.Errorf("expect %s, but got %s", expect, boolStr)
	}
}
