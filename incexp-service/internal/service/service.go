package service

import (
	"context"
	pb "incexp-service/genproto/incexpb"
	"incexp-service/internal/storage"

	"github.com/google/uuid"
)

type IncExpService struct {
	pb.UnimplementedIncExpServiceServer
	Queries *storage.Queries
}

func NewIncExpService(repo *storage.Queries) *IncExpService {
	return &IncExpService{
		Queries: repo,
	}
}

func (s *IncExpService) RegisterIncome(ctx context.Context, req *pb.RegisterIncomeRequest) (*pb.RegisterIncomeResponse, error) {

	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}
	income := storage.RegisterIncomeParams{
		Type:       "income",
		Amount:     req.Amount,
		Currency:   req.Currency,
		CategoryID: uuid.NullUUID{UUID: categoryId},
	}
	id, err := s.Queries.RegisterIncome(ctx, income)
	if err != nil {
		return nil, err
	}
	resp := pb.RegisterIncomeResponse{
		Message:       "Income logged successfully",
		TransactionId: id.String(),
	}
	return &resp, nil
}

func (s *IncExpService) RegisterExpense(ctx context.Context, req *pb.RegisterExpenseRequest) (*pb.RegisterExpenseResponse, error) {

	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}
	expense := storage.RegisterExpenseParams{
		Type:       "expense",
		Amount:     req.Amount,
		Currency:   req.Currency,
		CategoryID: uuid.NullUUID{UUID: categoryId},
	}
	
	id, err := s.Queries.RegisterExpense(ctx, expense)
	if err != nil {
		return nil, err
	}
	resp := pb.RegisterExpenseResponse{
		Message:       "Expense logged successfully",
		TransactionId: id.String(),
	}
	return &resp, nil
}

func (s *IncExpService) GetListIncomeVSExpense(ctx context.Context, req *pb.GetListIncomeVSExpenseRequest) (*pb.GetListIncomeVSExpenseResponse, error) {

	resp := pb.GetListIncomeVSExpenseResponse{}
	list, err := s.Queries.GetListIncomeVSExpense(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		item := &pb.IncomeVSExpense{
			TransactionId: v.ID.String(),
			Type:          string(v.Type),
			Amount:        v.Amount,
			Currency:      v.Currency,
			Category:      v.Category,
			Date:          v.Date.Time.String(),
		}
		resp.List = append(resp.List, item)
	}
	return &resp, nil
}
