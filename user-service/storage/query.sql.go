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

const checkUser = `-- name: CheckUser :one
SELECT id FROM users
WHERE username = $1 OR email = $2
`

type CheckUserParams struct {
	Username string
	Email    string
}

func (q *Queries) CheckUser(ctx context.Context, arg CheckUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, checkUser, arg.Username, arg.Email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password, email)
VALUES ($1, $2, $3)
RETURNING 
    id, 
    username, 
    email, 
    created_at, 
    updated_at, 
    deleted_at
`

type CreateUserParams struct {
	Username string
	Password string
	Email    string
}

type CreateUserRow struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	DeletedAt sql.NullInt32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Email)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users
SET deleted_at = $2
WHERE id = $1
`

type DeleteUserParams struct {
	ID        uuid.UUID
	DeletedAt sql.NullInt32
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteUser, arg.ID, arg.DeletedAt)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, password, email, created_at, updated_at, deleted_at FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const login = `-- name: Login :one
SELECT id, password FROM users
WHERE username = $1
`

type LoginRow struct {
	ID       uuid.UUID
	Password string
}

func (q *Queries) Login(ctx context.Context, username string) (LoginRow, error) {
	row := q.db.QueryRowContext(ctx, login, username)
	var i LoginRow
	err := row.Scan(&i.ID, &i.Password)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
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
    deleted_at
`

type UpdateUserParams struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Email     string
	UpdatedAt sql.NullString
}

type UpdateUserRow struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	DeletedAt sql.NullInt32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.UpdatedAt,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}