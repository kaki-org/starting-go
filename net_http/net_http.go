package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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
	htmlContents := string(body)

	f, err := os.Create("example.com.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(htmlContents)

	postRequest()
	upload()
}

func postRequest() {
	// POSTに送信するデータを生成
	vs := url.Values{}
	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode()) // id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83%BC%E3%82%B8

	// POSTメソッドを実行
	res, err := http.PostForm("https://example.com/comments/post", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode) // 404

}

func upload() {
	// 画像ファイルをオープン
	f, err := os.Open("../unnamed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// POSTメソッドによる画像ファイルのアップロード
	res, err := http.Post("https://example.com/upload", "image/jpeg", f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.StatusCode) // 404
}
