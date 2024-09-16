package models

type GetTotalReportsRequest struct{}

type GetTotalReportsResponse struct {
    TotalIncome   int64 `json:"total_income"`
    TotalExpenses int64 `json:"total_expenses"`
    NetSavings    int64 `json:"net_savings"`
}

type GetReportsByCategoryRequest struct{}

type SpendingByCategory struct {
    Category   string `json:"category"`
    TotalSpent int64  `json:"total_spent"`
}

type GetReportsByCategoryResponse struct {
    List []SpendingByCategory `json:"list"`
}
