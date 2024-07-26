package schema

type AiComputingOrder struct {
	ID            int64   `gorm:"column:id"`
	UserId        int64   `gorm:"column:user_id"`
	AiComputingId int64   `gorm:"column:ai_computing_id"`
	Amount        float64 `gorm:"column:amount"`
	ServiceMonth  int64   `gorm:"column:service_month"`
	EndTime       int64   `gorm:"column:end_time"`
	SendAmount    float64 `gorm:"column:send_amount"`
	CreatedAt     int64   `gorm:"column:created_at"`
	UpdatedAt     int64   `gorm:"column:updated_at"`
	DeletedAt     int64   `gorm:"column:deleted_at"`
}

func (m AiComputingOrder) TableName() string {
	return "ai_computing_order"
}
