package main

import (
	"fmt"
	"sync"
	"time"
)

// 各ゴルーチンが共有するパッケージ変数
var st struct{ A, B, C int }

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ロックを取得
	mutex.Lock()
	// stの各フィールドをスリープをはさみながら更新
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	// stの各フィールドを出力
	fmt.Println(st.A, st.B, st.C)
	// アンロック
	defer mutex.Unlock()
}
func main() {
	// ミューテックスの生成
	mutex = new(sync.Mutex)

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
