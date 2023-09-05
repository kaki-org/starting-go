package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "ABCDE")
	fmt.Println(h.Sum(nil))

	// MD5ハッシュ値から16進数の文字列を生成
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

	expect(t, s, "2ecdde3959051d913f61b14579ea136d")
}
