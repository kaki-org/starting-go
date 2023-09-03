package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// strings.Readerの例 {{{
	s := `このプログラムはstrings.NewReader()関数を使って与えられた文字列を
io.Readerインターフェースと互換性のあるstrings.Reader型を生成します
なにか文字列を入力してEnterキーをおしてみてください。終了はCtrl+Dです`

	/* 文字列からstrings.Readerを生成 */
	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// }}}

	// スキャン方法を切り替える例 {{{
	osTypes := `Linux Unix 
macOS Windows Solaris
FreeBSD OpenBSD Ubuntu`

	osReader := strings.NewReader(osTypes)
	scanner = bufio.NewScanner(osReader)
	/* スキャン関数をbufio.ScanWordsに変更 */
	scanner.Split(bufio.ScanWords)
	/* 他のスキャン方法は以下の通り
	 * bufio.ScanBytes
	 * bufio.ScanRunes
	 * bufio.ScanLines (デフォルト)
	 * bufio.ScanStrings
	 */

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// }}}

	/* 標準入力をソースにしたスキャナの生成 */
	scanner = bufio.NewScanner(os.Stdin)

	/* 入力のスキャンが成功する限り繰り返すループ */
	for scanner.Scan() {
		fmt.Println("=> ", scanner.Text()) // Ctrl+Dで終了
	}

	/* スキャン中にエラーが発生した場合の処理 */
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
