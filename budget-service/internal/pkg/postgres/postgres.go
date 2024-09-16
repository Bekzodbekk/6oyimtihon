package postgres

import (
	"budget-service/internal/pkg/load"
	dbs "budget-service/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(cfg load.Config) (*sql.DB, error) {
	target := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
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

func NewQueries(db *sql.DB) *dbs.Queries {
	return dbs.New(db)
}
