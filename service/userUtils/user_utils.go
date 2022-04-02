package userUtils

import (
	"github.com/gin-gonic/gin"
	"teacupapi/config"
	"teacupapi/db/mdata"
	"teacupapi/db/redisdb"
	glog "teacupapi/logs"
	"teacupapi/service/userService"
	"teacupapi/utils"
	"teacupapi/utils/constants"
	"teacupapi/utils/errUtils"
	"time"
)

func GetUserIdByGinContext(c *gin.Context) (int64, error) {
	tokenData, err := GetTokenDataByGinContext(c)
	if err != nil {
		return 0, err
	}
	return tokenData.ID, nil

}

func GetTokenDataByGinContext(c *gin.Context) (*userService.TokenData, error) {
	tokenData, err := userService.ParseToken(c.GetHeader(mdata.HeaderToken))
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}

/*******白名单校验******/
func IsWhiteIp(ip string, isLimit int) bool {
	//admin表中是否限制
	if isLimit == 2 {
		return true
	}
	//判断是否白名单
	bRet, err := redisdb.SIsMember(true, constants.REDIS_IP_WHITE2LIST, ip)
	if err != nil {
		glog.Errorf("get white ip list failure,error:%v", err)
		return false
	}
	return bRet
}

/*****密码校验*****/
func CheckPwd(pwd string) error {
	//参数校验-1:解密 2:校验
	data, err := utils.AesCBCPk7DecryptBase64(pwd, []byte(config.GetAesSert()), []byte(config.GetAesSert())) //解密
	if err != nil {
		glog.Errorf("密码解析错误,pwd:%s,error:%s", pwd, err)
		return errUtils.NewBizError("密码解析错误")
	}
	if err = utils.CheckFunc(constants.PwdLenMin, constants.PwdLenMax, string(data), constants.RegexpCheckPwd); err != nil { //校验密码
		return errUtils.NewBizError("密码不符合规则")
	}
	return nil
}

func RefreshTokenTime(tokenStr string) {
	userNo, err := redisdb.GetKey(true, tokenStr)
	if err != nil && err != redisdb.RedisNil {
		glog.Errorf("刷新token,err:%+v", err)
		return
	}
	err = redisdb.SetExpireKey(tokenStr, time.Hour*10*24)
	if err != nil {
		glog.Errorf("刷新token,err:%+v", err)
	}
	err = redisdb.SetExpireKey(userNo+constants.REDIS_SINGLE_LOGIN, time.Hour*10*24)
	if err != nil {
		glog.Errorf("刷新token,err:%+v", err)
	}
}
