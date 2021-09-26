package helper

import (
	"fmt"
	"time"
)

func MySqlTimeStamp() time.Time {
	t := time.Now() //.In(loc)
	loc, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		fmt.Println(err)
	}
	t = t.In(loc)
	mysqlf := "2006-01-02 15:04:05"
	ts := t.Format(mysqlf)
	myTime, _ := time.Parse(mysqlf, ts)

	return myTime
}

func UnixTimeStamp() int64 {
	t := time.Now()
	loc, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		fmt.Println(err)
	}
	t = t.In(loc)
	return t.Unix()
}

func MySqlTimeStampToUnixTimeStamp(mysqlTS string) int64 {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, mysqlTS)
	return t.Unix()
}

func UnixTimeStampToMySqlTimeStamp(unix int64) string {
	if unix < 1 {
		return "-"
	}
	tm := time.Unix(unix, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func FormatDate(t time.Time) string {
	//t.String()
	suffix := "th"
	switch t.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}
	//return t.Format("Mon 2" + suffix + " Jan 2006 3:4pm")
	return t.Format("2" + suffix + " Jan 2006 3:04pm")
}

func FormatDateDynamic(t time.Time, format string) string {
	return t.Format(format)
}

func UnixToDate(unix int64) string {
	if unix < 1 {
		return "-"
	}
	if unix < time.Now().Unix() {
		return "EXPIRED"
	}
	tm := time.Unix(unix, 0)
	return tm.Format("Jan 02 2006")
}

func DateFolderName() string {
	mysqlf := "2006_01"
	t := time.Now().Format(mysqlf)
	return t
}

func DateTimeFileName() string {
	mysqlf := "2006_01_02_15_04_05"
	t := time.Now().Format(mysqlf)
	return t
}

func Deleted(deleted *time.Time) bool {
	if deleted != nil {
		return true
	} else {
		return false
	}
}
