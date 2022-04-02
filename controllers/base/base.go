package base

import (
	"teacupapi/config"
	"teacupapi/db/mdata"
	"teacupapi/db/redisdb"
	glog "teacupapi/logs"
	"teacupapi/models"
	"teacupapi/service/userService"
	userUtils2 "teacupapi/service/userUtils"
	"teacupapi/utils"
	"teacupapi/utils/constants"

	"github.com/gin-gonic/gin"
)

// swagger:model basicResponse
type BasicResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}

// WebResp 发往前台的公共接口
// statusCode : http的状态码
// errCode: 错误码
// Msg: 信息
func WebResp(c *gin.Context, statusCode, errCode int, data interface{}, Msg string) {
	if data == nil {
		data = struct{}{}
	}
	respMap := BasicResponse{
		StatusCode: errCode,
		Data:       data,
		Message:    Msg,
	}
	c.JSON(statusCode, respMap)
	return
}

//WebStock返回数据
func WsResp(statusCode, errCode int, data interface{}, Msg string) ([]byte, error) {
	if data == nil {
		data = struct{}{}
	}
	respMap := BasicResponse{
		StatusCode: errCode,
		Data:       data,
		Message:    Msg,
	}
	return mdata.Cjson.Marshal(respMap)
}

// AuthRequired 针对请求进行鉴权与签名校验
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 判断鉴权开关是否打开
		if !config.GetAuthFlag() {
			c.Next()
			return
		}
		preURL := c.Request.URL.Path
		// 如果是过滤的路由检测接口直接跳过, 直接跳到 controllers
		_, exist := models.RouterFilterMap[preURL]
		if exist {
			c.Next()
			return
		}
		// 鉴权
		errCode, msg := AuthFunc(c)
		if errCode != utils.StatusOK {
			glog.Errorf("auth url=%s fail: errCode=%d msg=%s", preURL, errCode, msg)
			WebResp(c, 200, errCode, nil, msg)
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

// AuthFunc 鉴权函数
func AuthFunc(c *gin.Context) (int, string) {
	token := c.GetHeader(models.HeaderToken)
	if token == "" {
		return utils.ErrAccess, "鉴权失败,请从新登录"
	}

	tokenData, err := userService.ParseToken(token)
	exist, err := redisdb.KeyExist(true, tokenData.PhoneNo+constants.REDIS_SINGLE_LOGIN)
	if err != nil {
		glog.Errorf("redis KeyExist,key:%s,err:%+v", token, err)
		return utils.ErrAccess, "请重试"
	}

	if !exist {
		return utils.ErrToken, "token超时,请重新登录"
	}

	// 判断 token 在 redis 中是否存在
	exist, err = redisdb.KeyExist(true, token)
	if err != nil {
		glog.Errorf("redis KeyExist,key:%s,err:%+v", token, err)
		return utils.ErrAccess, "请重试"
	}
	if !exist {
		return utils.ErrAccess, "请重新登录"
	}

	//刷新token信息
	userUtils2.RefreshTokenTime(token)

	return utils.StatusOK, utils.Success
}

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	WebResp(c, 200, utils.StatusOK, nil, utils.Success)
	return
}
