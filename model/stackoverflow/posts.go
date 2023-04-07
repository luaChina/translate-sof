package stackoverflow

import (
	"context"
	"time"
)

// Posts .
type Posts struct {
	Id               int       `gorm:"column:Id;type:int(11);primary_key" json:"Id"`
	PostTypeId       int       `gorm:"column:PostTypeId;type:int(11);NOT NULL" json:"PostTypeId"`
	CreationDate     time.Time `gorm:"column:CreationDate;type:datetime;NOT NULL" json:"CreationDate"`
	Score            int       `gorm:"column:Score;type:int(11);NOT NULL" json:"Score"`
	ViewCount        int       `gorm:"column:ViewCount;type:int(11);NOT NULL" json:"ViewCount"`
	Body             string    `gorm:"column:Body;type:varchar(5000);NOT NULL" json:"Body"`
	OwnerUserId      int       `gorm:"column:OwnerUserId;type:int(11);NOT NULL" json:"OwnerUserId"`
	LastEditorUserId int       `gorm:"column:LastEditorUserId;type:int(11);NOT NULL" json:"LastEditorUserId"`
	LastEditDate     time.Time `gorm:"column:LastEditDate;type:datetime;NOT NULL" json:"LastEditDate"`
	LastActivityDate time.Time `gorm:"column:LastActivityDate;type:datetime;NOT NULL" json:"LastActivityDate"`
	Title            string    `gorm:"column:Title;type:varchar(255);NOT NULL" json:"Title"`
	Tags             string    `gorm:"column:Tags;type:varchar(255);NOT NULL" json:"Tags"`
	AnswerCount      int       `gorm:"column:AnswerCount;type:int(11);NOT NULL" json:"AnswerCount"`
	CommentCount     int       `gorm:"column:CommentCount;type:int(11);NOT NULL" json:"CommentCount"`
	FavoriteCount    int       `gorm:"column:FavoriteCount;type:int(11);NOT NULL" json:"FavoriteCount"`
	ContentLicense   string    `gorm:"column:ContentLicense;type:varchar(255);NOT NULL" json:"ContentLicense"`
}

// TableName .
func (Posts) TableName() string {
	return "Posts"
}

// Create .
func (m Posts) Create(ctx context.Context) error {
	return GetDB(ctx).Create(&m).Error
}

// GetPageByCondition .
func (m Posts) GetPageByCondition(ctx context.Context, sofPostIds []int, page, pageSize int) ([]Posts, error) {
	var list []Posts
	err := GetDB(ctx).Where("Id not in (?)", sofPostIds).Where("PostTypeId", 1).Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	return list, err
}

// GetAnswersPageByCondition .
func (m Posts) GetAnswersPageByCondition(ctx context.Context, page, pageSize int) ([]Posts, error) {
	var list []Posts
	err := GetDB(ctx).Joins("left join post_translate on Posts.Id = post_translate.id").Where("Posts.PostTypeId", 2).Where("post_translate.Id is NULL").Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error
	return list, err
}
