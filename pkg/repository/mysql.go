package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:qwerty@tcp(127.0.0.1:3306)/library")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
