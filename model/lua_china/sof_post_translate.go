package lua_china

import (
	"context"
	"gorm.io/gorm/clause"
)

// SofPostTranslate .
type SofPostTranslate struct {
	Id     int `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	PostId int `gorm:"column:post_id;type:int(11) unsigned;" json:"post_id"`
}

// TableName .
func (SofPostTranslate) TableName() string {
	return "sof_post_translate"
}

// Create .
func (m SofPostTranslate) Create(ctx context.Context) error {
	return GetDB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&m).Error
}

// GetAllSofPostId .
func (m SofPostTranslate) GetAllSofPostId(ctx context.Context) ([]int, error) {
	var list []int
	err := GetDB(ctx).Model(&m).Select("id").Find(&list).Error
	return list, err
}

// UpdatePostId .
func (m SofPostTranslate) UpdatePostId(ctx context.Context, id, postId int) error {
	return GetDB(ctx).Model(&m).Where("id", id).UpdateColumn("post_id", postId).Error
}

// GetBySofPostId .
func (m SofPostTranslate) GetBySofPostId(ctx context.Context, id int) (SofPostTranslate, error) {
	var t SofPostTranslate
	err := GetDB(ctx).Where("id", id).Take(&t).Error
	return t, err
}
