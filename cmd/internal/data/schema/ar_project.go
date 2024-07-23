package schema

import "gorm.io/plugin/soft_delete"

type AirPackageChild struct {
	ID        int64                 `gorm:"column:id"`
	Icon      string                `gorm:"column:icon"`
	Name      string                `gorm:"column:name"`
	Content   string                `gorm:"column:content"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(20);default:null"`
	CreatedAt int64                 `gorm:"column:created_at"`
	UpdatedAt int64                 `gorm:"column:updated_at"`
}

func (m AirPackageChild) TableName() string {
	return "ar_project"
}
