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
