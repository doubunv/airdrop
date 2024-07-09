package schema

type ArLinkReceive struct {
	ID        int64   `gorm:"column:id"`
	UserId    int64   `gorm:"column:user_id"`
	UAddress  string  `gorm:"column:u_address"`
	Amount    float64 `gorm:"column:amount"`
	Status    int64   `gorm:"column:status"`
	DropTime  int64   `gorm:"column:drop_time"`
	CreatedAt int64   `gorm:"column:created_at"`
	UpdatedAt int64   `gorm:"column:updated_at"`
}

func (ArLinkReceive) TableName() string { return "ar_link_apply" }
