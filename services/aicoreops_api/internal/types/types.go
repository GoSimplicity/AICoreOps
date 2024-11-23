// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type GetUserRequest struct {
	Id int `json:"id"`
}

type GetUserListRequest struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GeneralWithDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

