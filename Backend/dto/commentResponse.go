package dto

import "myblog.backend/model"

type CommentResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	ArticleID uint   `json:"article_id"`
	User      User   `json:"user"`
}

type User struct {
	UserID   uint   `json:"user_id"`
	FullName string `json:"full_name"`
}

func CommentToResponse(comment *model.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		ArticleID: comment.ArticleID,
		User: User{
			UserID:   comment.UserID,
			FullName: comment.User.FullName,
		},
	}
}

func CommentSliceToResponse(comments []model.Comment) []CommentResponse {
	var responses []CommentResponse
	for _, comment := range comments {
		response := CommentToResponse(&comment)
		responses = append(responses, response)
	}
	return responses
}
