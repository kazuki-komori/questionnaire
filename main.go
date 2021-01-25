package main

import (
	"fmt"
	"github.com/kazuki-komori/questionnaire_server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("fuga")
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/test", handler.Test)
	e.GET("/hoge", handler.Test)
	e.Logger.Fatal(e.Start(":8080"))
}
