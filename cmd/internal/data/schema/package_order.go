package schema

type PackageOrder struct {
	ID              int64   `gorm:"column:id"`
	UserId          int64   `gorm:"column:user_id"`
	UAddress        string  `gorm:"column:u_address"`
	PackageId       int64   `gorm:"column:package_id"`
	Amount          float64 `gorm:"column:amount"`
	BuyMonth        int64   `gorm:"column:buy_month"`
	EndTime         int64   `gorm:"column:end_time"`
	SendEarnings    float64 `gorm:"column:send_earnings"`
	MaxEarningsRate float64 `gorm:"column:max_earnings_rate"`
	PayStatus       int32   `gorm:"column:pay_status"`
	BlockTime       int64   `gorm:"column:block_time"`
	TxHash          string  `gorm:"column:tx_hash"`
	CreatedAt       int64   `gorm:"column:created_at"`
	UpdatedAt       int64   `gorm:"column:updated_at"`
}

func (PackageOrder) TableName() string { return "package_order" }
