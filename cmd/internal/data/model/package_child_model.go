package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type PackageChildModel struct {
	db *gorm.DB
}

func NewPackageChildModel(db *gorm.DB) *PackageChildModel {
	return &PackageChildModel{db: db}
}

func (m *PackageChildModel) Insert(res *schema.AirPackageChild) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *PackageChildModel) UpdateSchema(data *schema.AirPackageChild) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *PackageChildModel) FindById(id int64) (res schema.AirPackageChild, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *PackageChildModel) FindByIds(ids []int64) (res []schema.AirPackageChild, err error) {
	if len(ids) == 0 {
		return
	}
	err = m.db.Find(&res, "id in ?", ids).Error
	return
}

func (m *PackageChildModel) GetList(model *schema.AirPackageChild, startTime, endTime int64, page, pageSize int) (list []*schema.AirPackageChild, total int64, err error) {
	q := m.db.Model(&schema.AirPackageChild{})
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
