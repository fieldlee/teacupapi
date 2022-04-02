package userService

import (
	"errors"
	"teacupapi/config"
	"teacupapi/db/mdata"
	"teacupapi/db/redisdb"
	sqlDB "teacupapi/db/sqldb"
	glog "teacupapi/logs"
	"teacupapi/utils"
	"teacupapi/utils/constants"
	"time"
)

// swagger:model reqRegister
type ReqRegister struct {
	// example: 18616017954
	PhoneNo string `gorm:"column:user_phone;" json:"user_phone" binding:"required"`
	// cn 中国  us 美国
	Nation string `gorm:"column:nation;" json:"nation" binding:"required"`
	// 1 iOS  2 Android
	PhoneType int `gorm:"column:user_phone_type;" json:"user_phone_type" binding:"required"` // 1 iOS 2 android
	// example: 888888
	SMSCode string `json:"smscode" binding:"required"`
}

type ReqRegisterInfo struct {
	Id            int64  `gorm:"column:id;" json:"id"`
	UserPhone     string `gorm:"column:user_phone;" json:"user_phone" binding:"required"`
	Nation        string `gorm:"column:nation;" json:"nation" binding:"required"`
	UserPhoneType int    `gorm:"column:user_phone_type;" json:"user_phone_type" binding:"required"` // 1 iOS 2 android
}

func (a *ReqRegister) Register() (int, string, error) {

	var user UserInfo
	err := user.GetUserByPhone(a.PhoneNo)
	if err == nil && user.Id > 0 {
		return utils.ErrUserExistError, "", errors.New("用户已存在")
	}
	var saveRegister ReqRegisterInfo
	saveRegister.UserPhone = a.PhoneNo
	saveRegister.Nation = a.Nation
	saveRegister.UserPhoneType = a.PhoneType
	err = sqlDB.GetTDb().Table("user_info").Create(&saveRegister).Error
	if err != nil {
		return utils.ErrDataError, "", err
	}

	token := TokenData{
		ID:        saveRegister.Id,
		PhoneNo:   saveRegister.UserPhone,
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
	err = redisdb.SetExpireKV(saveRegister.UserPhone+constants.REDIS_SINGLE_LOGIN, tokenStr, time.Hour*10*24)
	if err != nil {
		glog.Errorf("setUserToken user=%s err: %v", saveRegister.UserPhone, err)
		return utils.ErrInternal, "", err
	}
	// 设置 token 的过期时间为10天
	err = redisdb.SetExpireKV(tokenStr, saveRegister.UserPhone, time.Hour*10*24)
	if err != nil {
		glog.Errorf("setUserToken user=%s err: %v", saveRegister.UserPhone, err)
		return utils.ErrInternal, "", err
	}
	return utils.StatusOK, tokenStr, nil
}
