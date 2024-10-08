// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package storage

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getListIncomeVSExpense = `-- name: GetListIncomeVSExpense :many
SELECT 
    t.id, 
    t.type, 
    t.amount, 
    t.currency, 
    c.name AS category, 
    t.date
FROM transactions AS t
INNER JOIN categories AS c 
ON t.category_id = c.id
ORDER BY t.date DESC
`

type GetListIncomeVSExpenseRow struct {
	ID       uuid.UUID
	Type     TransactionType
	Amount   int64
	Currency string
	Category string
	Date     sql.NullTime
}

func (q *Queries) GetListIncomeVSExpense(ctx context.Context) ([]GetListIncomeVSExpenseRow, error) {
	rows, err := q.db.QueryContext(ctx, getListIncomeVSExpense)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetListIncomeVSExpenseRow
	for rows.Next() {
		var i GetListIncomeVSExpenseRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Amount,
			&i.Currency,
			&i.Category,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const registerExpense = `-- name: RegisterExpense :one
INSERT INTO transactions(type, amount, currency, category_id)
VALUES($1, $2, $3, $4)
RETURNING id
`

type RegisterExpenseParams struct {
	Type       TransactionType
	Amount     int64
	Currency   string
	CategoryID uuid.NullUUID
}

func (q *Queries) RegisterExpense(ctx context.Context, arg RegisterExpenseParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, registerExpense,
		arg.Type,
		arg.Amount,
		arg.Currency,
		arg.CategoryID,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const registerIncome = `-- name: RegisterIncome :one
INSERT INTO transactions(type, amount, currency, category_id)
VALUES($1, $2, $3, $4)
RETURNING id
`

type RegisterIncomeParams struct {
	Type       TransactionType
	Amount     int64
	Currency   string
	CategoryID uuid.NullUUID
}

func (q *Queries) RegisterIncome(ctx context.Context, arg RegisterIncomeParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, registerIncome,
		arg.Type,
		arg.Amount,
		arg.Currency,
		arg.CategoryID,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
