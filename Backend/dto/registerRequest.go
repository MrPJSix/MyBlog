package dto

import "myblog.backend/model"

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
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
