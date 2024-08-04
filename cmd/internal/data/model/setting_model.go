package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type SettingModel struct {
	db *gorm.DB
}

func NewSettingModel(db *gorm.DB) *SettingModel {
	return &SettingModel{db: db}
}

func (m *SettingModel) Insert(res *schema.Setting) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *SettingModel) UpdateByKey(data *schema.Setting) error {
	return m.db.Where("setting_key = ?", data.Key).Save(data).Error
}

func (m *SettingModel) FindByKey(key string) string {
	res := schema.Setting{}
	_ = m.db.Find(&res, "setting_key = ?", key).Error
	return res.Value
}

func (m *SettingModel) GetList(model *schema.Setting, startTime, endTime int64, page, pageSize int) (list []*schema.Setting, total int64, err error) {
	q := m.db.Model(&schema.Setting{})
	q = q.Where("deleted_at is null")
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}
