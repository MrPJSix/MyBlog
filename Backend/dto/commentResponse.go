package dto

import (
	"myblog.backend/model"
)

type CommentResponse struct {
	ID              uint               `json:"id"`
	CreateAt        int64              `json:"create_at"`
	Content         string             `json:"content"`
	User            User               `json:"user"`
	ArticleID       uint               `json:"article_id"`
	RootCommentID   uint               `json:"root_comment_id"`
	ParentCommentID *uint              `json:"parent_comment_id"`
	RepliedUser     *User              `json:"replied_user"`
	Replies         []*CommentResponse `json:"replies"`
	TotalReplies    int                `json:"total_replies"`
}

type User struct {
	UserID    uint    `json:"user_id"`
	FullName  *string `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
}

func CommentToResponse(comment *model.Comment) *CommentResponse {
	var repliedUser *User
	if comment.RepliedUserID != nil {
		repliedUser = &User{
			UserID:    *comment.RepliedUserID,
			FullName:  comment.RepliedUser.FullName,
			AvatarURL: comment.RepliedUser.AvatarURL,
		}
	}
	replies := CommentSliceToResponse(comment.Replies)
	return &CommentResponse{
		ID:       comment.ID,
		CreateAt: comment.CreatedAt.Unix(),
		Content:  comment.Content,
		User: User{
			UserID:    comment.UserID,
			FullName:  comment.User.FullName,
			AvatarURL: comment.User.AvatarURL,
		},
		ArticleID:       comment.ArticleID,
		RootCommentID:   comment.RootCommentID,
		ParentCommentID: comment.ParentCommentID,
		RepliedUser:     repliedUser,
		Replies:         replies,
		TotalReplies:    len(replies),
	}
}

func CommentSliceToResponse(comments []model.Comment) []*CommentResponse {
	var responses []*CommentResponse
	for _, comment := range comments {
		response := CommentToResponse(&comment)
		responses = append(responses, response)
	}
	return responses
}
