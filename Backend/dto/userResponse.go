package dto

import "myblog.backend/model"

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Bio      string `json:"bio"`
	Role     uint8  `json:"role"`
}

func UserToResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Bio:      user.Bio,
		Role:     user.Role,
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
