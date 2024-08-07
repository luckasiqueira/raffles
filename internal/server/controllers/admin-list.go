package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"raffles/internal/frontend/components"
	"raffles/internal/frontend/components/parts"
	"raffles/internal/server/database"
	"raffles/utils/logger"
)

// DrawList handles /admin/list
func DrawList(c iris.Context) {
	draws := database.DrawList()
	c.RenderComponent(components.AdminList(draws))
}

// DrawListSingle handles /admin/list/{number}
func DrawListSingle(c iris.Context) {
	number, err := c.Params().GetInt("number")
	if err != nil {
		logger.NewLog(err.Error())
		c.StopWithStatus(http.StatusBadRequest)
	}
	draw := database.DrawSingle(number)
	c.RenderComponent(parts.DrawSingle(draw))
}
