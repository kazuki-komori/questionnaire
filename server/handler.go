package server

import (
	"github.com/kazuki-komori/questionnaire_server/database"
	"github.com/labstack/echo"
	"net/http"
)

func Test(c echo.Context) (err error) {
	return c.String(http.StatusOK, "ok!")
}

type CreateContext struct {
	Contents string `json:"contents"`
}

func PostCreateQuestion(c echo.Context) (error error) {
	err := database.CreateQuestion(c)
	if err != nil {
		return c.String(http.StatusBadGateway, "Failed to create question.")
	}
	contents := new(CreateContext)
	_ = c.Bind(contents)
	return c.String(http.StatusOK, "created question"+contents.Contents)
}
