package community

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"teacupapi/controllers/base"
	"teacupapi/models"
	"teacupapi/service/communityService"
	"teacupapi/service/userService"
	userUtils2 "teacupapi/service/userUtils"
	"teacupapi/utils"
	"teacupapi/utils/errUtils"
)

//创建社区
func NewCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.Community
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.CreatedUserId = userId
	//创建社区
	code, msg := req.CreateCommunity()
	if code != utils.StatusOK {
		base.WebResp(ctx, 200, code, nil, msg.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "成功")
	return
}

//更新社区
func UpdateCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.Community
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	req.CreatedUserId = userId
	//创建社区
	code, msg := req.UpdateCommunity()
	if code != utils.StatusOK {
		base.WebResp(ctx, 200, code, nil, msg.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "成功")
	return
}

//查找社区
func GetCommunity(ctx *gin.Context) {
	var req communityService.GetCommunity
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	var community communityService.Community
	community.ID = req.ID
	community2, err := community.GetCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, community2, "")
	return
}

// 热门社区
func GetHotsCommunity(ctx *gin.Context) {
	var req communityService.GetCommunityList
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	var community communityService.Community
	comList, count, err := community.GetHotsCommunities(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//我关注的社区
func GetFansCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.GetCommunityList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	var community communityService.Community
	comList, count, err := community.GetFansCommunities(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//嘉宾的社区
func GetVipsCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.GetCommunityList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	var community communityService.Community
	comList, count, err := community.GetVipsCommunities(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//我的社区
func GetMyCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.GetCommunityList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	var community communityService.Community
	comList, count, err := community.GetMyCommunities(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//我加入的社区
func GetJoinCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.GetCommunityList
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	var community communityService.Community
	comList, count, err := community.GetJoinsCommunities(userId, &req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//我加入的社区
func GetCommunityFans(ctx *gin.Context) {
	var req userService.GetTeacupVipList
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	comList, count, err := userService.GetCommunityOfFans(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//社区进入的成员
func GetCommunityJoins(ctx *gin.Context) {
	var req userService.GetTeacupVipList
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	comList, count, err := userService.GetCommunityOfJoins(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

//社区进入的嘉宾
func GetCommunityVips(ctx *gin.Context) {
	var req userService.GetTeacupVipList
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//创建社区
	comList, count, err := userService.GetCommunityOfVips(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//返回社区列表
	var resp models.PagingRep
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Count = count
	resp.List = comList

	base.WebResp(ctx, 200, utils.StatusOK, resp, "")
	return
}

// 关注社区
func Apply2Community(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	var req communityService.FansToCommunity
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//申请关注社区
	req.UserID = userId
	_, err = req.FansToCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

//取消关注
func CloseCommunity(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	var req communityService.FansToCommunity
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//申请关注社区
	req.UserID = userId
	_, err = req.CloseFansToCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 邀请
func Invite2Community(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	var req communityService.InviteToCommunity
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	if !(req.CommunityId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "社区id不能为空")
		return
	}
	if !(req.InviteUserId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "邀请人不能为空")
		return
	}

	var member communityService.MemberToCommunity
	member.UserID = userId
	member.CommunityId = req.CommunityId
	if m, err := member.GetMemberToCommunity(); err == nil && m != nil {
		base.WebResp(ctx, 200, utils.ErrDataExistFailed, nil, "已经是社区的成员")
		return
	}
	//邀请嘉宾
	req.UserID = userId
	req.InviteStatus = 1
	_, err = req.InviteToCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 获得邀请社区
func GetInvite2Community(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.InviteToCommunity
	code, com, err := req.GetInviteToCommunity(userId)
	if err != nil {
		base.WebResp(ctx, 200, code, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, com, "")
	return
}

// 审批邀请
func Update2Community(ctx *gin.Context) {
	var req communityService.InviteToCommunity
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
	preStatus := req.InviteStatus
	if req.InviteStatus == 2 { // 申请加入成员
		_, err = req.GetInviteCommunity()
		if err != nil {
			base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
			return
		}
		var member communityService.MemberToCommunity
		member.CommunityId = req.CommunityId
		member.UserID = req.UserID
		_, err := member.MemberToCommunity()
		if err != nil {
			base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
			return
		}
	}
	req.InviteStatus = preStatus
	_, err = req.UpdateToCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 申请加入
func Join2Community(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	var req communityService.JoinToCommunity
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}

	if !(req.CommunityId > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "申请id不能为空")
		return
	}

	var com communityService.Community
	getCom, err := com.GetCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataNotExistError, nil, "社区不存在")
		return
	}
	// 不可以加入
	if getCom.AnyJoin == 1 {
		base.WebResp(ctx, 200, utils.ErrAllowError, nil, "社区不允许加入")
		return
	}

	req.UserID = userId

	code, err := req.GetJoinForId()
	if err == nil && code == utils.StatusOK {
		base.WebResp(ctx, 200, utils.ErrDataExistFailed, nil, "已经申请")
		return
	}

	var member communityService.MemberToCommunity
	member.UserID = userId
	member.CommunityId = req.CommunityId
	if m, err := member.GetMemberToCommunity(); err == nil && m != nil {
		base.WebResp(ctx, 200, utils.ErrDataExistFailed, nil, "已经是社区的成员")
		return
	}
	//邀请嘉宾
	req.UserID = userId
	req.ApplyStatus = 1
	_, err = req.JoinToCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}

// 获得申请
func GetJoin2Community(ctx *gin.Context) {
	userId, err := userUtils2.GetUserIdByGinContext(ctx)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	var req communityService.JoinToCommunity
	code, com, err := req.GetJoinToCommunity(userId)
	if err != nil && err == gorm.ErrRecordNotFound {
		base.WebResp(ctx, 200, utils.StatusOK, nil, "")
		return
	}
	if err != nil {
		base.WebResp(ctx, 200, code, nil, err.Error())
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, com, "")
	return
}

// 审批申请
func UpdateJoinCommunity(ctx *gin.Context) {
	var req communityService.JoinToCommunity
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, errUtils.ErrorMsg(err))
		return
	}
	//邀请嘉宾
	if !(req.ID > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "申请id必填")
		return
	}
	if !(req.ApplyStatus > 0) {
		base.WebResp(ctx, 200, utils.ErrInvalidParams, nil, "申请状态必填")
		return
	}
	preStatus := req.ApplyStatus
	if req.ApplyStatus == 2 { // 申请加入成员
		_, err := req.GetJoin()
		if err != nil {
			base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
			return
		}
		var member communityService.MemberToCommunity
		member.CommunityId = req.CommunityId
		member.UserID = req.UserID
		_, err = member.MemberToCommunity()
		if err != nil {
			base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
			return
		}
	}
	req.ApplyStatus = preStatus
	_, err = req.UpdateToCommunity()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrDataError, nil, errUtils.ErrorMsg(err))
		return
	}
	base.WebResp(ctx, 200, utils.StatusOK, nil, "")
	return
}
