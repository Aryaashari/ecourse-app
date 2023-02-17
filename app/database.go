package app

import (
	"database/sql"
	"time"
)

func GetNewConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/ecourse_app")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
