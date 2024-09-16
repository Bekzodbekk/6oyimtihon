package main

import (
	budgetservice "budget-service/internal/pkg/budget-service"
	"budget-service/internal/pkg/load"
	"budget-service/internal/pkg/postgres"
	"budget-service/internal/pkg/redis"
	"budget-service/internal/repository"
	"budget-service/internal/service"
	"log"
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
	repo := repository.NewBudgetRepo(*cfg, queries, rdb)
	service := service.NewService(repo)
	runService := budgetservice.NewRunService(*service)

	log.Printf("Budget service running on :%d port", cfg.BudgetServicePort)
	if err := runService.RUN(*cfg); err != nil {
		log.Fatal(err)
	}
}
