package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPrintf(t *testing.T) {
	n := 4
	// outには出力したバイト数が入る
	out, _ := fmt.Printf("%d\n", n)
	if out != 2 {
		t.Errorf("%d != %d", out, 2)
	}

	// sには出力した文字列が入る
	s := fmt.Sprintf("%d\n", n)
	if s != "4\n" {
		t.Errorf("%s != %s", s, "4\n")
	}

	// ファイルへフォーマットした文字列を出力
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%d\n", n)

	// ファイルから読み込んで標準出力へ出力
	f, err = os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fscanf(f, "%d", &n)
	fmt.Println(n)
	if f != nil {
		os.Remove("test.txt")
	}
}
