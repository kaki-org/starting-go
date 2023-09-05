package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	// rand.Seed(256) golang 1.20以降では非推奨
	rand.NewSource(256)
	r1 := rand.Float64()
	r2 := rand.Float64()

	if r1 == r2 {
		t.Errorf("rand.NewSource() should return different values")
	}
}

func TestRandIntn(t *testing.T) {
	rand.NewSource(256)

	for i := 0; i < 100; i++ {
		r := rand.Intn(100)
		if r >= 100 {
			t.Errorf("rand.Intn() should return less than 100")
		}
	}
}

func TestRandNewSource(t *testing.T) {
	// 疑似乱数生成機のソースを現在時刻から生成
	src := rand.NewSource(time.Now().UnixNano())
	// 生成したソースを引数にして疑似乱数生成機を初期化
	rnd := rand.New(src)

	for i := 0; i < 100; i++ {
		r := rnd.Intn(100)
		if r >= 100 {
			t.Errorf("rand.Intn() should return less than 100")
		}
	}
}
