package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDb(config string) error {
	database, err := sql.Open("mysql", config)
	if err != nil {
		return err
	}
	DB = database
	return nil
}
