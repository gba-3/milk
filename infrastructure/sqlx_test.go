package infrastructure

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gba-3/milk/domain/entity"
	"github.com/jmoiron/sqlx"
)

func SetupDB() (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	return sqlxDB, mock, err
}

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

func TestExec(t *testing.T) {
	sqlxDB, mock, err := SetupDB()
	if err != nil {
		t.Fatal(err)
	}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`)).
		WithArgs("testname", "test@email.com", "testpass").
		WillReturnResult(sqlmock.NewResult(1, 6))
	mock.ExpectCommit()

	m := NewMySQL(sqlxDB)
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	name := "testname"
	email := "test@email.com"
	password := "testpass"
	err = m.Exec(query, name, email, password)
	if err != nil {
		t.Fatal(err)
	}
}
