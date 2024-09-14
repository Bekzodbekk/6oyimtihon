package main

import (
	"log"
	"user-service/internal/pkg/load"
	"user-service/internal/pkg/postgres"
	userservice "user-service/internal/pkg/user-service"
	"user-service/internal/repository"
	"user-service/internal/service"

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
	
	queries := postgres.NewQueries(db)

	repo := repository.NewUserRepo(queries)
	service := service.NewService(repo)
	runService := userservice.NewRunService(*service)

	if err := runService.RUN(*cfg); err != nil{
		log.Fatal(err)
	}
}