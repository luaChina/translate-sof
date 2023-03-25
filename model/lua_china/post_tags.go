package lua_china

import (
	"database/sql"
	"time"
)

type PostTags struct {
	Id        uint         `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Type      string       `gorm:"column:type;type:varchar(255);NOT NULL" json:"type"`
	CreatedAt time.Time    `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}
