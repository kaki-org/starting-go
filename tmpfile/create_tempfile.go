package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	/* プリフィックス"foo"でテンポラリファイルをオープン */
	f, err := ioutil.TempFile(os.TempDir(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// テンポラリファイルに書き込み
	f.WriteString("Hello, World\n")

	fmt.Println(f.Name())
}
