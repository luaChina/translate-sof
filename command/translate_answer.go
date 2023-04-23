package command

import (
	"context"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/luaChina/translate-sof/consts"
	"github.com/luaChina/translate-sof/logic"
	"github.com/luaChina/translate-sof/model/lua_china"
	"github.com/luaChina/translate-sof/model/stackoverflow"
	"github.com/luaChina/translate-sof/service"
	"time"
)

// TranslateAnswer .
func TranslateAnswer(ctx context.Context) error {
	page := 1
	pagesize := 100
	for {
		sofPostIds, err := lua_china.SofAnswerTranslate{}.GetAllSofPostId(ctx)
		if err != nil {
			return err
		}
		posts, err := stackoverflow.Answer{}.GetPageByCondition(ctx, sofPostIds, page, pagesize)
		if err != nil {
			return err
		}
		if len(posts) == 0 {
			break
		}
		for _, post := range posts {
			if err := processAnswer(ctx, post); err != nil {
				return err
			}
		}
	}
	return nil
}

// processAnswer .
func processAnswer(ctx context.Context, post stackoverflow.Answer) error {
	if len(post.Body) > 4096 {
		if err := (lua_china.SofAnswerTranslate{
			Id:       post.Id,
			AnswerId: 0,
		}).Create(ctx); err != nil {
			return err
		}
		return nil
	}
	converter := md.NewConverter("", true, nil)
	begin := time.Now()
	fmt.Println(begin, post.Id)
	markdownBody, err := converter.ConvertString(post.Body)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("将下面翻译成中文并且保留原本的 markdown 格式,\n %s", markdownBody)
	fmt.Println(time.Now(), query)
	result, err := service.SendChatMessage(ctx, query)
	if err != nil {
		return err
	}
	fmt.Println(time.Now(), result)
	if time.Now().Sub(begin).Seconds() < 20 {
		time.Sleep(time.Duration(20-time.Now().Sub(begin).Seconds()) * time.Second)
	}
	return saveAnswerToComment(ctx, post, result)
}

// saveAnswerToComment .
func saveAnswerToComment(ctx context.Context, post stackoverflow.Answer, result string) error {
	user, err := logic.FindOrCreateUser(ctx, post.OwnerUserId, post.LastEditorUserId)
	if err != nil {
		return err
	}
	fmt.Println(user)
	sofPost, err := lua_china.SofPostTranslate{}.GetBySofPostId(ctx, post.ParentId)
	if err != nil {
		return err
	}
	comment := lua_china.Comment{
		PostId:    sofPost.PostId,
		UserId:    user.Id,
		Content:   result,
		Source:    consts.SourceStackOverFlow,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := comment.Create(ctx); err != nil {
		return err
	}
	if err := (lua_china.SofAnswerTranslate{
		Id:       post.Id,
		AnswerId: comment.Id,
	}).Create(ctx); err != nil {
		return err
	}
	return nil
}
