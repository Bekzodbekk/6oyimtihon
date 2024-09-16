package reportingservice

import (
	pb "api-gateway/genproto/reportingpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithReportingService(cfg config.Config) (*pb.ReportingServiceClient, error) {

	target := fmt.Sprintf("%s:%s", cfg.ReportingService.Host, cfg.ReportingService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	reportingServiceClient := pb.NewReportingServiceClient(conn)
	return &reportingServiceClient, nil
}