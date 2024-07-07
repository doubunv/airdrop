package schema

type Setting struct {
	ID        int64  `gorm:"column:id"`
	Key       string `gorm:"column:key"`
	Value     string `gorm:"column:value"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (Setting) TableName() string { return "setting" }
