package helper

import (
	"github.com/google/uuid"
	"github.com/leekchan/accounting"
	"time"
)

func Uuid() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func MySqlTimeStamp() time.Time {
	t := time.Now() //.In(loc)
	mysqlf := "2006-01-02 15:04:05"
	ts := t.Format(mysqlf)
	myTime, _ := time.Parse(mysqlf, ts)

	return myTime
}

func FormatNumber(number int) string {
	ac := accounting.Accounting{Symbol: "", Precision: 0}
	return ac.FormatMoney(number)
}

func FormatMoney(number float64) string {
	ac := accounting.Accounting{Symbol: "", Precision: 2}
	return ac.FormatMoney(number)
}
