package service

import (
	"log"
	"mime/multipart"
	"myblog.backend/repository"
	"myblog.backend/utils/errmsg"
	"myblog.backend/utils/img"
	"net/url"
	"path"
)

/* ====================================== */

type IMinIOService interface {
	parseObjectNameFromURL(inputURL string) string
	UpLoadUserAvatar(userID uint, file *multipart.File) (string, int)
	DeleteUserAvatar(userID uint)
}

type MinIOService struct {
	minIORepo *repository.MinIORepo
	userRepo  *repository.UserRepo
}

func NewMinIOService() *MinIOService {
	minIORepo := repository.NewMinIORepo()
	userRepo := repository.NewUserRepo()
	return &MinIOService{minIORepo, userRepo}
}

/* ====================================== */

// 上传用户头像，并且删除旧头像（如有）
func (ms *MinIOService) UpLoadUserAvatar(userID uint, file *multipart.File) (string, int) {
	// 移除旧头像
	go ms.DeleteUserAvatar(userID)
	// 上传新头像
	avatarFile, avatarFileName, contentType, err := img.ProcessAvatar(userID, file)
	if err != nil {
		log.Fatalln("文件格式转换失败", err)
		return "", errmsg.ERROR_UPLOAD_USERAVT
	}
	avatarURL, code := ms.minIORepo.UpLoadImg(userID, avatarFileName, avatarFile, contentType)
	return avatarURL, code
}

// 删除用户旧头像
func (ms *MinIOService) DeleteUserAvatar(userID uint) {
	user, code := ms.userRepo.GetByID(userID)
	if code != errmsg.SUCCESS {
		log.Fatalln("删除用户旧头像出现的其他错误:", errmsg.GetErrMsg(code))
	}
	oldAvatarURL := user.AvatarURL
	if oldAvatarURL != nil {
		objectName := ms.parseObjectNameFromURL(oldAvatarURL)
		if objectName != "" {
			ms.minIORepo.DeleteImg(objectName)
		} else {
			log.Fatalln("解析用户头像URL失败！")
		}
	}
}

// 解析用户头像文件的objectName
func (ms *MinIOService) parseObjectNameFromURL(inputURL *string) string {
	parsedURL, err := url.Parse(*inputURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return ""
	}
	objectName := path.Base(parsedURL.Path)

	return objectName
}
