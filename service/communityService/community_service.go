package communityService

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	sqlDB "teacupapi/db/sqldb"
	glog "teacupapi/logs"
	"teacupapi/utils"
)

type GetCommunity struct {
	ID int64 `json:"id" binding:"required"`
}
type GetCommunityList struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required,lte=500"`
}

type Community struct {
	ID               int64  `gorm:"column:id;" json:"id"`
	CreatedUserId    int64  `gorm:"column:created_user_id;" json:"created_user_id" `
	CommunityName    string `gorm:"column:community_name;" json:"community_name" binding:"required"`
	CommunityComment string `gorm:"column:community_comment;" json:"community_comment" binding:"required"`
	CommunityTags    string `gorm:"column:community_tags;" json:"community_tags"`
	CommunityImage   string `gorm:"column:community_image;" json:"community_image"`
	OnlyInvite       int    `gorm:"column:is_only_invite;" json:"is_only_invite"` //'是否只有邀请人加入 1不是 2是',
	AnyJoin          int    `gorm:"column:is_any_join;" json:"is_any_join"`       //是否任何人都可以加入 1不可以 2可以',
}

func (community *Community) CreateCommunity() (int, error) {
	if err := sqlDB.GetTDb().Table("community").Create(&community).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

func (community *Community) UpdateCommunity() (int, error) {
	if err := sqlDB.GetTDb().Table("community").
		Where("id = ?", community.ID).
		Updates(map[string]interface{}{
			"community_name":    community.CommunityName,
			"community_comment": community.CommunityComment,
			"community_tags":    community.CommunityTags,
			"community_image":   community.CommunityImage,
			"is_only_invite":    community.OnlyInvite,
			"is_any_join":       community.AnyJoin,
		}).Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

func (community *Community) GetCommunity() (*Community, error) {
	var com Community
	if err := sqlDB.GetRDb().Table("community").
		Select(`id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Where("id = ?", community.ID).
		First(&com).Error; err != nil {
		return nil, err
	}
	return &com, nil
}

func (community *Community) GetHotsCommunities(req *GetCommunityList) ([]*Community, int, error) {
	tx := sqlDB.GetRDb().Table("community").
		Select(`id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Where("id = ?", community.ID)
	var count int
	dataList := make([]*Community, 0)
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

//我关注的社区
func (community *Community) GetFansCommunities(id int64, req *GetCommunityList) ([]*Community, int, error) {
	tx := sqlDB.GetRDb().Table("community").
		Select(`community.id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Joins("left join fans_community on community.id = fans_community.community_id").
		Where("fans_community.user_id = ?", id)
	var count int
	dataList := make([]*Community, 0)
	err := tx.Count(&count).Error
	if err != nil {
		glog.Errorf("GetFansCommunities  err: %v", err)
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

//我是嘉宾的社区
func (community *Community) GetVipsCommunities(id int64, req *GetCommunityList) ([]*Community, int, error) {
	tx := sqlDB.GetRDb().Table("community").
		Select(`community.id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Joins("left join invite_to_community on community.id = invite_to_community.community_id").
		Where("invite_to_community.invite_user_id = ? and invite_to_community.invite_status = 2", id)
	var count int
	dataList := make([]*Community, 0)
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

//我的社区
func (community *Community) GetMyCommunities(id int64, req *GetCommunityList) ([]*Community, int, error) {
	tx := sqlDB.GetRDb().Table("community").
		Select(`community.id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Where("created_user_id = ?", id)
	var count int
	dataList := make([]*Community, 0)
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

//我是加入的社区
func (community *Community) GetJoinsCommunities(id int64, req *GetCommunityList) ([]*Community, int, error) {
	tx := sqlDB.GetRDb().Table("community").
		Select(`community.id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Joins("left join apply_to_community on community.id = apply_to_community.community_id").
		Where("apply_to_community.user_id = ? and apply_to_community.apply_status = 2", id)
	var count int
	dataList := make([]*Community, 0)
	err := tx.Count(&count).Error
	if err != nil {
		glog.Errorf("GetJoinsCommunities err :%+v", err)
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

type FansToCommunity struct {
	UserID      int64 `gorm:"column:user_id;" json:"user_id"`
	CommunityId int64 `gorm:"column:community_id;" json:"community_id" binding:"required"`
}

func (fans *FansToCommunity) FansToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("fans_community").Create(fans).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (fans *FansToCommunity) CloseFansToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("fans_community").
		Delete(&FansToCommunity{}, "user_id = ? and community_id = ?", fans.UserID, fans.CommunityId).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type MemberToCommunity struct {
	UserID      int64 `gorm:"column:user_id;" json:"user_id"`
	CommunityId int64 `gorm:"column:community_id;" json:"community_id" binding:"required"`
}

func (member *MemberToCommunity) MemberToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("member_community").Create(member).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (member *MemberToCommunity) CloseMemberToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("member_community").
		Delete(&MemberToCommunity{}, "user_id = ? and community_id = ?", member.UserID, member.CommunityId).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (member *MemberToCommunity) GetMemberToCommunity() (*MemberToCommunity, error) {
	var m MemberToCommunity
	err := sqlDB.GetTDb().Table("member_community").
		Select("id,user_id,community_id").
		Where("user_id = ? and community_id = ?", member.UserID, member.CommunityId).
		First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

type InviteToCommunity struct {
	ID           int64 `gorm:"column:id;" json:"id"`
	UserID       int64 `gorm:"column:user_id;" json:"user_id"`
	CommunityId  int64 `gorm:"column:community_id;" json:"community_id"`
	InviteUserId int64 `gorm:"column:invite_user_id;" json:"invite_user_id"`
	InviteStatus int   `gorm:"column:invite_status;" json:"invite_status"`
}

type CommunityForInvite struct {
	ID               int64  `gorm:"column:id;" json:"id"`
	InviteId         int64  `gorm:"column:invite_id;" json:"invite_id"`
	CreatedUserId    int64  `gorm:"column:created_user_id;" json:"created_user_id" `
	CommunityName    string `gorm:"column:community_name;" json:"community_name" binding:"required"`
	CommunityComment string `gorm:"column:community_comment;" json:"community_comment" binding:"required"`
	CommunityTags    string `gorm:"column:community_tags;" json:"community_tags"`
	CommunityImage   string `gorm:"column:community_image;" json:"community_image"`
	OnlyInvite       int    `gorm:"column:is_only_invite;" json:"is_only_invite"` //'是否只有邀请人加入 1不是 2是',
	AnyJoin          int    `gorm:"column:is_any_join;" json:"is_any_join"`       //是否任何人都可以加入 1不可以 2可以',
}

func (invite *InviteToCommunity) InviteToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("invite_to_community").Create(invite).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (invite *InviteToCommunity) GetInviteToCommunity(userId int64) (int, *CommunityForInvite, error) {
	var com CommunityForInvite
	if err := sqlDB.GetRDb().Table("community").
		Select(`invite_to_community.id as invite_id,community.id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Joins("left join invite_to_community on community.id = invite_to_community.community_id").
		Where("invite_to_community.invite_user_id = ?", userId).First(&com).Error; err != nil {
		return utils.ErrDataError, nil, err
	}
	return utils.StatusOK, &com, nil
}
func (invite *InviteToCommunity) UpdateToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("invite_to_community").
		Where("id = ?", invite.ID).
		Update(map[string]interface{}{
			"invite_status": invite.InviteStatus,
		}).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (invite *InviteToCommunity) GetInviteCommunity() (int, error) {
	if err := sqlDB.GetTDb().Table("invite_to_community").
		Select(`id,user_id,community_id,invite_user_id,invite_status`).
		Where("id = ?", invite.ID).First(invite).
		Error; err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}

type JoinToCommunity struct {
	ID          int64 `gorm:"column:id;" json:"id"`
	UserID      int64 `gorm:"column:user_id;" json:"user_id"`
	CommunityId int64 `gorm:"column:community_id;" json:"community_id"`
	ApplyStatus int   `gorm:"column:apply_status;" json:"apply_status"`
}
type CommunityForJoin struct {
	ID               int64  `gorm:"column:id;" json:"id"`
	JoinId           int64  `gorm:"column:join_id;" json:"join_id"`
	CreatedUserId    int64  `gorm:"column:created_user_id;" json:"created_user_id" `
	CommunityName    string `gorm:"column:community_name;" json:"community_name" binding:"required"`
	CommunityComment string `gorm:"column:community_comment;" json:"community_comment" binding:"required"`
	CommunityTags    string `gorm:"column:community_tags;" json:"community_tags"`
	CommunityImage   string `gorm:"column:community_image;" json:"community_image"`
	OnlyInvite       int    `gorm:"column:is_only_invite;" json:"is_only_invite"` //'是否只有邀请人加入 1不是 2是',
	AnyJoin          int    `gorm:"column:is_any_join;" json:"is_any_join"`       //是否任何人都可以加入 1不可以 2可以',
}

func (join *JoinToCommunity) JoinToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("apply_to_community").Create(join).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (join *JoinToCommunity) GetJoin() (int, error) {
	if err := sqlDB.GetRDb().Table("apply_to_community").
		Select(`id,user_id,community_id,apply_status`).
		Where("id = ?", join.ID).First(join).Error; err != nil {
		glog.Errorf("GetJoin err : %+v", err)
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (join *JoinToCommunity) GetJoinForId() (int, error) {
	if err := sqlDB.GetRDb().Table("apply_to_community").
		Select(`id,user_id,community_id,apply_status`).
		Where("user_id = ? and community_id = ?", join.UserID, join.CommunityId).First(join).Error; err != nil {
		glog.Errorf("GetJoinForId err : %+v", err)
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
func (join *JoinToCommunity) GetJoinToCommunity(userId int64) (int, *CommunityForJoin, error) {
	var com CommunityForJoin
	if err := sqlDB.GetRDb().Table("community").
		Select(`apply_to_community.id as join_id,community.id,created_user_id,community_name,community_comment,community_tags,community_image,is_only_invite,is_any_join`).
		Joins("left join apply_to_community on community.id = apply_to_community.community_id").
		Where("community.created_user_id = ?", userId).First(&com).Error; err != nil {
		glog.Errorf("GetJoinToCommunity err : %+v", err)
		return utils.ErrDataError, nil, err
	}

	return utils.StatusOK, &com, nil
}
func (join *JoinToCommunity) UpdateToCommunity() (int, error) {
	err := sqlDB.GetTDb().Table("apply_to_community").
		Where("id = ?", join.ID).
		Update(map[string]interface{}{
			"apply_status": join.ApplyStatus,
		}).Error
	if err != nil {
		return utils.ErrDataError, err
	}
	return utils.StatusOK, nil
}
