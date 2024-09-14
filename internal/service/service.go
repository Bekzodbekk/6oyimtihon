package service

import (
	"context"
	"user-service/genproto/userpb"
	"user-service/internal/repository"
)

type Service struct {
	userpb.UnimplementedUserServiceServer
	repo repository.IUserRepository
}

func NewService(repo repository.IUserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (u *Service) Register(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserResp, error) {
	return u.repo.Register(ctx, req)
}
func (u *Service) UpdateUser(ctx context.Context, req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error) {
	return u.repo.UpdateUser(ctx, req)
}
func (u *Service) DeleteUser(ctx context.Context, req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error) {
	return u.repo.DeleteUser(ctx, req)
}
func (u *Service) GetUserById(ctx context.Context, req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error) {
	return u.repo.GetUserById(ctx, req)
}
func (u *Service) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error) {
	return u.repo.Login(ctx, req)
}
func (u *Service) VerifyUser(ctx context.Context, req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error) {
	return u.repo.VerifyUser(ctx, req)
}
