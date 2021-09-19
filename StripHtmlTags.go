package helper

import (
	strip "github.com/grokify/html-strip-tags-go"
	"strings"
)

func StripHtmlTags(cont string) string {
	strippedCont := strip.StripTags(cont)
	rmEmptySpace1 := strings.Replace(strippedCont, " &nbsp; ", "", -1 )
	rmEmptySpace2 := strings.Replace(rmEmptySpace1, "&nbsp;", "", -1 )

	return rmEmptySpace2
}
