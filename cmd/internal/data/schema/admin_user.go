package schema

type AdminUser struct {
	ID       int64  `gorm:"column:id"`
	UAddress string `gorm:"column:u_address"`
	Role     string `gorm:"column:role"`
}

func (*AdminUser) TableName() string {
	return "admin_user"
}
