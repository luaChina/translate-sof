package stackoverflow

import (
	"context"
	"fmt"
	"github.com/luaChina/translate-sof/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB

// GetDB .
func GetDB(ctx context.Context) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/stackoverflow?charset=utf8mb4&parseTime=True&loc=Local",
		config.SecretConfig.User, config.SecretConfig.Password, config.SecretConfig.Host, config.SecretConfig.Port)
	if instance == nil {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = db
	}
	return instance
}
