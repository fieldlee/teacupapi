package user

import (
	"fmt"
	"teacupapi/controllers/base"
	"teacupapi/db/mdata"
	"teacupapi/db/redisdb"
	"teacupapi/libs"
	glog "teacupapi/logs"
	"teacupapi/models"
	"teacupapi/service/userService"
	userUtils2 "teacupapi/service/userUtils"
	"teacupapi/utils"
	"teacupapi/utils/constants"
	"teacupapi/utils/errUtils"

	"github.com/gin-gonic/gin"
)

// swagger:route POST /user/register user register
//
// 注册
//
// _
//
// Parameters:
// + name: body
//   in: body
//   type: reqRegister
// Responses:
//   200: basicResponse
func Register(ctx *gin.Context) {
	var req userService.ReqRegister
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

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

	err = redisdb.DelKey(smsCodeKey)
	if err != nil {
		glog.Errorf("redis,DelKey key:%s,err:%+v", smsCodeKey, err)
	}

	//注册
	code, tokenStr, msg := req.Register()
	if msg != nil {
		base.WebResp(ctx, 200, code, nil, msg.Error())
		return
	}
	respMap := map[string]string{
		"token": tokenStr,
	}
	base.WebResp(ctx, 200, utils.StatusOK, respMap, "注册完成")
	return
}

func FillInformation(ctx *gin.Context) {
	var req userService.UserInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	if req.UserImage != "" {
		req.UserImage = libs.TrimStaticDomain(req.UserImage)
	}
	code, err := req.UpdateUserInfo()
	if err != nil {
		base.WebResp(ctx, 200, code, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
}

// swagger:route POST /user/login user login
//
// 登录
//
// _
//
// Parameters:
// + name: body
//   in: body
//   type: login
// Responses:
//   200: basicResponse
func Login(ctx *gin.Context) {
	var req userService.Login
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	//登录
	code, tokenStr, msg := req.Login()
	if code != utils.StatusOK {
		base.WebResp(ctx, 200, code, nil, msg.Error())
		return
	}
	respMap := map[string]string{
		"token": tokenStr,
	}
	base.WebResp(ctx, 200, utils.StatusOK, respMap, "登录成功")
	return
}

//用户退出
func LoginOut(ctx *gin.Context) {
	tokenData, _ := userUtils2.GetTokenDataByGinContext(ctx)
	if tokenData != nil {
		userNo := tokenData.PhoneNo
		_ = redisdb.DelKey(userNo + constants.REDIS_SINGLE_LOGIN)
		_ = redisdb.DelKey(ctx.GetHeader(mdata.HeaderToken))
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "退出成功")
	return
}

//身份检查
func GetUserInfo(ctx *gin.Context) {
	userId, _ := userUtils2.GetUserIdByGinContext(ctx)
	var req userService.ReqUser
	req.Id = userId
	updateInfo, err := req.GetUser()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInternal, nil, err.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, updateInfo, utils.Success)
	return
}

//申请vip
func Apply2Vip(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req userService.ReqAppToVip
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.UserID = userId
	req.ApplyStatus = 1
	_, err = req.AppToVip()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
	return
}

//审批
func AgreeRefuse(ctx *gin.Context) {
	var req userService.ReqAppToVip
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	_, err = req.AgreeRefuse()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
	return
}

//关注
func Fans(ctx *gin.Context) {
	var req userService.ReqFans
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	_, err = req.Fans()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
	return
}

//取消关注
func CloseFans(ctx *gin.Context) {
	var req userService.ReqFans
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	_, err = req.CloseFans()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, utils.Success)
	return
}

//我关注的人
func FansList(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req userService.GetUserList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	datalist, count, err := userService.Fans(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = datalist
	base.WebResp(ctx, 200, utils.StatusOK, resp, utils.Success)
	return
}

//关注我的人
func FollowsList(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req userService.GetUserList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	datalist, count, err := userService.Follows(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = datalist
	base.WebResp(ctx, 200, utils.StatusOK, resp, utils.Success)
	return
}
