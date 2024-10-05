package database

import (
	"database/sql"
	"fmt"
	"go-tweet-stream/errors"
)

const (
	DB_USER     = "parag"
	DB_PASSWORD = ""
	DB_NAME     = "parag"
)

// DB set up
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	errors.CheckErr(err)

	return db
}
