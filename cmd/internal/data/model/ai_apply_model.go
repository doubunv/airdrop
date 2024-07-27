package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type AiApplyModel struct {
	db *gorm.DB
}

func NewAiAiApplyModel(db *gorm.DB) *AiApplyModel {
	return &AiApplyModel{db: db}
}

func (m *AiApplyModel) Insert(res *schema.AiApply) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *AiApplyModel) UpdateSchema(data *schema.AiApply) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *AiApplyModel) FindById(id int64) (res schema.AiApply, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *AiApplyModel) GetList(model *schema.AiApply, startTime, endTime int64, page, pageSize int) (list []*schema.AiApply, total int64, err error) {
	q := m.db.Model(&schema.AiApply{})
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
