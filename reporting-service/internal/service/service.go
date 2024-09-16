package service

import (
	"context"
	pb "reporting-service/genproto/reportingpb"
	"reporting-service/internal/storage"
	"reporting-service/logger"
)

type ReportingService struct {
	pb.UnimplementedReportingServiceServer
	Queries *storage.Queries
}

func NewReportingService(repo *storage.Queries) *ReportingService {
	return &ReportingService{
		Queries: repo,
	}
}

func (s *ReportingService) GetTotalReports(ctx context.Context, req *pb.GetTotalReportsRequest) (*pb.GetTotalReportsResponse, error) {

	logger.Info("GetTotalReports: Fetching total reports")

	report, err := s.Queries.GetTotalReports(ctx)
	if err != nil {
		logger.Error("GetTotalReports: Error retrieving total reports - ", err)
		return nil, err
	}

	resp := pb.GetTotalReportsResponse{
		TotalIncome:   report.TotalIncome,
		TotalExpenses: report.TotalExpenses,
		NetSavings:    int64(report.NetSavings),
	}

	logger.Info("GetTotalReports: Successfully retrieved total reports - TotalIncome: ", resp.TotalIncome, ", TotalExpenses: ", resp.TotalExpenses, ", NetSavings: ", resp.NetSavings)
	return &resp, nil
}

func (s *ReportingService) GetReportsSpendingByCategory(ctx context.Context, req *pb.GetReportsByCategoryRequest) (*pb.GetReportsByCategoryResponse, error) {

	logger.Info("GetReportsSpendingByCategory: Fetching spending reports by category")

	resp := pb.GetReportsByCategoryResponse{}
	reports, err := s.Queries.GetReportsSpendingByCategory(ctx)
	if err != nil {
		logger.Error("GetReportsSpendingByCategory: Error retrieving spending reports by category - ", err)
		return nil, err
	}

	for _, report := range reports {
		item := &pb.SpendingByCategory{
			Category:   report.Category,
			TotalSpent: report.Totalspent,
		}
		resp.List = append(resp.List, item)
	}
	
	logger.Info("GetReportsSpendingByCategory: Successfully retrieved spending reports by category - Count: ", len(resp.List))
	return &resp, nil
}
