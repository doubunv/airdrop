package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type AiComputingModel struct {
	db *gorm.DB
}

func NewAiComputingModel(db *gorm.DB) *AiComputingModel {
	return &AiComputingModel{db: db}
}

func (m *AiComputingModel) Insert(res *schema.AiComputing) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *AiComputingModel) UpdateSchema(data *schema.AiComputing) error {
	data.CreatedAt = time.Now().Unix()
	return m.db.Where("id = ?", data.ID).Updates(data).Error
}

func (m *AiComputingModel) FindById(id int64) (res schema.AiComputing, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *AiComputingModel) GetList(model *schema.AiComputing, startTime, endTime int64, page, pageSize int) (list []*schema.AiComputing, total int64, err error) {
	q := m.db.Model(&schema.AiComputing{})
	if startTime != 0 {
		q = q.Where("created_at >= ?", startTime)
	}
	if endTime != 0 {
		q = q.Where("created_at <= ?", endTime)
	}
	q = q.Where("deleted_at is null")
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *AiComputingModel) FindByIds(ids []int64) (res []schema.AiComputing, err error) {
	if len(ids) == 0 {
		return
	}
	err = m.db.Find(&res, "id in ?", ids).Error
	return
}

func (m *AiComputingModel) FindDetailById(id int64) (res schema.AiComputing, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}
