package schema

type Setting struct {
	ID        int64  `gorm:"column:id"`
	Key       string `gorm:"column:key"`
	Value     string `gorm:"column:value"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

func (Setting) TableName() string { return "setting" }

// 一级推广用户佣金
const invite_one_level_commission = "invite_one_level_commission"

// 二级推广用户佣金
const invite_two_level_commission = "invite_two_level_commission"

// 三级推广用户佣金
const invite_three_level_commission = "invite_three_level_commission"

////////////////////////
//value 格式：{performance:xx,commission:xxx}

// ai 算力，初级代理
const ai_computing_primary_level_commission = "ai_computing_primary_level_commission"

// ai 算力，中级代理
const ai_computing_middle_level_commission = "ai_computing_middle_level_commission"

// ai 算力，高级代理
const ai_computing_high_level_commission = "ai_computing_high_level_commission"

// ai 算力，合伙人
const ai_computing_partner_commission = "ai_computing_partner_commission"

// ai 算力，平均提成
const ai_computing_average_commission = "ai_computing_average_commission"
