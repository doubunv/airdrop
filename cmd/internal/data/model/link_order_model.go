package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type LinkOrderModel struct {
	db *gorm.DB
}

func NewLinkOrderModel(db *gorm.DB) *LinkOrderModel {
	return &LinkOrderModel{db: db}
}

func (m *LinkOrderModel) Insert(res *schema.LinkOrder) error {
	return m.db.Create(res).Error
}

func (m *LinkOrderModel) UpdateSchema(data *schema.LinkOrder) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *LinkOrderModel) FindById(id int64) (res schema.LinkOrder, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *LinkOrderModel) GetList(model *schema.LinkOrder, startTime, endTime int64, page, pageSize int) (list []*schema.LinkOrder, total int64, err error) {
	q := m.db.Model(&schema.LinkOrder{})
	if model.UserId != 0 {
		q = q.Where("user_id = ?", model.UserId)
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
