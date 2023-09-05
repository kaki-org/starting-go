package main

import (
	"net/url"
	"testing"
)

func TestUrlParse(t *testing.T) {
	u, err := url.Parse("https://www.example.com/search?a=1&b=2#top")
	if err != nil {
		t.Fatal(err)
	}
	// url.URL型のフィールド
	expect(t, u.Scheme, "https")
	expect(t, u.Host, "www.example.com")
	expect(t, u.Path, "/search")
	expect(t, u.RawQuery, "a=1&b=2")
	expect(t, u.Fragment, "top")

	// url.URL型のメソッド
	expect(t, u.String(), "https://www.example.com/search?a=1&b=2#top")
	expect(t, u.IsAbs(), true)

	// TODO: 何故かエラーになる
	// expectEqual(t, u.Query(), map[string][]string{"a": []string{"1"}, "b": []string{"2"}})
}

func TestUrlCreate(t *testing.T) {
	u := &url.URL{} // url.URL型構造体のポインタ

	u.Scheme = "https"
	u.Host = "www.example.com"
	q := u.Query()
	q.Set("q", "Go言語")
	u.RawQuery = q.Encode()
	expect(t, u.String(), "https://www.example.com?q=Go%E8%A8%80%E8%AA%9E")
}
