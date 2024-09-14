package postgres

import (
	"database/sql"
	"fmt"
	"user-service/internal/pkg/load"
)

func InitDB(cfg load.Config) (*sql.DB, error) {
	target := fmt.Sprintf("host=%s port=%d dbname=%s username=%s password=%s sslmode=true",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database, cfg.Postgres.Username, cfg.Postgres.Password)

	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
