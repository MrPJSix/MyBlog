package service

import (
	"log"
	"mime/multipart"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/img"
)

/* ====================================== */

type IMinIOService interface {
	UpLoadUserAvatar(userID uint, file *multipart.File) (string, int)
}

type MinIOService struct {
	minIORepo *repository.MinIORepo
}

func NewMinIOService() *MinIOService {
	minIORepo := repository.NewMinIORepo()
	return &MinIOService{minIORepo}
}

/* ====================================== */

// Todo
func (ms *MinIOService) UpLoadUserAvatar(userID uint, file *multipart.File) (string, int) {
	avatarFile, avatarFileName, contentType, err := img.ProcessAvatar(userID, file)
	if err != nil {
		log.Println("文件格式转换失败", err)
		return "", errmsg.ERROR_UPLOAD_USERAVT
	}
	avatarURL, code := ms.minIORepo.UpLoadImg(userID, avatarFileName, avatarFile, contentType)
	return avatarURL, code
}
