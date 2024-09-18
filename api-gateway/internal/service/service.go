package service

import (
	pbBudget "api-gateway/genproto/budgetpb"
	pbincExp "api-gateway/genproto/incexpb"
	pbReporting "api-gateway/genproto/reportingpb"
	userpb "api-gateway/genproto/userpb"
	"context"
)

type ServiceRepositoryClient struct {
	userClient      userpb.UserServiceClient
	incexpClient    pbincExp.IncExpServiceClient
	budgetClient    pbBudget.BudgetServiceClient
	reportingClient pbReporting.ReportingServiceClient
}

func NewServiceRepositoryClient(
	conn1 *userpb.UserServiceClient,
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
func (s *ServiceRepositoryClient) RegisterUser(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserResp, error) {
	return s.userClient.Register(ctx, req)
}

func (s *ServiceRepositoryClient) LoginUser(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error) {
	return s.userClient.Login(ctx, req)
}

func (s *ServiceRepositoryClient) VerifyUser(ctx context.Context, req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error) {
	return s.userClient.VerifyUser(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateUser(ctx context.Context, req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error) {
	return s.userClient.UpdateUser(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteUser(ctx context.Context, req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error) {
	return s.userClient.DeleteUser(ctx, req)
}

func (s *ServiceRepositoryClient) GetUserById(ctx context.Context, req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error) {
	return s.userClient.GetUserById(ctx, req)
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
func (s *ServiceRepositoryClient) CreateBudget(ctx context.Context, req *pbBudget.CreateBudgetReq) (*pbBudget.CreateBudgetResp, error) {
	return s.budgetClient.CreateBudget(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateBudget(ctx context.Context, req *pbBudget.UpdateBudgetReq) (*pbBudget.UpdateBudgetResp, error) {
	return s.budgetClient.UpdateBudget(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteBudget(ctx context.Context, req *pbBudget.DeleteBudgetReq) (*pbBudget.DeleteBudgetResp, error) {
	return s.budgetClient.DeleteBudget(ctx, req)
}

func (s *ServiceRepositoryClient) GetBudgetById(ctx context.Context, req *pbBudget.GetBudgetByIdReq) (*pbBudget.GetBudgetByIdResp, error) {
	return s.budgetClient.GetBudgetById(ctx, req)
}

func (s *ServiceRepositoryClient) GetBudgets(ctx context.Context, req *pbBudget.GetBudgetsReq) (*pbBudget.GetBudgetsResp, error) {
	return s.budgetClient.GetBudgets(ctx, req)
}

// Reporting methods
func (s *ServiceRepositoryClient) GetTotalReports(ctx context.Context, req *pbReporting.GetTotalReportsRequest) (*pbReporting.GetTotalReportsResponse, error) {
	return s.reportingClient.GetTotalReports(ctx, req)
}

func (s *ServiceRepositoryClient) GetReportsSpendingByCategory(ctx context.Context, req *pbReporting.GetReportsByCategoryRequest) (*pbReporting.GetReportsByCategoryResponse, error) {
	return s.reportingClient.GetReportsSpendingByCategory(ctx, req)
}
