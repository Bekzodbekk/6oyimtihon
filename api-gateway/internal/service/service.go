package service

import (
	pbBudget "api-gateway/genproto/budgetpb"
	pbincExp "api-gateway/genproto/incexpb"
	pbReporting "api-gateway/genproto/reportingpb"
	pbUser "api-gateway/genproto/userpb"
	"context"
)

type ServiceRepositoryClient struct {
	userClient      pbUser.UserServiceClient
	incexpClient    pbincExp.IncExpServiceClient
	budgetClient    pbBudget.BudgetServiceClient
	reportingClient pbReporting.ReportingServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pbUser.UserServiceClient,
	conn2 *pbincExp.IncExpServiceClient,
	conn3 *pbBudget.BudgetServiceClient,
	conn4 *pbReporting.ReportingServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		userClient:      *conn1,
		incexpClient:    *conn2,
		budgetClient:    *conn3,
		reportingClient: *conn4,
	}
}

// User methods
func (s *ServiceRepositoryClient) RegisterUser(ctx context.Context, req *pbUser.RegisterUserRequest) (*pbUser.RegisterUserResponse, error) {
	return s.userClient.RegisterUser(ctx, req)
}

func (s *ServiceRepositoryClient) LoginUser(ctx context.Context, req *pbUser.LoginUserRequest) (*pbUser.LoginUserResponse, error) {
	return s.userClient.LoginUser(ctx, req)
}

func (s *ServiceRepositoryClient) GetUserById(ctx context.Context, req *pbUser.GetUserByIdRequest) (*pbUser.GetUserByIdResponse, error) {
	return s.userClient.GetUserById(ctx, req)
}

func (s *ServiceRepositoryClient) GetUsersList(ctx context.Context, req *pbUser.GetUsersRequest) (*pbUser.GetUsersResponse, error) {
	return s.userClient.GetUsersList(ctx, req)
} 

// Income Expense methods
func (s *ServiceRepositoryClient) RegisterIncome(ctx context.Context, req *pbincExp.RegisterIncomeRequest) (*pbincExp.RegisterIncomeResponse, error) {
	return s.incexpClient.RegisterIncome(ctx, req)
}

func (s *ServiceRepositoryClient) RegisterExpense(ctx context.Context, req *pbincExp.RegisterExpenseRequest) (*pbincExp.RegisterExpenseResponse, error) {
	return s.incexpClient.RegisterExpense(ctx, req)
}

func (s *ServiceRepositoryClient) GetListIncomeVSExpense(ctx context.Context, req *pbincExp.GetListIncomeVSExpenseRequest) (*pbincExp.GetListIncomeVSExpenseResponse, error) {
	return s.incexpClient.GetListIncomeVSExpense(ctx, req)
}

// Budget methods
func (s *ServiceRepositoryClient) CreateBudget(ctx context.Context, req *pbBudget.CreateBudgetRequest) (*pbBudget.CreateBudgetResponse, error) {
	return s.budgetClient.CreateBudget(ctx, req)
}

func (s *ServiceRepositoryClient) GetListBudget(ctx context.Context, req *pbBudget.GetListBudgetRequest) (*pbBudget.GetListBudgetResponse, error) {
	return s.budgetClient.GetListBudget(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateBudget(ctx context.Context, req *pbBudget.UpdateBudgetRequest) (*pbBudget.UpdateBudgetResponse, error) {
	return s.budgetClient.UpdateBudget(ctx, req)
}

// Reporting methods
func (s *ServiceRepositoryClient) GetTotalReports(ctx context.Context, req *pbReporting.GetTotalReportsRequest) (*pbReporting.GetTotalReportsResponse, error) {
	return s.reportingClient.GetTotalReports(ctx, req)
}

func (s *ServiceRepositoryClient) GetReportsSpendingByCategory(ctx context.Context, req *pbReporting.GetReportsByCategoryRequest) (*pbReporting.GetReportsByCategoryResponse, error) {
	return s.reportingClient.GetReportsSpendingByCategory(ctx, req)
}