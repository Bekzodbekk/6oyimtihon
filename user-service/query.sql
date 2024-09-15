-- name: CreateUser :one
INSERT INTO users (username, password, email)
VALUES ($1, $2, $3)
RETURNING 
    id, 
    username, 
    email, 
    created_at, 
    updated_at, 
    deleted_at;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    password = $3,
    email = $4,
    updated_at = $5
WHERE id = $1
RETURNING 
    id, 
    username, 
    email, 
    created_at, 
    updated_at, 
    deleted_at;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = $2
WHERE id = $1;


-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: CheckUser :one
SELECT id FROM users
WHERE username = $1 OR email = $2;

-- name: Login :one
SELECT id, password FROM users
WHERE username = $1;