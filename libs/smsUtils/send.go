package smsUtils

import (
	"net/url"
	"strconv"
	"strings"
	"teacupapi/db/mdata"
	"teacupapi/libs/httpclient"
	glog "teacupapi/logs"
	"teacupapi/utils"
	"teacupapi/utils/constants"
)

//发送短信验证码
func MockSendValidateCode(phone, validateCode, nation string) (int, string) {
	if nation == constants.REGION_CHINA {
		host := strings.TrimSpace("")
		if len(host) == 0 {
			glog.Errorf("发送短信获取，短信host配置错误，当前为空")
			return utils.ErrInternal, utils.InternalErr
		}
		uri := "/sms/v1/sendsmscode"

		values := url.Values{}
		values.Add("mobile", phone)
		values.Add("code", validateCode)
		values.Add("opt", "live")
		requestUrl := host + uri
		respByte, err := httpclient.POSTHTTP(requestUrl, []byte(values.Encode()), map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		})
		if err != nil {
			glog.Errorf("手机号：%s,request:%s,发送验证码错误,err:%+v", phone, requestUrl, err)
			return utils.ErrInternal, utils.InternalErr
		}

		respStruct := struct {
			Status  string `json:"status"`
			Data    string `json:"data"`
			Message string `json:"message"`
		}{}
		err = mdata.Cjson.Unmarshal(respByte, &respStruct)
		if err != nil {
			glog.Errorf("手机号：%s,request:%s,response:%s,反序列化错误,err:%+v",
				phone, requestUrl, string(respByte), err)
			return utils.ErrInternal, utils.InternalErr
		}
		status, err := strconv.Atoi(respStruct.Status)
		if err != nil {
			glog.Errorf("手机号：%s,request:%s,response:%s,反序列化错误,err:%+v",
				phone, requestUrl, string(respByte), err)
			return utils.ErrInternal, utils.InternalErr
		}

		if status != utils.StatusOK {
			return status, respStruct.Message
		}
		return utils.StatusOK, utils.Success
	}

	if nation == constants.REGION_USA {
		return utils.StatusOK, utils.Success
	}
	return utils.StatusOK, utils.Success
}

/**
//submail短信验证码
func SendValidateCode3(phone, validateCode string) (int, string) {
	config := make(map[string]string)
	config["appid"] = "63028"
	config["appkey"] = "0202997d964b01192f6796cbbeb5bc4a"
	config["signType"] = "normal"
	//创建 短信 Send 接口
	submail := sms.CreateXsend(config)
	//设置联系人 手机号码
	submail.SetTo(phone)
	//设置短信模板id
	submail.SetProject("zF0o53")
	//添加模板中的设置的动态变量。如模板为：【xxx】您的验证码是:@var(code),请在@var(time)分钟内输入。
	submail.AddVar("code", validateCode)
	submail.AddVar("time", "5")
	//执行 Xsend 方法发送短信
	xsend := submail.Xsend()
	//序列化
	respStruct := struct {
		Status string `json:"status"`
		SendId string `json:"send_id"`
	}{}
	err := mdata.Cjson.Unmarshal([]byte(xsend), &respStruct)
	if err != nil {
		glog.Errorf("手机号：%s,response:%s,反序列化错误,err:%+v",
			phone, xsend, err)
		return utils.ErrInternal, utils.InternalErr
	}
	if respStruct.Status != "success" {
		return utils.ErrInternal, respStruct.Status
	}
	return utils.StatusOK, respStruct.Status
}
 */