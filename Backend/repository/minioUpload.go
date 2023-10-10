package repository

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type IMinIORepo interface {
	UpLoadImg(userID uint, objectName string, file io.Reader, contentType string) int
}

type MinIORepo struct{}

func NewMinIORepo() *MinIORepo {
	return &MinIORepo{}
}

var bucketName = "myblog-img"

/* ====================================== */

func (mr *MinIORepo) UpLoadImg(userID uint, objectName string, file io.Reader, contentType string) int {
	info, err := minioClient.PutObject(
		context.Background(),
		// Bucket name
		bucketName,
		// Object name
		objectName,
		// File to be uploaded
		file,
		// File size
		-1,
		// Options (like content-type)
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		log.Println("上传头像失败！", err)
		return errmsg.ERROR_UPLOAD_USERAVT
	}

	err = db.Model(&model.User{}).Where("id = ?", userID).Update("avatar_url", info.Location).Error
	if err != nil {
		log.Println("头像URL保存失败", err)
		return errmsg.ERROR_UPLOAD_USERAVT
	}

	return errmsg.SUCCESS
}
