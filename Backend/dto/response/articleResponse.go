package dto

import (
	"myblog.backend/model"
	"myblog.backend/utils/totext"
	"strings"
)

type ArticleResponse struct {
	ID           uint     `json:"id"`
	CreatedAt    int64    `json:"created_at"`
	UpdatedAt    int64    `json:"updated_at"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	ContentType  string   `json:"content_type"`
	Img          *string  `json:"img"`
	CommentCount int      `json:"comment_count"`
	ReadCount    int      `json:"read_count"`
	Likes        int      `json:"likes"`
	Category     Category `json:"category"`
	Author       Author   `json:"author"`
}

type Author struct {
	UserID    uint    `json:"user_id"`
	FullName  string  `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
}

type Category struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func ArticleToResponse(article *model.Article) *ArticleResponse {
	return &ArticleResponse{
		ID:           article.ID,
		CreatedAt:    article.CreatedAt.Unix(),
		UpdatedAt:    article.UpdatedAt.Unix(),
		Title:        article.Title,
		Content:      article.Content,
		ContentType:  article.ContentType,
		Img:          article.Img,
		CommentCount: article.CommentCount,
		ReadCount:    article.ReadCount,
		Likes:        article.Likes,
		Category: Category{
			CategoryID:   article.CategoryID,
			CategoryName: article.Category.Name,
		},
		Author: Author{
			UserID:    article.UserID,
			FullName:  article.User.FullName,
			AvatarURL: article.User.AvatarURL,
		},
	}
}

func articleToResponseSlice(article *model.Article) *ArticleResponse {
	var content []rune
	if article.ContentType == "h" {
		content = []rune(totext.StripHTMLTags(&article.Content))
	} else {
		content = []rune(totext.MarkdownToText(&article.Content))
	}
	if len(content) > 200 {
		content = content[:200]
	}
	return &ArticleResponse{
		ID:           article.ID,
		CreatedAt:    article.CreatedAt.Unix(),
		UpdatedAt:    article.UpdatedAt.Unix(),
		Title:        article.Title,
		Content:      strings.ReplaceAll(string(content)+"...", "\n", " "),
		ContentType:  article.ContentType,
		Img:          article.Img,
		CommentCount: article.CommentCount,
		ReadCount:    article.ReadCount,
		Likes:        article.Likes,
		Category: Category{
			CategoryID:   article.CategoryID,
			CategoryName: article.Category.Name,
		},
		Author: Author{
			UserID:    article.UserID,
			FullName:  article.User.FullName,
			AvatarURL: article.User.AvatarURL,
		},
	}
}

func ArticleSliceToResponse(articles []model.Article) []*ArticleResponse {
	var responses []*ArticleResponse
	for _, article := range articles {
		response := articleToResponseSlice(&article)
		responses = append(responses, response)
	}
	return responses
}
