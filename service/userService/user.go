package userService

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	sqlDB "teacupapi/db/sqldb"
	"teacupapi/utils"
	"teacupapi/utils/errUtils"
)

type GetUserList struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required,lte=500"`
}

type GetTeacupVipList struct {
	ID       int64 `json:"id" binding:"required"`
	Page     int   `json:"page" binding:"required"`
	PageSize int   `json:"pageSize" binding:"required,lte=500"`
}

type UserStatistics struct {
	UserInfo
	Fans    int `gorm:"column:fans_count;" json:"fans_count"`
	Follows int `gorm:"column:follows_count;" json:"follows_count"`
}

func (userStatistics *UserStatistics) Statistics(id int64) (*UserStatistics, error) {
	tx := sqlDB.GetRDb().
		Table("user_info").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Where("id = ? ", id)
	tx = tx.First(userStatistics)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "database search error")
	}
	// fans
	var fansCount int
	tx = sqlDB.GetRDb().
		Table("fans").
		Select("count(1) as fans_count").
		Where("attention_user_id = ? ", id)
	err := tx.Count(&fansCount).Error
	if err != nil {
		return nil, errors.Wrap(err, "sqldb search group list count error")
	}
	// follows
	var followsCount int
	tx = sqlDB.GetRDb().
		Table("fans").
		Select("count(1) as follows_count").
		Where("user_id = ? ", id)
	err = tx.Count(&followsCount).Error
	if err != nil {
		return nil, errors.Wrap(err, "sqldb search group list count error")
	}
	userStatistics.Fans = fansCount
	userStatistics.Follows = followsCount
	return userStatistics, nil
}

type UserInfo struct {
	Id            int64  `gorm:"column:id;" json:"id"`
	UserPhone     string `gorm:"column:user_phone;" json:"user_phone"`
	Gender        int    `gorm:"column:gender;" json:"gender"`
	Birthday      string `gorm:"column:birthday;" json:"birthday"`
	Nation        string `gorm:"column:nation;" json:"nation"`
	UserPhoneType int    `gorm:"column:user_phone_type;" json:"user_phone_type"`
	UUID          string `gorm:"column:uuid;" json:"uuid"`
	UserName      string `gorm:"column:user_name;" json:"user_name"`
	UserAvator    string `gorm:"column:user_avator;" json:"user_avator"`
	UserUnion     string `gorm:"column:user_union;" json:"user_union"`
	UserImage     string `gorm:"column:user_image;" json:"user_image"`
	UserLvl       int    `gorm:"column:user_lvl;" json:"user_lvl"`
	UserBadges    string `gorm:"column:user_badges;" json:"user_badges"`
	UserTopics    string `gorm:"column:user_topics;" json:"user_topics"`
}

func (user *UserInfo) GetUserByPhone(phone string) error {
	tx := sqlDB.GetRDb().
		Table("user_info").
		Select("id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Where("user_phone = ? ", phone)
	var count int
	err := tx.Count(&count).Error
	if err != nil {
		return err
	}
	tx = tx.First(&user)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return errors.Wrap(err, "database search error")
	}
	return nil
}

func (user *UserInfo) GetUserById(id int64) error {
	tx := sqlDB.GetRDb().
		Table("user_info").
		Select("id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Where("id = ? ", id)
	var count int
	err := tx.Count(&count).Error
	if err != nil {
		return err
	}
	tx = tx.First(&user)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return errors.Wrap(err, "database search error")
	}
	return nil
}

