package main

import (
	"testing"
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
