package schema

type ArLink struct {
	ID        int64   `gorm:"column:id"`
	Name      string  `gorm:"column:name"`
	DropTime  int64   `gorm:"column:drop_time"`
	Price     float64 `gorm:"column:price"`
	CreatedAt int64   `gorm:"column:created_at"`
	UpdatedAt int64   `gorm:"column:updated_at"`
}

func (m ArLink) TableName() string {
	return "air_link"
}
