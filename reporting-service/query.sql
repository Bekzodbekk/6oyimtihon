-- name: GetTotalReports :one
SELECT
    (COALESCE(SUM(CASE WHEN type = 'income' THEN amount END), 0))::BIGINT AS total_income,
    (COALESCE(SUM(CASE WHEN type = 'expense' THEN amount END), 0))::BIGINT AS total_expenses,
    (COALESCE(SUM(CASE WHEN type = 'income' THEN amount END), 0) -
    COALESCE(SUM(CASE WHEN type = 'expense' THEN amount END), 0))::BIGINT AS net_savings
FROM
    transactions;


-- name: GetReportsSpendingByCategory :many
SELECT
    c.name AS category,
    (COALESCE(SUM(t.amount), 0))::BIGINT AS totalSpent
FROM
    transactions t
JOIN
    categories c ON t.category_id = c.id
GROUP BY
    c.name;
