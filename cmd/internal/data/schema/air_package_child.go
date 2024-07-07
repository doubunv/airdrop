package schema

import "time"

type AirPackageChild struct {
	ID        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (m AirPackageChild) TableName() string {
	return "air_package_child"
}
