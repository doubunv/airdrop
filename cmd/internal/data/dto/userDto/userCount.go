package userDto

import (
	"air-drop/pkg/systemType"
	"time"
)

type UserTeamCount struct {
	TeamLevel int   `json:"team_level"`
	Count     int64 `json:"count"`
}

type InviteData struct {
	ID                   int64             `gorm:"column:id"`
	UAddress             string            `gorm:"column:u_address"`
	TeamLevel            int64             `gorm:"column:team_level"`
	IsActive             int               `gorm:"column:is_active"`
	CreatedAt            time.Time         `gorm:"column:create_at"`
	TeamSpendAmount      systemType.Amount `gorm:"column:team_spend_amount"`
	SpendAmount          systemType.Amount `gorm:"column:spend_amount"`
	TotalTeamSpendAmount systemType.Amount `gorm:"column:total_team_spend_amount"`
}
