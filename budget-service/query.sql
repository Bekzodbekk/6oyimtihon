-- name: CreateBudget :one
INSERT INTO budgets (category, amount, currency)
VALUES ($1, $2, $3)
RETURNING
    id,
    category,
    amount,
    spent,
    currency,
    created_at,
    updated_at,
    deleted_at;

-- name: UpdateBudget :one
UPDATE budgets
SET category=$2, amount=$3, currency=$4, updated_at=$5
WHERE id=$1
RETURNING
    id,
    category,
    amount,
    spent,
    currency,
    created_at,
    updated_at,
    deleted_at;


-- name: DeleteBudget :exec
UPDATE budgets
SET deleted_at = $2
WHERE id = $1;

-- name: GetBudgetById :one
SELECT * FROM budgets
WHERE id = $1;

-- name: GetBudgets :many
SELECT * FROM budgets;