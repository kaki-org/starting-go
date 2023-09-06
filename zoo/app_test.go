package main

import (
	"fmt"
	"testing"
)

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := AppName()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func ExampleAppName() {
	fmt.Println(AppName())
	// Output: Zoo Application
}

func ExampleSalutations() {
	fmt.Println("Hello, and")
	fmt.Println("Goodbye.")
	// Output:
	// Hello, and
	// Goodbye.
}
