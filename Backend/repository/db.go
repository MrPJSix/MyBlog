package repository

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"myblog.backend/config"
	"myblog.backend/model"
)

var db *gorm.DB
var minioClient *minio.Client

func InitDB() {
	initMySQL()
	initMinIO()
}

// 初始化MySQL数据库
func initMySQL() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("连接数据库失败！", err)
	}

	err = db.AutoMigrate(
		&model.Article{},
		&model.Category{},
		&model.User{},
		&model.Comment{},
		&model.Notification{},
	)

	if err != nil {
		log.Println("数据库迁移出错！", err)
	}
}

// 初始化MinIO对象存储
func initMinIO() {
	var err error
	minioClient, err = minio.New(config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Println("连接MinIO客户端出错！", err)
	}
}
