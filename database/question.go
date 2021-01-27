package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"strconv"
	"time"
)

type Question struct {
	Id         int
	Contents   string `json:"contents"`
	IsAnswered bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// 質問を登録する関数
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

// 質問を取得する関数
func GetQuestion(db *gorm.DB, c echo.Context) (*Question, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	question := Question{}
	question.Id = id

	db.First(&question)
	return &question, nil
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
