package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Println("读取配置文件出错！", err)
	}
	InitServerConfig(file)
	InitDatabaseConfig(file)
}

func InitServerConfig(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("AppMode").MustString(":9000")
}

func InitDatabaseConfig(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassword = file.Section("database").Key("DbPassword").String()
	DbName = file.Section("database").Key("DbName").String()
}
