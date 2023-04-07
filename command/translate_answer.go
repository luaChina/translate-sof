package command

import (
	"context"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
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
		posts, err := stackoverflow.Posts{}.GetAnswersPageByCondition(ctx, page, pagesize)
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
func processAnswer(ctx context.Context, post stackoverflow.Posts) error {
	if len(post.Body) > 4096 {
		return nil
	}
	converter := md.NewConverter("", true, nil)
	fmt.Println(time.Now(), post.Id)
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
	return saveAnswerToComment(ctx, post, result)
}

// saveAnswerToComment .
func saveAnswerToComment(ctx context.Context, post stackoverflow.Posts, result string) error {
	user, err := lua_china.Users{}.FindOrCreate(ctx)
	return nil
}
