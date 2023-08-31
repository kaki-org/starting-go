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
