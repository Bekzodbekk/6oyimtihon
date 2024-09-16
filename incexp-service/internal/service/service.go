package service

import (
	"context"
	pb "incexp-service/genproto/incexpb"
	"incexp-service/internal/storage"
	"incexp-service/logger"

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
		logger.Error("RegisterIncome: Error parsing category ID - ", err)
		return nil, err
	}

	logger.Info("RegisterIncome: Category ID parsed - ", categoryId)

	income := storage.RegisterIncomeParams{
		Type:       "income",
		Amount:     req.Amount,
		Currency:   req.Currency,
		CategoryID: uuid.NullUUID{UUID: categoryId, Valid: true},
	}

	id, err := s.Queries.RegisterIncome(ctx, income)
	if err != nil {
		logger.Error("RegisterIncome: Error registering income - ", err)
		return nil, err
	}

	resp := pb.RegisterIncomeResponse{
		Message:       "Income logged successfully",
		TransactionId: id.String(),
	}

	logger.Info("RegisterIncome: Income registered successfully - Transaction ID: ", resp.TransactionId)
	return &resp, nil
}

func (s *IncExpService) RegisterExpense(ctx context.Context, req *pb.RegisterExpenseRequest) (*pb.RegisterExpenseResponse, error) {

	categoryId, err := uuid.Parse(req.CategoryId)
	if err != nil {
		logger.Error("RegisterExpense: Error parsing category ID - ", err)
		return nil, err
	}

	logger.Info("RegisterExpense: Category ID parsed - ", categoryId)

	expense := storage.RegisterExpenseParams{
		Type:       "expense",
		Amount:     req.Amount,
		Currency:   req.Currency,
		CategoryID: uuid.NullUUID{UUID: categoryId, Valid: true},
	}

	id, err := s.Queries.RegisterExpense(ctx, expense)
	if err != nil {
		logger.Error("RegisterExpense: Error registering expense - ", err)
		return nil, err
	}

	resp := pb.RegisterExpenseResponse{
		Message:       "Expense logged successfully",
		TransactionId: id.String(),
	}

	logger.Info("RegisterExpense: Expense registered successfully - Transaction ID: ", resp.TransactionId)
	return &resp, nil
}

func (s *IncExpService) GetListIncomeVSExpense(ctx context.Context, req *pb.GetListIncomeVSExpenseRequest) (*pb.GetListIncomeVSExpenseResponse, error) {

	logger.Info("GetListIncomeVSExpense: Fetching income vs expense list")

	resp := pb.GetListIncomeVSExpenseResponse{}
	list, err := s.Queries.GetListIncomeVSExpense(ctx)
	if err != nil {
		logger.Error("GetListIncomeVSExpense: Error retrieving income vs expense list - ", err)
		return nil, err
	}
	
	logger.Info("GetListIncomeVSExpense: Successfully retrieved income vs expense list - Count: ", len(list))

	for _, v := range list {
		item := &pb.IncomeVSExpense{
			TransactionId: v.ID.String(),
			Type:          string(v.Type),
			Amount:        v.Amount,
			Currency:      v.Currency,
			Category:      v.Category,
			Date:          v.Date.Time.Format("2006-01-02 15:04:05"),
		}
		resp.List = append(resp.List, item)
	}

	return &resp, nil
}
