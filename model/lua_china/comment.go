package lua_china

import (
	"context"
	"database/sql"
	"time"
)

type Comment struct {
	Id        int          `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int          `gorm:"column:user_id;type:int(10) unsigned;NOT NULL" json:"user_id"`
	PostId    int          `gorm:"column:post_id;type:int(10) unsigned;NOT NULL" json:"post_id"`
	Content   string       `gorm:"column:content;type:text;NOT NULL" json:"content"`
	Source    int          `gorm:"column:source;type:tinyint(1) unsigned;NOT NULL" json:"source"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

// TableName .
func (Comment) TableName() string {
	return "comments"
}

// Create .
func (m *Comment) Create(ctx context.Context) error {
	return GetDB(ctx).Create(&m).Error
}
