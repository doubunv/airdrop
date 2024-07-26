package schema

import "gorm.io/plugin/soft_delete"

type AiComputing struct {
	ID                  int64                 `gorm:"column:id"`
	Icon                string                `gorm:"column:icon"`
	Name                string                `gorm:"column:name"`
	Content             string                `gorm:"column:content"`
	Price               float64               `gorm:"column:price"`
	ComputingPowerValue int64                 `gorm:"column:computing_power_value"`
	ComputingPowerUnit  string                `gorm:"column:computing_power_unit"`
	ServiceMonth        int64                 `gorm:"column:service_month"`
	Status              int64                 `gorm:"column:status"`
	CreatedAt           int64                 `gorm:"column:created_at"`
	UpdatedAt           int64                 `gorm:"column:updated_at"`
	DeletedAt           soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(20);default:null"`
}

func (m AiComputing) TableName() string {
	return "ai_computing"
}
