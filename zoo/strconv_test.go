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

func TestFormatInt(t *testing.T) {
	n := 123

	expect := "123"
	intStr := strconv.FormatInt(int64(n), 10)
	if intStr != expect {
		t.Errorf("expect %s, but got %s", expect, intStr)
	}

	expect16 := "7b"
	intStr16 := strconv.FormatInt(int64(n), 16)
	if intStr16 != expect16 {
		t.Errorf("expect %s, but got %s", expect16, intStr16)
	}

	expectU2 := "1111011"
	intStrU2 := strconv.FormatUint(uint64(n), 2)
	if intStrU2 != expectU2 {
		t.Errorf("expect %s, but got %s", expectU2, intStrU2)
	}

	expectItoa := "123"
	itoaStr := strconv.Itoa(n)
	if itoaStr != expectItoa {
		t.Errorf("expect %s, but got %s", expectItoa, itoaStr)
	}
}
