package helper

import (
	"github.com/google/uuid"
	"github.com/leekchan/accounting"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
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

func MySqlTimeStampToUnixTimeStamp(mysqlTS string) int64  {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, mysqlTS)
	return t.Unix()
}

func FormatNumber(number int) string {
	ac := accounting.Accounting{Symbol: "", Precision: 0}
	return ac.FormatMoney(number)
}

func FormatMoney(number float64) string {
	ac := accounting.Accounting{Symbol: "", Precision: 2}
	return ac.FormatMoney(number)
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func IntToString(n int) string {
	s := strconv.Itoa(n)
	return s
}

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}

	return a
}

func PickIntRandomly(data []int) int {
	rand.Seed(time.Now().UnixNano())
	choosen := data[rand.Intn(len(data))]
	return choosen
}


func PickStringRandomly(data []string) string {
	rand.Seed(time.Now().UnixNano())
	choosen := data[rand.Intn(len(data))]
	return choosen
}

func PickIntRandomlyBetween(min, max int) int {
	data := MakeRange(min, max)
	return PickIntRandomly(data)
}
