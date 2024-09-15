package service

import (
	"context"
	pb "reporting-service/genproto/reportingpb"
	"reporting-service/internal/storage"
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

	report, err := s.Queries.GetTotalReports(ctx)
	if err != nil {
		return nil, err
	}
	resp := pb.GetTotalReportsResponse{
		TotalIncome:   report.TotalIncome.(int64),
		TotalExpenses: report.TotalExpenses.(int64),
		NetSavings:    int64(report.NetSavings),
	}
	return &resp, nil
}

func (s *ReportingService) GetReportsSpendingByCategory(ctx context.Context, req *pb.GetReportsByCategoryRequest) (*pb.GetReportsByCategoryResponse, error) {

	resp := pb.GetReportsByCategoryResponse{}
	reports, err := s.Queries.GetReportsSpendingByCategory(ctx)
	if err != nil {
		return nil, err
	}

	for _, report := range reports {
		item := &pb.SpendingByCategory{
			Category:   report.Category,
			TotalSpent: report.Totalspent.(int64),
		}
		resp.List = append(resp.List, item)
	}
	return &resp, nil
}
