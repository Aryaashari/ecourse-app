package app

import (
	"database/sql"
	"ecourse-app/helper"
	"time"
)

func GetNewConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/ecourse_app")
	helper.PanicError(&err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
