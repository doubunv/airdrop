package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type AmountChangeTypeModel struct {
	db *gorm.DB
}

func NewAmountChangeTypeModel(db *gorm.DB) *AmountChangeTypeModel {
	return &AmountChangeTypeModel{db: db}
}

func (m *AmountChangeTypeModel) Insert(res *schema.AmountChangeType) error {
	return m.db.Create(res).Error
}

func (m *AmountChangeTypeModel) UpdateSchema(data *schema.AmountChangeType) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *AmountChangeTypeModel) FindById(id int64) (res schema.AmountChangeType, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}
