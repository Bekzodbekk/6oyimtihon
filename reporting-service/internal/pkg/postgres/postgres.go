package postgres

import (
	"database/sql"
	"fmt"
	config "reporting-service/internal/pkg/load"
	"reporting-service/internal/storage"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	cfg.Postgres.Host, 
	cfg.Postgres.Port, 
	cfg.Postgres.User, 
	cfg.Postgres.Password, 
	cfg.Postgres.Database)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewReportingRepo(db *sql.DB) *storage.Queries {
	return storage.New(db) 
}