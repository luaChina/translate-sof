package stackoverflow

import (
	"context"
	"gorm.io/gorm/clause"
)

// PostTranslate .
type PostTranslate struct {
	Id                int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Title             string `gorm:"column:title;type:varchar(5000);NOT NULL" json:"title"`
	Body              string `gorm:"column:body;type:varchar(5000);NOT NULL" json:"body"`
	StackoverflowPost Posts  `gorm:"foreignKey:id;references:Id"`
}

// TableName .
func (m PostTranslate) TableName() string {
	return "post_translate"
}

// Create .
func (m PostTranslate) Create(ctx context.Context) error {
	return GetDB(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]any{"title": m.Title, "body": m.Body}),
	}).Create(&m).Error
}

// GetLatestRow .
func (m PostTranslate) GetLatestRow(ctx context.Context) (PostTranslate, error) {
	var t PostTranslate
	err := GetDB(ctx).Order("id desc").Take(&t).Error
	return t, err
}

// GetAll .
func (m PostTranslate) GetAll(ctx context.Context) ([]PostTranslate, error) {
	var list []PostTranslate
	err := GetDB(ctx).Preload("StackoverflowPost").Find(&list).Error
	return list, err
}

// GetAllNotInRelation .
func (m PostTranslate) GetAllNotInRelation(ctx context.Context, sofIds []int) ([]PostTranslate, error) {
	var list []PostTranslate
	err := GetDB(ctx).Where("id not in ?", sofIds).Preload("StackoverflowPost").Find(&list).Error
	return list, err
}
