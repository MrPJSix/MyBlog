package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	// 服务配置
	AppMode  string
	HttpPort string
	JwtKey   string

	// MySQL配置
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	// MinIO配置
	EndPoint  string
	AccessKey string
	SecretKey string

	// Redis配置
	RdsHost     string
	RdsPort     string
	RdsPassword string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("读取配置文件出错！", err)
	}
	InitServerConfig(file)
	InitDatabaseConfig(file)
	InitMinIOConfig(file)
	InitRedisConfit(file)
}

func InitServerConfig(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("Hello, my blog!")
}

func InitDatabaseConfig(file *ini.File) {
	database := "database-" + "production" //AppMode
	DbHost = file.Section(database).Key("DbHost").String()
	DbPort = file.Section(database).Key("DbPort").String()
	DbUser = file.Section(database).Key("DbUser").String()
	DbPassword = file.Section(database).Key("DbPassword").String()
	DbName = file.Section(database).Key("DbName").String()
}

func InitMinIOConfig(file *ini.File) {
	EndPoint = file.Section("minio").Key("EndPoint").String()
	AccessKey = file.Section("minio").Key("AccessKey").String()
	SecretKey = file.Section("minio").Key("SecretKey").String()
}

func InitRedisConfit(file *ini.File) {
	RdsHost = file.Section("redis").Key("RdsHost").String()
	RdsPort = file.Section("redis").Key("RdsPort").String()
	RdsPassword = file.Section("redis").Key("RdsPassword").String()
}
