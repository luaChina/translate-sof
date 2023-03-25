package lua_china

import (
	"context"
	"fmt"
	"github.com/luaChina/translate-sof/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

// GetDB .
func GetDB(ctx context.Context) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/lua_china?charset=utf8mb4&parseTime=True&loc=Local",
		config.SecretConfig.User, config.SecretConfig.Password, config.SecretConfig.Host, config.SecretConfig.Port)
	if instance == nil {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		instance = db
	}
	return instance
}
