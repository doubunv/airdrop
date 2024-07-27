package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type AiComputingOrderModel struct {
	db *gorm.DB
}

func NewAiComputingOrderModel(db *gorm.DB) *AiComputingOrderModel {
	return &AiComputingOrderModel{db: db}
}

func (m *AiComputingOrderModel) Insert(res *schema.AiComputingOrder) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *AiComputingOrderModel) UpdateSchema(data *schema.AiComputingOrder) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *AiComputingOrderModel) FindById(id int64) (res schema.AiComputingOrder, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *AiComputingOrderModel) GetList(model *schema.AiComputingOrder, startTime, endTime int64, page, pageSize int) (list []*schema.AiComputingOrder, total int64, err error) {
	q := m.db.Model(&schema.AiComputingOrder{})
	if startTime != 0 {
		q = q.Where("created_at >= ?", startTime)
	}
	if endTime != 0 {
		q = q.Where("created_at <= ?", endTime)
	}
	if model.UAddress != "" {
		q = q.Where("u_address = ?", model.UAddress)
	}

	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *AiComputingOrderModel) FindByIds(ids []int64) (res []schema.AiComputingOrder, err error) {
	if len(ids) == 0 {
		return
	}
	err = m.db.Find(&res, "id in ?", ids).Error
	return
}

func (m *AiComputingOrderModel) FindDetailById(id int64) (res schema.AiComputingOrder, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}
