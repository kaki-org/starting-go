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
	cmp(intStr, expect, t)

	expect16 := "7b"
	intStr16 := strconv.FormatInt(int64(n), 16)
	cmp(intStr16, expect16, t)

	expectU2 := "1111011"
	intStrU := strconv.FormatUint(uint64(n), 2)
	cmp(intStrU, expectU2, t)

	expectItoa := "123"
	itoaStr := strconv.Itoa(n)
	cmp(itoaStr, expectItoa, t)
}

func TestFormatFloat(t *testing.T) {
	// 指数表現による文字列化
	expect := "1.23456E+02"
	floatStr := strconv.FormatFloat(123.456, 'E', -1, 64)
	cmp(floatStr, expect, t)

	// 指数表現による文字列化(小数点以下2桁まで)
	expect = "1.23e+02"
	floatStr = strconv.FormatFloat(123.456, 'e', 2, 64)
	cmp(floatStr, expect, t)

	// 実数表現による文字列化
	expect = "123.456"
	floatStr = strconv.FormatFloat(123.456, 'f', -1, 64)
	cmp(floatStr, expect, t)

	// 実数表現による文字列化(小数点2桁まで)
	expect = "123.46"
	floatStr = strconv.FormatFloat(123.456, 'f', 2, 64)
	cmp(floatStr, expect, t)

	// 指数部の大きさで変動する表現による文字列化
	expect = "123.456"
	floatStr = strconv.FormatFloat(123.456, 'g', -1, 64)
	cmp(floatStr, expect, t)

	// 指数部の大きさで変動する表現による文字列化(仮数部全体が4桁まで)
	expect = "123.5"
	floatStr = strconv.FormatFloat(123.456, 'g', 4, 64)
	cmp(floatStr, expect, t)

	// 指数部の大きさで変動する表現による文字列化(仮数部全体が8桁まで)
	expect = "1.2345679E+08"
	floatStr = strconv.FormatFloat(123456789.123, 'G', 8, 64)
	cmp(floatStr, expect, t)
}

func TestParseBool(t *testing.T) {
	// trueへ変換できる文字列
	expect := true
	b, _ := strconv.ParseBool("true")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("1")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("t")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("T")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("TRUE")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("True")
	cmp(b, expect, t)

	// falseへ変換できる文字列
	expect = false
	b, _ = strconv.ParseBool("false")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("0")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("f")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("F")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("FALSE")
	cmp(b, expect, t)

	b, _ = strconv.ParseBool("False")
	cmp(b, expect, t)
}

func cmp(v1, v2 interface{}, t *testing.T) bool {
	if v1 == v2 {
		return true
	} else {
		t.Errorf("expect %s, but got %s", v1, v2)
	}
	return false
}
