package main

import (
	"testing"
)

func TestPointerInitialize(t *testing.T) {
	// 初期化のみ行ったポインタ変数はnil
	var p *int

	if p != nil {
		t.Errorf("ポインタ変数はnilではありません。")
	}
}
