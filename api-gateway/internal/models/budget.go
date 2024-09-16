package models

type CreateBudgetRequest struct {
    CategoryID string `json:"category_id"`
    Amount     int64  `json:"amount"`
    Currency   string `json:"currency"`
}

type CreateBudgetResponse struct {
    Message  string `json:"message"`
    BudgetID string `json:"budget_id"`
}

type GetListBudgetRequest struct{}

type Budget struct {
    BudgetID string `json:"budget_id"`
    Category string `json:"category"`
    Amount   int64  `json:"amount"`
    Spent    int64  `json:"spent"`
    Currency string `json:"currency"`
}

type GetListBudgetResponse struct {
    List []Budget `json:"list"`
}

type UpdateBudgetRequest struct {
    BudgetID string `json:"budget_id"`
    Amount   int64  `json:"amount"`
}

type UpdateBudgetResponse struct {
    Message string `json:"message"`
}
