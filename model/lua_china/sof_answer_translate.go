package lua_china

import (
	"context"
	"gorm.io/gorm/clause"
)

// SofAnswerTranslate .
type SofAnswerTranslate struct {
	Id       int `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	AnswerId int `gorm:"column:answer_id;type:int(11) unsigned;" json:"answer_id"`
}

// TableName .
func (SofAnswerTranslate) TableName() string {
	return "sof_answer_translate"
}

// Create .
func (m SofAnswerTranslate) Create(ctx context.Context) error {
	return GetDB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&m).Error
}

// GetAllSofPostId .
func (m SofAnswerTranslate) GetAllSofPostId(ctx context.Context) ([]int, error) {
	var list []int
	err := GetDB(ctx).Model(&m).Select("id").Find(&list).Error
	return list, err
}

// GetBySofPostId .
func (m SofAnswerTranslate) GetBySofPostId(ctx context.Context, postId int) (SofPostTranslate, error) {
	var t SofPostTranslate
	err := GetDB(ctx).Where("id", postId).Take(&t).Error
	return t, err
}
