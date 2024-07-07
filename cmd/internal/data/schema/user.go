package schema

import (
	"fmt"
	"time"
)

type UserPath string

func (p UserPath) Append(pid int64) UserPath {
	var res UserPath
	if p == "" {
		res = UserPath(fmt.Sprintf("%d", pid))
	} else {
		res = UserPath(fmt.Sprintf("%s,%d", p, pid))
	}

	return res
}

type User struct {
	ID                 int64     `gorm:"column:id"`
	UAddress           string    `gorm:"column:u_address"`
	Level              int64     `gorm:"column:level"`
	ParentAddress      string    `gorm:"column:parent_address"`
	Amount             float64   `gorm:"column:amount"`
	Version            int       `gorm:"column:version"`
	TeamLevel          int64     `gorm:"column:team_level"`
	InviteCode         string    `gorm:"column:invite_code"`
	Path               UserPath  `gorm:"column:path"`
	InvitePathDistance int       `gorm:"column:invite_path_distance"`
	CreateAt           time.Time `gorm:"column:create_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

func (u User) TableName() string {
	return "user"
}
