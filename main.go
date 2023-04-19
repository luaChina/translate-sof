package main

import (
	"context"
	"github.com/luaChina/translate-sof/command"
	_ "github.com/luaChina/translate-sof/config"
)

func main() {
	if err := command.TranslateAnswer(context.Background()); err != nil {
		panic(err)
	}
}
