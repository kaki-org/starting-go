package main

import (
	"math"
	"testing"
)

func TestSqrt2(t *testing.T) {
	expected := 1.4142135623730951
	actual := math.Sqrt2
	if actual != expected {
		t.Errorf("Sqrt2 = %f, want %f", actual, expected)
	}
}

func TestPi(t *testing.T) {
	// 円周率を求める
	expected := 3.141592653589793
	actual := math.Pi
	if actual != expected {
		t.Errorf("Pi = %f, want %f", actual, expected)
	}
}

func TestPhi(t *testing.T) {
	// 黄金比を求める
	expected := 1.618033988749895
	actual := math.Phi
	if actual != expected {
		t.Errorf("Phi = %f, want %f", actual, expected)
	}
}

func TestE(t *testing.T) {
	// 自然対数の底(ネイピア数)を求める
	expected := 2.718281828459045
	actual := math.E
	if actual != expected {
		t.Errorf("E = %f, want %f", actual, expected)
	}
}

func TestNumericFixedNumber(t *testing.T) {
	// float32で表現可能な最大値
	expectMaxFloat32 := 3.40282346638528859811704183484516925440e+38
	actualMaxFloat32 := math.MaxFloat32
	if actualMaxFloat32 != expectMaxFloat32 {
		t.Errorf("MaxFloat32 = %f, want %f", actualMaxFloat32, expectMaxFloat32)
	}

	// float32で表現可能な0ではない最小値
	expectSmallestNonZeroFloat32 := 1.401298464324817070923729583289916131280e-45
	actualSmallestNonZeroFloat32 := math.SmallestNonzeroFloat32
	if actualSmallestNonZeroFloat32 != expectSmallestNonZeroFloat32 {
		t.Errorf("SmallestNonzeroFloat32 = %f, want %f", actualSmallestNonZeroFloat32, expectSmallestNonZeroFloat32)
	}

	// float64で表現可能な最大値
	expectMaxFloat64 := 1.797693134862315708145274237317043567981e+308
	actualMaxFloat64 := math.MaxFloat64
	if actualMaxFloat64 != expectMaxFloat64 {
		t.Errorf("MaxFloat64 = %f, want %f", actualMaxFloat64, expectMaxFloat64)
	}

	// float64で表現可能な0ではない最小値
	expectSmallestNonZeroFloat64 := 4.940656458412465441765687928682213723651e-324
	actualSmallestNonZeroFloat64 := math.SmallestNonzeroFloat64
	if actualSmallestNonZeroFloat64 != expectSmallestNonZeroFloat64 {
		t.Errorf("SmallestNonzeroFloat64 = %f, want %f", actualSmallestNonZeroFloat64, expectSmallestNonZeroFloat64)
	}

	expectMaxInt8 := 127
	actualMaxInt8 := math.MaxInt8
	if actualMaxInt8 != expectMaxInt8 {
		t.Errorf("MaxInt8 = %d, want %d", actualMaxInt8, expectMaxInt8)
	}
	expectMinInt8 := -128
	actualMinInt8 := math.MinInt8
	if actualMinInt8 != expectMinInt8 {
		t.Errorf("MinInt8 = %d, want %d", actualMinInt8, expectMinInt8)
	}
	expectMaxInt16 := 32767
	actualMaxInt16 := math.MaxInt16
	if actualMaxInt16 != expectMaxInt16 {
		t.Errorf("MaxInt16 = %d, want %d", actualMaxInt16, expectMaxInt16)
	}
	expectMinInt16 := -32768
	actualMinInt16 := math.MinInt16
	if actualMinInt16 != expectMinInt16 {
		t.Errorf("MinInt16 = %d, want %d", actualMinInt16, expectMinInt16)
	}
	expectMaxInt32 := 2147483647
	actualMaxInt32 := math.MaxInt32
	if actualMaxInt32 != expectMaxInt32 {
		t.Errorf("MaxInt32 = %d, want %d", actualMaxInt32, expectMaxInt32)
	}
	expectMinInt32 := -2147483648
	actualMinInt32 := math.MinInt32
	if actualMinInt32 != expectMinInt32 {
		t.Errorf("MinInt32 = %d, want %d", actualMinInt32, expectMinInt32)
	}
	expectMaxInt64 := 9223372036854775807
	actualMaxInt64 := math.MaxInt64
	if actualMaxInt64 != expectMaxInt64 {
		t.Errorf("MaxInt64 = %d, want %d", actualMaxInt64, expectMaxInt64)
	}
	expectMinInt64 := -9223372036854775808
	actualMinInt64 := math.MinInt64
	if actualMinInt64 != expectMinInt64 {
		t.Errorf("MinInt64 = %d, want %d", actualMinInt64, expectMinInt64)
	}
	expectMaxUint8 := 255
	actualMaxUint8 := math.MaxUint8
	if actualMaxUint8 != expectMaxUint8 {
		t.Errorf("MaxUint8 = %d, want %d", actualMaxUint8, expectMaxUint8)
	}
	expectMaxUint16 := 65535
	actualMaxUint16 := math.MaxUint16
	if actualMaxUint16 != expectMaxUint16 {
		t.Errorf("MaxUint16 = %d, want %d", actualMaxUint16, expectMaxUint16)
	}
	expectMaxUint32 := 4294967295
	actualMaxUint32 := math.MaxUint32
	if actualMaxUint32 != expectMaxUint32 {
		t.Errorf("MaxUint32 = %d, want %d", actualMaxUint32, expectMaxUint32)
	}
	expectMaxUint64 := uint64(18446744073709551615)
	actualMaxUint64 := uint64(math.MaxUint64)
	if actualMaxUint64 != expectMaxUint64 {
		t.Errorf("MaxUint64 = %d, want %d", actualMaxUint64, expectMaxUint64)
	}
}

func TestAbs(t *testing.T) {
	expect := 3.14
	actual := math.Abs(-3.14)
	if actual != expect {
		t.Errorf("Abs(-3.14) = %f, want %f", actual, expect)
	}
}

func TestPow(t *testing.T) {
	expect := 0
	actual := int(math.Pow(0, 2))
	if actual != expect {
		t.Errorf("Pow(0, 2) = %d, want %d", actual, expect)
	}
	expect = 1
	actual = int(math.Pow(1, 2))
	if actual != expect {
		t.Errorf("Pow(1, 2) = %d, want %d", actual, expect)
	}
	expect = 256
	actual = int(math.Pow(2, 8))
	if actual != expect {
		t.Errorf("Pow(2, 8) = %d, want %d", actual, expect)
	}

}
