package model

import (
	"air-drop/cmd/internal/data/dto/userDto"
	"air-drop/cmd/internal/data/schema"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (m *UserModel) GetUserByUAddress(uAddress string) (res schema.User, err error) {
	err = m.db.Find(&res, "u_address = ?", uAddress).Error
	return
}

func (m *UserModel) GetUserById(id int64) (res schema.User, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *UserModel) GetUnique(inviteCode string) schema.User {
	var res schema.User
	err := m.db.First(&res, "invite_code = ?", inviteCode).Error
	if err != nil {
		logx.Errorf("NewUserModel GetUnique err:%v", err)
	}
	return res
}

func (m *UserModel) GetUserByParentAddress(parentAddress string, isOwner int) (res []schema.User, err error) {
	err = m.db.Find(&res, "parent_address = ? and is_owner = ?", parentAddress, isOwner).Error
	return
}

func (m *UserModel) GetIsOwnerByParentAddress(parentAddress string) (res []int64, err error) {
	err = m.db.Model(&schema.User{}).Select("is_owner").Find(&res, "parent_address = ?", parentAddress).Error
	return
}

func (m *UserModel) Insert(res *schema.User) error {
	return m.db.Create(res).Error
}

func (m *UserModel) UpdateSchema(data *schema.User) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *UserModel) CountInvite(userAddress string) int64 {
	var count int64
	m.db.Model(&schema.User{}).Where("parent_address = ?", userAddress).Count(&count)
	return count
}

func (m *UserModel) GetUserList(user *schema.User, startTime, endTime string, page, pageSize int) (list []*schema.User, total int64, err error) {
	q := m.db.Model(&schema.User{})
	if user.UAddress != "" {
		q = q.Where("u_address = ?", user.UAddress)
	}
	if user.ParentAddress != "" {
		q = q.Where("parent_address = ?", user.ParentAddress)
	}
	if user.InviteCode != "" {
		q = q.Where("invite_code = ?", user.InviteCode)
	}
	if startTime != "" {
		q = q.Where("create_at >= ?", startTime)
	}
	if endTime != "" {
		q = q.Where("create_at <= ?", endTime)
	}
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *UserModel) GetInviteList(userAddress string, page, pageSize int) (res []schema.User, total int64, err error) {
	q := m.db.Model(&schema.User{}).Where("parent_address = ?", userAddress)
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	return
}

func (m *UserModel) GetInviteCountList(userAddress string, page, pageSize int) (res []userDto.InviteData, total int64, err error) {
	q := m.db.Model(&schema.User{}).Where("parent_address = ?", userAddress).
		Joins("left join mm_user_team on mm_user.u_address = mm_user_team.user_address").Select(
		"mm_user.id,mm_user.u_address,mm_user.team_level,mm_user.create_at,mm_user_team.team_spend_amount,mm_user_team.spend_amount, mm_user_team.is_active, mm_user_team.total_team_spend_amount")
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error

	return
}

func (m *UserModel) GetAllUser() []schema.User {
	var res []schema.User

	err := m.db.Model(&schema.User{}).Order("id asc").Find(&res).Error
	if err != nil {
		logx.Errorf("GetAllUser err:%v", err)
	}

	return res
}

// not contain self
func (m *UserModel) GetInviteParentList(id []int64) (res []schema.User, err error) {
	err = m.db.Model(&schema.User{}).Where("id in (?)", id).Find(&res).Error
	return
}

type CountGroupByTime struct {
	DayTime  string `gorm:"column:day_time"`
	TodayAdd int64  `gorm:"column:today_add"`
}

func (m *UserModel) GetTeamGroupByTime(userId int64, startTime, endTime string) (res []CountGroupByTime, err error) {
	// err = m.db.Model(&User{}).Select("date(CONVERT_TZ(create_at, '+00:00','+08:00')) as day_time, count(1) as today_add").
	err = m.db.Model(&schema.User{}).Select("date(create_at) as day_time, count(1) as today_add").
		Where("id in (SELECT id FROM mm_user WHERE FIND_IN_SET(?,path) > 0) and create_at >= ? and create_at <= ?", userId, startTime, endTime).
		Group("day_time").Order("day_time asc").Scan(&res).Error
	return
}

func (m *UserModel) CountTeamMembersByTime(userId int64, startTime, endTime string) (res int64, err error) {
	err = m.db.Model(&schema.User{}).Where("id in (SELECT id FROM mm_user WHERE FIND_IN_SET(?,path) > 0) and create_at >= ? and create_at <= ?", userId, startTime, endTime).Count(&res).Error
	return
}

func (m *UserModel) GetInviteGroupByTime(address string, startTime, endTime string) (res []CountGroupByTime, err error) {
	err = m.db.Model(&schema.User{}).Select("date(create_at) as day_time, count(1) as today_add").
		Where("id in (SELECT id FROM mm_user WHERE parent_address = ?) and create_at >= ? and create_at <= ?", address, startTime, endTime).
		Group("day_time").Order("day_time asc").Scan(&res).Error
	return
}

func (m *UserModel) CountInviteByTime(address string, startTime, endTime string) (res int64, err error) {
	err = m.db.Model(&schema.User{}).Where("parent_address = ? and create_at >= ? and create_at <= ?", address, startTime, endTime).Count(&res).Error
	return
}

func (m *UserModel) GetByAddressList(address []string) []schema.User {
	var res []schema.User
	_ = m.db.Find(&res, "u_address in ?", address).Error
	return res
}

func (m *UserModel) CountTeamNum() []userDto.UserTeamCount {
	var res []userDto.UserTeamCount
	sql := "select team_level,count(*) as count" +
		" from mm_user group by team_level"

	if err := m.db.Raw(sql).Scan(&res).Error; err != nil {
		logx.Errorf("CountTeamNum err:%v", err.Error())
	}

	return res
}

func (m *UserModel) Count() int64 {
	var res int64

	if err := m.db.Model(&schema.User{}).Count(&res).Error; err != nil {
		logx.Errorf("Count err:%v", err.Error())
	}
	return res
}
