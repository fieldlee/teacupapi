package system

import (
	"github.com/gin-gonic/gin"
	"teacupapi/controllers/base"
	"teacupapi/db/mdata"
	"teacupapi/service/versionService"
	"teacupapi/utils"
)

//检查版本更新
func CheckUpdate(ctx *gin.Context) {
	var req versionService.VersionReq
	err := mdata.Cjson.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInternal, nil, "参数错误")
		return
	}
	// 判断是否需要更新
	_, version, err := req.CheckVersion()
	if err != nil {
		base.WebResp(ctx, 200, utils.ErrInternal, nil, err.Error())
		return
	}
	req.Version = version
	base.WebResp(ctx, 200, utils.StatusOK, req, utils.Success)
	return
}
