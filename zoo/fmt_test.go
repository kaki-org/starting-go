package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const errorExpectFormat = "%s != %s"
const dummyFileName = "test.txt"

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
		t.Errorf(errorExpectFormat, s, "4\n")
	}

	// ファイルへフォーマットした文字列を出力
	f, err := os.Create(dummyFileName)
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
		os.Remove(dummyFileName)
	}
}

type People struct {
	Id    int
	Email string
}

func (p *People) String() string {
	return fmt.Sprintf("<%d, %s>", p.Id, p.Email)
}

func TestPercentV(t *testing.T) {
	v := fmt.Sprintf("%v\n", [3]int{1, 2, 3})
	if v != "[1 2 3]\n" {
		t.Errorf(errorExpectFormat, v, "[1 2 3]\n")
	}

	v = fmt.Sprintf("%v\n", []string{"A", "B", "C"})
	if v != "[A B C]\n" {
		t.Errorf(errorExpectFormat, v, "[A B C]\n")
	}

	v = fmt.Sprintf("%v\n", map[int]float64{1: 1.0, 2: 4.0, 3: 32.0})
	if v != "map[1:1 2:4 3:32]\n" {
		t.Errorf(errorExpectFormat, v, "map[1:1 2:4 3:32]\n")
	}

	u := &People{Id: 123, Email: "mail@example.com"}
	fmt.Printf("%s\n", u)
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)

	// (p *People) String() string が定義されている場合は、
	// 出力内容が変わる。ただし、%#vは変わらない
	v = fmt.Sprintf("%v\n", u)
	if v != "<123, mail@example.com>\n" {
		t.Errorf(errorExpectFormat, v, "<123, mail@example.com>\n")
	}

	// printの場合は文字列と隣接しない場合のみスペースで区切られる
	s := fmt.Sprint(123, 3.14, "Golang", struct{ X, Y int }{1, 2})
	if s != "123 3.14Golang{1 2}" {
		t.Errorf(errorExpectFormat, s, "1233.14Golang{1 2}")
	}

	// printlnの場合は必ずスペースで区切られる
	s = fmt.Sprintln(123, 3.14, "Golang", struct{ X, Y int }{1, 2})
	if s != "123 3.14 Golang {1 2}\n" {
		t.Errorf(errorExpectFormat, s, "123 3.14 Golang {1 2}\n")
	}
}
