package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPointerInitialize(t *testing.T) {
	// 初期化のみ行ったポインタ変数はnil
	var p *int

	if p != nil {
		t.Errorf("ポインタ変数はnilではありません。")
	}
}

func TestAddressOperator(t *testing.T) {
	// &演算子で変数のアドレスを取得
	var i int
	p := &i // pは*iのポインタ型

	fmt.Printf("%T\n", p) // *int
	if p == nil {
		t.Errorf("ポインタ変数はnilです。")
	}

	pp := &p               // ppは**intのポインタ型
	fmt.Printf("%T\n", pp) // **int
	if pp == nil {
		t.Errorf("ポインタ変数はnilです。")
	}

	i = 5
	expect := 5
	actual := *p // *pはiのエイリアス
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}

	*p = 10
	expect = 10
	actual = i
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}
}

func inc(p *int) {
	// pをデリファレンスして、値をインクリメント
	*p++
}

func TestPointerArgument(t *testing.T) {
	i := 1
	inc(&i)
	inc(&i)
	inc(&i)
	expect := 4
	actual := i
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}
}

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		/* 各要素を累乗する */
		p[i] = p[i] * p[i] // 通常の書き方だと→(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func TestPointerArray(t *testing.T) {
	/* 変数pは*[3]int型 */
	p := &[3]int{1, 2, 3}
	pow(p)
	fmt.Println(p) // [1 4 9]
	expect := [3]int{1, 4, 9}
	actual := *p
	if expect != actual {
		t.Errorf(decimalExpectFormat, expect, actual)
	}
}

func TestPointArrayDereference(t *testing.T) {
	a := [3]string{"Apple", "Banana", "Cherry"}
	ap := &a

	if a[1] != (*ap)[1] {
		t.Errorf("%s != %s", a[1], (*ap)[1])
	}
	if a[2] != (*ap)[2] {
		t.Errorf("%s != %s", a[2], (*ap)[2])
	}

	// ポインタ型に対してもrangeが使える
	for i, v := range ap { // for i, v := range *ap {
		fmt.Println(i, v)
	}

	// lenやcapはポインタ型に対しても使える
	expect2 := 3
	if expect2 != len(ap) {
		t.Errorf("%d != %d", expect2, len(ap))
	}
	if expect2 != cap(ap) {
		t.Errorf("%d != %d", expect2, cap(ap))
	}

	expect3 := [2]string{"Apple", "Banana"}
	fmt.Println(ap[0:2])
	if reflect.DeepEqual(expect3, ap[0:2]) {
		t.Errorf("%v != %v", expect3, ap[0:2])
	}

}
