package dto

import (
	"myblog.backend/model"
	"myblog.backend/service"
)

var commentServ = service.NewCommentService()

type CommentResponse struct {
	ID           uint               `json:"id"`
	CreateAt     int64              `json:"create_at"`
	Content      string             `json:"content"`
	User         User               `json:"user"`
	Likes        int                `json:"likes"`
	RepliedUser  *User              `json:"replied_user"`
	Replies      []*CommentResponse `json:"replies"`
	TotalReplies int                `json:"total_replies"`
	IsLiked      bool               `json:"is_liked"`
}

type User struct {
	UserID    uint    `json:"user_id"`
	FullName  string  `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
}

func CommentToResponse(comment *model.Comment) *CommentResponse {
	var repliedUser *User
	if comment.RepliedUserID != nil && comment.RepliedUser != nil {
		repliedUser = &User{
			UserID:    *comment.RepliedUserID,
			FullName:  comment.RepliedUser.FullName,
			AvatarURL: nil,
		}
	}
	replies := CommentSliceToResponse(comment.Replies)
	totalReplies := len(replies)
	if totalReplies >= 3 {
		replies = replies[:3]
	}
	return &CommentResponse{
		ID:       comment.ID,
		CreateAt: comment.CreatedAt.Unix(),
		Content:  comment.Content,
		User: User{
			UserID:    comment.UserID,
			FullName:  comment.User.FullName,
			AvatarURL: comment.User.AvatarURL,
		},
		Likes:        comment.Likes,
		RepliedUser:  repliedUser,
		Replies:      replies,
		TotalReplies: totalReplies,
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

// 加上点赞状态
func CommentLikeToResponse(comment *model.Comment, userID uint) *CommentResponse {
	var repliedUser *User
	if comment.RepliedUserID != nil && comment.RepliedUser != nil {
		repliedUser = &User{
			UserID:    *comment.RepliedUserID,
			FullName:  comment.RepliedUser.FullName,
			AvatarURL: nil,
		}
	}
	isLiked, _ := commentServ.UserIsLiked(comment.ID, userID)
	replies := CommentSliceToResponse(comment.Replies)
	totalReplies := len(replies)
	if totalReplies >= 3 {
		replies = replies[:3]
	}
	return &CommentResponse{
		ID:       comment.ID,
		CreateAt: comment.CreatedAt.Unix(),
		Content:  comment.Content,
		User: User{
			UserID:    comment.UserID,
			FullName:  comment.User.FullName,
			AvatarURL: comment.User.AvatarURL,
		},
		Likes:        comment.Likes,
		RepliedUser:  repliedUser,
		Replies:      replies,
		TotalReplies: totalReplies,
		IsLiked:      isLiked,
	}
}

func CommentLikeSliceToResponse(comments []model.Comment, userID uint) []*CommentResponse {
	var responses []*CommentResponse
	for _, comment := range comments {
		response := CommentLikeToResponse(&comment, userID)
		responses = append(responses, response)
	}
	return responses
}
