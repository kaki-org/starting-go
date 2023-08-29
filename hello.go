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

	_, err := os.Open("no-such-file")
	if err != nil {
		log.Fatal(err) // os.Exit(1)も呼び出される
	}
}
