package helper

import (
	"os"
	"strings"

	"github.com/kataras/iris"
)

func BaseURL() string {
	return os.Getenv("APP_URL")
}

func InitApp(ctx iris.Context) {

	paths := strings.Split(ctx.Path(), "/")
	ctx.ViewData("Title", strings.Title(paths[1]))
	ctx.ViewData("PageTitle", os.Getenv("APP_NAME"))

	if os.Getenv("APP_ENV") == "production" {
		ctx.ViewData("BaseURL", os.Getenv("APP_URL"))
	} else {
		ctx.ViewData("BaseURL", BaseURL())
	}

	ctx.Next()
}
