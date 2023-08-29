package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	_, err := os.Open("no-such-file")
	if err != nil {
		log.Fatal(err)
	}
}
