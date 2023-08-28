package main

import (
	"fmt"
	"testing"

	"github.com/kakikubo/starting-go/zoo/animals"
)

/* 独自定義のエラーを表す型 */
type MyError struct {
	Message string
	ErrCode int
}

/* errorインターフェースのメソッドを実装 */
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func TestMyError(t *testing.T) {
	err := RaiseError()

	expect := "エラーが発生しました"
	actual := err.Error()
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}

	/* 型アサーションによって本来の型を取り出す */
	e, ok := err.(*MyError)
	if ok {
		expect := 1234
		actual := e.ErrCode
		if actual != expect {
			t.Errorf("%d != %d", actual, expect)
		}
	} else {
		t.Errorf("Type assertion error")
	}
}

/* 文字列化できることを示すインターフェース */
type Stringify interface {
	ToString() string
}

/* 構造体型Personの定義 */
type Person2 struct {
	Name string
	Age  int
}

func (p *Person2) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

/* 構造体型	Carの定義 */
type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

func TestStringify(t *testing.T) {
	vs := []Stringify{
		&Person2{Name: "Taro", Age: 21},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}

	/* vsの要素それぞれに対して、ToString()を呼び出す */
	for _, v := range vs {
		// fmt.Println(v.ToString())
		Println(v) // 上記と同じ
	}

	Println(&Person2{Name: "Hanako", Age: 30})
	Println(&Car{Number: "YYY-0123", Model: "PX513"})
}

type TS struct {
	Id   int
	Name string
}

func (t *TS) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

func TestStringer(t *testing.T) {
	ts := &TS{Id: 1, Name: "Taro"}
	fmt.Println(ts)
}

func TestPackageInterface(t *testing.T) {
	// animals.go参照
	i := animals.NewI()

	expect := "Method1()"
	actual := i.Method1()
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}

	// i.method2() // コンパイルエラー
}
