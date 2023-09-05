package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "ABCDE")
	fmt.Println(h.Sum(nil))

	// MD5ハッシュ値から16進数の文字列を生成
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

	expect(t, s, "2ecdde3959051d913f61b14579ea136d")
}

func TestSha1(t *testing.T) {
	s1 := sha1.New()
	io.WriteString(s1, "ABCDE")
	fmt.Println(s1.Sum(nil))

	s1hash := fmt.Sprintf("%x", s1.Sum(nil))
	expect(t, s1hash, "7be07aaf460d593a323d0db33da05b64bfdcb3a5")
}

func TestSha256(t *testing.T) {
	s256 := sha256.New()
	io.WriteString(s256, "ABCDE")
	fmt.Println(s256.Sum(nil))

	s256hash := fmt.Sprintf("%x", s256.Sum(nil))
	expect(t, s256hash, "f0393febe8baaa55e32f7be2a7cc180bf34e52137d99e056c817a9c07b8f239a")
}

func TestSha512(t *testing.T) {
	s512 := sha512.New()
	io.WriteString(s512, "ABCDE")
	fmt.Println(s512.Sum(nil))

	s512hash := fmt.Sprintf("%x", s512.Sum(nil))
	expect(t, s512hash, "9989a8fcbc29044b5883a0a36c146fe7415b1439e995b4d806ea0af7da9ca4390eb92a604b3ecfa3d75f9911c768fbe2aecc59eff1e48dcaeca1957bdde01dfb")
}
