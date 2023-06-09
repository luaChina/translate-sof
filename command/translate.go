package command

import (
	"context"
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/luaChina/translate-sof/entity"
	"github.com/luaChina/translate-sof/model/lua_china"
	"github.com/luaChina/translate-sof/model/stackoverflow"
	"github.com/luaChina/translate-sof/service"
	"gorm.io/gorm"
	"time"
)

// TranslateSofAndSave .
func TranslateSofAndSave(ctx context.Context) error {
	page := 1
	pagesize := 100
	for {
		sofPostIds, err := lua_china.SofPostTranslate{}.GetAllSofPostId(ctx)
		if err != nil {
			return err
		}
		posts, err := stackoverflow.Posts{}.GetPageByCondition(ctx, sofPostIds, page, pagesize)
		if err != nil {
			return err
		}
		if len(posts) == 0 {
			break
		}
		for _, post := range posts {
			if err := processItem(ctx, post); err != nil {
				return err
			}
		}
	}
	return nil
}

func processItem(ctx context.Context, post stackoverflow.Posts) error {
	if len(post.Title) > 4096 || len(post.Body) > 4096 {
		if err := (lua_china.SofPostTranslate{
			Id:     post.Id,
			PostId: 0,
		}).Create(ctx); err != nil {
			return err
		}
		return nil
	}
	converter := md.NewConverter("", true, nil)
	fmt.Println(time.Now(), post.Id)
	markdownBody, err := converter.ConvertString(post.Body)
	if err != nil {
		return err
	}
	needTransMessage, err := json.Marshal(entity.TitleAndContent{
		Title:   post.Title,
		Content: markdownBody,
	})
	if err != nil {
		return err
	}
	query := fmt.Sprintf("将下面的 json 中 title 和 content 字段翻译成中文并且保留原本的 markdown 格式，然后json返回,\n %s", string(needTransMessage))
	fmt.Println(time.Now(), query)
	result, err := service.SendChatMessage(ctx, query)
	if err != nil {
		return err
	}
	var response entity.TitleAndContent
	if err := json.Unmarshal([]byte(result), &response); err != nil {
		return err
	}
	fmt.Println(time.Now(), response)
	return saveToLuaChina(ctx, post, response.Title, response.Content)
}

// save to luachina
func saveToLuaChina(ctx context.Context, sofPost stackoverflow.Posts, title, body string) error {
	_, err := lua_china.SofPostTranslate{}.GetBySofPostId(ctx, sofPost.Id)
	if err == gorm.ErrRecordNotFound {
		userId := sofPost.OwnerUserId
		if userId == 0 {
			userId = sofPost.LastEditorUserId
		}
		if userId == 0 || userId == -1 {
			userId = Stackoverflow0IdMap
		}
		_, err = lua_china.Users{}.GetByUserId(ctx, userId)
		if err == gorm.ErrRecordNotFound {
			if err := CreateSofUser(ctx, userId); err != nil {
				return err
			}
		} else if err == nil {
		} else {
			return err
		}
		article := lua_china.Posts{
			PostTagId: 2,
			UserId:    userId,
			Title:     title,
			Content:   body,
			CreatedAt: sofPost.CreationDate,
			UpdatedAt: sofPost.LastEditDate,
			Source:    1,
		}
		if err := article.Create(ctx); err != nil {
			return err
		}
		if err := (lua_china.SofPostTranslate{
			Id:     sofPost.Id,
			PostId: article.Id,
		}).Create(ctx); err != nil {
			return err
		}
		return nil
	} else if err == nil {
		return nil
	} else {
		return err
	}
}
