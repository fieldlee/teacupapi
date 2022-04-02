package teacupService

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	sqlDB "teacupapi/db/sqldb"
	glog "teacupapi/logs"
	"teacupapi/utils"
)

type GetTeacup struct {
	ID int64 `json:"id" binding:"required"`
}
type GetTeacupList struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required,lte=500"`
}
type Teacup struct {
	ID            int64  `gorm:"column:id;" json:"id"`
	CreatedUserId int64  `gorm:"column:user_id;" json:"user_id" `
	CommunityId   int64  `gorm:"column:community_id;" json:"community_id" binding:"required"`
	TeacupName    string `gorm:"column:teacup_name;" json:"teacup_name" binding:"required"`
	TeacupComment string `gorm:"column:teacup_comment;" json:"teacup_comment"`
	TeacupImage   string `gorm:"column:teacup_image;" json:"teacup_image"`
	TeacupTags    string `gorm:"column:teacup_tags;" json:"teacup_tags"`
}

//创建
func (teacup *Teacup) CreateTeacup() (int, error) {
	if err := sqlDB.GetTDb().Table("teacup").Create(teacup).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

//更新
func (teacup *Teacup) UpdateTeacup() (int, error) {
	if err := sqlDB.GetTDb().Table("teacup").
		Where("id = ?", teacup.ID).Update(map[string]interface{}{
		"teacup_name":    teacup.TeacupName,
		"teacup_comment": teacup.TeacupComment,
		"teacup_image":   teacup.TeacupImage,
		"teacup_tags":    teacup.TeacupTags,
	}).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

//查找
func (teacup *Teacup) GetTeacup() (*Teacup, error) {
	var tea Teacup
	if err := sqlDB.GetRDb().Table("teacup").
		Select(`id,user_id,community_id,teacup_name,teacup_comment,teacup_image,teacup_tags`).
		Where("id = ?", teacup.ID).
		First(&tea).Error; err != nil {
		return nil, err
	}
	return &tea, nil
}

//查找
func (teacup *Teacup) GetTeacupForUserId(userId int64, req *GetTeacupList) ([]*Teacup, int, error) {
	tx := sqlDB.GetRDb().Table("teacup").
		Select(`id,user_id,community_id,teacup_name,teacup_comment,teacup_image,teacup_tags`).
		Where("user_id = ?", userId)
	var count int
	dataList := make([]*Teacup, 0)
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

//查找
func (teacup *Teacup) GetTeacupForVip(userId int64, req *GetTeacupList) ([]*Teacup, int, error) {
	tx := sqlDB.GetRDb().Table("teacup").
		Select(`teacup.id,teacup.user_id,community_id,teacup_name,teacup_comment,teacup_image,teacup_tags`).
		Joins("left join teacup_vip on teacup.id = teacup_vip.teacup_id").
		Where("teacup_vip.user_id = ?", userId)
	var count int
	dataList := make([]*Teacup, 0)
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

type TeacupVip struct {
	ID       int64 `gorm:"column:id;" json:"id"`
	UserId   int64 `gorm:"column:user_id;" json:"user_id" `
	TeacupId int64 `gorm:"column:teacup_id;" json:"teacup_id"`
}

//创建
func (teacupVip *TeacupVip) CreateTeacupVip() (int, error) {
	if err := sqlDB.GetTDb().Table("teacup_vip").Create(teacupVip).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

//创建
func (teacupVip *TeacupVip) GetTeacupVip() (*TeacupVip, error) {
	var vip TeacupVip
	if err := sqlDB.GetRDb().Table("teacup_vip").Select("id,user_id,teacup_id").
		Where("user_id = ? and teacup_id = ?", teacupVip.UserId, teacupVip.TeacupId).First(&vip).Error; err != nil {
		return nil, err
	}
	return &vip, nil
}

type InviteToTeacup struct {
	ID           int64 `gorm:"column:id;" json:"id"`
	UserID       int64 `gorm:"column:user_id;" json:"user_id"`
	TeacupId     int64 `gorm:"column:teacup_id;" json:"teacup_id"`
	InviteUserId int64 `gorm:"column:invited_user_id;" json:"invited_user_id"`
	InviteStatus int   `gorm:"column:invited_status;" json:"invited_status"`
}
type TeacupForInvite struct {
	ID            int64  `gorm:"column:id;" json:"id"`
	InvitedId     int64  `gorm:"column:invite_id;" json:"invite_id"`
	CreatedUserId int64  `gorm:"column:user_id;" json:"user_id" `
	CommunityId   int64  `gorm:"column:community_id;" json:"community_id" binding:"required"`
	TeacupName    string `gorm:"column:teacup_name;" json:"teacup_name" binding:"required"`
	TeacupComment string `gorm:"column:teacup_comment;" json:"teacup_comment"`
	TeacupImage   string `gorm:"column:teacup_image;" json:"teacup_image"`
	TeacupTags    string `gorm:"column:teacup_tags;" json:"teacup_tags"`
}

func (invite *InviteToTeacup) InviteToTeacup() (int, error) {
	err := sqlDB.GetTDb().Table("invite_to_teacup").Create(invite).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (invite *InviteToTeacup) GetInviteTeacup(userId int64) (*TeacupForInvite, error) {
	var teacup TeacupForInvite
	if err := sqlDB.GetRDb().Table("invite_to_teacup").
		Select(`invite_to_teacup.id as invite_id,teacup.id,teacup.user_id,teacup.community_id,teacup.teacup_name,teacup.teacup_comment,teacup.teacup_image,teacup.teacup_tags`).
		Joins("left join teacup on invite_to_teacup.teacup_id = teacup.id").
		Where("invite_to_teacup.invited_user_id = ? ", userId).First(&teacup).Error; err != nil {
		glog.Errorf("GetInviteTeacup err : %+v", err)
		return nil, err
	}
	return &teacup, nil
}
func (invite *InviteToTeacup) UpdateToTeacup() (int, error) {
	err := sqlDB.GetTDb().Table("invite_to_teacup").
		Where("id = ?", invite.ID).
		Update(map[string]interface{}{
			"invited_status": invite.InviteStatus,
		}).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type JoinToTeacup struct {
	ID         int64  `gorm:"column:id;" json:"id"`
	UserID     int64  `gorm:"column:user_id;" json:"user_id"`
	TeacupId   int64  `gorm:"column:teacup_id;" json:"teacup_id" binding:"required"`
	JoinStatus int    `gorm:"column:join_status;" json:"join_status"`
	JoinAt     string `gorm:"column:join_at;" json:"join_at"`
	LeaveAt    string `gorm:"column:leave_at;" json:"leave_at"`
}

func (join *JoinToTeacup) JoinToTeacup() (int, error) {
	join.JoinStatus = 1
	err := sqlDB.GetTDb().Table("join_teacup").Create(join).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

func (join *JoinToTeacup) LeaveToTeacup() (int, error) {
	err := sqlDB.GetTDb().Table("join_teacup").
		Where("user_id = ? and teacup_id = ?", join.UserID, join.TeacupId).
		Order("id desc").Limit(1).
		Update(map[string]interface{}{
			"join_status": 2,
			"leave_at":    join.LeaveAt,
		}).Error

	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type TeacupHistory struct {
	ID            int64  `gorm:"column:id;" json:"id"`
	TeacupId      int64  `gorm:"column:teacup_id;" json:"teacup_id" `
	CreatedUserId int64  `gorm:"column:user_id;" json:"user_id" `
	CommunityId   int64  `gorm:"column:community_id;" json:"community_id"`
	TeacupStartAt string `gorm:"column:teacup_start_at;" json:"teacup_start_at"`
	TeacupEndAt   string `gorm:"column:teacup_end_at;" json:"teacup_end_at"`
	SpeechRoom    string `gorm:"column:speech_room;" json:"speech_room"`
	SpeechChannel string `gorm:"column:speech_channel;" json:"speech_channel"`
	SpeechStatus  int    `gorm:"column:speech_status;" json:"speech_status"` //开播 1未开播 2 开播中 3 已开播
}

//创建
func (history *TeacupHistory) CreateTeacupHistory() (int, error) {
	if err := sqlDB.GetTDb().Table("teacup_history").Create(history).Error; err != nil {
		glog.Errorf("CreateTeacupHistory error:%+v", err)
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

//创建
func (history *TeacupHistory) GetTeacupHistory() (int, *TeacupHistory, error) {
	var preHistory TeacupHistory
	if err := sqlDB.GetTDb().Table("teacup_history").
		Select("id,teacup_id,user_id,community_id,teacup_start_at,teacup_end_at,speech_room,speech_channel,speech_status").
		Where("id = ?", history.ID).First(&preHistory).Error; err != nil {
		glog.Errorf("CreateTeacupHistory error:%+v", err)
		return utils.ErrDataError, nil, err
	}
	return utils.StatusOK, &preHistory, nil
}

//更新
func (history *TeacupHistory) UpdateTeacupHistory() (int, error) {

	if history.SpeechStatus == 2 {
		if err := sqlDB.GetTDb().Table("teacup_history").Where("id = ?", history.ID).
			Update(map[string]interface{}{
				"speech_room":    history.SpeechRoom,
				"speech_channel": history.SpeechChannel,
				"speech_status":  history.SpeechStatus,
			}).Error; err != nil {
			return utils.ErrDataError, err
		}
	}
	if history.SpeechStatus == 3 {
		if err := sqlDB.GetTDb().Table("teacup_history").Where("id = ?", history.ID).
			Update(map[string]interface{}{
				"speech_status": history.SpeechStatus,
			}).Error; err != nil {
			return utils.ErrDataError, err
		}
	}

	return utils.StatusOK, nil
}

type TeacupChat struct {
	ID           int64  `gorm:"column:id;" json:"id"`
	TeacupId     int64  `gorm:"column:teacup_id;" json:"teacup_id" binding:"required"`
	UserId       int64  `gorm:"column:user_id;" json:"user_id" `
	CommunityId  int64  `gorm:"column:community_id;" json:"community_id"`
	ChatCategory int    `gorm:"column:chat_category;" json:"chat_category"` //聊天类型 1 心情 2举手 3 发言语音
	ChatComment  string `gorm:"column:chat_comment;" json:"chat_comment"`
}

//创建
func (chat *TeacupChat) CreateChat() (int, error) {
	if err := sqlDB.GetTDb().Table("teacup_chat").Create(chat).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type TeacupMark struct {
	ID          int64  `gorm:"column:id;" json:"id"`
	TeacupId    int64  `gorm:"column:teacup_id;" json:"teacup_id" binding:"required"`
	UserId      int64  `gorm:"column:user_id;" json:"user_id" `
	MarkValue   int    `gorm:"column:mark_value;" json:"mark_value"`
	MarkComment string `gorm:"column:mark_comment;" json:"mark_comment"`
}

//创建
func (mark *TeacupMark) CreateMark() (int, error) {
	if err := sqlDB.GetTDb().Table("teacup_mark").Create(mark).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
