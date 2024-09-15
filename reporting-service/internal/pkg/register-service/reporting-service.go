package registerservice

import (
	"fmt"
	"net"
	config "reporting-service/internal/pkg/load"
	"reporting-service/internal/service"
	"reporting-service/logger"

	pb "reporting-service/genproto/reportingpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.ReportingService
}

func NewGRPCReportingService(service *service.ReportingService) *Service {
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

	logger.Info("Reporting service listening on :", cfg.Service.Port)

	pb.RegisterReportingServiceServer(grpcServer, s.Service)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}
