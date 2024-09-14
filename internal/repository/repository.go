package repository

import "user-service/genproto/userpb"

type IUserRepository interface{
	Register(req *userpb.CreateUserReq)(*userpb.CreateUserResp, error)
    UpdateUser(req *userpb.UpdateUserReq)(*userpb.UpdateUserResp, error)
    DeleteUser(req *userpb.DeleteUserReq)(*userpb.DeleteUserResp, error)
    GetUserById(req *userpb.GetUserByIdReq)(*userpb.GetUserByIdResp, error)
    Login(req *userpb.LoginReq)(*userpb.LoginResp, error)
    VerifyUser(req *userpb.VerifyUserReq)(*userpb.VerifyUserResp, error)
}