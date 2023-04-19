package command

import (
	"context"
	"github.com/luaChina/translate-sof/model/lua_china"
	"github.com/luaChina/translate-sof/model/stackoverflow"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strconv"
)

const Stackoverflow0IdMap = 88888888

// MigratePost .
func MigratePost(ctx context.Context) error {
	sofIds, err := lua_china.SofPostTranslate{}.GetAllSofPostId(ctx)
	if err != nil {
		return err
	}
	translates, err := stackoverflow.PostTranslate{}.GetAllNotInRelation(ctx, sofIds)
	if err != nil {
		return err
	}
	for _, translate := range translates {
		userId := translate.StackoverflowPost.OwnerUserId
		if userId == 0 {
			userId = translate.StackoverflowPost.LastEditorUserId
		}
		if userId == 0 || userId == -1 {
			userId = Stackoverflow0IdMap
		}
		_, err := lua_china.Users{}.GetByUserId(ctx, userId)
		if err == gorm.ErrRecordNotFound {
			if err := CreateSofUser(ctx, userId); err != nil {
				return err
			}
		} else if err == nil {
			post := lua_china.Posts{
				PostTagId: 2,
				UserId:    userId,
				Title:     translate.Title,
				Content:   translate.Body,
				CreatedAt: translate.StackoverflowPost.CreationDate,
				UpdatedAt: translate.StackoverflowPost.LastEditDate,
				Source:    1,
			}
			if err := post.Create(ctx); err != nil {
				return err
			}
			if err := (lua_china.SofPostTranslate{
				Id:     translate.StackoverflowPost.Id,
				PostId: post.Id,
			}).Create(ctx); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func MigrateUsers(ctx context.Context) error {
	translates, err := stackoverflow.PostTranslate{}.GetAll(ctx)
	if err != nil {
		return err
	}
	userIds := make([]int, 0)
	for _, translate := range translates {
		userIds = append(userIds, TransformSofUserId(translate))
	}
	userIds = lo.Uniq(userIds)
	for _, userId := range userIds {
		if err := CreateSofUser(ctx, userId); err != nil {
			return err
		}
	}
	return nil
}

func TransformSofUserId(translate stackoverflow.PostTranslate) int {
	userId := translate.StackoverflowPost.OwnerUserId
	if userId == 0 {
		userId = translate.StackoverflowPost.LastEditorUserId
	}
	if userId == 0 {
		userId = Stackoverflow0IdMap
	}
	return userId
}

func CreateSofUser(ctx context.Context, userId int) error {
	users := lua_china.Users{
		Id:       userId,
		Name:     "stackoverflow用户" + strconv.Itoa(userId),
		Phone:    strconv.Itoa(userId),
		Source:   1,
		Password: "",
	}
	if err := users.Create(ctx); err != nil {
		return err
	}
	return nil
}
