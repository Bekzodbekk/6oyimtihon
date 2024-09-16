package incexpservice

import (
	pb "api-gateway/genproto/incexpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithIncExpService(cfg config.Config) (*pb.IncExpServiceClient, error) {

	target := fmt.Sprintf("%s:%s", cfg.IncExpService.Host, cfg.IncExpService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	incExpServiceClient := pb.NewIncExpServiceClient(conn)
	return &incExpServiceClient, nil
}