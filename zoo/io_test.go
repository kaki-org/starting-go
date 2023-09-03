package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

/*
	 os.Fileにはio.Readerインタフェースが実装されている

		type Reader interface {
			Read(p []byte) (n int, err error)
		}
*/
func TestIoReader(t *testing.T) {
	/* fは*os.File型 */
	f, err := os.Open("../README.md")
	if err != nil {
		log.Fatal(err)
	}
	/* 入力のための[]byte型を要素数128で初期化 */
	bs := make([]byte, 128)
	/* バイト列の読み込み */
	n, _ := f.Read(bs) // nは実際に読み込んだバイト数
	fmt.Printf("%d bytes read\n", n)

	expect(t, n, 128)
}
