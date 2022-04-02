package teacup

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"teacupapi/controllers/base"
	"teacupapi/models"
	"teacupapi/service/teacupService"
	"teacupapi/service/userService"
	userUtils2 "teacupapi/service/userUtils"
	"teacupapi/utils"
	"teacupapi/utils/errUtils"
)

//创建茶室
func NewTeacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.Teacup
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.CreatedUserId = userId
	//创建茶室
	code, msg := req.CreateTeacup()
	if code != utils.StatusOK {
		base.WebResp(ctx, 200, code, nil, msg.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "成功")
	return
}

//更新茶室
func UpdateTeacup(ctx *gin.Context) {
	var req teacupService.Teacup
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//更新茶室
	code, msg := req.UpdateTeacup()
	if code != utils.StatusOK {
		base.WebResp(ctx, 200, code, nil, msg.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "成功")
	return
}

//我的茶室
func MyTeacups(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.GetTeacupList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//更新茶室
	var teacup teacupService.Teacup
	list, count, err := teacup.GetTeacupForUserId(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, err.Error())
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = list

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//查找茶室
func GetTeacup(ctx *gin.Context) {
	var req teacupService.GetTeacup
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//更新茶室
	var teacup teacupService.Teacup
	teacup.ID = req.ID
	t, err := teacup.GetTeacup()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, err.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, t, "成功")
	return
}

// 邀请
func Invite2Teacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	var req teacupService.InviteToTeacup
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	// 判断是否已经是嘉宾
	{
		var vip teacupService.TeacupVip
		vip.UserId = req.InviteUserId
		vip.TeacupId = req.TeacupId
		v, err := vip.GetTeacupVip()
		if err == nil && v != nil {
			base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "已是嘉宾")
			return
		}
	}

	//邀请嘉宾
	req.UserID = userId
	req.InviteStatus = 1
	_, err = req.InviteToTeacup()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

func GetInviteTeacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.InviteToTeacup
	invite, err := req.GetInviteTeacup(userId)
	if err != nil && err == gorm.ErrRecordNotFound {
		base.WebResp(ctx, 200, utils.StatusOK, nil, "")
		return
	}
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataNotExistError, nil, err.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, invite, "")
	return
}

// 审批邀请
func Update2Teacup(ctx *gin.Context) {
	var req teacupService.InviteToTeacup
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//邀请嘉宾
	if !(req.ID > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "邀请id必填")
		return
	}
	if !(req.InviteStatus > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "邀请状态必填")
		return
	}
	if req.InviteStatus == 2 { //嘉宾通过
		var vip teacupService.TeacupVip
		vip.UserId = req.InviteUserId
		vip.TeacupId = req.TeacupId
		_, err := vip.CreateTeacupVip()
		if err != nil {
			base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
			return
		}
	}
	_, err = req.UpdateToTeacup()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// JoinTeacup
func Join2Teacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	var req teacupService.JoinToTeacup
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.UserID = userId
	req.JoinAt = utils.CurTimeToStr()
	req.LeaveAt = utils.CurTimeToStr()
	_, err = req.JoinToTeacup()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// leave Teacup
func LeaveTeacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.JoinToTeacup
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	if !(req.TeacupId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "茶室ID必填")
		return
	}
	req.UserID = userId
	req.LeaveAt = utils.CurTimeToStr()
	_, err = req.LeaveToTeacup()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 创建茶室记录
func CreateTeacupHistory(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.TeacupHistory
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	if !(req.CommunityId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "社区ID必填")
		return
	}
	if !(req.TeacupId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "茶室ID必填")
		return
	}
	//if req.TeacupStartAt.IsZero() {
	//	base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "茶室开播时间")
	//	return
	//}
	//if req.TeacupEndAt.IsZero() {
	//	base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "茶室结束时间")
	//	return
	//}
	req.CreatedUserId = userId
	req.SpeechStatus = 1
	_, err = req.CreateTeacupHistory()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 创建茶室记录
func StartTeacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.TeacupHistory
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	_, preHistory, err := req.GetTeacupHistory()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	if preHistory.CreatedUserId != userId {
		base.WebResp(ctx, 200, utils.ErrUserError, nil, "只有创建人才可以开播")
		return
	}
	if req.SpeechChannel == "" {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "演讲的channel不能为空")
		return
	}
	req.SpeechStatus = 2  // 茶室开播中
	_, err = req.UpdateTeacupHistory()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 创建茶室记录
func CloseTeacup(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.TeacupHistory
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	_, preHistory, err := req.GetTeacupHistory()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	if preHistory.CreatedUserId != userId {
		base.WebResp(ctx, 200, utils.ErrUserError, nil, "只有创建人才可以开播")
		return
	}
	req.SpeechStatus = 3 //已开播
	_, err = req.UpdateTeacupHistory()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 创建茶室记录
func CreateChat(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.TeacupChat
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	if !(req.ChatCategory > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "类型")
		return
	}
	req.UserId = userId
	_, err = req.CreateChat()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 创建茶室记录
func CreateMark(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.TeacupMark
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	if !(req.TeacupId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "茶室")
		return
	}
	if !(req.MarkValue > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "评分值")
		return
	}
	req.UserId = userId
	_, err = req.CreateMark()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

//作为嘉宾的茶室
func VipTeacups(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req teacupService.GetTeacupList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//作为嘉宾的茶室
	var teacup teacupService.Teacup
	list, count, err := teacup.GetTeacupForVip(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, err.Error())
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = list

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//茶室的嘉宾
func TeacupsVip(ctx *gin.Context) {

	var req userService.GetTeacupVipList
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//作为嘉宾的茶室
	list, count, err := userService.GetTeacupOfVip(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, err.Error())
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = list
	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}
