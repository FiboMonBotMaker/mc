package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"mc.com/mc/apis"
	"mc.com/mc/db"
)

func main() {
	// 定数のセット
	db.SetInit(os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_USER"), "db", os.Getenv("MYSQL_DATABASE"))

	// APIを設定
	e := echo.New()
	apis.SetApis(*e)

	// APIを常駐させる
	e.Logger.Fatal(e.Start(":80"))
}
