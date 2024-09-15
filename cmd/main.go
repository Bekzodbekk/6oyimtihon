package main

import (
	"log"
	"user-service/internal/pkg/load"
	"user-service/internal/pkg/postgres"
	"user-service/internal/pkg/redis"
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

	rdb, err := redis.ConnectRedis(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	queries := postgres.NewQueries(db)

	repo := repository.NewUserRepo(*cfg, queries, rdb)
	service := service.NewService(repo)
	runService := userservice.NewRunService(*service)

	log.Printf("User Service running on :%d port", cfg.UserServicePort)
	if err := runService.RUN(*cfg); err != nil {
		log.Fatal(err)
	}
}
