package db

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const SCHEMA = "biblioteca"

func OpenConnection() (*sqlx.DB, error) {
	conf := config.GetDB()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("db.OpenConnection: %w", err)
	}

	err = db.Ping()
	return db, err
}
