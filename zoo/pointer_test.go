package main

import (
	"fmt"
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
		t.Errorf("%d != %d", expect, actual)
	}

	*p = 10
	expect = 10
	actual = i
	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
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
		t.Errorf("%d != %d", expect, actual)
	}
}
