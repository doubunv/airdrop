package schema

type AiComputingOrder struct {
	ID             int64   `gorm:"column:id"`
	UserId         int64   `gorm:"column:user_id"`
	UAddress       string  `gorm:"column:u_address"`
	AiComputingId  int64   `gorm:"column:ai_computing_id"`
	Amount         float64 `gorm:"column:amount"`
	ServiceMonth   int64   `gorm:"column:service_month"`
	EndTime        int64   `gorm:"column:end_time"`
	WithdrawAmount float64 `gorm:"column:withdraw_amount"`
	SendAmount     float64 `gorm:"column:send_amount"`
	Status         int32   `gorm:"column:status"`
	BlockTime      int64   `gorm:"column:block_time"`
	TxHash         string  `gorm:"column:tx_hash"`
	CreatedAt      int64   `gorm:"column:created_at"`
	UpdatedAt      int64   `gorm:"column:updated_at"`
	DeletedAt      int64   `gorm:"column:deleted_at"`
}

func (m AiComputingOrder) TableName() string {
	return "ai_computing_order"
}
