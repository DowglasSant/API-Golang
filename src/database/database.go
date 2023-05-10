package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	database, erro := sql.Open("mysql", config.StringConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = database.Ping(); erro != nil {
		database.Close()
		return nil, erro
	}

	return database, nil
}
