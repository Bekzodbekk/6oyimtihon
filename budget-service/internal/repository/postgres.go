package repository

import (
	"budget-service/genproto/budgetpb"
	"budget-service/internal/pkg/load"
	dbs "budget-service/storage"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type BudgetRepo struct {
	db  *dbs.Queries
	rdb *redis.Client
	cfg load.Config
}

func NewBudgetRepo(cfg load.Config, db *dbs.Queries, rdb *redis.Client) IBudgetRepository {
	return &BudgetRepo{
		db:  db,
		rdb: rdb,
		cfg: cfg,
	}
}

func (s *BudgetRepo) CreateBudget(ctx context.Context, req *budgetpb.CreateBudgetReq) (*budgetpb.CreateBudgetResp, error) {
	budget, err := s.db.CreateBudget(ctx, dbs.CreateBudgetParams{
		Category: req.Category,
		Amount:   float64(req.Amount),
		Currency: req.Currency,
	})
	if err != nil {
		return nil, err
	}

	return &budgetpb.CreateBudgetResp{
		Status:  true,
		Message: "Budget Successfuly created",
		Budget: &budgetpb.Budget{
			Id:       budget.ID.String(),
			Category: budget.Category,
			Amount:   float32(budget.Amount),
			Spent:    float32(budget.Spent.Float64),
			Currency: budget.Currency,
			Cud: &budgetpb.CUDBudget{
				CreatedAt: budget.CreatedAt.String,
				UpdatedAt: budget.UpdatedAt.String,
				DeletedAt: int64(budget.DeletedAt.Int32),
			},
		},
	}, nil
}
func (s *BudgetRepo) UpdateBudget(ctx context.Context, req *budgetpb.UpdateBudgetReq) (*budgetpb.UpdateBudgetResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	result, err := s.db.UpdateBudget(ctx, dbs.UpdateBudgetParams{
		ID:        id,
		Category:  req.Category,
		Amount:    float64(req.Amount),
		Currency:  req.Currency,
		UpdatedAt: sql.NullString{String: time.Now().String(), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &budgetpb.UpdateBudgetResp{
		Status:  true,
		Message: "Budget Successfuly created",
		Budget: &budgetpb.Budget{
			Id:       result.ID.String(),
			Category: result.Category,
			Amount:   float32(result.Amount),
			Spent:    float32(result.Spent.Float64),
			Currency: result.Currency,
			Cud: &budgetpb.CUDBudget{
				CreatedAt: result.CreatedAt.String,
				UpdatedAt: result.UpdatedAt.String,
				DeletedAt: int64(result.DeletedAt.Int32),
			},
		},
	}, nil
}
func (s *BudgetRepo) DeleteBudget(ctx context.Context, req *budgetpb.DeleteBudgetReq) (*budgetpb.DeleteBudgetResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = s.db.DeleteBudget(ctx, dbs.DeleteBudgetParams{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &budgetpb.DeleteBudgetResp{
		Status:  true,
		Message: "Budget Successfuly deleted",
	}, nil
}
func (s *BudgetRepo) GetBudgetById(ctx context.Context, req *budgetpb.GetBudgetByIdReq) (*budgetpb.GetBudgetByIdResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	result, err := s.db.GetBudgetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &budgetpb.GetBudgetByIdResp{
		Status:  true,
		Message: "Budget Successfuly created",
		Budget: &budgetpb.Budget{
			Id:       result.ID.String(),
			Category: result.Category,
			Amount:   float32(result.Amount),
			Spent:    float32(result.Spent.Float64),
			Currency: result.Currency,
			Cud: &budgetpb.CUDBudget{
				CreatedAt: result.CreatedAt.String,
				UpdatedAt: result.UpdatedAt.String,
				DeletedAt: int64(result.DeletedAt.Int32),
			},
		},
	}, nil
}
func (s *BudgetRepo) GetBudgets(ctx context.Context, req *budgetpb.GetBudgetsReq) (*budgetpb.GetBudgetsResp, error) {
	result, err := s.db.GetBudgets(ctx)
	if err != nil {
		return nil, err
	}

	budgets := []*budgetpb.Budget{}
	for _, elm := range result {
		budgets = append(budgets, &budgetpb.Budget{
			Id:       elm.ID.String(),
			Category: elm.Category,
			Amount:   float32(elm.Amount),
			Spent:    float32(elm.Spent.Float64),
			Currency: elm.Currency,
			Cud: &budgetpb.CUDBudget{
				CreatedAt: elm.CreatedAt.String,
				UpdatedAt: elm.UpdatedAt.String,
				DeletedAt: int64(elm.DeletedAt.Int32),
			},
		})
	}
	return &budgetpb.GetBudgetsResp{
		Status:  true,
		Message: "Get Budgets successfuly",
		Count:   int64(len(budgets)),
		Budget:  budgets,
	}, nil
}
