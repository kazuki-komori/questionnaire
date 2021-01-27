package database

import (
	"database/sql/driver"
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

// 正常系
func TestCreateQuestion(t *testing.T) {
	// Setup Echo
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(contentsJson))
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
	defer gdb.Close()

	mock.ExpectBegin()
	mock.ExpectExec(
		"INSERT INTO `questions` (`contents`,`is_answered`,`updated_at`,`created_at`) VALUES (?,?,?,?)").
		WithArgs("質問", false, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = CreateQuestion(gdb, c); err != nil {
		t.Fatalf("failed to process=\n%v", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("db incorrect")
	}
}
