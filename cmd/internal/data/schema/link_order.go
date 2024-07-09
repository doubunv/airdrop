package schema

type LinkOrder struct {
	ID         int64   `gorm:"column:id"`
	UserId     int64   `gorm:"column:user_id"`
	Type       int64   `gorm:"column:type"`
	TargetId   int64   `gorm:"column:target_id"`
	Amount     float64 `gorm:"column:amount"`
	DropTime   int64   `gorm:"column:drop_time"`
	BuyNumber  int64   `gorm:"column:buy_number"`
	LaveNumber int64   `gorm:"column:lave_number"`
	Status     int64   `gorm:"column:status"`
	CreatedAt  int64   `gorm:"column:created_at"`
	UpdatedAt  int64   `gorm:"column:updated_at"`
}

func (LinkOrder) TableName() string { return "link_order" }
