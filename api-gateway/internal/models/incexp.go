package models

type RegisterIncomeRequest struct {
    Amount     int64  `json:"amount"`
    Currency   string `json:"currency"`
    CategoryID string `json:"category_id"`
    Date       string `json:"date"`
}

type RegisterIncomeResponse struct {
    Message       string `json:"message"`
    TransactionID string `json:"transaction_id"`
}

type RegisterExpenseRequest struct {
    Amount     int64  `json:"amount"`
    Currency   string `json:"currency"`
    CategoryID string `json:"category_id"`
    Date       string `json:"date"`
}

type RegisterExpenseResponse struct {
    Message       string `json:"message"`
    TransactionID string `json:"transaction_id"`
}

type IncomeVSExpense struct {
    TransactionID string `json:"transaction_id"`
    Type          string `json:"type"`
    Amount        int64  `json:"amount"`
    Currency      string `json:"currency"`
    Category      string `json:"category"`
    Date          string `json:"date"`
}

type GetListIncomeVSExpenseRequest struct{}

type GetListIncomeVSExpenseResponse struct {
    List []IncomeVSExpense `json:"list"`
}
