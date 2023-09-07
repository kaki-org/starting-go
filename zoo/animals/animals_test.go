package animals

import (
	"testing"
)

func TestElephantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
func TestMonkeyFeed(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestRabbitFeed(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestFooFunc(t *testing.T) {
	expect := 3
	actual := FooFunc(2)

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestConstant(t *testing.T) {
	if MAX != 100 {
		t.Errorf("MAX should be 100, but has %d", MAX)
	}
	if internal_const != 1 {
		t.Errorf("internal_const should be 1, but has %d", internal_const)
	}
}
