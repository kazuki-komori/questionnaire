package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func Test(c echo.Context) (err error) {
	return c.String(http.StatusOK, "ok!")
}
