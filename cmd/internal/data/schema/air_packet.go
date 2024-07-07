package schema

type AirPackage struct {
	ID        int64   `gorm:"column:id"`
	Name      string  `gorm:"column:name"`
	ChildId   string  `gorm:"column:child_id"`
	Price     float64 `gorm:"column:price"`
	Month     int64   `gorm:"column:month"`
	CreatedAt int64   `gorm:"column:created_at"`
	UpdatedAt int64   `gorm:"column:updated_at"`
}

func (m AirPackage) TableName() string {
	return "air_package"
}
