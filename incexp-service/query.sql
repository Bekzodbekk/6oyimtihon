-- name: RegisterIncome :one 
INSERT INTO transactions(type, amount, currency, category_id)
VALUES($1, $2, $3, $4)
RETURNING id;

-- name: RegisterExpense :one 
INSERT INTO transactions(type, amount, currency, category_id)
VALUES($1, $2, $3, $4)
RETURNING id;


-- name: GetListIncomeVSExpense :many
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
ORDER BY t.date DESC;
