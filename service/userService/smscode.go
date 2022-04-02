package userService

import (
	"errors"
	"strings"
	sqlDB "teacupapi/db/sqldb"
	"teacupapi/libs/smsUtils"
	glog "teacupapi/logs"
	"teacupapi/utils"
	"teacupapi/utils/constants"
)

const (
	LoginType    = 1
	ReigsterType = 2
)

type ReqSMSCode struct {
	PhoneNo   string `gorm:"column:user_phone;"  json:"user_phone" binding:"required"`
	Nation    string `gorm:"column:nation;"  json:"nation" binding:"required"`
	PhoneType int    `gorm:"column:user_phone_type;"  json:"user_phone_type" binding:"required"` // 手机类型 1 iOS 2 android
	Type      int    `gorm:"column:type;"  json:"type" binding:"required"`                       // type=1 login type=2 register
}

func (a *ReqSMSCode) CheckPhone(ip string) (int, error) {
	phone := strings.TrimSpace(a.PhoneNo)
	//手机号验证-针对大陆
	if a.Nation == constants.REGION_CHINA {
		if err := utils.ChinaPhoneCheck(constants.ChinaPhoneLen, constants.ChinaPhoneLen, phone); err != nil {
			glog.Errorf("name is not rule err: %s", utils.WrongPhoneNo)
			return utils.ErrPhoneError, errors.New(utils.WrongPhoneNo)
		}
	}
	//手机号验证-针对美国
	if a.Nation == constants.REGION_USA {
		if err := utils.USAPhoneCheck(constants.USAPhoneLen, constants.USAPhoneLen, phone); err != nil {
			glog.Errorf("name is not rule err: %s", utils.WrongPhoneNo)
			return utils.ErrPhoneError, errors.New(utils.WrongPhoneNo)
		}
	}
	// 注册
	if a.Type == ReigsterType {
		var user UserInfo
		if err := user.GetUserByPhone(phone); err == nil && user.Id > 0 {
			return utils.ErrUserExistError, errors.New("user exist")
		}
	}

	return utils.StatusOK, nil
}

func (a *ReqSMSCode) SendValidateCode(code string, nation string) (int, string) {
	err := sqlDB.GetTDb().Table("sms_log").Create(a).Error
	if err != nil {
		return utils.ErrDataError, err.Error()
	}
	return smsUtils.MockSendValidateCode(a.PhoneNo, code, nation)
}

type ReqCheckSMSCode struct {
	PhoneNo string `json:"user_phone" binding:"required"`
	Nation  string `json:"nation" binding:"required"`
	SMSCode string `json:"smscode" binding:"required"`
}
