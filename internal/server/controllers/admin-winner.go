package controllers

import (
	"github.com/kataras/iris/v12"
	"raffles/internal/frontend/components/parts"
)

func DrawRun(c iris.Context) {
	// DB QUERY
	c.RenderComponent(parts.DrawWinnerLoading())
}

func DrawWinner(c iris.Context) {
	c.RenderComponent(parts.DrawWinnerInfo())
}
