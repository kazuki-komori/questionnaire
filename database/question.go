package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
)

type Question struct {
	Id         int
	Contents   string `json:"contents"`
	IsAnswered bool
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

func CreateQuestion(db *gorm.DB, c echo.Context) error {

	question := new(Question)
	err := c.Bind(question)
	if err != nil {
		return fmt.Errorf("params error=%w", err)
	}

	err = insert(question, db)
	if err != nil {
		return err
	}
	return nil
}

func insert(question *Question, db *gorm.DB) error {
	if !db.NewRecord(*question) {
		return fmt.Errorf("failed to create new record")
	}
	res := db.Create(question)
	if res.Error != nil {
		return fmt.Errorf("failed to create new question=%w", res.Error)
	}
	return nil
}
