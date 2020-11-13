package tool

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func GetValidateErr(err validator.ValidationErrors) string  {
	fmt.Println(err)
	return "参数错误"
}
