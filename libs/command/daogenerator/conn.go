package main

import (
	"database/sql"
)

var db *sql.DB
var DSN string

func conn() *sql.DB {
	if db == nil {
		db, _ = sql.Open("mysql", DSN)
	}
	return db
}
