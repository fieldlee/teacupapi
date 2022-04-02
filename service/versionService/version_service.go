package versionService

import (
	"github.com/jinzhu/gorm"
	sqlDB "teacupapi/db/sqldb"
	"teacupapi/utils"
)

type VersionReq struct {
	PhoneType int    `gorm:"column:phone_type;" json:"phone_type"` //版本
	Version   string `gorm:"column:version;" json:"version"`       //版本
}

type Version struct {
	ID        int    `gorm:"column:id;" json:"id"`
	Version   string `gorm:"column:version;" json:"version"`       //版本
	PhoneType int    `gorm:"column:phone_type;" json:"phone_type"` //版本
}

func (v *VersionReq) CreateVersion() (int, error) {
	err := sqlDB.GetTDb().Table("version").Create(v).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

func (v *VersionReq) CheckVersion() (int, string, error) {
	var version Version
	tx := sqlDB.GetRDb().Table("version").
		Select("id,`version`,`phone_type`").Where("phone_type = ?", version.PhoneType).Order(`id desc`)
	tx = tx.First(&version)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return utils.ErrDataNotExistError, "", err
	}
	if v.Version == version.Version {
		return utils.StatusOK, "", nil
	}
	return utils.ErrWarning, version.Version, nil
}
