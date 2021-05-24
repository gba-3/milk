package infrastructure

import "github.com/jmoiron/sqlx"

type MySQL struct {
	DB *sqlx.DB
}

func NewMySQL(DB *sqlx.DB) *MySQL {
	return &MySQL{DB}
}

func (m *MySQL) Select(dest interface{}, query string, args ...interface{}) error {
	return m.DB.Select(dest, query, args...)
}
