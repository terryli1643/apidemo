package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var DSN string

func conn() *sql.DB {
	if db == nil {
		d, err := sql.Open("mysql", DSN)
		if err != nil {
			panic(err)
		}
		db = d
	}
	return db
}
