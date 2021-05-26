package infrastructure

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gba-3/milk/domain/entity"
	"github.com/jmoiron/sqlx"
)

func TestSelectNoArgs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).
			AddRow(1, "testname", "test@email.com", "testpass"))

	m := NewMySQL(sqlxDB)

	users := []entity.User{}
	query := "SELECT * FROM users"
	err = m.Select(&users, query)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Fatal("users table is empty in db.")
	}
}

func TestSelectWithArgs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE group=?`)).
		WithArgs(10).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password"}).AddRow(1, "testname", "test@email.com", "testpass"))

	m := NewMySQL(sqlxDB)
	users := []entity.User{}
	query := "SELECT * FROM users WHERE group=?"
	args := []interface{}{10}
	err = m.Select(&users, query, args...)
	if err != nil {
		t.Fatal(err)
	}
}
