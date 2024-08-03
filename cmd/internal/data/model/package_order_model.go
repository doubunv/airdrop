package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type PackageOrderModel struct {
	db *gorm.DB
}

func NewPackageOrderModel(db *gorm.DB) *PackageOrderModel {
	return &PackageOrderModel{db: db}
}

func (m *PackageOrderModel) Insert(res *schema.PackageOrder) error {
	return m.db.Create(res).Error
}

func (m *PackageOrderModel) UpdateSchema(data *schema.PackageOrder) error {
	data.CreatedAt = time.Now().Unix()
	return m.db.Where("id = ?", data.ID).Updates(data).Error
}

func (m *PackageOrderModel) FindById(id int64) (res schema.PackageOrder, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *PackageOrderModel) GetList(user *schema.PackageOrder, startTime, endTime int64, page, pageSize int) (list []*schema.PackageOrder, total int64, err error) {
	q := m.db.Model(&schema.PackageOrder{})
	if user.UserId != 0 {
		q = q.Where("user_id = ?", user.UserId)
	}
	if startTime != 0 {
		q = q.Where("created_at >= ?", startTime)
	}
	if endTime != 0 {
		q = q.Where("created_at <= ?", endTime)
	}
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *PackageOrderModel) SumSendEarnings(userId int64) (resp *schema.PackageOrder, err error) {
	q := m.db.Model(&schema.PackageOrder{}).Where("user_id", userId)
	err = q.Select("sum(send_earnings) as send_earnings").Find(&resp).Error
	return
}
