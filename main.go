package main

import (
	"github.com/kataras/iris/v12"
	"option_get/server"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)
	server.InitRoute(app)
	if err := app.Run(iris.Addr(":8080")); err != nil {
		panic(err)
	}
}
