package dto

import (
	"myblog.backend/model"
)

type UserResponse struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username"`
	FullName     string  `json:"full_name"`
	RegisterDate string  `json:"register_date"`
	Bio          string  `json:"bio"`
	Role         uint8   `json:"role"`
	AvatarURL    *string `json:"avatar_url"`
	Follows      int     `json:"followers"`
	Fans         int     `json:"fans"`
}

func UserToResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		FullName:     user.FullName,
		RegisterDate: user.CreatedAt.Format("2006-01-02"),
		Bio:          user.Bio,
		Role:         user.Role,
		AvatarURL:    user.AvatarURL,
		Follows:      user.Follows,
		Fans:         user.Fans,
	}
}

func UserSliceToResponse(users []model.User) []*UserResponse {
	var responses []*UserResponse
	for _, user := range users {
		response := UserToResponse(&user)
		responses = append(responses, response)
	}
	return responses
}
