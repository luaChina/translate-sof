package command

import (
	"context"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/luaChina/translate-sof/model/lua_china"
	"github.com/luaChina/translate-sof/model/stackoverflow"
	"github.com/luaChina/translate-sof/service"
	"gorm.io/gorm"
	"strings"
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
	converter := md.NewConverter("", true, nil)
	fmt.Println(post.Id)
	query := fmt.Sprintf("将下面内容翻译成中文只显示翻译内容 %s", post.Title)
	result, err := service.SendChatMessage(ctx, query)
	if err != nil {
		return err
	}
	title := strings.Trim(result, "\n")
	fmt.Println(title)
	markdownBody, err := converter.ConvertString(post.Body)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("将下面内容翻译成中文只显示翻译内容，保留原本的 markdown 格式 %s", markdownBody)
	result, err = service.SendChatMessage(ctx, query)
	if err != nil {
		return err
	}
	body := strings.Trim(result, "\n")
	fmt.Println(body)
	if err := saveToLuaChina(ctx, post, title, body); err != nil {
		return err
	}
	return nil
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
		} else {
			return err
		}
		if err := (stackoverflow.PostTranslate{
			Id:    sofPost.Id,
			Title: "",
			Body:  "",
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
