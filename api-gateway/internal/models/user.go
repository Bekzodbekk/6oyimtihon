package models

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type GetUserByIdRequest struct {
	UserID string `json:"user_id"`
}

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserByIdResponse struct {
	User User `json:"user"`
}

type GetUsersRequest struct{}

type GetUsersResponse struct {
	List []User `json:"list"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}
type DeleteUserRes struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type UpdateUserResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type VerifyUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type VerifyUserResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
