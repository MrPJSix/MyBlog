package repository

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"myblog.backend/model"
	"myblog.backend/utils/errmsg"
)

/* ====================================== */

type IMinIORepo interface {
	UpLoadImg(userID uint, objectName string, file *bytes.Buffer, contentType string) int
	DeleteImg(objectName string)
}

type MinIORepo struct{}

func NewMinIORepo() *MinIORepo {
	return &MinIORepo{}
}

var bucketName = "myblog-img"

/* ====================================== */

func (mr *MinIORepo) UpLoadImg(userID uint, objectName string, file *bytes.Buffer, contentType string) (string, int) {
	info, err := minioClient.PutObject(
		context.Background(),
		// Bucket name
		bucketName,
		// Object name
		objectName,
		// File to be uploaded
		bytes.NewReader((*file).Bytes()),
		// File size
		-1,
		// Options (like content-type)
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		log.Println("上传头像失败！", err)
		return "", errmsg.ERROR_UPLOAD_USERAVT
	}

	err = db.Model(&model.User{}).Where("id = ?", userID).Update("avatar_url", info.Location).Error
	if err != nil {
		log.Println("头像URL保存失败", err)
		return "", errmsg.ERROR_UPLOAD_USERAVT
	}

	return info.Location, errmsg.SUCCESS
}

func (mr *MinIORepo) DeleteImg(objectName string) {

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err := minioClient.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		log.Fatalln("移除旧用户头像失败！", err)

	}

	log.Println("移除旧用户头像成功！")
}
