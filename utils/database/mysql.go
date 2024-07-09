package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dbVar *DBServiceVar) *sql.DB {
	db, err := sql.Open(*dbVar.Dialect, *dbVar.DbUri)
	l := dbVar.Logger.LogWithContext("database", "NewDB")

	if err != nil {
		l.Error(err)
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	l.Info("Database successfully connected ....")

	return db
}
