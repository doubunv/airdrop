package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type LinkReceiveModel struct {
	db *gorm.DB
}

const LinkReceiveStatus1 = 1
const LinkReceiveStatus2 = 2
const LinkReceiveStatus3 = 3

func NewLinkReceiveModel(db *gorm.DB) *LinkReceiveModel {
	return &LinkReceiveModel{db: db}
}

func (m *LinkReceiveModel) Insert(res *schema.ArLinkReceive) error {
	return m.db.Create(res).Error
}

func (m *LinkReceiveModel) UpdateSchema(data *schema.ArLinkReceive) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *LinkReceiveModel) FindById(id int64) (res schema.ArLinkReceive, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *LinkReceiveModel) GetList(model *schema.ArLinkReceive, startTime, endTime int64, page, pageSize int) (list []*schema.ArLinkReceive, total int64, err error) {
	q := m.db.Model(&schema.ArLinkReceive{})
	if model.UserId != 0 {
		q = q.Where("user_id = ?", model.UserId)
	}
	if model.UAddress != "" {
		q = q.Where("u_address = ?", model.UAddress)
	}
	if startTime != 0 {
		q = q.Where("create_at >= ?", startTime)
	}
	if endTime != 0 {
		q = q.Where("create_at <= ?", endTime)
	}
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *LinkReceiveModel) SumLinkReceive(data *schema.ArLinkReceive) (resp *schema.ArLinkReceive, err error) {
	q := m.db.Model(&schema.ArLinkReceive{})
	if data.UserId != 0 {
		q.Where("user_id", data.UserId)
	}
	if data.Status != 0 {
		q.Where("status", data.Status)
	}
	err = q.Select("sum(amount) as amount").Find(&resp).Error
	return
}

func (m *LinkReceiveModel) FindByOrderId(orderId int64) (res schema.ArLinkReceive, err error) {
	err = m.db.Find(&res, "order_id = ?", orderId).Error
	return
}
