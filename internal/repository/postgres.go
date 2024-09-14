package repository

import (
	"database/sql"
	"user-service/genproto/userpb"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Register(req *userpb.CreateUserReq) (*userpb.CreateUserResp, error) {
	return nil, nil
}
func (u *UserRepo) UpdateUser(req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error) {
	return nil, nil
}
func (u *UserRepo) DeleteUser(req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error) {
	return nil, nil
}
func (u *UserRepo) GetUserById(req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error) {
	return nil, nil
}
func (u *UserRepo) Login(req *userpb.LoginReq) (*userpb.LoginResp, error) {
	return nil, nil
}
func (u *UserRepo) VerifyUser(req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error) {
	return nil, nil
}
