package ginValidator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

//去除字符串首尾空格
var TrimHeaderTailBlank = func(fl validator.FieldLevel) bool {
	fl.Field().SetString(strings.TrimSpace(fl.Field().String()))
	return true
}
