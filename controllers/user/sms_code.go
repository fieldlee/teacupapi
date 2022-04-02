package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"teacupapi/config"
	"teacupapi/controllers/base"
	"teacupapi/db/redisdb"
	glog "teacupapi/logs"
	"teacupapi/service/userService"
	"teacupapi/utils"
	"teacupapi/utils/constants"
	"teacupapi/utils/errUtils"
	"time"
)

func GetSmsCode(ctx *gin.Context) {
	var req userService.ReqSMSCode
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.PhoneNo = strings.TrimSpace(req.PhoneNo)
	ip := ctx.ClientIP()
	code, err := req.CheckPhone(ip)
	if err != nil {
		base.WebResp(ctx, 200, code, nil, err.Error())
		return
	}

	smsCode := "888888"
	if config.GetBaseConf().Env != "dev" {
		code := utils.RealNumRand(6)
		errorCode, msg := req.SendValidateCode(code, req.Nation)
		if errorCode != utils.StatusOK {
			base.WebResp(ctx, 200, errorCode, nil, msg)
			return
		}
		smsCode = code
	}
	smsCodeKey := fmt.Sprintf(constants.SmsCode, strings.TrimSpace(req.PhoneNo))

	smsCodeCheck, err := redisdb.GetKey(true, smsCodeKey)
	if err == nil && smsCodeCheck != "" {
		base.WebResp(ctx, 200, utils.ErrCodeTimeError, nil, "验证码五分钟后再试")
		return
	}

	err = redisdb.SetExpireKV(smsCodeKey, smsCode, time.Minute*5)
	if err != nil {
		glog.Errorf("redis,set key:%s,value:%s,err:%+v", smsCodeKey, smsCode, err)
		base.WebResp(ctx, 200, utils.ErrInternal, nil, utils.InternalErr)
		return
	}
	glog.Infof("手机号：%s,当前验证码：%s", req.PhoneNo, smsCode)
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
}

func CheckSmsCode(ctx *gin.Context) {
	var req userService.ReqCheckSMSCode
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.PhoneNo = strings.TrimSpace(req.PhoneNo)
	smsCodeKey := fmt.Sprintf(constants.SmsCode, req.PhoneNo)
	smsCode, err := redisdb.GetKey(true, smsCodeKey)
	if err != nil {
		glog.Errorf("redis,GetKey key:%s,err:%+v", smsCodeKey, err)
		base.WebResp(ctx, 200, utils.ErrCodeError, nil, "验证码错误")
		return
	}

	if smsCode != req.SMSCode {
		glog.Infof("手机号：%s,请求验证码：%s,缓存验证码：%s", req.PhoneNo, req.SMSCode, smsCode)
		base.WebResp(ctx, 200, utils.ErrCodeError, nil, "验证码错误")
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
}
