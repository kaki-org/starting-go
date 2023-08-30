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
