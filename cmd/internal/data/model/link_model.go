package model

import (
	"air-drop/cmd/internal/data/dto"
	"air-drop/cmd/internal/data/schema"
	"gorm.io/gorm"
	"time"
)

type LinkModel struct {
	db *gorm.DB
}

func NewLinkModel(db *gorm.DB) *LinkModel {
	return &LinkModel{db: db}
}

func (m *LinkModel) Insert(res *schema.ArLink) error {
	res.CreatedAt = time.Now().Unix()
	return m.db.Create(res).Error
}

func (m *LinkModel) UpdateSchema(data *schema.ArLink) error {
	return m.db.Where("id = ?", data.ID).Save(data).Error
}

func (m *LinkModel) FindById(id int64) (res schema.ArLink, err error) {
	err = m.db.Find(&res, "id = ?", id).Error
	return
}

func (m *LinkModel) GetList(model *schema.ArLink, startTime, endTime int64, page, pageSize int) (list []*schema.ArLink, total int64, err error) {
	q := m.db.Model(&schema.ArLink{})
	if startTime != 0 {
		q = q.Where("created_at >= ?", startTime)
	}
	if endTime != 0 {
		q = q.Where("created_at <= ?", endTime)
	}
	q = q.Where("deleted_at is null")
	err = q.Count(&total).Error
	err = q.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func (m *LinkModel) FindByIds(ids []int64) (res []dto.LinkDetail, err error) {
	if len(ids) == 0 {
		return
	}
	err = m.db.Model(&schema.ArLink{}).
		Unscoped().
		Select("ar_link.*,ar_project.icon as project_icon,ar_project.name as project_name").
		Joins("left join ar_project on ar_link.project_ids=ar_project.id").
		Find(&res, "ar_link.id in ?", ids).Error
	return
}

func (m *LinkModel) FindDetailById(id int64) (res dto.LinkDetail, err error) {
	err = m.db.Model(&schema.ArLink{}).
		Unscoped().
		Select("ar_link.*,ar_project.icon as project_icon,ar_project.name as project_name").
		Joins("left join ar_project on ar_link.project_ids=ar_project.id").
		Find(&res, "ar_link.id = ?", id).Error
	return
}
