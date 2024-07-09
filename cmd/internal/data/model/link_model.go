package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type LinkModel struct {
	db *gorm.DB
}

func NewLinkModel(db *gorm.DB) *LinkModel {
	return &LinkModel{db: db}
}

func (m *LinkModel) Insert(res *schema.ArLink) error {
	return m.db.Create(res).Error
}

func (m *LinkModel) UpdateSchema(data *schema.ArLink) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *LinkModel) FindById(id int64) (res schema.ArLink, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *LinkModel) GetList(model *schema.ArLink, startTime, endTime string, page, pageSize int) (list []*schema.ArLink, total int64, err error) {
	q := m.db.Model(&schema.ArLink{})
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
