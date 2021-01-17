package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(database:3306)/world")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
