package valid

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"myblog.backend/utils/errmsg"
	"reflect"
	"regexp"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{8,25}$`)
	passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+{}\[\]:;<>,.?~\\/-]{8,25}$`)
)

func ValidateCredentials(username, password string) int {
	isUsernameValid := usernameRegex.MatchString(username)
	if !isUsernameValid {
		return errmsg.ERROR_BAD_USERNAME
	}
	isPasswordValid := passwordRegex.MatchString(password)
	if !isPasswordValid {
		return errmsg.ERROR_BAD_PASSWORD
	}
	return errmsg.SUCCESS
}

func ValidateRegister(data interface{}) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
