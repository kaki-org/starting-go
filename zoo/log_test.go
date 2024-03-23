package main

import (
	"log"
	"os"
	"testing"
)

func TestLogging(t *testing.T) {
	log.Print("ログの1行目\n")
	log.Println("ログの2行目")
	log.Printf(logNumberFormat, 3)

	// テストのロギングにはlog.Printlnではなくt.Logを使った方が良い理由
	// https://qiita.com/croquette0212/items/be2e19c4af45d6dab3ab
	t.Log("テストのログ")
	t.Logf("テストの%d行目", 2)

	// ログの出力先を変更
	log.SetOutput(os.Stdout)
	log.Printf(logNumberFormat, 4)

	/* test.logファイルを作成 */
	f, err := os.Create("test.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	log.Printf(logNumberFormat, 5)
	defer f.Close()
}

func TestLoggingFlags(t *testing.T) {
	// 標準ログのフォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("A")

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	// 時刻とファイルの行番号(フルパス)
	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")

	// 標準ログのフォーマットに戻す
	log.SetFlags(log.LstdFlags)

	// ログのプリフィックスを設定
	log.SetPrefix("[LOG] ")
	log.Println("E")
}

func TestNewLogger(t *testing.T) {
	logger := log.New(os.Stdout, "[LOG] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("This is a log message with new logger")
}
