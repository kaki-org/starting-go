package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	// sync.WaitGroupを利用するかどうかをコマンドライン引数から取得
	// -wを指定しない場合はmain()関数が終了するとゴルーチンも終了する
	// 指定した場合はwaitGroupを利用してゴルーチンの終了を待つ
	// 起動option
	var w bool
	flag.BoolVar(&w, "w", false, "waitGroupを利用するかどうか")
	flag.Parse()

	// sync.WaitGroupを生成
	wg := new(sync.WaitGroup)
	if w {
		// 待ち受けするゴルーチンの数は3
		wg.Add(3)
	}
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("A:", i)
		}
		if w {
			// ゴルーチンの終了を通知
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("B:", i)
		}
		if w {
			// ゴルーチンの終了を通知
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("C:", i)
		}
		if w {
			// ゴルーチンの終了を通知
			wg.Done()
		}
	}()
	if w {
		// ゴルーチンの終了を待つ
		wg.Wait()
	}
}
