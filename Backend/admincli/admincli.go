package admincli

import (
	"flag"
	"fmt"
	"myblog.backend/model"
	"myblog.backend/service"
	"myblog.backend/utils/errmsg"
)

func InitAdminCli() bool {
	var createAdmin bool
	flag.BoolVar(&createAdmin, "create-admin", false, "Create an admin user")
	flag.Parse()
	if createAdmin {
		createAdminUser()
	}
	return createAdmin
}

func createAdminUser() {
	fmt.Print("Enter admin username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Enter admin password: ")
	var password string
	fmt.Scanln(&password)

	var user model.User
	user.Username = username
	user.Password = password
	user.FullName = username
	user.Role = 1 // 授权管理员
	userService := service.NewUserService()
	code := userService.CreateUser(&user)

	fmt.Println(errmsg.GetErrMsg(code))
}
