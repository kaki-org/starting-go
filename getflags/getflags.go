package main

import (
	"flag"
	"fmt"
)

func main() {
	/* オプションの値を格納する変数を定義 */
	var (
		max int
		msg string
		x   bool
	)

	/* コマンドラインオプションの定義 */
	flag.IntVar(&max, "n", 32, "処理回数の最大値")
	flag.StringVar(&msg, "m", "Hello, World!", "出力するメッセージ")
	flag.BoolVar(&x, "x", false, "拡張オプション")
	/* コマンドラインをパース */
	flag.Parse()

	fmt.Println("処理数の最大値 =", max)
	fmt.Println("処理メッセージ =", msg)
	fmt.Println("拡張オプション = ", x)

}
