package service

import (
	"mime/multipart"
	"myblog.backend/repository"
)

/* ====================================== */

type IMinIOService interface {
	UpLoadUserAvatar(userID uint, file *multipart.File) (string, int)
}

type MinIOService struct {
	userRepo *repository.MinIORepo
}

func NewMinIOService() *MinIOService {
	MinIORepo := repository.NewMinIORepo()
	return &MinIOService{MinIORepo}
}

/* ====================================== */

// Todo
func (ms *MinIOService) UpLoadUserAvatar(userID uint, file *multipart.File) (string, int) {

	return "", 0
}
