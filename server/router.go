package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	v1 := e.Group("/api/v1")

	v1.GET("/test", Test)
	v1.GET("/get/:id", GetQuestion)
	v1.POST("/create", PostCreateQuestion)

	e.Logger.Fatal(e.Start(":8080"))
}
