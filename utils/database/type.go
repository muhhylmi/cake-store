package database

import (
	"cake-store/utils/logger"
	"database/sql"
)

type DBServiceVar struct {
	Logger  *logger.Logger
	DbUri   *string
	Dialect *string
}

type DBService struct {
	DB *sql.DB
}
