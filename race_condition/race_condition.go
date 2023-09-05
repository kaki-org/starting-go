package main

import (
	"fmt"
	"time"
)

// 各ゴルーチンが共有するパッケージ変数
var st struct{ A, B, C int }

func UpdateAndPrint(n int) {
	// stの各フィールドをスリープをはさみながら更新
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	// stの各フィールドを出力
	fmt.Println(st.A, st.B, st.C)
}
func main() {
	// 複数のゴルーチンを起動する
	// この関数を実行すると、stの各値が常に等しくなることが期待される。が、実際はそうならない。
	// なぜなら、各ゴルーチンがstのフィールドを更新する.順序が不定だから
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {

	}

}
