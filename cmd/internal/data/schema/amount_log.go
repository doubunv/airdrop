package schema

type AmountLog struct {
	ID        int64   `gorm:"column:id"`
	UserId    int64   `gorm:"column:user_id"`
	UAddress  string  `gorm:"column:u_address"`
	Balance   float64 `gorm:"column:balance"`
	TargetId  int64   `gorm:"column:target_id"`
	TypeId    int64   `gorm:"column:type_id"`
	Mark      string  `gorm:"column:mark"`
	CreatedAt int64   `gorm:"column:created_at"`
	UpdatedAt int64   `gorm:"column:updated_at"`
}

func (m AmountLog) TableName() string {
	return "amount_log"
}
