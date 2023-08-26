// type_test.go
package main

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)

	fmt.Printf("%v %T\n", pair, pair) // [1 2] main.IntPair
	fmt.Printf("%v %T\n", strs, strs) // [Apple Banana Cherry] main.Strings
	fmt.Printf("%v %T\n", amap, amap) // map[Tokyo:[35.689488 139.691706]] main.AreaMap
	fmt.Printf("%v %T\n", ich, ich)   // 0xc0000a4000 chan []int
}

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func TestTypeCallback(t *testing.T) {
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)

	expect := 30
	actual := n
	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}
