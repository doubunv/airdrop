package model

import "gorm.io/gorm"

type P3Price struct {
	ID         int64   `gorm:"id"`
	TxHash     string  `gorm:"tx_hash"`
	Price      float64 `gorm:"price"`
	CreateTime int64   `gorm:"create_time"`
}

func (m P3Price) TableName() string {
	return "mm_p3_price"
}

type ContractModel struct {
	db *gorm.DB
}

func NewContractModel(db *gorm.DB) *ContractModel {
	return &ContractModel{db: db}
}

func (m *ContractModel) GetPage() (res []P3Price, err error) {
	err = m.db.Find(&res).Error
	return
}
func (m *ContractModel) GetLastP3() (price float64, err error) {
	err = m.db.Table("mm_p3_price").Where("price > 0").Order("id desc").Limit(1).Select("price").Find(&price).Error
	return
}

func (m *ContractModel) InsertP3Price(res *P3Price) error {
	return m.db.Create(res).Error
}
