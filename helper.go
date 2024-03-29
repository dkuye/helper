package helper

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dkuye/random"
	"github.com/google/uuid"
	"github.com/leekchan/accounting"
	"golang.org/x/crypto/bcrypt"
)

func Uuid() string {
	return uuid.Must(uuid.NewRandom()).String()
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

func Uint64ToSting(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
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

func ShortNumber(number int) string {
	count := countDigits(number)
	if count > 3 && count < 7 {
		// between 1,000 and 999,999
		return IntToString(number/1000) + "K"
	} else if count > 6 && count < 10 {
		// between 1,000,000  and 999,999,999
		return IntToString(number/1000000) + "M"
	} else if count > 9 && count < 13 {
		// between 1,000,000,000  and 999,999,999,999
		return IntToString(number/1000000000) + "B"
	} else if count > 12 && count < 16 {
		// between 1,000,000,000,000  and 999,999,999,999,999
		return IntToString(number/1000000000) + "T"
	}
	return IntToString(number)
}

func countDigits(i int) (count int) {
	for i != 0 {

		i /= 10
		count = count + 1
	}
	return count
}

func Shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := make([]int, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func ToJson(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func Sum(a, b int) int {
	return a + b
}

func HTMLDisplay(str string) template.HTML {
	return template.HTML(str)
}

func Due(number int64) bool {
	if number > time.Now().Unix() {
		return true
	}
	return false
}

// It secure session and CSRF in production
func SecureInProduction() bool {
	if os.Getenv("APP_ENV") == "production" {
		return true
	} else {
		return false
	}
}

func JBRef() string {
	t := time.Now()
	d := t.Format("060102")
	ref := d + "-" + random.String{Upper: true, Number: true}.Gen(6)
	return ref
}

func RandomPassword() string {
	str := random.String{Lower: true, Upper: true, Number: true}.Gen(6)
	return str
}

func TitleCase(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
