package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("ファイル名を入力してください")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		f, err := os.Open(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		// 入力されたファイルをすべて読み込む
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bs))
		f.Close()
		fmt.Println("ファイル名を入力してください")
	}
}
