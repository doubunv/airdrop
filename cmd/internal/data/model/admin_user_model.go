package model

import (
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
)

type AdminUserModel struct {
	db *gorm.DB
}

func NewAdminUserModel(db *gorm.DB) *AdminUserModel {
	return &AdminUserModel{db: db}
}

func (m *AdminUserModel) GetUserByUAddress(uAddress string) (res schema.AdminUser, err error) {
	err = m.db.Find(&res, "u_address = ?", uAddress).Error
	return
}
