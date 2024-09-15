-- name: GetTotalReports :one
SELECT
    COALESCE(SUM(CASE WHEN type = 'income' THEN amount END), 0) AS total_income,
    COALESCE(SUM(CASE WHEN type = 'expense' THEN amount END), 0) AS total_expenses,
    COALESCE(SUM(CASE WHEN type = 'income' THEN amount END), 0) -
    COALESCE(SUM(CASE WHEN type = 'expense' THEN amount END), 0) AS net_savings
FROM
    transactions;

-- name: GetReportsSpendingByCategory :many
SELECT
    c.name AS category,
    COALESCE(SUM(t.amount), 0) AS totalSpent
FROM
    transactions t
JOIN
    categories c ON t.category_id = c.id
WHERE
    t.type = 'expense'
GROUP BY
    c.name;
