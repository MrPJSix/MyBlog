package main

import (
	"myblog.backend/repository"
	"myblog.backend/routes"
)

func main() {
	repository.InitDB()
	routes.InitRouter()
}
