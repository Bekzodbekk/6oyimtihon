package models

type RegisterUserRequest struct {
    UserName string `json:"user_name"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

type RegisterUserResponse struct {
    UserID   string `json:"user_id"`
    UserName string `json:"user_name"`
    Email    string `json:"email"`
}

type LoginUserRequest struct {
    UserName string `json:"user_name"`
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
    UserName string `json:"user_name"`
    Email    string `json:"email"`
}

type GetUserByIdResponse struct {
    User User `json:"user"`
}

type GetUsersRequest struct {}

type GetUsersResponse struct {
    List []User `json:"list"`
}
