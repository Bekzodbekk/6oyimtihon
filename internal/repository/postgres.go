package repository

import (
	"context"
	"database/sql"
	"time"
	"user-service/genproto/userpb"
	"user-service/storage"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *storage.Queries
}

func NewUserRepo(db *storage.Queries) IUserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Register(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserResp, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = u.db.CheckUser(ctx, storage.CheckUserParams{
		Username: req.Username,
		Email:    req.Email,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return &userpb.CreateUserResp{
				Status:  false,
				Message: "Bunday user mavjud",
			}, nil
		}
	}

	user, err := u.db.CreateUser(ctx, storage.CreateUserParams{
		Username: req.Username,
		Password: string(passwordHash),
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResp{
		Status:  true,
		Message: "User successfuly created",
		User: &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
			Cud: &userpb.CUD{
				CreatedAt: user.CreatedAt.String,
				UpdatedAt: user.UpdatedAt.String,
				DeletedAt: int64(user.DeletedAt.Int32),
			},
		},
	}, nil
}
func (u *UserRepo) UpdateUser(ctx context.Context, req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	result, err := u.db.UpdateUser(ctx, storage.UpdateUserParams{
		ID: id,
		Username: req.Username,
		Password: req.Password,
		Email: req.Email,
		UpdatedAt: time.Now().String(),
	})
	return nil, nil
}
func (u *UserRepo) DeleteUser(ctx context.Context, req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error) {
	return nil, nil
}
func (u *UserRepo) GetUserById(ctx context.Context, req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error) {
	return nil, nil
}
func (u *UserRepo) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error) {
	return nil, nil
}
func (u *UserRepo) VerifyUser(ctx context.Context, req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error) {
	return nil, nil
}
