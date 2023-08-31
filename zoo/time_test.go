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

func TestTimeBeforeAfterEqual(t *testing.T) {
	t0 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.Local)
	t1 := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)

	// t0 < t1 → true
	if t1.Before(t0) {
		t.Errorf("t1.Before(t0) is false")
	}
	// t1 < t0 → false
	if !t1.After(t0) {
		t.Errorf("!t1.After(t0) is false")
	}
	// t0 > t1 → false
	if !t0.Before(t1) {
		t.Errorf("!t0.Before(t1) is false")
	}
	// t0 < t1 → true
	if t0.After(t1) {
		t.Errorf("t0.After(t1) is false")
	}

	jst9 := time.Date(2015, 10, 1, 9, 0, 0, 0, time.Local) // 2015/10/1 9:00:00(JST)
	utc0 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)   // 2015/10/1 0:00:00(UTC)

	if !jst9.Equal(utc0) {
		t.Errorf("!jst9.Equal(utc0) is false")
	}
}

func TestTimeAddDate(t *testing.T) {
	tm := time.Date(2023, 9, 1, 7, 0, 0, 0, time.Local)
	tm2 := tm.AddDate(1, 2, 3)

	expect := "2024-11-04 07:00:00"
	actual := tm2.Format("2006-01-02 15:04:05")
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}

	tm3 := tm2.AddDate(-1, -2, -3)
	if tm != tm3 {
		t.Errorf("tm != tm3")
	}
}
