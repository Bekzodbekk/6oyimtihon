package repository

import (
	"budget-service/genproto/budgetpb"
	"context"
)

type IBudgetRepository interface {
	CreateBudget(ctx context.Context, req *budgetpb.CreateBudgetReq) (*budgetpb.CreateBudgetResp, error)
	UpdateBudget(ctx context.Context, req *budgetpb.UpdateBudgetReq) (*budgetpb.UpdateBudgetResp, error)
	DeleteBudget(ctx context.Context, req *budgetpb.DeleteBudgetReq) (*budgetpb.DeleteBudgetResp, error)
	GetBudgetById(ctx context.Context, req *budgetpb.GetBudgetByIdReq) (*budgetpb.GetBudgetByIdResp, error)
	GetBudgets(ctx context.Context, req *budgetpb.GetBudgetsReq) (*budgetpb.GetBudgetsResp, error)
}
