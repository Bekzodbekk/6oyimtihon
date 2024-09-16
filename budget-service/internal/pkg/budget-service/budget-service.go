package budgetservice

import (
	"budget-service/genproto/budgetpb"
	"budget-service/internal/pkg/load"
	"budget-service/internal/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type RunService struct {
	serv *service.Service
}

func NewRunService(serv service.Service) *RunService {
	return &RunService{
		serv: &serv,
	}
}

func (s *RunService) RUN(cfg load.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.BudgetServiceHost, cfg.BudgetServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	newServer := grpc.NewServer()
	budgetpb.RegisterBudgetServiceServer(newServer, s.serv)

	return newServer.Serve(listener)
}
