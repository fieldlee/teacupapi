package errUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	glog "teacupapi/logs"
	"teacupapi/utils"
)

type BizError struct {
	Message string `json:"message"`
	Err     error
}

// 实现接口
func (e *BizError) Error() string {
	return e.Message
}

//可以预判的错误使用
//err错误
//message:发生错误用于前台展示的错误消息提示
func Wrapper(err error, message string) *BizError {
	return &BizError{
		Message: message,
		Err:     err,
	}
}

//一个新的bizError
func NewBizError(message string) *BizError {
	return &BizError{
		Message: message,
		Err:     errors.New(message),
	}
}

func ValidationErrorToText(e *validator.FieldError) string {
	switch (*e).Tag() {
	case "required":
		return fmt.Sprintf("%s不能为空", (*e).Field())
	case "max":
		return fmt.Sprintf("%s不能大于%s", (*e).Field(), (*e).Param())
	case "min":
		return fmt.Sprintf("%s不能小于%s", (*e).Field(), (*e).Param())
	case "email":
		return fmt.Sprintf("email格式错误")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", (*e).Param(), (*e).Field())
	case "gte":
		return fmt.Sprintf("%s必须大于等于%s", (*e).Field(), (*e).Param())
	case "lte":
		return fmt.Sprintf("%s必须小于等于%s", (*e).Field(), (*e).Param())
	case "uppercase":
		return fmt.Sprintf("%s必须为大写字母%s", (*e).Field(), (*e).Param())
	}
	return fmt.Sprintf("%s校验错误", (*e).Field())
}

func ErrorMsg(err error) string {

	// 如果业务代码执行出错
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			//gin validator 参数校验
			glog.Infof("请求参数校验不通过: %s", err)

			msgSlice := make([]string, 0)
			for _, err := range errs {
				msgSlice = append(msgSlice, ValidationErrorToText(&err))
			}
			return strings.Join(msgSlice, ",")
		}

		if errs, ok := err.(*json.UnmarshalTypeError); ok {
			glog.Errorf("json返序列化异常：%+v", errs)
			return errs.Field + "反序列化错误"
		}
		if err, ok := err.(*BizError); ok {
			return err.Message
		}

		//其他错误统一后台定位处理
		glog.Errorf("服务器内部错误 ,error: %+v", err)
		return utils.InternalErr
	}

	return "success11"

}
