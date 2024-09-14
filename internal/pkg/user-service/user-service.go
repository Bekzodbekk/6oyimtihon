package userservice

import (
	"fmt"
	"net"
	"user-service/genproto/userpb"
	"user-service/internal/pkg/load"
	"user-service/internal/service"

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
	target := fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	newServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(newServer, s.serv)

	return newServer.Serve(listener)
}
