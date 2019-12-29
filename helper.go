package helper

import (
	"github.com/google/uuid"
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
