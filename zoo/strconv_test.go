package main

import (
	"math"
	"strconv"
	"testing"
)

const errorFormatValue = "err = %v"

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
	b, err := strconv.ParseBool("true")
	if err != nil {
		t.Errorf("%v\n", err)
	}
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

func TestParseInt(t *testing.T) {
	intValue, _ := strconv.ParseInt("12345", 10, 0)
	cmp(intValue, int64(12345), t)

	uIntValue, _ := strconv.ParseUint("12345", 10, 0)
	cmp(uIntValue, uint64(12345), t)

	intValue, _ = strconv.ParseInt("-12", 10, 0)
	cmp(intValue, int64(-12), t)

	// 8進数
	intValue, _ = strconv.ParseInt("740", 8, 0)
	cmp(intValue, int64(480), t)
	// 16進数
	intValue, _ = strconv.ParseInt("7F4D", 16, 0)
	cmp(intValue, int64(32589), t)

	//負の値はエラー
	_, err := strconv.ParseUint("-1", 10, 0)
	if err != nil {
		t.Logf(errorFormatValue, err)
	}
	//指定した精度を超える整数表現はエラー
	_, err = strconv.ParseInt("123456", 10, 16)
	if err != nil {
		t.Logf(errorFormatValue, err)
	}

	// 0が指定された場合は文字列のプリフィックスが作用する
	oValue, _ := strconv.ParseInt("0123", 0, 0)
	cmp(oValue, int64(83), t)
	hValue, _ := strconv.ParseInt("0x123", 0, 0)
	cmp(hValue, int64(291), t)
}

func TestParseFloat(t *testing.T) {
	floatValue, _ := strconv.ParseFloat("3.14", 64)
	cmp(floatValue, 3.14, t)
	floatValue, _ = strconv.ParseFloat("0.2", 64)
	cmp(floatValue, 0.2, t)
	floatValue, _ = strconv.ParseFloat("-2", 64)
	cmp(floatValue, float64(-2), t)
	floatValue, _ = strconv.ParseFloat("1.2345e8", 64)
	cmp(floatValue, float64(1.2345e+08), t)
	floatValue, _ = strconv.ParseFloat("1.2345E8", 64)
	cmp(floatValue, float64(1.2345e+08), t) // Eでもeでも同じ

	// 指定した精度に合わせて丸められる
	floatValue, _ = strconv.ParseFloat("1.000000000001", 32)
	cmp(floatValue, float64(1), t)
	floatValue, _ = strconv.ParseFloat("1.000000000001", 64)
	cmp(floatValue, float64(1.000000000001), t)
	// 精度を超過する範囲の表現は無限大とエラーを返す
	floatValue, err := strconv.ParseFloat("1E500", 64)
	cmp(floatValue, math.Inf(0), t)
	if err != nil {
		t.Logf(errorFormatValue, err)
	}
	floatValue, err = strconv.ParseFloat("-1E500", 64)
	cmp(floatValue, math.Inf(-1), t)
	if err != nil {
		t.Logf(errorFormatValue, err)
	}

}

func cmp(v1, v2 interface{}, t *testing.T) bool {
	if v1 == v2 {
		return true
	} else {
		t.Errorf("expect %v, but got %v", v1, v2)
	}
	return false
}
