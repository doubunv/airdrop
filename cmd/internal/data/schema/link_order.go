package schema

type LinkOrder struct {
	ID        int64   `gorm:"column:id"`
	UserId    int64   `gorm:"column:user_id"`
	LinkId    int64   `gorm:"column:link_id"`
	BuyAmount float64 `gorm:"column:buy_amount"`
	DropTime  int64   `gorm:"column:drop_time"`
	BuyNumber int64   `gorm:"column:buy_number"`
	Status    int64   `gorm:"column:status"`
	CreatedAt int64   `gorm:"column:created_at"`
	UpdatedAt int64   `gorm:"column:updated_at"`
}

func (LinkOrder) TableName() string { return "link_order" }
