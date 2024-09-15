package repository

import (
	"context"
	"user-service/genproto/userpb"
)

type IUserRepository interface {
	Register(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserResp, error)
	UpdateUser(ctx context.Context, req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error)
	DeleteUser(ctx context.Context, req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error)
	GetUserById(ctx context.Context, req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error)
	Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error)
	VerifyUser(ctx context.Context, req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error)
}
