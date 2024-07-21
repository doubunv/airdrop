package schema

import "gorm.io/plugin/soft_delete"

type ArLink struct {
	ID          int64                 `gorm:"column:id"`
	ProjectIds  string                `gorm:"column:project_ids"`
	DropTime    int64                 `gorm:"column:drop_time"`
	Price       float64               `gorm:"column:price"`
	Status      int64                 `gorm:"column:status"`
	SellEndTime int64                 `gorm:"column:sell_end_time"`
	DropAmount  float64               `gorm:"column:drop_amount"`
	CreatedAt   int64                 `gorm:"column:created_at"`
	UpdatedAt   int64                 `gorm:"column:updated_at"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(20);default:null"`
}

func (m ArLink) TableName() string {
	return "air_link"
}
