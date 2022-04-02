package userService

import (
	"errors"
	"teacupapi/config"
	"teacupapi/db/mdata"
	"teacupapi/db/redisdb"
	glog "teacupapi/logs"
	"teacupapi/utils"
	"teacupapi/utils/constants"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

type TokenData struct {
	ID        int64  `json:"id"`
	PhoneNo   string `json:"user_phone"`
	Timestamp int64  `json:"timestamp"`
}

// swagger:model login
type Login struct {
	// example: 18616017954
	PhoneNo string `json:"user_phone" binding:"required"`
	// cn 中国  us 美国
	Nation string `json:"nation" binding:"required"`
	// 1 iOS  2 Android
	PhoneType int `json:"user_phone_type" binding:"required"`
	// example: 888888
	SMSCode string `json:"smscode" binding:"required"`
}

func (login *Login) Login() (int, string, error) {
	var user UserInfo
	err := user.GetUserByPhone(login.PhoneNo)
	if err != nil {
		return utils.ErrDataNotExistError, "", err
	}
	if user.Id <= 0 {
		return utils.ErrDataNotExistError, "", errors.New("user not exist")
	}
	token := TokenData{
		ID:        user.Id,
		PhoneNo:   user.UserPhone,
		Timestamp: time.Now().Unix(),
	}
	tokenByte, err := mdata.Cjson.Marshal(&token)
	if err != nil {
		glog.Errorf("jsonMarshal user=%s token err: %v", token.ID, err)
		return utils.ErrInternal, "", err
	}
	keyBytes := []byte(config.GetTokenkey())
	tokenStr, err := utils.AesCBCPk7EncryptBase64(tokenByte, keyBytes, keyBytes)
	if err != nil {
		glog.Errorf("encrypt user=%s token err: %v", token.PhoneNo, err)
		return utils.ErrInternal, "", err
	}
	// 设置 token 的过期时间为10天
	err = redisdb.SetExpireKV(user.UserPhone+constants.REDIS_SINGLE_LOGIN, tokenStr, time.Hour*10*24)
	if err != nil {
		glog.Errorf("setUserToken user=%s err: %v", user.UserPhone, err)
		return utils.ErrInternal, "", err
	}
	// 设置 token 的过期时间为10天
	err = redisdb.SetExpireKV(tokenStr, user.UserPhone, time.Hour*10*24)
	if err != nil {
		glog.Errorf("setUserToken user=%s err: %v", user.UserPhone, err)
		return utils.ErrInternal, "", err
	}
	return utils.StatusOK, tokenStr, nil
}

// 解析 token
func ParseToken(tokenStr string) (*TokenData, error) {
	keyBytes := []byte(config.GetTokenkey())
	tokenBytes, err := utils.AesCBCPk7DecryptBase64(tokenStr, keyBytes, keyBytes)
	if err != nil {
		tmpStr := "decrypt token err: " + err.Error()
		return nil, errors.New(tmpStr)
	}

	var tokenData TokenData
	err = mdata.Cjson.Unmarshal(tokenBytes, &tokenData)
	if err != nil {
		tmpStr := "jsonUnmarshal err: " + err.Error()
		return nil, errors.New(tmpStr)
	}
	return &tokenData, nil
}
