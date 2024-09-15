package registerservice

import (
	"fmt"
	"net"
	config "incexp-service/internal/pkg/load"
	"incexp-service/internal/service"
	"incexp-service/logger"

	pb "incexp-service/genproto/incexpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.IncExpService
}

func NewGRPCIncExpService(service *service.IncExpService) *Service {
	return &Service{
		Service: service,
	}
}

var grpcServer = grpc.NewServer()

func (s *Service) RUN(cfg *config.Config) error {

	addr := fmt.Sprintf(":%s", cfg.Service.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Info("Income expense service listening on :", cfg.Service.Port)

	pb.RegisterIncExpServiceServer(grpcServer, s.Service)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}
