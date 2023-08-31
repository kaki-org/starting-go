package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeStandard(t *testing.T) {
	tm := time.Date(2023, 8, 31, 7, 15, 30, 0, time.Local)
	expect := "2023-08-31 07:15:30"
	actual := tm.Format("2006-01-02 15:04:05")
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
	fmt.Println(tm)
	fmt.Println(tm.Year())
	fmt.Println(tm.YearDay())
	fmt.Println(tm.Month())
	fmt.Println(tm.Day())
	fmt.Println(tm.Hour())
	fmt.Println(tm.Minute())
	fmt.Println(tm.Second())
	fmt.Println(tm.Nanosecond())
	fmt.Println(tm.Location())
	fmt.Println(tm.Weekday())
	fmt.Println(tm.Zone())

	fmt.Println(time.July.String())
	fmt.Println(time.Sunday.String())
}

func TestTimeDuration(t *testing.T) {
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
	/* 「2時間30分」を表すDurationを文字列から作成 */
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	/* 特定の日時を表すTimeを作成 */
	tm := time.Date(2023, 9, 1, 4, 44, 30, 0, time.Local)
	tm = tm.Add(2*time.Hour + 15*time.Minute + 30*time.Second)

	expect := "2023-09-01 07:00:00"
	actual := tm.Format("2006-01-02 15:04:05")

	// 日本時間2023/09/01 7:00:00を表すTimeが生成された事を確認
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestTimeDurationSub(t *testing.T) {
	tm := time.Date(2023, 9, 1, 7, 0, 0, 0, time.Local)
	tm2 := time.Date(2018, 8, 8, 7, 0, 0, 0, time.Local)
	d := tm.Sub(tm2)

	expect := "44400h0m0s"
	actual := d.String()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
