package main

import (
	"encoding/json"
	"testing"
)

type JsonUser struct {
	Id    int
	Name  string
	Email string
}

func TestJsonEncode(t *testing.T) {
	// JsonUser構造体の初期化
	u := new(JsonUser)
	u.Id = 1
	u.Name = "山田太郎"
	u.Email = "yamada@example.com"

	// JSONエンコード
	b, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}

	expect(t, string(b), `{"Id":1,"Name":"山田太郎","Email":"yamada@example.com"}`)

	// 新たにJsonUser型を初期化
	u2 := new(JsonUser)

	// JSONデコード
	err = json.Unmarshal([]byte(b), u2)
	if err != nil {
		t.Fatal(err)
	}

	expect(t, u2.Id, 1)
	expect(t, u2.Name, "山田太郎")
	expect(t, u2.Email, "yamada@example.com")
}
