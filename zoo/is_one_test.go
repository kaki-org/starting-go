package main

import (
	"testing"
)

func TestIsOne(t *testing.T) {
	b := IsOne(1)
	expect(t, b, true)
}

func TestIsOneFalse(t *testing.T) {
	b := IsOne(100)
	expect(t, b, false)
}
