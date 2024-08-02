package schema

const (
	AmountLogTypeId1 = 1 //'1-佣金记录，
	AmountLogTypeId2 = 2 // 2-回报记录'
)

type AmountLog struct {
	ID            int64   `gorm:"column:id"`
	UserId        int64   `gorm:"column:user_id"`
	UAddress      string  `gorm:"column:u_address"`
	Balance       float64 `gorm:"column:balance"`
	TargetId      int64   `gorm:"column:target_id"`
	TargetUid     int64   `gorm:"column:target_uid"`
	TargetAddress string  `gorm:"column:target_address"`
	TypeId        int64   `gorm:"column:type_id"`
	Mark          string  `gorm:"column:mark"`
	CreatedAt     int64   `gorm:"column:created_at"`
	UpdatedAt     int64   `gorm:"column:updated_at"`
}

func (m AmountLog) TableName() string {
	return "amount_log"
}
