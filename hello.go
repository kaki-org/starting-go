package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	// go build -o getargs hello.go
	/* os.Argsの要素数を表示 */
	fmt.Printf("length=%d\n", len(os.Args))

	/* os.Argsの内容を表示 */
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}

	/* ファイルの読み込み */
	bs, err := readFile("./go.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}

func readFile(filename string) ([]byte, error) {
	/* 変数fは*os.File型 */
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	/* []byte型のスライスにファイルの内容を読み込む */
	bs := make([]byte, 128)
	n, err := f.Read(bs) // nは実際に読み込んだバイト数
	fmt.Printf("%d bytes read\n", n)

	/* ファイルのオフセットを指定して読み込む */
	// n, err := f.ReadAt(bs, 10) // 10バイト目から読み込む

	/* ファイル内のシーク */
	offset, err := f.Seek(10, os.SEEK_SET) // ファイルの先頭から10バイト目にシーク
	offset, err = f.Seek(-2, os.SEEK_CUR)  // 現在位置から-2バイト先にシーク
	offset, err = f.Seek(0, os.SEEK_END)   // ファイルの末尾から0バイト目にシーク
	fmt.Printf("offset = %d\n", offset)

	/* ファイルのステータスを取得 */
	fi, err := f.Stat() // fiはos.FileInfo型
	fi.Name()           // ファイル名(string型)
	fi.Size()           // ファイルサイズ(int64型)
	fi.Mode()           // ファイルのモード(os.FileMode型)
	fi.ModTime()        // 最終更新日時(time.Time型)
	fi.IsDir()          // ディレクトリかどうか(bool型)
	fmt.Printf("name = %s\n", fi.Name())
	fmt.Printf("size = %d\n", fi.Size())
	fmt.Printf("mode = %v\n", fi.Mode())
	fmt.Printf("modtime = %v\n", fi.ModTime())
	fmt.Printf("isdir = %v\n", fi.IsDir())
	return bs, err
}
