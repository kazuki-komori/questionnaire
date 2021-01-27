package server

import (
	"encoding/json"
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
	db, err := database.NewDB()
	if err != nil {
		return c.String(http.StatusBadGateway, "Failed to connect DB")
	}
	defer db.Close()

	err = database.CreateQuestion(db, c)
	if err != nil {
		return c.String(http.StatusBadGateway, "Failed to create question.")
	}
	contents := new(CreateContext)
	_ = c.Bind(contents)
	return c.String(http.StatusOK, "created question"+contents.Contents)
}

func GetQuestion(c echo.Context) (error error) {
	db, err := database.NewDB()
	if err != nil {
		return c.String(http.StatusBadGateway, "Failed to connect DB")
	}
	defer db.Close()

	question, err := database.GetQuestion(db, c)
	if err != nil {
		return c.String(http.StatusBadGateway, "Failed to create question.")
	}
	res, _ := json.Marshal(question)
	return c.String(http.StatusOK, string(res))
}
