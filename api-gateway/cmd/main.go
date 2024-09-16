package main

import (
	api "api-gateway/internal/https"
	config "api-gateway/internal/pkg/load"
	_ "api-gateway/docs"

	userService "api-gateway/internal/pkg/user-service"
	incexpService "api-gateway/internal/pkg/incexp-service"
	budgetService "api-gateway/internal/pkg/budget-service"
	reportingService "api-gateway/internal/pkg/reporting-service"
	service "api-gateway/internal/service"
	"api-gateway/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}
	logger.Info("Configuration loaded successfully")

	conn1, err := userService.DialWithUserService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to User Service:", err)
	}
	logger.Info("Connected to User Service")

	conn2, err := incexpService.DialWithIncExpService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Income Expense Service:", err)
	}
	logger.Info("Connected to Income Expense Service")

	conn3, err := budgetService.DialWithBudgetService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Budget Service:", err)
	}
	logger.Info("Connected to Budget Service")

	conn4, err := reportingService.DialWithReportingService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Reporting Service:", err)
	}
	logger.Info("Connected to Reporting Service")

	clientService := service.NewServiceRepositoryClient(conn1, conn2, conn3, conn4)
	logger.Info("Service clients initialized")

	srv := api.NewGin(clientService)
	addr := fmt.Sprintf(":%s", cfg.ApiGateway.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Starting API Gateway on: ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			logger.Fatal(err)
		}
	}()
	logger.Info("Starting API Gateway on address:", addr)

	signalReceived := <-sigChan
	logger.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown error: ", err)
	}
	logger.Info("Graceful shutdown complete.")
}
