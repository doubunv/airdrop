package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type PackageChildModel struct {
	db *gorm.DB
}

func NewPackageChildModel(db *gorm.DB) *PackageChildModel {
	return &PackageChildModel{db: db}
}

func (m *PackageChildModel) Insert(res *schema.AirPackageChild) error {
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
