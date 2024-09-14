package service

import (
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

func (u *Service) Register(req *userpb.CreateUserReq) (*userpb.CreateUserResp, error) {
	return u.repo.Register(req)
}
func (u *Service) UpdateUser(req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error) {
	return u.repo.UpdateUser(req)
}
func (u *Service) DeleteUser(req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error) {
	return u.repo.DeleteUser(req)
}
func (u *Service) GetUserById(req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error) {
	return u.repo.GetUserById(req)
}
func (u *Service) Login(req *userpb.LoginReq) (*userpb.LoginResp, error) {
	return u.repo.Login(req)
}
func (u *Service) VerifyUser(req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error) {
	return u.repo.VerifyUser(req)
}
