package main

import (
	"fmt"

	"github.com/kakikubo/starting-go/zoo/animals"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	// コメントです。
	a := [3]string{
		"Yamada Taro",
		"Sato Hanako",
		"Suzuki Kenji",
	}
	fmt.Println(a[0], a[1], a[2])
	// 10進数の形式で数値5を%dの箇所へ埋め込む
	fmt.Printf("%d\n", 5) // => "数値=5"
	// 数値用の書式いろいろ
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17) // => "10進数=17 2進数=10001 8進数=21 16進数=11"
	// 埋め込むパラメータが足りない
	fmt.Printf("%d年%d月%d日\n", 2015, 7) // => "2015年7月%!d(MISSING)日"
	// 埋め込むパラメータが過剰
	fmt.Printf("%d年%d月%d日\n", 2015, 7, 15, 23)
	/* => "2015年7月15日
	%!(EXTRA int=23)" */

	// %vはさまざまな型のデータを埋め込む
	fmt.Printf("\n数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列=Golang 配列=[1 2 3]"
	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列="Golang" 配列=[3]int{1, 2, 3}"
	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=int 文字列=string 配列=[3]int"

	print("Hello, World!\n") // => "Hello, World!" を標準エラー出力へ出力
	println("Hello, World!") // => "Hello, World!\n" を標準エラー出力へ出力
	print(1, 2, 3, "\n")     // => "123\n" を標準エラー出力へ出力
	println(1, 2, 3)         // => "1 2 3\n" を標準エラー出力へ出力

	var n int
	n = 5
	// n == 5
	// n = "string value" // コンパイルエラー
	fmt.Println(n)
	// int型の変数iを定義して1を代入(型推論)
	i := 1
	fmt.Println(i)
	// bool型の変数bを定義して真偽値trueを代入(型推論)
	b := true
	fmt.Println(b)
	// float64型の変数fを定義して浮動小数点数3.14を代入(型推論)
	f := 3.14
	fmt.Println(f)
	// string型の変数sを定義して文字列"abc"を代入(型推論)
	s := "abc"
	fmt.Println(s)

	// 関数の戻り値を元に型推論を行って変数を定義、初期化
	n2 := one()
	fmt.Println(n2)

	// varで変数定義をまとめる書き方
	var (
		n3 = 1
		s3 = "string value"
		b3 = true
	)
	// 暗黙的な定義を並べる書き方だと以下
	// n3 := 1
	// s3 := "string value"
	// b3 := true
	fmt.Println(n3, s3, b3)

	// パッケージ変数n1の値を用いて計算
	n = n1 + 1
	fmt.Printf("n=%d\n", n)

	var_sample()
}

func one() int {
	return 1
}

func var_sample() {
	a := 1
	// 以下を定義しているとエラーになってしまう
	// b := 2
	// c := 3
	// # command-line-arguments
	// ./main.go:93:2: b declared and not used
	// ./main.go:94:2: c declared and not used

	fmt.Println(a)
}
