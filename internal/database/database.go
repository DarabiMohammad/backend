package database

import (
	"database/sql"
	"github.com/DarabiMohammad/backend-app/internal/config"
	_ "github.com/lib/pq"
)

func Connect() error {

	db, err := sql.Open(config.ProvideDBConfig())

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
