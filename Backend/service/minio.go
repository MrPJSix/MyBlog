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
	userRepo *repository.UserRepo
}

func NewMinIOService() *MinIOService {
	userRepo := repository.NewUserRepo()
	return &MinIOService{userRepo}
}

/* ====================================== */

// Todo
func (ms *MinIOService) UpLoadUserAvatar(userID uint, file *multipart.File) (string, int) {

	return "", 0
}
