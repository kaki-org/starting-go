package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// GETメソッドでURLにアクセス
	res, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	// *http.Response型のフィールド
	fmt.Println(res.StatusCode) // 200
	fmt.Println(res.Proto)      // HTTP/2.0
	// レスポンスヘッダの情報
	fmt.Println(res.Header["Date"])         // [Tue, 05 Sep 2023 20:48:04 GMT]
	fmt.Println(res.Header["Content-Type"]) // [text/html; charset=UTF-8]
	// リクエスト情報
	fmt.Println(res.Request.Method) // GET
	fmt.Println(res.Request.URL)    // https://example.com
	// レスポンスボディの読み込み
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
