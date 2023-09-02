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

func TestFormatFloat(t *testing.T) {
	// 指数表現による文字列化
	expect := "1.23456E+02"
	floatStr := strconv.FormatFloat(123.456, 'E', -1, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

	// 指数表現による文字列化(小数点以下2桁まで)
	expect = "1.23e+02"
	floatStr = strconv.FormatFloat(123.456, 'e', 2, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

	// 実数表現による文字列化
	expect = "123.456"
	floatStr = strconv.FormatFloat(123.456, 'f', -1, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

	// 実数表現による文字列化(小数点2桁まで)
	expect = "123.46"
	floatStr = strconv.FormatFloat(123.456, 'f', 2, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

	// 指数部の大きさで変動する表現による文字列化
	expect = "123.456"
	floatStr = strconv.FormatFloat(123.456, 'g', -1, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

	// 指数部の大きさで変動する表現による文字列化(仮数部全体が4桁まで)
	expect = "123.5"
	floatStr = strconv.FormatFloat(123.456, 'g', 4, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

	// 指数部の大きさで変動する表現による文字列化(仮数部全体が8桁まで)
	expect = "1.2345679E+08"
	floatStr = strconv.FormatFloat(123456789.123, 'G', 8, 64)
	if floatStr != expect {
		t.Errorf("expect %s, but got %s", expect, floatStr)
	}

}
