package schema

type AmountChangeType struct {
	ID        int64  `gorm:"column:id"`
	TypeId    int64  `gorm:"column:type_id"`
	TypeName  string `gorm:"column:type_name"`
	Type      int64  `gorm:"column:type"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (AmountChangeType) TableName() string { return "amount_change_type" }
