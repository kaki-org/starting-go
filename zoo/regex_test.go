package main

import (
	"regexp"
	"testing"
)

func TestRegexMatchString(t *testing.T) {
	m, _ := regexp.MatchString("A", "ABC")
	expect(t, m, true)
	m, _ = regexp.MatchString("A", "XYZ")
	expect(t, m, false)
}
