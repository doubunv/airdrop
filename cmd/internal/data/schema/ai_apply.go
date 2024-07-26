package schema

type AiApply struct {
	ID                 int64   `gorm:"column:id"`
	AiComputingOrderId int64   `gorm:"column:ai_computing_order_id"`
	UserId             int64   `gorm:"column:user_id"`
	UAddress           string  `gorm:"column:u_address"`
	Amount             float64 `gorm:"column:amount"`
	CreatedAt          int64   `gorm:"column:created_at"`
	UpdatedAt          int64   `gorm:"column:updated_at"`
}

func (m AiApply) TableName() string {
	return "ai_apply"
}
