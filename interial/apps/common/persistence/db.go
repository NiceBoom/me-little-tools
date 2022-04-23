package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDb(driver, uri string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, uri)
	if err != nil {
		return nil, err
	}
	return db, err
}
