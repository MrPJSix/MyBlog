package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
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
var rdb *redis.Client

func InitDB() {
	initMySQL()
	initMinIO()
	initRedis()
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
		log.Println("连接MySQL数据库失败！", err)
	}

	err = db.AutoMigrate(
		&model.Article{},
		&model.Category{},
		&model.User{},
		&model.Comment{},
		&model.Notification{},
		&model.ArtileLike{},
		&model.CommentLike{},
	)

	if err != nil {
		log.Println("MySQL数据库迁移出错！", err)
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

// 初始化Redis数据库
func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RdsHost, config.RdsPort),
		Password: config.RdsPassword,
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接Redis数据库出错:", err)
		return
	}
}
