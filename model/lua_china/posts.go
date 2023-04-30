package lua_china

import (
	"context"
	"database/sql"
	"time"
)

// Posts .
type Posts struct {
	Id        int          `gorm:"primaryKey" json:"id"`
	PostTagId uint         `gorm:"column:post_tag_id;type:int(11) unsigned;NOT NULL" json:"post_tag_id"`
	UserId    int          `gorm:"column:user_id;type:int(11) unsigned;NOT NULL" json:"user_id"`
	Title     string       `gorm:"column:title;type:text;NOT NULL" json:"title"`
	Thumbnail string       `gorm:"column:thumbnail;type:varchar(255);NOT NULL" json:"thumbnail"`
	Content   string       `gorm:"column:content;type:text;NOT NULL" json:"content"`
	ReadCount uint         `gorm:"column:read_count;type:int(11) unsigned;default:0;NOT NULL" json:"read_count"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	Excellent int          `gorm:"column:excellent;type:tinyint(1);default:0;NOT NULL" json:"excellent"`
	Source    int          `gorm:"column:source"`
}

// Create .
func (m *Posts) Create(ctx context.Context) error {
	return GetDB(ctx).Create(&m).Error
}

// BatchCreate .
func (m Posts) BatchCreate(ctx context.Context, posts []Posts) error {
	return GetDB(ctx).Create(&posts).Error
}

// GetByTitle .
func (m Posts) GetByTitle(ctx context.Context, title string) (Posts, error) {
	var t Posts
	err := GetDB(ctx).Where("title", title).Take(&t).Error
	return t, err
}

// GetByPostId .
func (m Posts) GetByPostId(ctx context.Context, postId int) (Posts, error) {
	var post Posts
	err := GetDB(ctx).Where("id = ?", postId).Take(&post).Error
	return post, err
}

// Updates .
func (m Posts) UpdateDate(ctx context.Context, postId int, createdAt, updatedAt time.Time) error {
	return GetDB(ctx).Model(&m).Where("id", postId).UpdateColumns(map[string]any{
		"created_at": createdAt,
		"updated_at": updatedAt,
	}).Error
}
