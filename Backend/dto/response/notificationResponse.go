package dto

import "myblog.backend/model"

type NotificationResponse struct {
	ID              uint             `json:"id"`
	CreateAt        int64            `json:"create_at"`
	Sender          Sender           `json:"sender"`
	OriginalComment *OriginalComment `json:"original_comment"`
	Reply           Reply            `json:"reply"`
	ArticleID       uint             `json:"article_id"`
	IsRead          bool             `json:"is_read"`
}

type Receiver struct {
	ID        uint    `json:"id"`
	FullName  string  `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
}

type Sender Receiver

type OriginalComment struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

type Reply OriginalComment

func NotificationToReponse(notif *model.Notification) *NotificationResponse {
	var originalComment *OriginalComment
	if notif.CommentID != nil {
		originalComment = &OriginalComment{
			ID:      *notif.CommentID,
			Content: notif.Comment.Content,
		}
	}
	return &NotificationResponse{
		ID:       notif.ID,
		CreateAt: notif.CreateAt.Unix(),
		Sender: Sender{
			ID:        notif.SenderID,
			FullName:  notif.Sender.FullName,
			AvatarURL: notif.Sender.AvatarURL,
		},
		OriginalComment: originalComment,
		Reply: Reply{
			ID:      notif.ReplyID,
			Content: notif.Reply.Content,
		},
		ArticleID: notif.ArticleID,
		IsRead:    notif.IsRead,
	}
}

func NotificationSliceToResponse(notifs []model.Notification) []*NotificationResponse {
	var responses []*NotificationResponse
	for _, notif := range notifs {
		response := NotificationToReponse(&notif)
		responses = append(responses, response)
	}
	return responses
}
