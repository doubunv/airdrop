package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type AmountLogModel struct {
	db *gorm.DB
}

func NewAmountLogModel(db *gorm.DB) *AmountLogModel {
	return &AmountLogModel{db: db}
}

func (m *AmountLogModel) Insert(res *schema.AmountLog) error {
	return m.db.Create(res).Error
}

func (m *AmountLogModel) UpdateSchema(data *schema.AmountLog) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *AmountLogModel) FindById(id int64) (res schema.AmountLog, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *AmountLogModel) GetList(model *schema.AmountLog, startTime, endTime string, page, pageSize int) (list []*schema.AmountLog, total int64, err error) {
	q := m.db.Model(&schema.AmountLog{})
	if model.UAddress != "" {
		q = q.Where("u_address = ?", model.UAddress)
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
