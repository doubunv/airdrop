package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type PackageModel struct {
	db *gorm.DB
}

func NewPackageModel(db *gorm.DB) *PackageModel {
	return &PackageModel{db: db}
}

func (m *PackageModel) Insert(res *schema.AirPackage) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *PackageModel) UpdateSchema(data *schema.AirPackage) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *PackageModel) FindById(id int64) (res schema.AirPackage, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *PackageModel) GetList(model *schema.AirPackage, startTime, endTime int64, page, pageSize int) (list []*schema.AirPackage, total int64, err error) {
	q := m.db.Model(&schema.AirPackage{})
	if startTime != 0 {
		q = q.Where("create_at >= ?", startTime)
	}
	if endTime != 0 {
		q = q.Where("create_at <= ?", endTime)
	}
	q = q.Where("deleted_at is null")
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *PackageModel) FindByIds(ids []int64) (res []schema.AirPackage, err error) {
	if len(ids) == 0 {
		return
	}
	err = m.db.Find(&res, "id in ?", ids).Error
	return
}
