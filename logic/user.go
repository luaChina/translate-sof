package logic

import (
	"context"
	"github.com/luaChina/translate-sof/consts"
	"github.com/luaChina/translate-sof/model/lua_china"
	"gorm.io/gorm"
	"strconv"
)

func FindOrCreateUser(ctx context.Context, ownerUserId, lastEditorUserId int) (*lua_china.Users, error) {
	userId := ownerUserId
	if userId == 0 {
		userId = lastEditorUserId
	}
	if userId == 0 || userId == -1 {
		userId = consts.Stackoverflow0IdMap
	}
	user, err := lua_china.Users{}.GetByUserId(ctx, userId)
	if err == gorm.ErrRecordNotFound {
		user := lua_china.Users{
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
