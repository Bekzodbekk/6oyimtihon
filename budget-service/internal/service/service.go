package service

import (
	"budget-service/genproto/budgetpb"
	"budget-service/internal/repository"
	"context"
)

type Service struct {
	budgetpb.UnimplementedBudgetServiceServer
	repo repository.IBudgetRepository
}

func NewService(repo repository.IBudgetRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateBudget(ctx context.Context, req *budgetpb.CreateBudgetReq) (*budgetpb.CreateBudgetResp, error) {
	return s.repo.CreateBudget(ctx, req)
}
func (s *Service) UpdateBudget(ctx context.Context, req *budgetpb.UpdateBudgetReq) (*budgetpb.UpdateBudgetResp, error) {
	return s.repo.UpdateBudget(ctx, req)
}
func (s *Service) DeleteBudget(ctx context.Context, req *budgetpb.DeleteBudgetReq) (*budgetpb.DeleteBudgetResp, error) {
	return s.repo.DeleteBudget(ctx, req)
}
func (s *Service) GetBudgetById(ctx context.Context, req *budgetpb.GetBudgetByIdReq) (*budgetpb.GetBudgetByIdResp, error) {
	return s.repo.GetBudgetById(ctx, req)
}
func (s *Service) GetBudgets(ctx context.Context, req *budgetpb.GetBudgetsReq) (*budgetpb.GetBudgetsResp, error) {
	return s.repo.GetBudgets(ctx, req)
}
