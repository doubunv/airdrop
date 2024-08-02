package schema

import (
	"fmt"
	"strconv"
	"strings"
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

// 向上返回邀请顺序层级，1->n级
func (p UserPath) GetParentIdByNumber(input string, number int) []int64 {
	stringSlice := strings.Split(input, ",")
	if len(stringSlice) > number {
		stringSlice = stringSlice[len(stringSlice)-number:]
	}
	var intSlice []int64
	for _, str := range stringSlice {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			continue
		}
		intSlice = append(intSlice, num)
	}

	reversedMap := make(map[int]int64)
	for i := len(intSlice) - 1; i >= 0; i-- {
		reversedMap[len(intSlice)-i] = intSlice[i]
	}

	return intSlice
}

type User struct {
	ID                 int64    `gorm:"column:id"`
	Pid                int64    `gorm:"column:pid"`
	UAddress           string   `gorm:"column:u_address"`
	ParentAddress      string   `gorm:"column:parent_address"`
	Amount             float64  `gorm:"column:amount"`
	PayAmount          float64  `gorm:"column:pay_amount"`
	TotalCommission    float64  `gorm:"column:total_commission"`
	Version            int      `gorm:"column:version"`
	InviteCode         string   `gorm:"column:invite_code"`
	Path               UserPath `gorm:"column:path"`
	InvitePathDistance int      `gorm:"column:invite_path_distance"`
	CreatedAt          int64    `gorm:"column:created_at"`
	UpdatedAt          int64    `gorm:"column:updated_at"`
}

func (u User) TableName() string {
	return "user"
}
