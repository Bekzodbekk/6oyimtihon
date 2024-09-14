package main

import (
	"log"
	"user-service/internal/pkg/load"
	"user-service/internal/pkg/postgres"
	"user-service/internal/repository"
)

func main() {
	cfg, err := load.Load("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.InitDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepo()
}