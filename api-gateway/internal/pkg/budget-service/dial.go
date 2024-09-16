package budgetservice

import (
	pb "api-gateway/genproto/budgetpb"
	config "api-gateway/internal/pkg/load"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithBudgetService(cfg config.Config) (*pb.BudgetServiceClient, error) {

	target := fmt.Sprintf("%s:%s", cfg.BudgetService.Host, cfg.BudgetService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	budgetServiceClient := pb.NewBudgetServiceClient(conn)
	return &budgetServiceClient, nil
}