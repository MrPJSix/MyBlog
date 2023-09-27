package main

import (
	"myblog.backend/admincli"
	"myblog.backend/repository"
	"myblog.backend/routes"
)

func main() {
	repository.InitDB()
	if admincli.InitAdminCli() {
		return
	}
	routes.InitRouter()
}
