package lua_china

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type Users struct {
	Id            int          `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name          string       `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Avatar        string       `gorm:"column:avatar;type:varchar(255);NOT NULL" json:"avatar"`
	Email         string       `gorm:"column:email;type:varchar(255);NOT NULL" json:"email"`
	City          string       `gorm:"column:city;type:varchar(255);NOT NULL" json:"city"`
	Phone         string       `gorm:"column:phone;type:varchar(50);NOT NULL" json:"phone"`
	Password      string       `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	RememberToken string       `gorm:"column:remember_token;type:varchar(100);NOT NULL" json:"remember_token"`
	CreatedAt     time.Time    `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt     time.Time    `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt     sql.NullTime `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	OauthId       string       `gorm:"column:oauth_id;type:varchar(255);NOT NULL" json:"oauth_id"`
	OauthFrom     string       `gorm:"column:oauth_from;type:varchar(255);NOT NULL" json:"oauth_from"`
	Source        int          `gorm:"column:source;type:int(11);default:0;NOT NULL" json:"source"`
}

// TableName .
func (Users) TableName() string {
	return "users"
}

// Create .
func (m *Users) Create(ctx context.Context) error {
	return GetDB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&m).Error
}

// BatchCreate .
func (m Users) BatchCreate(ctx context.Context, users []Users) error {
	return GetDB(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&users).Error
}

func (m Users) GetByUserId(ctx context.Context, userId int) (Users, error) {
	var user Users
	err := GetDB(ctx).Where("id = ?", userId).Take(&user).Error
	return user, err
}

// FindOrCreate .
func (m Users) FindOrCreate(ctx context.Context, userId int) (*Users, error) {
	user, err := m.GetByUserId(ctx, userId)
	if err == gorm.ErrRecordNotFound {
		user := Users{
			Id:       userId,
			Name:     "stackoverflow用户" + strconv.Itoa(userId),
			Phone:    strconv.Itoa(userId),
			Source:   1,
			Password: "",
		}
		if err := user.Create(ctx); err != nil {
			return nil, err
		}
		return &user, nil
	} else if err == nil {
		return &user, nil
	} else {
		return nil, err
	}
}
