package dto

import "myblog.backend/model"

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,min=8,max=25" lable:"用户名"`
	Password        string `json:"password" validate:"required,min=8,max=25" lable:"密码"`
	FullName        string `json:"full_name"`
	ConfirmPassword string `json:"confirm_password"`
}

func RegisterRequestToUser(rq *RegisterRequest) *model.User {
	return &model.User{
		Username: rq.Username,
		Password: rq.Password,
		FullName: rq.FullName,
	}
}
