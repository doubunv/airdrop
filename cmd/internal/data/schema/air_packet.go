package schema

type AirPackage struct {
	ID         int64   `gorm:"column:id"`
	Name       string  `gorm:"column:name"`
	ProjectIds string  `gorm:"column:project_ids"`
	Price      float64 `gorm:"column:price"`
	Month      int64   `gorm:"column:month"`
	Status     int32   `gorm:"column:status"`
	CreatedAt  int64   `gorm:"column:created_at"`
	UpdatedAt  int64   `gorm:"column:updated_at"`
}

func (m AirPackage) TableName() string {
	return "air_package"
}
