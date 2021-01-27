package database

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var contentsJson = `{"contents": "質問"}`

func setup(t *testing.T, req *http.Request) (*gorm.DB, sqlmock.Sqlmock, echo.Context) {
	// Setup Echo
	e := echo.New()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Setup DB
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("cannnot create gorm DB")
	}
	return gdb, mock, c
}

// 正常系
func TestCreateQuestion(t *testing.T) {
	//setup
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(contentsJson))
	gdb, mock, c := setup(t, req)
	defer gdb.Close()

	mock.ExpectBegin()
	mock.ExpectExec(
		"INSERT INTO `questions` (`contents`,`is_answered`,`created_at`,`updated_at`) VALUES (?,?,?,?)").
		WithArgs("質問", false, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := CreateQuestion(gdb, c); err != nil {
		t.Fatalf("failed to process=\n%v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("db incorrect")
	}
}

// 正常系
func TestGetQuestion(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	gdb, _, c := setup(t, req)
	c.SetPath("/get/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	defer gdb.Close()

	db, err := NewDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	question, err := GetQuestion(db, c)
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(question)
}