func (user *UserInfo) UpdateUserInfo() (int, error) {
	var searchUser UserInfo
	if bRet := sqlDB.GetRDb().Table("user_info").Where("id = ? ", user.Id).Find(&searchUser).RecordNotFound(); bRet {
		return utils.ErrDataNotExistError, errUtils.NewBizError("该用户不存在")
	}
	if err := sqlDB.GetTDb().Table("user_info").
		Where("id = ?", user.Id).
		Updates(user).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type ReqUser struct {
	Id        int64  `gorm:"column:id;" json:"id"`
	UserPhone string `gorm:"column:user_phone;" json:"user_phone"`
}

func (reqUser *ReqUser) GetUser() (*UserInfo, error) {
	var user UserInfo
	if reqUser.Id > 0 {
		err := user.GetUserById(reqUser.Id)
		if err != nil {
			return nil, err
		}
	}
	if reqUser.UserPhone != "" {
		err := user.GetUserByPhone(reqUser.UserPhone)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

type ReqAppToVip struct {
	Id           int64  `gorm:"column:id;" json:"id"`
	UserID       int64  `gorm:"column:user_id;" json:"user_id"`
	ApplyComment string `gorm:"column:apply_comment;" json:"apply_comment" binding:"required"`
	ApplyStatus  int64  `gorm:"column:apply_status;" json:"apply_status"`
}

func (app *ReqAppToVip) AppToVip() (int, error) {
	err := sqlDB.GetTDb().Table("apply_to_vip").Create(app).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (app *ReqAppToVip) AgreeRefuse() (int, error) {
	err := sqlDB.GetTDb().Table("apply_to_vip").
		Where("id = ?", app.Id).
		Update(map[string]interface{}{
			"apply_status": app.ApplyStatus,
		}).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type ReqFans struct {
	UserID          int64 `gorm:"column:user_id;" json:"user_id"`
	AttentionUserId int64 `gorm:"column:attention_user_id;" json:"attention_user_id" binding:"required"`
}

func (fans *ReqFans) Fans() (int, error) {
	err := sqlDB.GetTDb().Table("fans").Create(fans).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

func (fans *ReqFans) CloseFans() (int, error) {
	err := sqlDB.GetTDb().Table("fans").
		Delete(&ReqFans{}, "user_id = ? and attention_user_id = ?", fans.UserID, fans.AttentionUserId).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

//关注我的人
func Fans(id int64, req *GetUserList) ([]*UserInfo, int, error) {
	tx := sqlDB.GetRDb().
		Table("user_info").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Joins("left join fans on user_info.id = fans.user_id").
		Where("fans.attention_user_id = ?", id)
	var count int
	dataList := make([]*UserInfo, 0)
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "sqldb search community list count error")
	}
	if req.Page > 0 && req.PageSize > 0 {
		tx = tx.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	tx = tx.Find(&dataList)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errors.Wrap(err, "sqldb  search video list  error")
	}
	return dataList, count, nil
}

// 我关注的人
func Follows(id int64, req *GetUserList) ([]*UserInfo, int, error) {
	tx := sqlDB.GetRDb().
		Table("user_info").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Joins("left join fans on user_info.id = fans.attention_user_id").
		Where("fans.user_id = ?", id)
	var count int
	dataList := make([]*UserInfo, 0)
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "sqldb search community list count error")
	}
	if req.Page > 0 && req.PageSize > 0 {
		tx = tx.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	tx = tx.Find(&dataList)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errors.Wrap(err, "sqldb  search video list  error")
	}
	return dataList, count, nil
}

//茶室的嘉宾
func GetTeacupOfVip(req *GetTeacupVipList) ([]*UserInfo, int, error) {
	tx := sqlDB.GetRDb().
		Table("teacup_vip").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Joins("left join user_info on teacup_vip.user_id = user_info.id").
		Where("teacup_vip.teacup_id = ? ", req.ID)
	var count int
	dataList := make([]*UserInfo, 0)
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "sqldb search community list count error")
	}
	if req.Page > 0 && req.PageSize > 0 {
		tx = tx.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	tx = tx.Find(&dataList)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errors.Wrap(err, "sqldb  search video list  error")
	}
	return dataList, count, nil
}

//社区的嘉宾
func GetCommunityOfFans(req *GetTeacupVipList) ([]*UserInfo, int, error) {
	tx := sqlDB.GetRDb().
		Table("fans_community").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Joins("left join user_info on fans_community.user_id = user_info.id").
		Where("fans_community.community_id = ? ", req.ID)
	var count int
	dataList := make([]*UserInfo, 0)
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "sqldb search community list count error")
	}
	if req.Page > 0 && req.PageSize > 0 {
		tx = tx.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	tx = tx.Find(&dataList)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errors.Wrap(err, "sqldb  search video list  error")
	}
	return dataList, count, nil
}

//社区的嘉宾
func GetCommunityOfVips(req *GetTeacupVipList) ([]*UserInfo, int, error) {
	tx := sqlDB.GetRDb().
		Table("invite_to_community").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Joins("left join user_info on invite_to_community.invite_user_id = user_info.id").
		Where("invite_to_community.community_id = ?  and invite_to_community.invite_status = 2", req.ID)

	var count int
	dataList := make([]*UserInfo, 0)
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "sqldb search community list count error")
	}
	if req.Page > 0 && req.PageSize > 0 {
		tx = tx.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	tx = tx.Find(&dataList)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errors.Wrap(err, "sqldb  search video list  error")
	}
	return dataList, count, nil
}

//社区的成员
func GetCommunityOfJoins(req *GetTeacupVipList) ([]*UserInfo, int, error) {
	tx := sqlDB.GetRDb().
		Table("apply_to_community").
		Select("user_info.id,`user_phone`,`gender`,`birthday`,`nation`,`user_phone_type`,`uuid`,`user_name`,"+
			"`user_avator`,`user_union`,`user_image`,`user_lvl`,`user_badges`,`user_topics`").
		Joins("left join user_info on apply_to_community.user_id = user_info.id").
		Where("apply_to_community.community_id = ?  and apply_to_community.apply_status = 2", req.ID)

	var count int
	dataList := make([]*UserInfo, 0)
	err := tx.Count(&count).Error
	if err != nil {
		return nil, 0, errors.Wrap(err, "sqldb search community list count error")
	}
	if req.Page > 0 && req.PageSize > 0 {
		tx = tx.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)
	}
	tx = tx.Find(&dataList)
	if err := tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errors.Wrap(err, "sqldb  search video list  error")
	}
	return dataList, count, nil
}
