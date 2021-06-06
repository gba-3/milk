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

func (m *MySQL) Exec(query string, args ...interface{}) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		return tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		return tx.Rollback()
	}
	return err
}
